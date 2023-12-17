package game

type IComponent interface {
	Type() int16
}

type IUpdatable interface {
	Update(world *World)
}
