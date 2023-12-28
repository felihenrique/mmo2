package ecs

import (
	"fmt"
)

type Entity struct {
	id         EntityID
	components map[ComponentID]IComponent
	world      *World
}

func (e *Entity) ID() int16 {
	return e.id
}

func (e *Entity) Add(components ...IComponent) {
	for _, c := range components {
		e.components[c.Type()] = c
	}
	e.world.updateSystems(e)
}

func (e *Entity) Remove(id ComponentID) {
	delete(e.components, id)
	e.world.updateSystems(e)
}

func (e *Entity) Has(id ComponentID) bool {
	_, ok := e.components[id]
	return ok
}

func (e *Entity) Get(componentId int16) IComponent {
	return e.components[componentId]
}

func (e *Entity) String() string {
	str := "Entity { Id: %d, Components: { %s } }"
	compsStr := ""
	for _, component := range e.components {
		compsStr = compsStr + component.String()
	}
	return fmt.Sprintf(str, e.id, compsStr)
}
