package events

type Move struct {
	Dx int32
	Dy int32
}

type EntityMoved struct {
	NewX     int32
	NewY     int32
	EntityId int32
	Velocity int8
}
