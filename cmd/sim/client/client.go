package main

import (
	"mmo2/pkg/events"
	"net"
	"time"
)

var readed int32
var sent int32

var writing bool = true
var reading bool = true

func read(conn net.Conn) {
	reader := events.NewReader()
	for reading {
		readed += 1
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
		if evId != events.TypeMoveRequest {
			panic("wrong type!")
		}
		event := events.MoveRequest{}
		events.Unserialize(eventBytes, &event)
		if event.Dx != 5 && event.Dy != 2 {
			panic("wrong data")
		}
		reader.Pop()
	}
}

func write(conn net.Conn) {
	for writing {
		event := events.MoveRequest{
			Dx: 5,
			Dy: 2,
		}
		data := events.Serialize(&event, 123)
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
		sent += 1
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 0; i < 500; i++ {
		dialer := net.Dialer{}
		conn, err := dialer.Dial("tcp4", "192.168.0.9:5555")
		if err != nil {
			panic(err)
		}
		go read(conn)
		go write(conn)
	}
	time.Sleep(time.Second * 10)
	writing = false
	time.Sleep(time.Second * 3)
	reading = false
	println("total sent", sent)
	println("total received", readed)
}
