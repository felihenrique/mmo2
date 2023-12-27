package ecs

import rl "github.com/gen2brain/raylib-go/raylib"

//go:generate ../../bin/serialize-generator

type Transform struct {
	X        int32
	Y        int32
	Rotation float32
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

func ToEcsColor(color rl.Color) *Color {
	return &Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}
}

type Circle struct {
	Radius float32
	Color  *Color
}

type MoveTo struct {
	X int32
	Y int32
}
