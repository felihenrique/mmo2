package packets

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/pkg/serialization"
)

const (
	TypeNone = int16(iota - 1)
	TypePing
	TypeAckRequest
	TypeRequestError
	TypeMoveRequest
	TypeJoinShardRequest
	TypeJoinShardResponse
	TypePlayerJoined
	TypeEntityMoved
	TypeEntityRemoved
)

func NewPing() *Ping {
	return &Ping{}
}

func ParsePing(event []byte) (*Ping, int16) {
	str := Ping{}
	n := str.FromBytes(event)
	return &str, n
}

func (str *Ping) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypePing)
	buffer = serialization.Append(buffer, eventId)

	return buffer
}

func (str *Ping) FromBytes(data []byte) int16 {
	var n int16 = 4

	return n
}

func (str *Ping) Type() int16 {
	return TypePing
}

func (str *Ping) String() string {
	return fmt.Sprintf("Ping: {  }")
}

func NewAckRequest() *AckRequest {
	return &AckRequest{}
}

func ParseAckRequest(event []byte) (*AckRequest, int16) {
	str := AckRequest{}
	n := str.FromBytes(event)
	return &str, n
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

func ParseRequestError(event []byte) (*RequestError, int16) {
	str := RequestError{}
	n := str.FromBytes(event)
	return &str, n
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

func NewMoveRequest(Dx float64, Dy float64) *MoveRequest {
	return &MoveRequest{
		Dx: Dx,
		Dy: Dy,
	}
}

func ParseMoveRequest(event []byte) (*MoveRequest, int16) {
	str := MoveRequest{}
	n := str.FromBytes(event)
	return &str, n
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

func NewJoinShardRequest(Name string, Color *ecs.Color, Portal int8) *JoinShardRequest {
	return &JoinShardRequest{
		Name:   Name,
		Color:  Color,
		Portal: Portal,
	}
}

func ParseJoinShardRequest(event []byte) (*JoinShardRequest, int16) {
	str := JoinShardRequest{}
	n := str.FromBytes(event)
	return &str, n
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

func NewJoinShardResponse(PlayerEntity []byte) *JoinShardResponse {
	return &JoinShardResponse{
		PlayerEntity: PlayerEntity,
	}
}

func ParseJoinShardResponse(event []byte) (*JoinShardResponse, int16) {
	str := JoinShardResponse{}
	n := str.FromBytes(event)
	return &str, n
}

func (str *JoinShardResponse) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeJoinShardResponse)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.PlayerEntity)

	return buffer
}

func (str *JoinShardResponse) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.PlayerEntity)

	return n
}

func (str *JoinShardResponse) Type() int16 {
	return TypeJoinShardResponse
}

func (str *JoinShardResponse) String() string {
	return fmt.Sprintf("JoinShardResponse: { PlayerEntity: %v,  }", str.PlayerEntity)
}

func NewPlayerJoined(Entity []byte) *PlayerJoined {
	return &PlayerJoined{
		Entity: Entity,
	}
}

func ParsePlayerJoined(event []byte) (*PlayerJoined, int16) {
	str := PlayerJoined{}
	n := str.FromBytes(event)
	return &str, n
}

func (str *PlayerJoined) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypePlayerJoined)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Entity)

	return buffer
}

func (str *PlayerJoined) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Entity)

	return n
}

func (str *PlayerJoined) Type() int16 {
	return TypePlayerJoined
}

func (str *PlayerJoined) String() string {
	return fmt.Sprintf("PlayerJoined: { Entity: %v,  }", str.Entity)
}

func NewEntityMoved(EntityId int16, Move *ecs.MoveTo) *EntityMoved {
	return &EntityMoved{
		EntityId: EntityId,
		Move:     Move,
	}
}

func ParseEntityMoved(event []byte) (*EntityMoved, int16) {
	str := EntityMoved{}
	n := str.FromBytes(event)
	return &str, n
}

func (str *EntityMoved) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeEntityMoved)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.EntityId)
	buffer = serialization.Append(buffer, str.Move)

	return buffer
}

func (str *EntityMoved) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.EntityId)

	str.Move = &ecs.MoveTo{}
	n += serialization.Read(data[n:], str.Move)

	return n
}

func (str *EntityMoved) Type() int16 {
	return TypeEntityMoved
}

func (str *EntityMoved) String() string {
	return fmt.Sprintf("EntityMoved: { EntityId: %v, Move: %v,  }", str.EntityId, str.Move)
}

func NewEntityRemoved(EntityId int16) *EntityRemoved {
	return &EntityRemoved{
		EntityId: EntityId,
	}
}

func ParseEntityRemoved(event []byte) (*EntityRemoved, int16) {
	str := EntityRemoved{}
	n := str.FromBytes(event)
	return &str, n
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
