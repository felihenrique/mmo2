package game

import (
	"fmt"
	"mmo2/pkg/serialization"
)

type Entity struct {
	id         int16
	components map[int16]serialization.ISerializable
}

func (e *Entity) ID() int16 {
	return e.id
}

func (e *Entity) Add(components ...serialization.ISerializable) {
	for _, c := range components {
		e.components[c.Type()] = c
	}
}

func (e *Entity) Remove(componentId int16) {
	delete(e.components, componentId)
}

func (e *Entity) Has(componentId int16) bool {
	_, ok := e.components[componentId]
	return ok
}

func (e *Entity) Get(componentId int16) (serialization.ISerializable, bool) {
	c, ok := e.components[componentId]
	return c, ok
}

func (e *Entity) String() string {
	str := "Entity { Id: %d, Components: { %s } }"
	compsStr := ""
	for _, component := range e.components {
		compsStr = compsStr + component.String()
	}
	return fmt.Sprintf(str, e.id, compsStr)
}
