package game

type World struct {
	entities   [10000]*Entity
	currentPos int
}

func NewWorld() *World {
	w := World{}
	w.entities = [10000]*Entity{}
	w.currentPos = 0
	return &w
}

func (w *World) nextPos() int16 {
	for i, entity := range w.entities {
		if entity == nil {
			return int16(i)
		}
	}
	return -1
}

func (w *World) NewEntity() *Entity {
	entity := Entity{}
	entity.components = make(map[uint8]IComponent)
	entity.world = w
	id := w.nextPos()
	if id == -1 {
		return nil
	}
	entity.id = id
	w.entities[entity.id] = &entity
	return &entity
}
