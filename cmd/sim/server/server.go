package main

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"os"
	"runtime/pprof"
	"sync"
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

	server := gsp.NewTcpServer()
	peers := sync.Map{}

	server.OnPeerConnect(func(peer *gsp.TcpPeer) {
		peers.Store(peer.Addr(), peer)
	})

	server.OnPeerDisconnect(func(peer *gsp.TcpPeer) {
		peers.Delete(peer.Addr())
	})

	server.OnEvent(events.TypeMove, func(peer *gsp.TcpPeer, eventBytes []byte) {
		if len(eventBytes) > 12 {
			panic("dfdsfdsfdsfdfsdfs")
		}
		event := events.Move{}
		event.FromBytes(eventBytes)
		if event.Dx != 111 || event.Dy != 656 {
			panic("DIVERGENT")
		}
		peers.Range(func(key, value any) bool {
			peer := value.(*gsp.TcpPeer)
			peer.SendEvent(eventBytes)
			return true
		})
	})

	println("Listening on port 5555")
	err = server.Listen("", 5555)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 20)
}
