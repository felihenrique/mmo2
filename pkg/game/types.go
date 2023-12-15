package game

type IComponent interface {
	ID() int16
}

type IUpdatable interface {
	Update(world *World)
}
