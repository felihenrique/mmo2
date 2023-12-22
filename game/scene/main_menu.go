package scene

import (
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/internal/shard-client"
	"mmo2/pkg/serialization"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	ShardClient *shard.Client
	isJoining   bool
	joined      bool
}

func (m *MainMenu) Init(world *ecs.World) {

}

func (m *MainMenu) handleJoin(response serialization.ISerializable) {
	m.joined = true
	m.isJoining = false
}

func (m *MainMenu) Update(world *ecs.World, timeStep float32) {
	if !m.isJoining && !m.joined {
		rl.DrawText("PRESS J TO JOIN", 100, 100, 50, rl.White)
		if rl.IsKeyPressed(rl.KeyJ) {
			println("SENT REQUEST")
			m.ShardClient.SendRequest(
				packets.NewJoinShardRequest("player", ecs.ToEcsColor(rl.DarkBlue), 0),
				m.handleJoin,
			)
		}
	} else if m.isJoining {
		rl.DrawText("JOINING...", 100, 100, 50, rl.White)
	} else {
		rl.DrawText("JOINED!", 100, 100, 50, rl.White)
	}
}
