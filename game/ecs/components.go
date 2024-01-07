package ecs

//go:generate sh ../../bin/gen-all.sh

type Transform struct {
	X        float64
	Y        float64
	Rotation float64
}

func (t *Transform) Move(x float64, y float64) {
	t.X += x
	t.Y += y
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

func (m *MoveTo) Add(x float64, y float64) {
	m.X += x
	m.Y += y
}
