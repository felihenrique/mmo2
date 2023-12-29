package events

import (
	"bytes"
	"errors"
	"mmo2/game/packets"
	"mmo2/pkg/event_utils"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := packets.NewMoveRequest(
		111, 244,
	)
	ev2 := packets.NewMoveRequest(
		123, 656,
	)
	writer := NewWriter()
	writer.Append(ev1.ToBytes(0))
	writer.Append(ev2.ToBytes(0))
	buffer := bytes.Buffer{}
	writer.Send(&buffer)
	if buffer.Len() != 52 {
		panic("wrong")
	}
	eventReader := NewReader()
	eventReader.FillFrom(&buffer)
	readedBytes1, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if event_utils.GetType(readedBytes1) != packets.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent := packets.MoveRequest{}
	readedEvent.FromBytes(readedBytes1)
	if readedEvent.Dx != 111 || readedEvent.Dy != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if event_utils.GetType(readedBytes2) != packets.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent2 := packets.MoveRequest{}
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
