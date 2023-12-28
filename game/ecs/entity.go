package ecs

import (
	"fmt"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/serialization"
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

func (e *Entity) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, e.id)
	buffer = serialization.Append(buffer, int16(len(e.components)))
	for _, item := range e.components {
		component := item.(serialization.ISerializable)
		buffer = append(buffer, component.ToBytes(0)...)
	}
	return buffer
}

func (e *Entity) FromBytes(data []byte) {
	n := serialization.Read(data, &e.id)
	var numComponents int16
	n += serialization.Read(data[n:], &numComponents)
	for i := int16(0); i < numComponents; i++ {
		component, readed := Mapper[event_utils.GetType(data[n:])](data[n:])
		n += readed
		e.Add(component)
	}
}
