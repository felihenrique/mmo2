package events

import (
	"bytes"
	"errors"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := Serialize(&Move{
		Dx: 111,
		Dy: 244,
	}, 123)
	ev2 := Serialize(&Move{
		Dx: 123,
		Dy: 656,
	}, 123)
	byteReader := bytes.NewReader(append(ev1, ev2...))
	eventReader := NewReader()
	eventReader.FillFrom(byteReader)
	readedBytes1, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetType(readedBytes1) != TypeMove {
		panic("wrong type")
	}
	readedEvent := Move{}
	Unserialize(readedBytes1, &readedEvent)
	if readedEvent.Dx != 111 || readedEvent.Dy != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetType(readedBytes2) != TypeMove {
		panic("wrong type")
	}
	readedEvent2 := Move{}
	Unserialize(readedBytes1, &readedEvent2)
	if readedEvent2.Dx != 123 || readedEvent2.Dy != 656 {
		panic("wrong data")
	}
	eventReader.Pop()
	_, err = eventReader.Next()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("wrong result")
	}
}
