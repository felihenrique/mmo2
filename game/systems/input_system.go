package systems

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/internal/shard-client"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var lastAdded time.Time
var passed time.Duration
var numPerSec int

var InputSystem = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypePlayer, ecs.TypeTransform},
	func(deltaTime time.Duration, entities map[int16]*ecs.Entity) {
		if len(entities) == 0 {
			return
		}
		var player *ecs.Entity
		for _, ent := range entities {
			player = ent
			break
		}
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
		fmt.Printf("added move to: %f, %f \n", transform.X+moveX, transform.Y+moveY)
		passed += time.Since(lastAdded)
		lastAdded = time.Now()
		numPerSec += 1
		if passed.Seconds() >= 1 {
			println("num per sec: ", numPerSec)
			numPerSec = 0
			passed = 0
		}
		move := ecs.NewMoveTo(transform.X+moveX, transform.Y+moveY)
		shard.SendEventsChan <- packets.NewMoveRequest(moveX, moveY)
		player.Add(move)
	},
)
