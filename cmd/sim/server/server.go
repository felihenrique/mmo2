package main

import (
	"fmt"
	"log"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
	"os"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"time"
)

var sent atomic.Int32
var readed atomic.Int32

func main() {
	f, err := os.Create("cpu_server.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	var peers = make(map[string]gsp.IPeer)
	var peersLock sync.Mutex

	server := gsp.NewTcpServer()
	go handleChans(server, &peers, &peersLock)

	println("Listening on port 5555")
	err = server.Listen("", 5555)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 20)
	fmt.Printf("Sent: %d, Readed: %d", sent.Load(), readed.Load())
}

func handleChans(server *gsp.TcpServer, peers *map[string]gsp.IPeer, peersLock *sync.Mutex) {
	peerConnected := server.NewConnectionChan()
	peerDisconnected := server.DisconnectedChan()
	go readEvents(server, peers, peersLock)
	go readEvents(server, peers, peersLock)
	for {
		select {
		case peer := <-peerConnected:
			peersLock.Lock()
			(*peers)[peer.Addr()] = peer
			peersLock.Unlock()
		case peer := <-peerDisconnected:
			delete(*peers, peer.Addr())
		}
	}
}

func readEvents(server *gsp.TcpServer, peers *map[string]gsp.IPeer, peersLock *sync.Mutex) {
	newEvents := server.EventsChan()
	for peerEvent := range newEvents {
		readed.Add(1)
		rawEvent := peerEvent.Event
		if len(rawEvent) != 14 {
			panic("WRONG")
		}
		event := packets.MoveRequest{}
		event.FromBytes(rawEvent)
		if event.Dx != 5 || event.Dy != 2 {
			panic("DIVERGENT")
		}
		peersLock.Lock()
		for _, peer := range *peers {
			sent.Add(1)
			peer.SendBytes(rawEvent)
		}
		peersLock.Unlock()
	}
}
