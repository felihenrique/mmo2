package packets

import "mmo2/game/ecs"

//go:generate sh ../../bin/gen-all.sh

// INPUTS
type Ping struct {
}
type AckRequest struct {
}
type RequestError struct {
	Message string
}
type MoveRequest struct {
	Dx float64
	Dy float64
}

// REQUESTS
type JoinShardRequest struct {
	Name   string
	Color  *ecs.Color
	Portal int8
}
type JoinShardResponse struct {
	PlayerEntity []byte
}

// EVENTS
type EntityCreated struct {
	Entity []byte
}
type EntityMoved struct {
	EntityId int16
	Move     *ecs.MoveTo
}
type EntityRemoved struct {
	EntityId int16
}
