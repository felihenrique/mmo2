package assets_scene

import (
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/game/scene"
	"mmo2/game/systems"
	"mmo2/internal/shard-client"
	"mmo2/pkg/serialization"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type mainMenu struct{}

func (mainMenu) Init() {
	ecs.MainWorld.AddSystem(systems.CircleRenderer)
	ecs.MainWorld.AddSystem(systems.InputSystem)
	ecs.MainWorld.AddSystem(systems.MoveSystem)
}

func (mainMenu) RenderGUI(client *shard.Client) {
	rl.DrawText("Press J to join", 0, 0, 30, rl.White)
	if rl.IsKeyPressed(rl.KeyJ) {
		client.SendRequest(
			packets.NewJoinShardRequest("Player", (*ecs.Color)(&rl.White), 0),
			func(response serialization.ISerializable) {
				scene.ChangeTo(BattleGround)
			},
		)
	}
}

func (mainMenu) Finalize() {}

var MainMenu = mainMenu{}
