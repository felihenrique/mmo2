package systems

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/internal/shard-client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var InputSystem = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypePlayer, ecs.TypeTransform},
	func(timeStep float32, entities map[int16]*ecs.Entity) {
		if len(entities) == 0 {
			return
		}
		var player *ecs.Entity
		for _, ent := range entities {
			player = ent
			break
		}
		if player.Has(ecs.TypeMove) {
			return
		}
		var axisX float32
		var axisY float32
		if rl.IsKeyDown(rl.KeyRight) {
			axisX += 1
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			axisX -= 1
		}
		if rl.IsKeyDown(rl.KeyUp) {
			axisY -= 1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			axisY += 1
		}
		if axisX == 0 && axisY == 0 {
			return
		}
		fmt.Printf("Sent: %f, %f \n", axisX, axisY)
		transform := ecs.Get[*ecs.Transform](player, ecs.TypeTransform)
		moveX, moveY := axisX*16, axisY*16
		move := ecs.NewMove(
			moveX,
			moveY,
			transform.X+moveX,
			transform.Y+moveY,
		)
		shard.SendEventsChan <- packets.NewMoveRequest(move)
		player.Add(move)
	},
)
