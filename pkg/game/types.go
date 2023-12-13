package game

type IComponent interface {
	ID() uint8
	Update(world *World)
}
