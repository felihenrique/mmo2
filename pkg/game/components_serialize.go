package game

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	TypeTransform
	TypeRotation
	TypeMovable
)

func (str *Transform) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.X)
	buffer = serialization.WriteBinary(buffer, str.Y)

	return buffer
}

func (str *Transform) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.X)
	n += serialization.ReadBinary(data[n:], &str.Y)

	return n
}

func (str *Transform) EvType() int16 {
	return TypeTransform
}

func (str *Rotation) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Rot)

	return buffer
}

func (str *Rotation) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Rot)

	return n
}

func (str *Rotation) EvType() int16 {
	return TypeRotation
}

func (str *Movable) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.WriteBinary(buffer, str.Velocity)

	return buffer
}

func (str *Movable) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.ReadBinary(data[n:], &str.Velocity)

	return n
}

func (str *Movable) EvType() int16 {
	return TypeMovable
}
