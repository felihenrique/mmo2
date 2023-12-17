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
	TypeEntityRemoved
)

func (str *AckInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.EventId)

	return buffer
}

func (str *AckInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.EventId)

	return n
}

func (str *AckInput) Type() int16 {
	return TypeAck
}

func (str *MoveInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Dx)
	buffer = serialization.Write(buffer, str.Dy)

	return buffer
}

func (str *MoveInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Dx)
	n += serialization.Read(data[n:], &str.Dy)

	return n
}

func (str *MoveInput) Type() int16 {
	return TypeMoveRequest
}

func (str *RotateInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Quantity)

	return buffer
}

func (str *RotateInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Quantity)

	return n
}

func (str *RotateInput) Type() int16 {
	return TypeRotateRequest
}

func (str *JoinShardRequest) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) Type() int16 {
	return TypeJoinShardRequest
}

func (str *EntityCreated) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Entity)

	return buffer
}

func (str *EntityCreated) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Entity)

	return n
}

func (str *EntityCreated) Type() int16 {
	return TypeEntityCreated
}

func (str *EntityUpdated) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Entity)

	return buffer
}

func (str *EntityUpdated) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Entity)

	return n
}

func (str *EntityUpdated) Type() int16 {
	return TypeEntityUpdated
}

func (str *EntityRemoved) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.EntityId)

	return buffer
}

func (str *EntityRemoved) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.EntityId)

	return n
}

func (str *EntityRemoved) Type() int16 {
	return TypeEntityRemoved
}
