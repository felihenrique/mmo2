package events

import (
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
	err := binary.Write(writer, binary.BigEndian, TypeMove)
	if err != nil {
		return err
	}
	err = binary.Write(writer, binary.BigEndian, payload.Dx)
	if err != nil {
		return err
	}
	err = binary.Write(writer, binary.BigEndian, payload.Dy)
	if err != nil {
		return err
	}
	return nil
}
func ReadMove(reader io.Reader) (*moveEventPacket, error) {
	payload := MovePayload{}
	err := binary.Read(reader, binary.BigEndian, &payload.Dx)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.BigEndian, &payload.Dy)
	if err != nil {
		return nil, err
	}
	packet := moveEventPacket{}
	packet.Type = TypeMove
	packet.Payload = payload
	return &packet, nil
}
