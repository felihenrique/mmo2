package shard

import (
	"fmt"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
	"mmo2/pkg/serialization"
	"time"
)

type Server struct {
	gspServer gsp.TcpServer
	handlers  map[int16]EventHandler
	world     *game.World
	players   map[string]*Player
}

type Player struct {
	entity *game.Entity
	peer   *gsp.TcpPeer
}

func (s *Server) handleEvent(pe gsp.PeerEvent) {
	player, ok := s.players[pe.Peer.Addr()]
	if !ok {
		fmt.Printf("wrong: player %s not found \n", player.peer.Addr())
		return
	}
	evType := events.GetType(pe.Event)
	handler := s.handlers[evType]
	if handler == nil {
		fmt.Printf("Handler for event %d not found \n", evType)
		return
	}
	handler(player, pe.Event)
}

func (s *Server) handleChans() {
	ticker := time.NewTicker(time.Millisecond * 100)
	peerConnChan := s.gspServer.PeerConnChan()
	peerDisChan := s.gspServer.PeerDisChan()
	newEventsChan := s.gspServer.NewEventsChan()
	for s.gspServer.Listening() {
		select {
		case peer := <-peerDisChan:
			player := s.players[peer.Addr()]
			if player.entity != nil {
				s.world.RemoveEntity(player.entity.ID())
			}
			delete(s.players, peer.Addr())
		case peer := <-peerConnChan:
			s.players[peer.Addr()] = &Player{
				entity: nil,
				peer:   peer,
			}
		case newEvent := <-newEventsChan:
			s.handleEvent(newEvent)
		case <-ticker.C:
			// DO TICK
		}
	}
	ticker.Stop()
}

func New() *Server {
	server := Server{}
	server.world = game.NewWorld()
	server.gspServer = gsp.NewTcpServer()
	server.players = make(map[string]*Player)
	server.handlers = make(map[int16]EventHandler)

	server.handlers[packets.TypeJoinShardRequest] = server.joinShardRequest
	server.handlers[packets.TypeMoveRequest] = server.moveRequest

	return &server
}

func (s *Server) Broadcast(event serialization.ISerializable) {
	data := event.ToBytes(0)
	for _, player := range s.players {
		if player.entity == nil {
			continue
		}
		player.peer.SendBytes(data)
	}
}

func (s *Server) BroadcastFiltered(event serialization.ISerializable, filterPeer *gsp.TcpPeer) {
	data := event.ToBytes(0)
	for _, player := range s.players {
		if player.entity == nil || player.peer.Addr() == filterPeer.Addr() {
			continue
		}
		player.peer.SendBytes(data)
	}
}

func (s *Server) Start(host string, port int) error {
	err := s.gspServer.Listen(host, port)
	if err != nil {
		return err
	}
	go s.handleChans()
	return nil
}

func (s *Server) World() *game.World {
	return s.world
}
