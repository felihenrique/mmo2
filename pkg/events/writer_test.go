package events

import (
	"bytes"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"testing"
)

func TestWriter(t *testing.T) {
	writer := NewWriter()
	writer.Append((&packets.MoveRequest{
		Move: ecs.NewMove(11, 12, 0, 0),
	}).ToBytes(0))
	writer.Append((&packets.MoveRequest{
		Move: ecs.NewMove(44, 78, 0, 0),
	}).ToBytes(0))
	writer.Append((&packets.MoveRequest{
		Move: ecs.NewMove(78, 90, 0, 0),
	}).ToBytes(0))
	buffer := bytes.Buffer{}
	n, err := writer.Send(&buffer)
	if n < 42 || err != nil {
		panic(err)
	}
	buffer.Reset()
	n, err = writer.Send(&buffer)
	if n > 0 || err != nil {
		panic(err)
	}
}
