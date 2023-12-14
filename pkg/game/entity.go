package game

type Entity struct {
	id         int16
	components map[uint8]IComponent
	world      *World
}

func (e *Entity) ID() int16 {
	return e.id
}

func (e *Entity) Add(component IComponent) {
	e.components[component.ID()] = component
}

func (e *Entity) Remove(componentId uint8) {
	delete(e.components, componentId)
}

func (e *Entity) Get(componentId uint8) (IComponent, bool) {
	c, ok := e.components[componentId]
	return c, ok
}

func (e *Entity) Has(componentId uint8) bool {
	_, ok := e.components[componentId]
	return ok
}

func (e *Entity) Update() {
	for _, component := range e.components {
		component.Update(e.world)
	}
}
