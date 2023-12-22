package packets

import "mmo2/pkg/game"

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
	Portal int8
}
type JoinShardResponse struct {
	EntityId int16
	Position *game.Position
	Movable  *game.Movable
	Name     *game.Name
}

// EVENTS
type PlayerJoined struct {
	EntityId int16
	Position *game.Position
	Name     *game.Name
	Movable  *game.Movable
}
type EntityMoved struct {
	EntityId    int16
	NewPosition *game.Position
}
type EntityRemoved struct {
	EntityId int16
}
