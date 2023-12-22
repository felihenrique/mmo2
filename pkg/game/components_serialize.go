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
	TypeName
)

func NewPosition(X int32, Y int32) *Position {
	return &Position{
		X: X,
		Y: Y,
	}
}

func ParsePosition(event []byte) *Position {
	str := Position{}
	str.FromBytes(event)
	return &str
}

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

func NewRotation(Rot float32) *Rotation {
	return &Rotation{
		Rot: Rot,
	}
}

func ParseRotation(event []byte) *Rotation {
	str := Rotation{}
	str.FromBytes(event)
	return &str
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

func NewMovable(Velocity float32) *Movable {
	return &Movable{
		Velocity: Velocity,
	}
}

func ParseMovable(event []byte) *Movable {
	str := Movable{}
	str.FromBytes(event)
	return &str
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

func NewName(Value string) *Name {
	return &Name{
		Value: Value,
	}
}

func ParseName(event []byte) *Name {
	str := Name{}
	str.FromBytes(event)
	return &str
}

func (str *Name) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeName)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Value)

	return buffer
}

func (str *Name) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Value)

	return n
}

func (str *Name) Type() int16 {
	return TypeName
}

func (str *Name) String() string {
	return fmt.Sprintf("Name: { Value: %v,  }", str.Value)
}
