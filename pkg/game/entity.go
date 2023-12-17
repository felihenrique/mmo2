package game

import (
	"encoding/binary"
	"mmo2/pkg/serialization"
)

type Entity struct {
	id         int16
	components map[int16]any
	world      *World
}

func (e *Entity) ID() int16 {
	return e.id
}

func (e *Entity) Add(component IComponent) {
	e.components[component.ID()] = component
}

func (e *Entity) Remove(componentId int16) {
	delete(e.components, componentId)
}

func (e *Entity) Get(componentId int16) (IComponent, bool) {
	c, ok := e.components[componentId]
	return c.(IComponent), ok
}

func (e *Entity) Has(componentId int16) bool {
	_, ok := e.components[componentId]
	return ok
}

func (e *Entity) Update() {
	for _, item := range e.components {
		updatable := item.(IUpdatable)
		updatable.Update(e.world)
	}
}

func (e *Entity) Serialize() []byte {
	data := make([]byte, 4)
	binary.BigEndian.PutUint16(data, uint16(e.id))
	for _, value := range e.components {
		item, ok := value.(serialization.ISerializable)
		if !ok {
			continue
		}
		data = serialization.Write(data, item.ID())
		data = append(data, item.ToBytes()...)
	}
	binary.BigEndian.PutUint16(data, uint16(len(data)))
	return data
}

func (e *Entity) Unserialize(data []byte) {

}
