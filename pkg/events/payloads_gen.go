package events

import (
	"encoding/binary"
)

const (
	TypeAck = int16(iota + 1)
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeEntityCreated
	TypeEntityUpdated
)

func (str *MoveRequest) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dx))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dy))
	return buffer
}

func (str *MoveRequest) fromBytes(data []byte) {
	position := 0
	str.Dx = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	str.Dy = int32(binary.BigEndian.Uint32(data[position:]))
}

func (str *MoveRequest) evType() int16 {
	return TypeMoveRequest
}
