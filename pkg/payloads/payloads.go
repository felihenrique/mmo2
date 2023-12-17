package payloads

//go:generate ../../bin/serialize-generator

// INPUTS
type AckInput struct {
	EventId int16
}
type MoveInput struct {
	Dx int32
	Dy int32
}
type RotateInput struct {
	Quantity float32
}

// REQUESTS
type JoinShardRequest struct {
	Name   string
	Portal uint8
}
type JoinShardResponse struct {
	Entity []byte
}

// EVENTS
type EntityCreated struct {
	Entity []byte
}
type EntityUpdated struct {
	Entity []byte
}
type EntityRemoved struct {
	EntityId int16
}
