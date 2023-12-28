package main

import (
	"log"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/gsp"
	"os"
	"runtime/pprof"
	"sync/atomic"
	"time"
)

var readed atomic.Int32
var sent atomic.Int32

var writing atomic.Bool

func reader(client *gsp.TcpClient) {
	eventsChan := client.EventsChan()
	for eventBytes := range eventsChan {
		evId := event_utils.GetType(eventBytes)
		if evId != packets.TypeMoveRequest {
			panic("wrong type!")
		}
		event := packets.MoveRequest{}
		event.FromBytes(eventBytes)
		if event.Move.QuantityX != 5 && event.Move.QuantityY != 2 {
			panic("wrong data")
		}
		readed.Add(1)
	}
}

func writer(client *gsp.TcpClient) {
	for writing.Load() {
		event := packets.MoveRequest{
			Move: ecs.NewMove(5, 2, 0, 0),
		}
		client.SendRequest(&event)
		sent.Add(1)
		time.Sleep(time.Millisecond * 100)
	}
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
	var client *gsp.TcpClient
	for i := 0; i < 1000; i++ {
		client = gsp.NewTcpClient()
		err := client.Connect("", 5555)
		if err != nil {
			panic(err)
		}
		go writer(client)
		go reader(client)
	}
	time.Sleep(time.Second * 10)
	writing.Store(false)
	time.Sleep(time.Second * 5)
	println("total sent", sent.Load())
	println("total received", readed.Load())
}
