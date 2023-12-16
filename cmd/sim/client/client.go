package main

import (
	"mmo2/pkg/events"
	"mmo2/pkg/payloads"
	"net"
	"sync/atomic"
	"time"
)

var readed atomic.Int32
var sent atomic.Int32

var writing atomic.Bool
var reading atomic.Bool

func read(conn net.Conn) {
	reader := events.NewReader()
	for reading.Load() {
		readed.Add(1)
		err := reader.FillFrom(conn)
		if err != nil {
			println(err.Error())
			continue
		}
		eventBytes, err := reader.Next()
		if err != nil {
			println(err.Error())
			continue
		}
		evId := events.GetType(eventBytes)
		if evId != payloads.TypeMoveRequest {
			panic("wrong type!")
		}
		event := payloads.MoveRequest{}
		events.Unserialize(eventBytes, &event)
		if event.Dx != 5 && event.Dy != 2 {
			panic("wrong data")
		}
		reader.Pop()
	}
}

func write(conn net.Conn) {
	for writing.Load() {
		event := payloads.MoveRequest{
			Dx: 5,
			Dy: 2,
		}
		data := events.Serialize(&event)
		offset := 0
		written := 0
		for retries := 0; retries < 3; retries++ {
			n, err := conn.Write(data[offset:])
			written += n
			if err != nil {
				panic(err)
			}
			if n < len(data) {
				offset = n
			} else {
				break
			}
			retries += 1
		}
		sent.Add(1)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	writing.Store(true)
	reading.Store(true)
	for i := 0; i < 900; i++ {
		dialer := net.Dialer{}
		conn, err := dialer.Dial("tcp4", "192.168.0.9:5555")
		if err != nil {
			panic(err)
		}
		go read(conn)
		go write(conn)
	}
	time.Sleep(time.Second * 10)
	writing.Store(false)
	time.Sleep(time.Second * 3)
	reading.Store(false)
	println("total sent", sent.Load())
	println("total received", readed.Load())
}
