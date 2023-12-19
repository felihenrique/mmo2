package packets

//go:generate ../../bin/serialize-generator

// INPUTS
type AckRequest struct {
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
	RequestId int16
	Entity    []byte
}

// EVENTS
type EntityRemoved struct {
	EntityId int16
}
