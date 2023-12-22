package scene

import (
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
	"mmo2/pkg/packets"
	"mmo2/pkg/serialization"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	ShardClient *shard.Client
	isJoining   bool
	joined      bool
}

func (m *MainMenu) Init(world *game.World) {

}

func (m *MainMenu) handleJoin(response serialization.ISerializable) {
	m.joined = true
}

func (m *MainMenu) Update(world *game.World, timeStep float32) {
	if !m.isJoining {
		rl.DrawText("PRESS J TO JOIN", 100, 100, 50, rl.White)
		if rl.IsKeyPressed(rl.KeyJ) {
			println("SENT REQUEST")
			m.ShardClient.SendRequest(packets.NewJoinShardRequest("player", 0), m.handleJoin)
		}
	} else if !m.joined {
		rl.DrawText("JOINING...", 100, 100, 50, rl.White)
	} else {
		rl.DrawText("JOINED!", 100, 100, 50, rl.White)
	}
}
