package events

import (
	"bytes"
	"encoding/binary"
)

type MoveEvent struct {
	Dx int32
	Dy int32
}

func (str *MoveEvent) ToBytes() ([]byte, error) {
	bufferData := bytes.Buffer{}
	err := binary.Write(&bufferData, binary.BigEndian, TypeMove)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&bufferData, binary.BigEndian, str.Dx)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&bufferData, binary.BigEndian, str.Dy)
	if err != nil {
		return nil, err
	}
	finalBuffer := bytes.Buffer{}
	err = binary.Write(&finalBuffer, binary.BigEndian, int16(bufferData.Len()))
	if err != nil {
		return nil, err
	}
	n, err := finalBuffer.Write(bufferData.Bytes())
	if n < len(bufferData.Bytes()) {
		panic(err)
	}
	if err != nil {
		return nil, err
	}
	return finalBuffer.Bytes(), nil
}

func (str *MoveEvent) FromBytes(data []byte) error {
	reader := bytes.NewReader(data[4:])
	err := binary.Read(reader, binary.BigEndian, &str.Dx)
	if err != nil {
		return nil
	}
	err = binary.Read(reader, binary.BigEndian, &str.Dy)
	if err != nil {
		return nil
	}
	return nil
}
