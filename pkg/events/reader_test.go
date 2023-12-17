package events

import (
	"bytes"
	"errors"
	"mmo2/pkg/packets"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := Serialize(&packets.MoveInput{
		Dx: 111,
		Dy: 244,
	})
	ev2 := Serialize(&packets.MoveInput{
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
	if GetType(readedBytes1) != packets.TypeMoveInput {
		panic("wrong type")
	}
	readedEvent := packets.MoveInput{}
	Unserialize(readedBytes1, &readedEvent)
	if readedEvent.Dx != 111 || readedEvent.Dy != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetType(readedBytes2) != packets.TypeMoveInput {
		panic("wrong type")
	}
	readedEvent2 := packets.MoveInput{}
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
