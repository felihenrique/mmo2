package ecs

import (
	"testing"
)

func TestEntityBytes(t *testing.T) {
	w := newWorld()
	entity := w.NewEntityFrom(1, NewTransform(1, 2, 3), NewLiving("teste", 11))
	bytes := entity.ToBytes()
	ent2 := w.NewEntityFromBytes(bytes)
	if ent2.id != 1 {
		panic("wrong")
	}
}
