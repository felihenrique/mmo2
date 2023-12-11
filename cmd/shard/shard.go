package main

import (
	"fmt"
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	server := gsp.NewTcpServer(1000)
	peers := make(map[int32]*gsp.TcpPeer)

	server.OnPeerConnect(func(peer *gsp.TcpPeer) {
		peers[peer.Id()] = peer
	})

	server.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		delete(peers, peer.Id())
	})

	received := 0
	sent := 0
	server.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer) {
		payload, err := events.ReadMove(peer.Reader())
		if err != nil {
			println(err.Error())
			return
		}
		for _, p := range peers {
			sent += 1
			err = events.WriteMove(p.Writer(), payload.Payload)
			if err != nil {
				println(err.Error())
				return
			}
		}
		received += 1
	})

	println("listening")
	err = server.Listen("", 5555)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 10)
	fmt.Printf("sent: %d, received: %d", sent, received)

}
