package ecs

import "mmo2/game/ecs"

type IUpdatable interface {
	Update(world *ecs.World)
}
