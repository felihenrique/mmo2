package game

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
