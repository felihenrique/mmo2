package events

type Ack struct{}

func (Ack) toBytes() []byte       { return []byte{} }
func (Ack) fromBytes(data []byte) {}
func (Ack) evType() int16         { return TypeAck }

type Move struct {
	Dx int32
	Dy int32
}

type EntityMoved struct {
	NewX     int32
	NewY     int32
	EntityId int16
	Velocity float32
}
