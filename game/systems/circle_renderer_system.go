package systems

import (
	"image/color"
	"mmo2/game/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var CircleRenderer = ecs.NewSystem(
	[]ecs.ComponentID{ecs.TypeCircle, ecs.TypeTransform},
	func(timeStep float32, entities map[ecs.EntityID]*ecs.Entity) {
		for _, entity := range entities {
			circle := ecs.Get[*ecs.Circle](entity, ecs.TypeCircle)
			transform := ecs.Get[*ecs.Transform](entity, ecs.TypeTransform)
			rl.DrawEllipse(
				int32(transform.X), int32(transform.Y), circle.Radius,
				circle.Radius, color.RGBA(*circle.Color),
			)
		}
	},
)
