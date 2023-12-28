package ecs

import (
	"mmo2/pkg/ds"
)

type World struct {
	entities     map[int16]*Entity
	currentPos   EntityID
	availablePos ds.Queue[EntityID]
	systems      []*System
}

func newWorld() *World {
	w := World{}
	w.systems = make([]*System, 0)
	w.entities = make(map[int16]*Entity)
	w.currentPos = 0
	return &w
}

var MainWorld = newWorld()

func (w *World) nextPos() int16 {
	if w.availablePos.Len() > 0 {
		return w.availablePos.Pop()
	}
	return -1
}

func (w *World) NewEntity() *Entity {
	entity := Entity{}
	entity.components = make(map[int16]IComponent)
	entity.world = w
	nextPos := w.nextPos()
	if nextPos == -1 {
		entity.id = int16(len(w.entities))
	} else {
		entity.id = nextPos
	}
	w.entities[entity.id] = &entity
	w.updateSystems(&entity)
	return &entity
}

func (w *World) NewEntityFrom(id EntityID, components ...IComponent) *Entity {
	entity := Entity{}
	entity.id = id
	entity.components = make(map[int16]IComponent)
	entity.world = w
	entity.Add(components...)
	w.entities[id] = &entity
	w.updateSystems(&entity)
	return &entity
}

func (w *World) GetEntity(id int16) *Entity {
	return w.entities[id]
}

func (w *World) RemoveEntity(entityId int16) {
	delete(w.entities, entityId)
	w.availablePos.Push(entityId)
	for _, system := range w.systems {
		system.RemoveEntity(entityId)
	}
}

func (w *World) Entities() map[int16]*Entity {
	return w.entities
}

func (s *World) AddSystem(system *System) {
	s.systems = append(s.systems, system)
}

func (s *World) Update() {
	for _, s := range s.systems {
		s.Update(0)
	}
}

func (s *World) updateSystems(entity *Entity) {
	for _, system := range s.systems {
		system.CheckEntity(entity)
	}
}
