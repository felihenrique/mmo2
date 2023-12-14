package main

import (
	"mmo2/pkg/events"
	"net"
	"time"
)

var writing bool = true
var reading bool = true

type Player struct {
	X        int32
	Y        int32
	Velocity float32
}

func read(conn net.Conn) {
	reader := events.NewReader()
	players := make([]Player, 0)
	for reading {
		err := reader.FillFrom(conn)
		if err != nil {
			println(err.Error())
			break
		}
		eventBytes, err := reader.Next()
		if err != nil {
			println(err.Error())
			break
		}
		evId := events.GetType(eventBytes)

		reader.Pop()
	}
}

func write(conn net.Conn) {
	writer := events.NewWriter()
	for writing {

		time.Sleep(time.Second)
	}
}

func main() {
	for i := 0; i < 1; i++ {
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
	time.Sleep(time.Millisecond * 100)
	reading = false
	time.Sleep(time.Millisecond * 100)
}
