package events

import (
	"bytes"
	"errors"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"testing"
)

func TestReader(t *testing.T) {
	ev1 := packets.NewMoveRequest(
		ecs.NewMove(111, 244, 0, 0),
	)
	ev2 := packets.NewMoveRequest(
		ecs.NewMove(123, 656, 0, 0),
	)
	writer := NewWriter()
	writer.Append(ev1.ToBytes(0))
	writer.Append(ev2.ToBytes(0))
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
	if GetType(readedBytes1) != packets.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent := packets.MoveRequest{}
	readedEvent.FromBytes(readedBytes1)
	if readedEvent.Move.QuantityX != 111 || readedEvent.Move.QuantityY != 244 {
		panic("wrong data")
	}
	eventReader.Pop()
	readedBytes2, err := eventReader.Next()
	if err != nil {
		panic(err)
	}
	if GetType(readedBytes2) != packets.TypeMoveRequest {
		panic("wrong type")
	}
	readedEvent2 := packets.MoveRequest{}
	readedEvent2.FromBytes(readedBytes2)
	if readedEvent2.Move.QuantityX != 123 || readedEvent2.Move.QuantityY != 656 {
		panic("wrong data")
	}
	eventReader.Pop()
	_, err = eventReader.Next()
	if !errors.Is(err, ErrNotEnoughBytes) {
		panic("wrong result")
	}
}
