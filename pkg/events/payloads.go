package events

//go:generate ../../bin/serialize-generator

// ACK REQUESTS
type Ack struct{}

func (Ack) toBytes() []byte             { return []byte{} }
func (Ack) fromBytes(data []byte) int16 { return 0 }
func (Ack) evType() int16               { return TypeAck }

// REQUEST
type MoveRequest struct {
	Dx int32
	Dy int32
}
type RotateRequest struct {
	Quantity float32
}
type JoinShardRequest struct {
	Portal uint8
}

// EVENTS
type EntityCreated struct {
	Data []byte
}
type EntityUpdated struct {
	Data []byte
}
type EntityRemoved struct {
	EntityId int16
}
