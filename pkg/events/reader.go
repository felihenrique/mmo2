package events

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var ErrNotEnoughBytes = errors.New("not enough bytes in the buffer. please fill it")

type Reader struct {
	buffer []byte
	length int16
}

func GetEventSize(data []byte) int16 {
	return int16(uint16(data[1]) | uint16(data[0])<<8)
}

func GetEventType(data []byte) int16 {
	return int16(uint16(data[3]) | uint16(data[2])<<8)
}

func NewReader() *Reader {
	reader := Reader{}
	reader.buffer = make([]byte, 2048)
	reader.length = 0
	return &reader
}

func (r *Reader) FillFrom(reader io.Reader) error {
	n, err := reader.Read(r.buffer[r.length:])
	r.length += int16(n)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) NextEvent() ([]byte, error) {
	reader := bytes.NewReader(r.buffer)
	var packetSize int16
	err := binary.Read(reader, binary.BigEndian, &packetSize)
	if err != nil {
		return nil, err
	}
	if r.length < packetSize {
		return nil, ErrNotEnoughBytes
	}
	nextLength := r.length - packetSize - 2
	toReturn := make([]byte, packetSize+2)
	copy(toReturn, r.buffer[0:packetSize+2])
	if nextLength > 0 {
		copy(r.buffer, r.buffer[packetSize+2:r.length])
	}
	r.length = nextLength
	return toReturn, nil
}
