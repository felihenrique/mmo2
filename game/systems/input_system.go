package systems

import (
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/internal/shard-client"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var InputSystem = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypePlayer, ecs.TypeTransform},
	func(deltaTime time.Duration, entities map[int16]*ecs.Entity) {
		for _, player := range entities {
			if player.Has(ecs.TypeMoveTo) {
				return
			}
			var axisX float64
			var axisY float64
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
			transform := ecs.Get[*ecs.Transform](player, ecs.TypeTransform)
			moveX, moveY := axisX*16, axisY*16
			move := ecs.NewMoveTo(transform.X+moveX, transform.Y+moveY)
			shard.SendEventsChan <- packets.NewMoveRequest(moveX, moveY)
			player.Add(move)
		}
	},
)
