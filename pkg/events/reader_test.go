package events

import (
	"bytes"
	"errors"
	"mmo2/pkg/packets"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := packets.MoveInput{
		Dx:      111,
		Dy:      244,
		InputId: 123,
	}
	ev2 := packets.MoveInput{
		Dx:      123,
		Dy:      656,
		InputId: 777,
	}
	writer := NewWriter()
	writer.Append(&ev1)
	writer.Append(&ev2)
	buffer := bytes.Buffer{}
	writer.Send(&buffer)
	if buffer.Len() != 28 {
		panic("wrong")
	}
	eventReader := NewReader()
	eventReader.FillFrom(&buffer)
	readedBytes1, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetType(readedBytes1) != packets.TypeMoveInput {
		panic("wrong type")
	}
	readedEvent := packets.MoveInput{}
	readedEvent.FromBytes(readedBytes1)
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
	readedEvent2.FromBytes(readedBytes2)
	if readedEvent2.Dx != 123 || readedEvent2.Dy != 656 {
		panic("wrong data")
	}
	eventReader.Pop()
	_, err = eventReader.Next()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("wrong result")
	}
}
