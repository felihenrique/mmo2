package events

import (
	"bytes"
	"mmo2/game/packets"
	"testing"
)

func TestWriter(t *testing.T) {
	writer := NewWriter()
	writer.Append((&packets.MoveRequest{
		Dx: 11,
		Dy: 12,
	}).ToBytes(0))
	writer.Append((&packets.MoveRequest{
		Dx: 44,
		Dy: 78,
	}).ToBytes(0))
	writer.Append((&packets.MoveRequest{
		Dx: 78,
		Dy: 90,
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
