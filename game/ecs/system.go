package ecs

type IProcessor = func(timeStep float32, entities map[int16]*Entity)

type System struct {
	filter    []ComponentID
	entities  map[EntityID]*Entity
	processor IProcessor
}

func NewSystem(filter []ComponentID, processor IProcessor) *System {
	s := System{}
	s.filter = filter
	s.processor = processor
	s.entities = make(map[int16]*Entity)
	return &s
}

func (s *System) CheckEntity(entity *Entity) {
	hasAll := true
	for _, f := range s.filter {
		hasAll = entity.Has(f) && hasAll
	}
	if hasAll {
		s.entities[entity.id] = entity
	}
}

func (s *System) RemoveEntity(entityId int16) {
	delete(s.entities, entityId)
}

func (s *System) Update(timeStep float32) {
	s.processor(timeStep, s.entities)
}
