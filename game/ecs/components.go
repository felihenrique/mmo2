package ecs

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

type Circle struct {
	Radius float32
	Color  *Color
}

type MoveTo struct {
	X int32
	Y int32
}
