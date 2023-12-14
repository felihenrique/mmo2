package game

const (
	TypeTransform = uint8(iota + 1)
	TypeMovable
)

func (c *Transform) ID() uint8 {
	return TypeTransform
}

func (c *Movable) ID() uint8 {
	return TypeMovable
}

func (c *Transform) Update(world *World) {}
func (c *Movable) Update(world *World)   {}
