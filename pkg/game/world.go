package game

import (
	"fmt"
	"mmo2/pkg/ds"
	"mmo2/pkg/serialization"
)

type World struct {
	entities     []*Entity
	maxEntites   int16
	currentPos   int16
	availablePos ds.Queue[int16]
}

func NewWorld(maxEntites int16) *World {
	w := World{}
	w.entities = make([]*Entity, 0, maxEntites)
	w.currentPos = 0
	w.maxEntites = maxEntites
	return &w
}

func (w *World) nextPos() int16 {
	if w.availablePos.Len() > 0 {
		return w.availablePos.Pop()
	}
	return -1
}

func (w *World) NewEntity() *Entity {
	if w.currentPos == w.maxEntites && w.availablePos.Len() == 0 {
		return nil
	}
	entity := Entity{}
	entity.components = make(map[int16]serialization.ISerializable)
	nextPos := w.nextPos()
	if nextPos == -1 {
		entity.id = int16(len(w.entities))
		w.entities = append(w.entities, &entity)
	} else {
		entity.id = nextPos
		w.entities[entity.id] = &entity
	}

	return &entity
}

func (w *World) GetEntity(id int16) *Entity {
	if id >= w.maxEntites {
		return nil
	}
	return w.entities[id]
}

func (w *World) RemoveEntity(entityId int16) {
	if w.entities[entityId] == nil {
		fmt.Printf("WRONG situation, removing entity with id: %d \n", entityId)
		return
	}
	w.entities[entityId] = nil
	w.availablePos.Push(entityId)
}

func (w *World) Entites() []*Entity {
	return w.entities
}
