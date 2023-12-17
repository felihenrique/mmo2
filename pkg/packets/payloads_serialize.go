package packets

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	TypeAckInput
	TypeMoveInput
	TypeRotateInput
	TypeJoinShardRequest
	TypeJoinShardResponse
	TypeEntityCreated
	TypeEntityUpdated
	TypeEntityRemoved
)

func Read(data []byte) (serialization.ISerializable, int16) {
	var strType int16
	n := serialization.Read(data, &strType)
	switch strType {

	case TypeAckInput:
		var str AckInput
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeMoveInput:
		var str MoveInput
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeRotateInput:
		var str RotateInput
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeJoinShardRequest:
		var str JoinShardRequest
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeJoinShardResponse:
		var str JoinShardResponse
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeEntityCreated:
		var str EntityCreated
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeEntityUpdated:
		var str EntityUpdated
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeEntityRemoved:
		var str EntityRemoved
		n += str.FromBytes(data[n:])
		return &str, n

	default:
		panic("wrong type")
	}
}

func (str *AckInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.InputId)

	return buffer
}

func (str *AckInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.InputId)

	return n
}

func (str *AckInput) Type() int16 {
	return TypeAckInput
}

func (str *MoveInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Dx)
	buffer = serialization.Append(buffer, str.Dy)

	return buffer
}

func (str *MoveInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Dx)
	n += serialization.Read(data[n:], &str.Dy)

	return n
}

func (str *MoveInput) Type() int16 {
	return TypeMoveInput
}

func (str *RotateInput) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Quantity)

	return buffer
}

func (str *RotateInput) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Quantity)

	return n
}

func (str *RotateInput) Type() int16 {
	return TypeRotateInput
}

func (str *JoinShardRequest) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Name)
	buffer = serialization.Append(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Name)
	n += serialization.Read(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) Type() int16 {
	return TypeJoinShardRequest
}

func (str *JoinShardResponse) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Entity)

	return buffer
}

func (str *JoinShardResponse) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Entity)

	return n
}

func (str *JoinShardResponse) Type() int16 {
	return TypeJoinShardResponse
}

func (str *EntityCreated) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Entity)

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
	buffer = serialization.Append(buffer, str.Entity)

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
	buffer = serialization.Append(buffer, str.EntityId)

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
