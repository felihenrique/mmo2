package game

const (
	TypeTransform = uint8(iota + 1)
)

func (Transform) ID() uint8 {
	return TypeTransform
}

func (Transform) Update(world *World) {}
