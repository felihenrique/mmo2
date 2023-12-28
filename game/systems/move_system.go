package systems

import (
	"fmt"
	"math"
	"mmo2/game/ecs"
)

var MoveSystem = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypeMove, ecs.TypeTransform},
	func(timeStep float32, entities map[int16]*ecs.Entity) {
		for _, entity := range entities {
			move := ecs.Get[*ecs.Move](entity, ecs.TypeMove)
			stepX := move.QuantityX / 15
			stepY := move.QuantityY / 15
			fmt.Printf("Step: %f, %f. Final: %f, %f \n", stepX, stepY, move.FinalX, move.FinalY)
			transform := ecs.Get[*ecs.Transform](entity, ecs.TypeTransform)
			transform.X += stepX
			transform.Y += stepY

			if math.Abs(float64(transform.X-move.FinalX)) <= 0.01 && math.Abs(float64(transform.Y-move.FinalY)) <= 0.01 {
				entity.Remove(ecs.TypeMove)
			}
		}
	},
)
