package payloads

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	TypeAck
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeEntityCreated
	TypeEntityUpdated
)

func (str *Ack) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.EventId)

	return buffer
}

func (str *Ack) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.EventId)

	return n
}

func (str *Ack) EvType() int16 {
	return TypeAck
}

func (str *MoveRequest) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Dx)
	buffer = serialization.WriteBinary(buffer, str.Dy)

	return buffer
}

func (str *MoveRequest) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Dx)
	n += serialization.ReadBinary(data[n:], &str.Dy)

	return n
}

func (str *MoveRequest) EvType() int16 {
	return TypeMoveRequest
}

func (str *RotateRequest) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Quantity)

	return buffer
}

func (str *RotateRequest) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Quantity)

	return n
}

func (str *RotateRequest) EvType() int16 {
	return TypeRotateRequest
}

func (str *JoinShardRequest) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) EvType() int16 {
	return TypeJoinShardRequest
}

func (str *EntityCreated) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Data)

	return buffer
}

func (str *EntityCreated) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Data)

	return n
}

func (str *EntityCreated) EvType() int16 {
	return TypeEntityCreated
}

func (str *EntityUpdated) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Data)

	return buffer
}

func (str *EntityUpdated) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Data)

	return n
}

func (str *EntityUpdated) EvType() int16 {
	return TypeEntityUpdated
}
