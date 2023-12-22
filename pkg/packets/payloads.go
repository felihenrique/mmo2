package packets

import "mmo2/pkg/ecs"

//go:generate ../../bin/serialize-generator

// INPUTS
type AckRequest struct {
}
type RequestError struct {
	Message string
}
type MoveRequest struct {
	Dx int32
	Dy int32
}
type RotateRequest struct {
	Quantity float32
}

// REQUESTS
type JoinShardRequest struct {
	Name   string
	Color  *ecs.Color
	Portal int8
}
type JoinShardResponse struct {
	EntityId     int16
	Position     *ecs.Position
	Movable      *ecs.Movable
	Name         *ecs.Name
	PlayerCircle *ecs.PlayerCircle
}

// EVENTS
type PlayerJoined struct {
	EntityId     int16
	Position     *ecs.Position
	Name         *ecs.Name
	Movable      *ecs.Movable
	PlayerCircle *ecs.PlayerCircle
}
type EntityMoved struct {
	EntityId    int16
	NewPosition *ecs.Position
}
type EntityRemoved struct {
	EntityId int16
}
