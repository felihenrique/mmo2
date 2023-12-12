package events

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestEventReader(t *testing.T) {
	eventReader := NewReader()
	event := MoveEvent{
		Dx: 111,
		Dy: 656,
	}
	eventBytes, err := event.ToBytes()
	if err != nil {
		panic(err)
	}
	byteReader := bytes.NewReader(eventBytes)
	eventReader.FillFrom(byteReader)
	readedEventBytes, err := eventReader.NextEvent()
	if err != nil {
		panic(err)
	}
	readedEvent := MoveEvent{}
	err = readedEvent.FromBytes(readedEventBytes)
	if err != nil {
		panic(err)
	}
	if readedEvent.Dx != 111 || readedEvent.Dy != 543 {
		panic(fmt.Sprintf("different values: %d, %d", readedEvent.Dx, readedEvent.Dy))
	}
	_, err = eventReader.NextEvent()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("err should occur")
	}
}

func TestEventFuncs(t *testing.T) {
	event := MoveEvent{
		Dx: 111,
		Dy: 543,
	}
	eventBytes, err := event.ToBytes()
	if err != nil {
		panic(err)
	}
	size := GetEventSize(eventBytes)
	if size != 10 {
		panic("mismatched")
	}
	evType := GetEventType(eventBytes)
	if evType != TypeMove {
		panic("mismatch")
	}
}
