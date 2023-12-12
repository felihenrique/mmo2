package main

import (
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"sync"
)

func main() {
	server := gsp.NewTcpServer(1000)
	peers := sync.Map{}

	server.OnPeerConnect(func(peer *gsp.TcpPeer) {
		peers.Store(peer.Id(), peer)
	})

	server.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		peers.Delete(peer.Id())
	})

	server.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer, eventBytes []byte) {
		if len(eventBytes) > 12 {
			panic("dfdsfdsfdsfdfsdfs")
		}
		event := events.MoveEvent{}
		err := event.FromBytes(eventBytes)
		if event.Dx != 111 || event.Dy != 656 {
			panic("DIVERGENT")
		}
		if err != nil {
			println(err.Error())
			return
		}
		peers.Range(func(key, value any) bool {
			peer := value.(*gsp.TcpPeer)
			peer.Writer().Write(eventBytes)
			return true
		})
		peer.Writer().Write(eventBytes)
	})

	println("Listening on port 5555")
	err := server.Listen("", 5555)
	if err != nil {
		panic(err)
	}
}
