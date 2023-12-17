package game

import (
	"encoding/binary"
	"mmo2/pkg/serialization"
)

type Entity struct {
	id         int16
	components map[int16]serialization.ISerializable
}

func (e *Entity) ID() int16 {
	return e.id
}

func (e *Entity) Add(component serialization.ISerializable) {
	e.components[component.Type()] = component
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

func (e *Entity) ToBytes() []byte {
	data := make([]byte, 4)
	binary.BigEndian.PutUint16(data, uint16(e.id))
	binary.BigEndian.PutUint16(data, uint16(len(e.components)))
	for _, value := range e.components {
		data = serialization.Append(data, value.Type())
		data = append(data, value.ToBytes()...)
	}
	return data
}

func (e *Entity) FromBytes(data []byte) int16 {
	n := serialization.Read(data, &e.id)
	var cNumber int16
	n += serialization.Read(data, &cNumber)
	for i := int16(0); i < cNumber; i++ {
		var cType int16
		n += serialization.Read(data[n:], cType)
		component, readed := Read(data[n:])
		n += readed
		e.components[component.Type()] = component
	}
	return n
}
