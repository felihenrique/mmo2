package events

import (
	"errors"
	"io"
	"mmo2/pkg/serialization"
)

var ErrNotEnoughBytes = errors.New("not enough bytes in the buffer. please fill it")

type Reader struct {
	buffer []byte
	length int32
}

func getSize(data Raw) int16 {
	var size int16
	serialization.ReadInt16(data, &size)
	return size
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
	if r.length < 5 {
		return nil, ErrNotEnoughBytes
	}
	eventLength := getSize(r.buffer)
	if r.length < int32(eventLength) {
		return nil, ErrNotEnoughBytes
	}
	event := make([]byte, eventLength)
	copy(event, r.buffer[2:eventLength])
	return event, nil
}

func (r *Reader) Pop() {
	eventLength := getSize(r.buffer)
	nextLength := r.length - int32(eventLength)
	copy(r.buffer, r.buffer[eventLength:r.length])
	r.length = nextLength
}
