package events

import (
	"bytes"
	"errors"
	"mmo2/pkg/payloads"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := Serialize(&payloads.MoveRequest{
		Dx: 111,
		Dy: 244,
	})
	ev2 := Serialize(&payloads.MoveRequest{
		Dx: 123,
		Dy: 656,
	})
	byteReader := bytes.NewReader(append(ev1, ev2...))
	eventReader := NewReader()
	eventReader.FillFrom(byteReader)
	readedBytes1, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetID(readedBytes1) != payloads.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent := payloads.MoveRequest{}
	Unserialize(readedBytes1, &readedEvent)
	if readedEvent.Dx != 111 || readedEvent.Dy != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetID(readedBytes2) != payloads.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent2 := payloads.MoveRequest{}
	Unserialize(readedBytes2, &readedEvent2)
	if readedEvent2.Dx != 123 || readedEvent2.Dy != 656 {
		panic("wrong data")
	}
	eventReader.Pop()
	_, err = eventReader.Next()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("wrong result")
	}
}
