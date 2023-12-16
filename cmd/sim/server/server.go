package main

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"mmo2/pkg/payloads"
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

	server := gsp.NewTcpServer()
	go handleChans(server)

	println("Listening on port 5555")
	err = server.Listen("", 5555)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 60)
}

func handleChans(server gsp.TcpServer) {
	peers := make(map[string]*gsp.TcpPeer)
	peerConnected := server.PeerConnChan()
	peerDisconnected := server.PeerDisChan()
	newEvents := server.NewEventsChan()
	for {
		select {
		case peer := <-peerConnected:
			peers[peer.Addr()] = peer
		case peer := <-peerDisconnected:
			delete(peers, peer.Addr())
		case peerEvent := <-newEvents:
			handleEvent(peerEvent.Event, peers)
		}
	}
}

func handleEvent(rawEvent events.RawEvent, peers map[string]*gsp.TcpPeer) {
	if len(rawEvent) != 14 {
		panic("WRONG")
	}
	event := payloads.MoveRequest{}
	events.Unserialize(rawEvent, &event)
	if event.Dx != 5 || event.Dy != 2 {
		panic("DIVERGENT")
	}
	for _, peer := range peers {
		peer.SendEvent(rawEvent)
	}
}
