package assets_scene

import (
	"mmo2/game/ecs"
	"mmo2/game/systems"
	"mmo2/internal/shard-client"
)

type battleGround struct{}

func (battleGround) Init() {
	ecs.MainWorld.AddSystem(systems.CircleRenderer)
}

func (battleGround) RenderGUI(client *shard.Client) {

}

func (battleGround) Finalize() {}

var BattleGround = battleGround{}
