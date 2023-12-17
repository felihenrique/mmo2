package game

import "mmo2/pkg/serialization"

const (
	TypeNone = int16(iota)
	TypePosition
	TypeRotation
	TypeMovable
)

func Read(data []byte) (serialization.ISerializable, int16) {
	var strType int16
	n := serialization.Read(data, &strType)
	switch strType {

	case TypePosition:
		var str Position
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeRotation:
		var str Rotation
		n += str.FromBytes(data[n:])
		return &str, n

	case TypeMovable:
		var str Movable
		n += str.FromBytes(data[n:])
		return &str, n

	default:
		panic("wrong type")
	}
}

func (str *Position) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.X)
	buffer = serialization.Append(buffer, str.Y)

	return buffer
}

func (str *Position) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.X)
	n += serialization.Read(data[n:], &str.Y)

	return n
}

func (str *Position) Type() int16 {
	return TypePosition
}

func (str *Rotation) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Rot)

	return buffer
}

func (str *Rotation) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Rot)

	return n
}

func (str *Rotation) Type() int16 {
	return TypeRotation
}

func (str *Movable) ToBytes() []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, str.Velocity)

	return buffer
}

func (str *Movable) FromBytes(data []byte) int16 {
	var n int16 = 0
	n += serialization.Read(data[n:], &str.Velocity)

	return n
}

func (str *Movable) Type() int16 {
	return TypeMovable
}
