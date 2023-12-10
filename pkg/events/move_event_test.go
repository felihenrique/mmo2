package events

import (
	"bytes"
	"testing"
)

func BenchmarkMoveEvent(b *testing.B) {
	buffer := bytes.Buffer{}
	payload := MovePayload{
		Dx: 111,
		Dy: 123,
	}
	for i := 0; i < b.N; i++ {
		err := WriteMove(&buffer, payload)
		if err != nil {
			panic(err)
		}
		_, err = ReadMove(&buffer)
		if err != nil {
			panic(err)
		}
	}
}
