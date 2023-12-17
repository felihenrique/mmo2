package game

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	TypePosition
	TypeRotation
	TypeMovable
)

func (str *Position) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.X)
	buffer = serialization.Write(buffer, str.Y)

	return buffer
}

func (str *Position) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.X)
	n += serialization.Read(data[n:], &str.Y)

	return n
}

func (str *Position) ID() int16 {
	return TypePosition
}

func (str *Rotation) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Rot)

	return buffer
}

func (str *Rotation) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Rot)

	return n
}

func (str *Rotation) ID() int16 {
	return TypeRotation
}

func (str *Movable) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Write(buffer, str.Velocity)

	return buffer
}

func (str *Movable) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Velocity)

	return n
}

func (str *Movable) ID() int16 {
	return TypeMovable
}
