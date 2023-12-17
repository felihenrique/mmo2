package events

import (
	"bytes"
	"mmo2/pkg/packets"
	"testing"
)

func TestWriter(t *testing.T) {
	writer := NewWriter()
	writer.Append(&packets.MoveInput{
		InputId: 123,
		Dx:      11,
		Dy:      12,
	})
	writer.Append(&packets.MoveInput{
		InputId: 565,
		Dx:      44,
		Dy:      78,
	})
	writer.Append(&packets.MoveInput{
		InputId: 999,
		Dx:      78,
		Dy:      90,
	})
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
