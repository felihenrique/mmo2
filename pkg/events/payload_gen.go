package events

import (
	"encoding/binary"
	"math"
)

const (
	TypeAck = int16(iota + 1)
	TypeMove
	TypeEntityMoved
)

func (str *Move) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dx))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dy))
	return buffer
}

func (str *Move) fromBytes(data []byte) {
	position := 0
	str.Dx = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	str.Dy = int32(binary.BigEndian.Uint32(data[position:]))
}

func (str *Move) evType() int16 {
	return TypeMove
}

func (m *EntityMoved) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(m.NewX))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(m.NewY))
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(m.EntityId))
	buffer = binary.BigEndian.AppendUint32(buffer, math.Float32bits(m.Velocity))
	return buffer
}

func (m *EntityMoved) fromBytes(data []byte) {
	position := 0
	m.NewX = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	m.NewY = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	m.EntityId = int16(binary.BigEndian.Uint16(data[position:]))
	position += 2
	m.Velocity = math.Float32frombits(binary.BigEndian.Uint32(data[position:]))
}

func (str *EntityMoved) evType() int16 {
	return TypeEntityMoved
}
