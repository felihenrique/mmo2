package packets

import "mmo2/game/ecs"

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

// EVENTS
type PlayerJoined struct {
	EntityId     int16
	Transform    *ecs.Transform
	Living       *ecs.Living
	PlayerCircle *ecs.Circle
}
type EntityMoved struct {
	EntityId    int16
	NewPosition *ecs.Transform
}
type EntityRemoved struct {
	EntityId int16
}
