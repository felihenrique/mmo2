package ecs

import rl "github.com/gen2brain/raylib-go/raylib"

//go:generate ../../bin/serialize-generator

type Position struct {
	X int32
	Y int32
}

type Rotation struct {
	Rot float32
}

type Movable struct {
	Velocity float32
}

type Name struct {
	Value string
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func ToEcsColor(color rl.Color) *Color {
	return &Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}
}

type PlayerCircle struct {
	Radius float32
	Color  *Color
}
