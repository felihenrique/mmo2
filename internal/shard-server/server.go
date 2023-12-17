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
	world     *game.World
	players   map[string]*Player
	host      string
	port      int
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
	switch evType {
	case packets.TypeMoveInput:
		s.moveRequest(player, pe)
	case packets.TypeJoinShardRequest:
		s.joinShardRequest(player, pe)
	default:
		fmt.Printf("wrong request: %d", evType)
		return
	}
}

func (s *Server) handleChans() {
	ticker := time.NewTicker(time.Millisecond * 100)
	peerConnChan := s.gspServer.PeerConnChan()
	peerDisChan := s.gspServer.PeerDisChan()
	newEventsChan := s.gspServer.NewEventsChan()
	for s.gspServer.Listening() {
		select {
		case peer := <-peerDisChan:
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

func New(host string, port int) *Server {
	server := Server{}
	server.host = host
	server.port = port
	server.world = game.NewWorld(1000)
	server.gspServer = gsp.NewTcpServer()

	return &server
}

func (s *Server) ackInput(input events.Raw, peer *gsp.TcpPeer) {
	data := events.Serialize(&packets.AckInput{
		InputId: events.GetID(input),
	})
	peer.SendEvent(data)
}

func (s *Server) Broadcast(event serialization.ISerializable) {
	data := events.Serialize(event)
	for _, player := range s.players {
		if player.entity == nil {
			continue
		}
		player.peer.SendEvent(data)
	}
}

func (s *Server) BroadcastFiltered(event serialization.ISerializable, filterPeer *gsp.TcpPeer) {
	data := events.Serialize(event)
	for _, player := range s.players {
		if player.entity == nil {
			continue
		}
		player.peer.SendEvent(data)
	}
}

func (s *Server) Start() error {
	err := s.gspServer.Listen(s.host, s.port)
	if err != nil {
		return err
	}
	s.handleChans()
	return nil
}
