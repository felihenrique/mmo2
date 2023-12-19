package packets

import (
	"fmt"
	"mmo2/pkg/serialization"
)

const (
	TypeNone = int16(iota)
	TypeAckRequest
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeJoinShardResponse
	TypeEntityRemoved
)

func (str *AckRequest) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeAckRequest)
	buffer = serialization.Append(buffer, eventId)

	return buffer
}

func (str *AckRequest) FromBytes(data []byte) int16 {
	var n int16 = 4

	return n
}

func (str *AckRequest) Type() int16 {
	return TypeAckRequest
}

func (str *AckRequest) String() string {
	return fmt.Sprintf("AckRequest: {  }")
}

func (str *MoveRequest) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeMoveRequest)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Dx)
	buffer = serialization.Append(buffer, str.Dy)

	return buffer
}

func (str *MoveRequest) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Dx)
	n += serialization.Read(data[n:], &str.Dy)

	return n
}

func (str *MoveRequest) Type() int16 {
	return TypeMoveRequest
}

func (str *MoveRequest) String() string {
	return fmt.Sprintf("MoveRequest: { Dx: %v, Dy: %v,  }", str.Dx, str.Dy)
}

func (str *RotateRequest) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeRotateRequest)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Quantity)

	return buffer
}

func (str *RotateRequest) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Quantity)

	return n
}

func (str *RotateRequest) Type() int16 {
	return TypeRotateRequest
}

func (str *RotateRequest) String() string {
	return fmt.Sprintf("RotateRequest: { Quantity: %v,  }", str.Quantity)
}

func (str *JoinShardRequest) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeJoinShardRequest)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Name)
	buffer = serialization.Append(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Name)
	n += serialization.Read(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) Type() int16 {
	return TypeJoinShardRequest
}

func (str *JoinShardRequest) String() string {
	return fmt.Sprintf("JoinShardRequest: { Name: %v, Portal: %v,  }", str.Name, str.Portal)
}

func (str *JoinShardResponse) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeJoinShardResponse)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.RequestId)
	buffer = serialization.Append(buffer, str.Entity)

	return buffer
}

func (str *JoinShardResponse) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.RequestId)
	n += serialization.Read(data[n:], &str.Entity)

	return n
}

func (str *JoinShardResponse) Type() int16 {
	return TypeJoinShardResponse
}

func (str *JoinShardResponse) String() string {
	return fmt.Sprintf("JoinShardResponse: { RequestId: %v, Entity: %v,  }", str.RequestId, str.Entity)
}

func (str *EntityRemoved) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeEntityRemoved)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.EntityId)

	return buffer
}

func (str *EntityRemoved) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	return n
}

func (str *EntityRemoved) Type() int16 {
	return TypeEntityRemoved
}

func (str *EntityRemoved) String() string {
	return fmt.Sprintf("EntityRemoved: { EntityId: %v,  }", str.EntityId)
}
