package main

import (
	"bytes"
	"encoding/binary"
	"mmo2/pkg/events"
)

func main() {
	buffer := bytes.Buffer{}

	ev1 := events.MovePayload{
		Dx: 111,
		Dy: 123,
	}
	err := events.WriteMove(&buffer, ev1)
	if err != nil {
		panic(err)
	}
	var evType uint16
	binary.Read(&buffer, binary.BigEndian, &evType)
	ev2, err := events.ReadMove(&buffer)
	if err != nil {
		panic(err)
	}
	println(ev2.Type)
}
