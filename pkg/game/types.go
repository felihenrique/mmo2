package game

type IComponent interface {
	ID() uint8
}

type IUpdatable interface {
	Update(world *World)
}
