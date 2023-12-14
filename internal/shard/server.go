package shard

import (
	"mmo2/pkg/ds"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"sync"
	"time"
)

type Server struct {
	gspServer    gsp.TcpServer
	commandQueue ds.ConcurrentQueue[ICommand]
	world        *game.World
	players      sync.Map
	host         string
	port         int
}

type Player struct {
	entityId int16
	peer     *gsp.TcpPeer
}

func New(host string, port int) *Server {
	server := Server{}
	server.host = host
	server.port = port
	server.world = game.NewWorld(1000)
	server.gspServer = gsp.NewTcpServer()

	server.gspServer.OnPeerConnect(func(peer *gsp.TcpPeer) {
		entity := server.world.NewEntity()
		entity.Add(&game.Transform{
			X:        0,
			Y:        0,
			Rotation: 0,
		})
		if entity == nil {
			peer.Close()
			return
		}
		player := Player{
			entityId: entity.ID(),
			peer:     peer,
		}
		server.players.Store(peer.Addr(), player)
	})

	server.gspServer.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		server.players.Delete(peer.Addr())
	})

	server.gspServer.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer, rawEvent events.RawEvent) {
		entry, ok := server.players.Load(peer.Addr())
		if !ok {
			return
		}
		move := events.Move{}
		events.Unserialize(rawEvent, &move)
		player := entry.(Player)
		server.commandQueue.Push(&MoveCommand{
			event:             move,
			eventId:           events.GetId(rawEvent),
			player:            player,
			world:             server.world,
			broadcastFiltered: server.BroadcastFiltered,
		})
	})

	return &server
}

func (s *Server) Broadcast(event events.ISerializable, id int16) {
	s.players.Range(func(key, value any) bool {
		player := value.(Player)
		player.peer.SendEvent(event, id)
		return true
	})
}

func (s *Server) BroadcastFiltered(event events.ISerializable, eventId int16, filterPeer *gsp.TcpPeer) {
	s.players.Range(func(key, value any) bool {
		player := value.(Player)
		if player.peer.Addr() == filterPeer.Addr() {
			return true
		}
		player.peer.SendEvent(event, eventId)
		return true
	})
}

func (s *Server) Start() error {
	err := s.gspServer.Listen(s.host, s.port)
	if err != nil {
		return err
	}
	for {
		commands := s.commandQueue.PopAll()
		for _, command := range commands {
			command.Execute()
		}
		// AQUI SERÁ CHAMADA A LóGICA PARA SIMULAR O MUNDO
		// NO CASO DO ECS, EXECUTAR OS SYSTEMS
		// E ENVIAR PARA OS JOGADORES OS EVENTOS RESULTANTES DA SIMULAÇÃO
		time.Sleep(time.Millisecond * 50)
	}
}

/*
Haverá um commando para cada tipo de evento. Na instancia do command (no metodo OnEvent)
serão colocados todas as referências que aquele comando precisa, como por exemplo o estado
do jogo, jogador que executou, entre outras coisas.
No comando de movimento, serão checados também as colisões com obstaculos e entidades.
Primeiramente serão processados todos os inputs, depois será processada a lógica do jogo
Por exemplo, checar se um monstro está próximo do jogador e caso esteja atacar ele.
*/
