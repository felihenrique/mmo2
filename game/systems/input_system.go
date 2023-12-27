package systems

import (
	"mmo2/game/ecs"
)

var InputSystem = ecs.NewSystem(
	[]ecs.ComponentID{},
	func(timeStep float32, entities map[int16]*ecs.Entity) {

	},
)
