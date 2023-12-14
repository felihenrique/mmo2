package game

type Transform struct {
	X        int32
	Y        int32
	Rotation float32
}

type Movable struct {
	Velocity float32
}
