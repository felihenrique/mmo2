package events

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	writer := NewWriter()
	writer.Append([]byte("ola mundo"))
	writer.Append([]byte("aaaaa"))
	writer.Append([]byte("uity"))
	buffer := bytes.Buffer{}
	n, err := writer.Send(&buffer)
	if n < 18 || err != nil {
		panic(err)
	}
	if buffer.String() != "ola mundoaaaaauity" {
		panic("wrong result")
	}
	writer.Append([]byte("abcde"))
	buffer.Reset()
	n, err = writer.Send(&buffer)
	if n < 5 || err != nil {
		panic(err)
	}
	if buffer.String() != "abcde" {
		panic("wrong result")
	}
	n, err = writer.Send(&buffer)
	if n != 0 || err != nil {
		panic(err)
	}
}
