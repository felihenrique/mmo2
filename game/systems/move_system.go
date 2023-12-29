package systems

import (
	"math"
	"mmo2/game/ecs"
	"mmo2/pkg/ds"
	"time"
)

var MoveSystem = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypeMoveTo, ecs.TypeTransform, ecs.TypeLiving},
	func(deltaTime time.Duration, entities map[int16]*ecs.Entity) {
		for _, entity := range entities {
			moveTo := ecs.Get[*ecs.MoveTo](entity, ecs.TypeMoveTo)
			living := ecs.Get[*ecs.Living](entity, ecs.TypeLiving)
			transform := ecs.Get[*ecs.Transform](entity, ecs.TypeTransform)
			distanceX, distanceY := moveTo.X-transform.X, moveTo.Y-transform.Y
			normalizedX, normalizedY := ds.NormalizeVec(distanceX, distanceY)
			stepX := living.Velocity * deltaTime.Seconds() * normalizedX
			stepY := living.Velocity * deltaTime.Seconds() * normalizedY
			if math.Abs(stepX) > math.Abs(distanceX) || math.Abs(stepY) > math.Abs(distanceY) {
				stepX = distanceX
				stepY = distanceY
			}
			transform.X += stepX
			transform.Y += stepY
			if ds.Distance(transform.X, transform.Y, moveTo.X, moveTo.Y) <= 1 {
				entity.Remove(ecs.TypeMoveTo)
			}
		}
	},
)
