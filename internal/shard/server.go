package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"sync"
)

type Server struct {
	gspServer gsp.TcpServer
	queue     gsp.Queue[ICommand]
	peers     sync.Map
	host      string
	port      int
}

func New(host string, port int) *Server {
	server := Server{}
	server.host = host
	server.port = port
	server.gspServer = gsp.NewTcpServer()

	server.gspServer.OnPeerConnect(func(peer *gsp.TcpPeer) {
		server.peers.Store(peer.Id(), peer)
	})

	server.gspServer.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		server.peers.Delete(peer.Id())
	})

	server.gspServer.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer, eventBytes []byte) {
		event := events.MoveEvent{}
		err := event.FromBytes(eventBytes)
		if err != nil {
			println(err.Error())
			return
		}
		server.queue.Push(MoveCommand{
			payload: event,
		})
	})

	return &server
}

func (s *Server) Start() error {
	err := s.gspServer.Listen(s.host, s.port)
	if err != nil {
		return err
	}
	for {
		commands := s.queue.PopAll()
		for _, command := range commands {
			command.Execute()
		}
		// AQUI SERÁ CHAMADA A LóGICA PARA SIMULAR O MUNDO
		// NO CASO DO ECS, EXECUTAR OS SYSTEMS
		// E ENVIAR PARA OS JOGADORES OS EVENTOS RESULTANTES DA SIMULAÇÃO
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
