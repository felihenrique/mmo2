package ecs

type EntityID = int16
type ComponentID = int16

type IComponent interface {
	Type() ComponentID
	String() string
}

func Get[T IComponent](entity *Entity, id ComponentID) T {
	comp := entity.components[id]
	return comp.(T)
}
