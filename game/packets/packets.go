package packets

import "mmo2/game/ecs"

//go:generate sh ../../bin/gen-all.sh

// INPUTS
type AckRequest struct {
}
type RequestError struct {
	Message string
}
type MoveRequest struct {
	Move *ecs.Move
}

// REQUESTS
type JoinShardRequest struct {
	Name   string
	Color  *ecs.Color
	Portal int8
}
type JoinShardResponse struct {
	EntityId     int16
	Transform    *ecs.Transform
	Living       *ecs.Living
	PlayerCircle *ecs.Circle
}

// EVENTS
type PlayerJoined struct {
	EntityId     int16
	Transform    *ecs.Transform
	Living       *ecs.Living
	PlayerCircle *ecs.Circle
}
type EntityMoved struct {
	EntityId int16
	Move     *ecs.Move
}
type EntityRemoved struct {
	EntityId int16
}
