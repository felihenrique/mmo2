package ecs

type IProcessor interface {
	Update(entities map[int16]*Entity)
}

type System struct {
	filter    int16
	entities  map[int16]*Entity
	processor IProcessor
}

func NewSystem(filter int16, processor IProcessor) *System {
	s := System{}
	s.filter = filter
	s.processor = processor
	s.entities = make(map[int16]*Entity)
	return &s
}

func (s *System) AddEntity(entity *Entity) {
	if entity.Has(s.filter) {
		s.entities[entity.ID()] = entity
	}
}

func (s *System) Update() {
	s.processor.Update(s.entities)
}
