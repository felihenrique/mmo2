package scene

import "mmo2/pkg/game"

type IScene interface {
	Init(world *game.World)
	Update(world *game.World, timeStep float32)
}
