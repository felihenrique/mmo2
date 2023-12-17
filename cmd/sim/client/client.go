package main

import (
	"log"
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
	"os"
	"runtime/pprof"
	"sync/atomic"
	"time"
)

var readed atomic.Int32
var sent atomic.Int32

var writing atomic.Bool
var reading atomic.Bool

func manageClient(client *gsp.TcpClient) {
	eventsChan := client.EventsChan()
	disconChan := client.DisconnectedChan()
	ticker := time.NewTicker(time.Millisecond * 100)
main:
	for {
		select {
		case eventBytes := <-eventsChan:
			evId := events.GetType(eventBytes)
			if evId != packets.TypeMoveInput {
				panic("wrong type!")
			}
			event := packets.MoveInput{}
			event.FromBytes(eventBytes)
			if event.Dx != 5 && event.Dy != 2 {
				panic("wrong data")
			}
			readed.Add(1)
		case <-ticker.C:
			event := packets.MoveInput{
				Dx: 5,
				Dy: 2,
			}
			client.SendEvent(&event)
			sent.Add(1)
		case <-disconChan:
			break main
		}
	}
	ticker.Stop()
}

func main() {
	f, err := os.Create("cpu_client.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	writing.Store(true)
	reading.Store(true)
	for i := 0; i < 100; i++ {
		client := gsp.NewTcpClient()
		err := client.Connect("", 5555)
		if err != nil {
			panic(err)
		}
		go manageClient(client)
	}
	time.Sleep(time.Second * 10)
	writing.Store(false)
	time.Sleep(time.Second * 5)
	reading.Store(false)
	println("total sent", sent.Load())
	println("total received", readed.Load())
}
