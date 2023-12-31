package shard

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/game/systems"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/gsp"
	"mmo2/pkg/serialization"
	"time"
)

type Server struct {
	gspServer gsp.IServer
	handlers  map[int16]EventHandler
	players   map[string]*Player
}

type Player struct {
	entity *ecs.Entity
	peer   gsp.IPeer
}

func (s *Server) handleEvent(pe gsp.PeerEvent) {
	player, ok := s.players[pe.Peer.Addr()]
	if !ok {
		fmt.Printf("wrong: player %s not found \n", player.peer.Addr())
		return
	}
	evType := event_utils.GetType(pe.Event)
	handler := s.handlers[evType]
	if handler == nil {
		fmt.Printf("Handler for event %d not found \n", evType)
		return
	}
	handler(player, pe.Event)
}

func (s *Server) handleChans() {
	ticker := time.NewTicker(time.Millisecond * 50)
	peerConnChan := s.gspServer.NewConnectionChan()
	peerDisChan := s.gspServer.DisconnectedChan()
	newEventsChan := s.gspServer.EventsChan()
	for s.gspServer.Listening() {
		select {
		case peer := <-peerDisChan:
			player := s.players[peer.Addr()]
			if player.entity != nil { // Player disconnect before join shard
				s.Broadcast(packets.NewEntityRemoved(player.entity.ID()))
				ecs.MainWorld.RemoveEntity(player.entity.ID())
			}
			delete(s.players, peer.Addr())
		case peer := <-peerConnChan:
			s.players[peer.Addr()] = &Player{
				entity: nil,
				peer:   peer,
			}
		case newEvent := <-newEventsChan:
			if event_utils.GetType(newEvent.Event) == packets.TypePing {
				newEvent.Peer.SendResponse(newEvent.Event, packets.NewAckRequest())
				continue
			}
			player := s.players[newEvent.Peer.Addr()]
			if player.entity == nil && event_utils.GetType(newEvent.Event) != packets.TypeJoinShardRequest {
				fmt.Printf("Peer %s cant send events. Need to join shard first \n", newEvent.Peer.Addr())
				newEvent.Peer.SendResponse(
					newEvent.Event,
					packets.NewRequestError("Can't send requests. Should join the shard first"),
				)
				continue
			}
			s.handleEvent(newEvent)
		case <-ticker.C:
			ecs.MainWorld.Update()
		}
	}
	ticker.Stop()
}

func New() *Server {
	server := Server{}
	server.gspServer = gsp.NewTcpServer()
	server.players = make(map[string]*Player)
	server.handlers = make(map[int16]EventHandler)
	ecs.MainWorld.AddSystem(systems.MoveSystem)

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

func (s *Server) BroadcastFiltered(event serialization.ISerializable, filterPeer gsp.IPeer) {
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
