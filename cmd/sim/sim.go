package main

import (
	"mmo2/pkg/events"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	dialer := net.Dialer{}
	conn, err := dialer.Dial("tcp4", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}
	var counter atomic.Int32
	timeNow := time.Now()
	for time.Since(timeNow) < time.Second*10 {
		events.WriteMove(conn, events.MovePayload{
			Dx: 111,
			Dy: 656,
		})
		counter.Add(1)
	}
	println("total ", counter.Load())
}
