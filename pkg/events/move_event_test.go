package events

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func BenchmarkMoveEvent(b *testing.B) {
	payload := MoveEvent{
		Dx: 111,
		Dy: 123,
	}
	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		data, err := payload.ToBytes()
		if err != nil {
			panic(err)
		}
		err = binary.Write(&buffer, binary.BigEndian, int16(len(data)))
		if err != nil {
			panic(err)
		}
		_, err = buffer.Write(data)
		if err != nil {
			panic(err)
		}

	}
}
