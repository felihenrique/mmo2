package packets

import (
	"fmt"
	"mmo2/pkg/game"
	"mmo2/pkg/serialization"
)

const (
	TypeNone = int16(iota)
	TypeAckRequest
	TypeRequestError
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeJoinShardResponse
	TypePlayerJoined
	TypeEntityMoved
	TypeEntityRemoved
)

func NewAckRequest() *AckRequest {
	return &AckRequest{}
}

func ParseAckRequest(event []byte) *AckRequest {
	str := AckRequest{}
	str.FromBytes(event)
	return &str
}

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

func NewRequestError(Message string) *RequestError {
	return &RequestError{
		Message: Message,
	}
}

func ParseRequestError(event []byte) *RequestError {
	str := RequestError{}
	str.FromBytes(event)
	return &str
}

func (str *RequestError) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeRequestError)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Message)

	return buffer
}

func (str *RequestError) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Message)

	return n
}

func (str *RequestError) Type() int16 {
	return TypeRequestError
}

func (str *RequestError) String() string {
	return fmt.Sprintf("RequestError: { Message: %v,  }", str.Message)
}

func NewMoveRequest(Dx int32, Dy int32) *MoveRequest {
	return &MoveRequest{
		Dx: Dx,
		Dy: Dy,
	}
}

func ParseMoveRequest(event []byte) *MoveRequest {
	str := MoveRequest{}
	str.FromBytes(event)
	return &str
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

func NewRotateRequest(Quantity float32) *RotateRequest {
	return &RotateRequest{
		Quantity: Quantity,
	}
}

func ParseRotateRequest(event []byte) *RotateRequest {
	str := RotateRequest{}
	str.FromBytes(event)
	return &str
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

func NewJoinShardRequest(Name string, Portal int8) *JoinShardRequest {
	return &JoinShardRequest{
		Name:   Name,
		Portal: Portal,
	}
}

func ParseJoinShardRequest(event []byte) *JoinShardRequest {
	str := JoinShardRequest{}
	str.FromBytes(event)
	return &str
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

func NewJoinShardResponse(EntityId int16, Position *game.Position, Movable *game.Movable, Name *game.Name) *JoinShardResponse {
	return &JoinShardResponse{
		EntityId: EntityId,
		Position: Position,
		Movable:  Movable,
		Name:     Name,
	}
}

func ParseJoinShardResponse(event []byte) *JoinShardResponse {
	str := JoinShardResponse{}
	str.FromBytes(event)
	return &str
}

func (str *JoinShardResponse) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeJoinShardResponse)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.EntityId)
	buffer = serialization.Append(buffer, str.Position)
	buffer = serialization.Append(buffer, str.Movable)
	buffer = serialization.Append(buffer, str.Name)

	return buffer
}

func (str *JoinShardResponse) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.Position = &game.Position{}
	n += serialization.Read(data[n:], str.Position)
	str.Movable = &game.Movable{}
	n += serialization.Read(data[n:], str.Movable)
	str.Name = &game.Name{}
	n += serialization.Read(data[n:], str.Name)

	return n
}

func (str *JoinShardResponse) Type() int16 {
	return TypeJoinShardResponse
}

func (str *JoinShardResponse) String() string {
	return fmt.Sprintf("JoinShardResponse: { EntityId: %v, Position: %v, Movable: %v, Name: %v,  }", str.EntityId, str.Position, str.Movable, str.Name)
}

func NewPlayerJoined(EntityId int16, Position *game.Position, Name *game.Name, Movable *game.Movable) *PlayerJoined {
	return &PlayerJoined{
		EntityId: EntityId,
		Position: Position,
		Name:     Name,
		Movable:  Movable,
	}
}

func ParsePlayerJoined(event []byte) *PlayerJoined {
	str := PlayerJoined{}
	str.FromBytes(event)
	return &str
}

func (str *PlayerJoined) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypePlayerJoined)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.EntityId)
	buffer = serialization.Append(buffer, str.Position)
	buffer = serialization.Append(buffer, str.Name)
	buffer = serialization.Append(buffer, str.Movable)

	return buffer
}

func (str *PlayerJoined) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.Position = &game.Position{}
	n += serialization.Read(data[n:], str.Position)
	str.Name = &game.Name{}
	n += serialization.Read(data[n:], str.Name)
	str.Movable = &game.Movable{}
	n += serialization.Read(data[n:], str.Movable)

	return n
}

func (str *PlayerJoined) Type() int16 {
	return TypePlayerJoined
}

func (str *PlayerJoined) String() string {
	return fmt.Sprintf("PlayerJoined: { EntityId: %v, Position: %v, Name: %v, Movable: %v,  }", str.EntityId, str.Position, str.Name, str.Movable)
}

func NewEntityMoved(EntityId int16, NewPosition *game.Position) *EntityMoved {
	return &EntityMoved{
		EntityId:    EntityId,
		NewPosition: NewPosition,
	}
}

func ParseEntityMoved(event []byte) *EntityMoved {
	str := EntityMoved{}
	str.FromBytes(event)
	return &str
}

func (str *EntityMoved) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeEntityMoved)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.EntityId)
	buffer = serialization.Append(buffer, str.NewPosition)

	return buffer
}

func (str *EntityMoved) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.NewPosition = &game.Position{}
	n += serialization.Read(data[n:], str.NewPosition)

	return n
}

func (str *EntityMoved) Type() int16 {
	return TypeEntityMoved
}

func (str *EntityMoved) String() string {
	return fmt.Sprintf("EntityMoved: { EntityId: %v, NewPosition: %v,  }", str.EntityId, str.NewPosition)
}

func NewEntityRemoved(EntityId int16) *EntityRemoved {
	return &EntityRemoved{
		EntityId: EntityId,
	}
}

func ParseEntityRemoved(event []byte) *EntityRemoved {
	str := EntityRemoved{}
	str.FromBytes(event)
	return &str
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
