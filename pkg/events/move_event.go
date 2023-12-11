package events

import (
	"bytes"
	"encoding/binary"
	"io"
)

type moveEventPacket struct {
	Type    uint16
	Payload MovePayload
}

func (m *moveEventPacket) GetType() uint16 {
	return m.Type
}

type MovePayload struct {
	Dx int32
	Dy int32
}

func WriteMove(writer io.Writer, payload MovePayload) error {
	buffer := bytes.Buffer{}
	err := binary.Write(&buffer, binary.BigEndian, TypeMove)
	if err != nil {
		return err
	}
	err = binary.Write(&buffer, binary.BigEndian, payload.Dx)
	if err != nil {
		return err
	}
	err = binary.Write(&buffer, binary.BigEndian, payload.Dy)
	if err != nil {
		return err
	}
	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
func ReadMove(reader io.Reader) (*moveEventPacket, error) {
	buffer := make([]byte, 8)
	_, err := reader.Read(buffer)
	if err != nil {
		return nil, err
	}
	bufferReader := bytes.NewReader(buffer)
	payload := MovePayload{}
	err = binary.Read(bufferReader, binary.BigEndian, &payload.Dx)
	if err != nil {
		return nil, err
	}
	err = binary.Read(bufferReader, binary.BigEndian, &payload.Dy)
	if err != nil {
		return nil, err
	}
	packet := moveEventPacket{}
	packet.Type = TypeMove
	packet.Payload = payload
	return &packet, nil
}
