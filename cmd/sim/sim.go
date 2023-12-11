package main

import (
	"encoding/binary"
	"mmo2/pkg/events"
	"net"
	"time"
)

var readed int32
var sent int32

var writing bool = true
var reading bool = true

func read(conn net.Conn) {
	for reading {
		var evId uint16
		binary.Read(conn, binary.BigEndian, &evId)
		payload, err := events.ReadMove(conn)
		if evId != events.TypeMove {
			panic("wrong type!")
		}
		if err != nil {
			println(err.Error())
			break
		}
		payload.Payload.Dx += payload.Payload.Dy
		readed += 1
	}
}

func write(conn net.Conn) {
	for writing {
		err := events.WriteMove(conn, events.MovePayload{
			Dx: 111,
			Dy: 656,
		})
		if err != nil {
			println(err.Error())
			continue
		}
		sent += 1
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 0; i < 900; i++ {
		dialer := net.Dialer{}
		conn, err := dialer.Dial("tcp4", "192.168.0.9:5555")
		if err != nil {
			panic(err)
		}
		go read(conn)
		go write(conn)
	}
	time.Sleep(time.Second * 5)
	writing = false
	time.Sleep(time.Second * 3)
	reading = false
	println("total sent", sent)
	println("total received", readed)
}
