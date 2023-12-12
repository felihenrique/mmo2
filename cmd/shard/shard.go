package main

import (
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"sync"
)

func main() {
	server := gsp.NewTcpServer(1000)
	peers := sync.Map{}
	queue := gsp.Queue[ICommand]{}

	server.OnPeerConnect(func(peer *gsp.TcpPeer) {
		peers.Store(peer.Id(), peer)
	})

	server.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		peers.Delete(peer.Id())
	})

	server.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer, eventBytes []byte) {
		event := events.MoveEvent{}
		err := event.FromBytes(eventBytes)
		if err != nil {
			println(err.Error())
			return
		}
		queue.Push(MoveCommand{
			payload: event,
		})
	})

	go func() {
		for {
			commands := queue.PopAll()
			for _, command := range commands {
				command.Execute()
			}
			// AQUI SERÁ CHAMADA A LóGICA PARA SIMULAR O MUNDO
			// NO CASO DO ECS, EXECUTAR OS SYSTEMS
			// E ENVIAR PARA OS JOGADORES OS EVENTOS RESULTANTES DA SIMULAÇÃO
		}
	}()

	println("Listening on port 5555")
	err := server.Listen("", 5555)
	if err != nil {
		panic(err)
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
