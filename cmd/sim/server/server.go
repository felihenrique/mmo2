package main

import (
	"fmt"
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
	"os"
	"runtime/pprof"
	"sync/atomic"
	"time"
)

var sent atomic.Int32
var readed atomic.Int32

var peers = make(map[string]gsp.IPeer)

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

	time.Sleep(time.Second * 30)
	fmt.Printf("Sent: %d, Readed: %d", sent.Load(), readed.Load())
}

func handleChans(server gsp.TcpServer) {
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

func handleEvent(rawEvent events.Raw, peers map[string]gsp.IPeer) {
	readed.Add(1)
	if len(rawEvent) != 14 {
		panic("WRONG")
	}
	event := packets.MoveRequest{}
	event.FromBytes(rawEvent)
	if event.Dx != 5 || event.Dy != 2 {
		panic("DIVERGENT")
	}
	for _, peer := range peers {
		sent.Add(1)
		peer.SendBytes(rawEvent)
	}
}
