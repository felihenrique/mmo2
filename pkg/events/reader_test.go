package events

import (
	"bytes"
	"errors"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := Move{
		Dx: 111,
		Dy: 244,
	}
	ev1Bytes := ev1.ToBytes()
	ev2 := Move{
		Dx: 123,
		Dy: 656,
	}
	ev2Bytes := ev2.ToBytes()
	byteReader := bytes.NewReader(append(ev1Bytes, ev2Bytes...))
	eventReader := NewReader()
	eventReader.FillFrom(byteReader)
	readedBytes1, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetEventType(readedBytes1) != TypeMove {
		panic("wrong type")
	}
	readedEvent := Move{}
	readedEvent.FromBytes(readedBytes1)
	if readedEvent.Dx != 111 || readedEvent.Dy != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetEventType(readedBytes2) != TypeMove {
		panic("wrong type")
	}
	readedEvent2 := Move{}
	readedEvent2.FromBytes(readedBytes1)
	if readedEvent2.Dx != 123 || readedEvent2.Dy != 656 {
		panic("wrong data")
	}
	eventReader.Pop()
	_, err = eventReader.Next()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("wrong result")
	}
}
