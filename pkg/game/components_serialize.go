package game

import (
	"fmt"
	"mmo2/pkg/serialization"
)

const (
	TypeNone = int16(iota)
	TypePosition
	TypeRotation
	TypeMovable
)

func (str *Position) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypePosition)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.X)
	buffer = serialization.Append(buffer, str.Y)

	return buffer
}

func (str *Position) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.X)
	n += serialization.Read(data[n:], &str.Y)

	return n
}

func (str *Position) Type() int16 {
	return TypePosition
}

func (str *Position) String() string {
	return fmt.Sprintf("Position: { X: %v, Y: %v,  }", str.X, str.Y)
}

func (str *Rotation) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeRotation)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Rot)

	return buffer
}

func (str *Rotation) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Rot)

	return n
}

func (str *Rotation) Type() int16 {
	return TypeRotation
}

func (str *Rotation) String() string {
	return fmt.Sprintf("Rotation: { Rot: %v,  }", str.Rot)
}

func (str *Movable) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeMovable)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Velocity)

	return buffer
}

func (str *Movable) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Velocity)

	return n
}

func (str *Movable) Type() int16 {
	return TypeMovable
}

func (str *Movable) String() string {
	return fmt.Sprintf("Movable: { Velocity: %v,  }", str.Velocity)
}
