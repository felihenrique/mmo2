package assets_scene

import (
	"mmo2/internal/shard-client"
)

type battleGround struct{}

func (battleGround) Init() {
}

func (battleGround) RenderGUI(client *shard.Client) {

}

func (battleGround) Finalize() {}

var BattleGround = battleGround{}
