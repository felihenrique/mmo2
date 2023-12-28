package ecs

//go:generate sh ../../bin/gen-all.sh

type Transform struct {
	X        float32
	Y        float32
	Rotation float32
}

type Player struct {
}

type Living struct {
	Name     string
	Velocity float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Circle struct {
	Radius float32
	Color  *Color
}

type Move struct {
	QuantityX float32
	QuantityY float32
	FinalX    float32
	FinalY    float32
}
