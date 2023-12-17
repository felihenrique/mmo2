package packets

//go:generate ../../bin/serialize-generator

// INPUTS
type AckInput struct {
	InputId int16
}
type MoveInput struct {
	InputId int16
	Dx      int32
	Dy      int32
}
type RotateInput struct {
	InputId  int16
	Quantity float32
}

// REQUESTS
type JoinShardRequest struct {
	RequestId int16
	Name      string
	Portal    uint8
}
type JoinShardResponse struct {
	RequestId int16
	Entity    []byte
}

// EVENTS
type EntityCreated struct {
	Entity []byte
}
type EntityUpdated struct {
	EntityId   int16
	Components []byte
}
type EntityRemoved struct {
	EntityId int16
}
