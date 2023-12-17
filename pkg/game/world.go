package game

import (
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
	w.entities = make([]*Entity, maxEntites)
	w.currentPos = 0
	w.maxEntites = maxEntites
	return &w
}

func (w *World) nextPos() int16 {
	if w.availablePos.Len() > 0 {
		return w.availablePos.Pop()
	}
	pos := w.currentPos
	w.currentPos += 1

	return pos
}

func (w *World) NewEntity() *Entity {
	if w.currentPos == w.maxEntites && w.availablePos.Len() == 0 {
		return nil
	}
	entity := Entity{}
	entity.components = make(map[int16]serialization.ISerializable)
	entity.id = w.nextPos()
	w.entities[entity.id] = &entity
	return &entity
}

func (w *World) GetEntity(id int16) *Entity {
	if id >= w.maxEntites {
		return nil
	}
	return w.entities[id]
}

func (w *World) RemoveEntity(entity *Entity) {
	w.entities[entity.id] = nil
	w.availablePos.Push(entity.id)
}
