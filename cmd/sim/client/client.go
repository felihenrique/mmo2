package main

import (
	"mmo2/pkg/events"
	"mmo2/pkg/gsp"
	"mmo2/pkg/payloads"
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
			evId := events.GetID(eventBytes)
			if evId != payloads.TypeMoveRequest {
				panic("wrong type!")
			}
			event := payloads.MoveRequest{}
			events.Unserialize(eventBytes, &event)
			if event.Dx != 5 && event.Dy != 2 {
				panic("wrong data")
			}
			readed.Add(1)
		case <-ticker.C:
			event := payloads.MoveRequest{
				Dx: 5,
				Dy: 2,
			}
			data := events.Serialize(&event)
			client.SendEvent(data)
			sent.Add(1)
		case <-disconChan:
			break main
		}
	}
	ticker.Stop()
}

func main() {
	writing.Store(true)
	reading.Store(true)
	for i := 0; i < 1000; i++ {
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
