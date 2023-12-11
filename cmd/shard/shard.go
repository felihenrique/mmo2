package main

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("prof/cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	f2, err := os.Create("prof/mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f2.Close() // error handling omitted for example
	runtime.GC()     // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	server := gsp.NewTcpServer(1000)

	server.OnPeerConnect(func(peer gsp.TcpPeer) {

	})

	server.OnPeerDisconnect(func(peer gsp.TcpPeer) {

	})

	received := 0
	server.OnEvent(events.TypeMove, func(peer gsp.TcpPeer) {
		payload, err := events.ReadMove(peer.Reader())
		if err != nil {
			println(err.Error())
			return
		}
		err = events.WriteMove(peer.Writer(), payload.Payload)
		if err != nil {
			println(err.Error())
			return
		}
		received += 1
	})

	println("listening")
	err = server.Listen("", 5555)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 60)
	println(received)
}
