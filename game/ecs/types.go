package ecs

type IUpdatable interface {
	Update(world *World)
}
