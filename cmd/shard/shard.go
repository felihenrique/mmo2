package main

import (
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
)

func main() {
	server := gsp.NewTcpServer(1000)

	server.OnPeerConnect(func(peer gsp.TcpPeer) {

	})

	server.OnPeerDisconnect(func(peer gsp.TcpPeer) {

	})

	server.OnEvent(events.TypeMove, func(peer gsp.TcpPeer) {
		_, err := events.ReadMove(peer.Reader())
		if err != nil {
			println("read move error: ", err.Error())
			return
		}
	})

	println("listening")
	err := server.Listen("127.0.0.1", 5555)
	if err != nil {
		panic(err)
	}
}
