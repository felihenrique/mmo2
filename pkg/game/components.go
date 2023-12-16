package game

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
