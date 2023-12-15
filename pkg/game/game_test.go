package game

import (
	"testing"
)

func TestWorld(t *testing.T) {
	world := NewWorld(1000)
	entity := world.NewEntity()
	if world.entities[0] == nil {
		panic("wrong")
	}
	if entity.id != 0 {
		panic("wrong")
	}
	entity.Add(&Transform{
		X: 123,
		Y: 111,
	})
	if entity.components[TypeTransform] == nil {
		panic("wrong")
	}
	if !entity.Has(TypeTransform) {
		panic("wrong")
	}
	entity.Remove(TypeTransform)
	if entity.components[TypeTransform] != nil {
		panic("wrong")
	}
	if entity.Has(TypeTransform) {
		panic("wrong")
	}
}
