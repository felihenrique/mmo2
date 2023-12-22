package ecs

import (
	"fmt"
	"mmo2/pkg/ds"
	"mmo2/pkg/serialization"
)

type World struct {
	entities     map[int16]*Entity
	currentPos   int16
	availablePos ds.Queue[int16]
}

func NewWorld() *World {
	w := World{}
	w.entities = make(map[int16]*Entity)
	w.currentPos = 0
	return &w
}

func (w *World) nextPos() int16 {
	if w.availablePos.Len() > 0 {
		return w.availablePos.Pop()
	}
	return -1
}

func (w *World) NewEntity() *Entity {
	entity := Entity{}
	entity.components = make(map[int16]serialization.ISerializable)
	nextPos := w.nextPos()
	if nextPos == -1 {
		entity.id = int16(len(w.entities))
		w.entities[entity.id] = &entity
	} else {
		entity.id = nextPos
		w.entities[entity.id] = &entity
	}

	return &entity
}

func (w *World) NewEntityFrom(id int16, components []serialization.ISerializable) *Entity {
	entity := Entity{}
	entity.id = id
	entity.components = make(map[int16]serialization.ISerializable)
	for _, c := range components {
		entity.Add(c)
	}
	w.entities[id] = &entity
	return &entity
}

func (w *World) GetEntity(id int16) *Entity {
	return w.entities[id]
}

func (w *World) RemoveEntity(entityId int16) {
	if w.entities[entityId] == nil {
		fmt.Printf("WRONG situation, removing entity with id: %d \n", entityId)
		return
	}
	delete(w.entities, entityId)
	w.availablePos.Push(entityId)
}

func (w *World) Entites() map[int16]*Entity {
	return w.entities
}
