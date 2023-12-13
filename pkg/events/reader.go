package events

import (
	"errors"
	"io"
)

var ErrNotEnoughBytes = errors.New("not enough bytes in the buffer. please fill it")

type Reader struct {
	buffer []byte
	length int32
}

func NewReader() *Reader {
	reader := Reader{}
	reader.buffer = make([]byte, 2048)
	reader.length = 0
	return &reader
}

func (r *Reader) FillFrom(reader io.Reader) error {
	n, err := reader.Read(r.buffer[r.length:])
	r.length += int32(n)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) Next() ([]byte, error) {
	if r.length < 7 {
		return nil, ErrNotEnoughBytes
	}
	eventLength := GetSize(r.buffer)
	if r.length < int32(eventLength) {
		return nil, ErrNotEnoughBytes
	}
	return r.buffer[0:eventLength], nil
}

func (r *Reader) Pop() {
	eventLength := GetSize(r.buffer)
	nextLength := r.length - int32(eventLength)
	copy(r.buffer, r.buffer[eventLength:r.length])
	r.length = nextLength
}
