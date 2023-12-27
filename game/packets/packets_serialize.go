package packets

import (
	"fmt"
	"mmo2/game/ecs"
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

func NewJoinShardRequest(Name string, Color *ecs.Color, Portal int8) *JoinShardRequest {
	return &JoinShardRequest{
		Name:   Name,
		Color:  Color,
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
	buffer = serialization.Append(buffer, str.Color)
	buffer = serialization.Append(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Name)

	str.Color = &ecs.Color{}
	n += serialization.Read(data[n:], str.Color)
	n += serialization.Read(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) Type() int16 {
	return TypeJoinShardRequest
}

func (str *JoinShardRequest) String() string {
	return fmt.Sprintf("JoinShardRequest: { Name: %v, Color: %v, Portal: %v,  }", str.Name, str.Color, str.Portal)
}

func NewJoinShardResponse(EntityId int16, Transform *ecs.Transform, Living *ecs.Living, PlayerCircle *ecs.Circle) *JoinShardResponse {
	return &JoinShardResponse{
		EntityId:     EntityId,
		Transform:    Transform,
		Living:       Living,
		PlayerCircle: PlayerCircle,
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
	buffer = serialization.Append(buffer, str.Transform)
	buffer = serialization.Append(buffer, str.Living)
	buffer = serialization.Append(buffer, str.PlayerCircle)

	return buffer
}

func (str *JoinShardResponse) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.Transform = &ecs.Transform{}
	n += serialization.Read(data[n:], str.Transform)
	str.Living = &ecs.Living{}
	n += serialization.Read(data[n:], str.Living)
	str.PlayerCircle = &ecs.Circle{}
	n += serialization.Read(data[n:], str.PlayerCircle)

	return n
}

func (str *JoinShardResponse) Type() int16 {
	return TypeJoinShardResponse
}

func (str *JoinShardResponse) String() string {
	return fmt.Sprintf("JoinShardResponse: { EntityId: %v, Transform: %v, Living: %v, PlayerCircle: %v,  }", str.EntityId, str.Transform, str.Living, str.PlayerCircle)
}

func NewPlayerJoined(EntityId int16, Transform *ecs.Transform, Living *ecs.Living, PlayerCircle *ecs.Circle) *PlayerJoined {
	return &PlayerJoined{
		EntityId:     EntityId,
		Transform:    Transform,
		Living:       Living,
		PlayerCircle: PlayerCircle,
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
	buffer = serialization.Append(buffer, str.Transform)
	buffer = serialization.Append(buffer, str.Living)
	buffer = serialization.Append(buffer, str.PlayerCircle)

	return buffer
}

func (str *PlayerJoined) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.Transform = &ecs.Transform{}
	n += serialization.Read(data[n:], str.Transform)
	str.Living = &ecs.Living{}
	n += serialization.Read(data[n:], str.Living)
	str.PlayerCircle = &ecs.Circle{}
	n += serialization.Read(data[n:], str.PlayerCircle)

	return n
}

func (str *PlayerJoined) Type() int16 {
	return TypePlayerJoined
}

func (str *PlayerJoined) String() string {
	return fmt.Sprintf("PlayerJoined: { EntityId: %v, Transform: %v, Living: %v, PlayerCircle: %v,  }", str.EntityId, str.Transform, str.Living, str.PlayerCircle)
}

func NewEntityMoved(EntityId int16, NewPosition *ecs.Transform) *EntityMoved {
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

	str.NewPosition = &ecs.Transform{}
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
