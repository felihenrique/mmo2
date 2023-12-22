package scene

import "mmo2/game/ecs"

type IScene interface {
	Init(world *ecs.World)
	Update(world *ecs.World, timeStep float32)
}
