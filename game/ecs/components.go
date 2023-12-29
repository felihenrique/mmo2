package ecs

//go:generate sh ../../bin/gen-all.sh

type Transform struct {
	X        float64
	Y        float64
	Rotation float64
}

type Player struct {
}

type Living struct {
	Name     string
	Velocity float64
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Circle struct {
	Radius float64
	Color  *Color
}

type MoveTo struct {
	X float64
	Y float64
}
