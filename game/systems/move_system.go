package systems

import (
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
			if ds.SquaredDistance(transform.X, transform.Y, moveTo.X, moveTo.Y) == 0 {
				entity.Remove(ecs.TypeMoveTo)
				return
			}
			distanceX, distanceY := moveTo.X-transform.X, moveTo.Y-transform.Y
			normalizedX, normalizedY := ds.Normalize(distanceX, distanceY)
			stepX, stepY := ds.Scale(normalizedX, normalizedY, living.Velocity*deltaTime.Seconds())
			if ds.SquaredLength(stepX, stepY) > ds.SquaredLength(distanceX, distanceY) {
				stepX, stepY = distanceX, distanceY
			}
			transform.Move(stepX, stepY)
		}
	},
)
