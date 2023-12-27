package ecs

import (
	"fmt"
	"mmo2/pkg/serialization"
)

const (
	TypeNone = int16(iota)
	TypeTransform
	TypeLiving
	TypeColor
	TypeCircle
	TypeMoveTo
)

func NewTransform(X int32, Y int32, Rotation float32) *Transform {
	return &Transform{
		X:        X,
		Y:        Y,
		Rotation: Rotation,
	}
}

func ParseTransform(event []byte) *Transform {
	str := Transform{}
	str.FromBytes(event)
	return &str
}

func (str *Transform) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeTransform)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.X)
	buffer = serialization.Append(buffer, str.Y)
	buffer = serialization.Append(buffer, str.Rotation)

	return buffer
}

func (str *Transform) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.X)

	n += serialization.Read(data[n:], &str.Y)

	n += serialization.Read(data[n:], &str.Rotation)

	return n
}

func (str *Transform) Type() int16 {
	return TypeTransform
}

func (str *Transform) String() string {
	return fmt.Sprintf("Transform: { X: %v, Y: %v, Rotation: %v,  }", str.X, str.Y, str.Rotation)
}

func NewLiving(Name string, Velocity float32) *Living {
	return &Living{
		Name:     Name,
		Velocity: Velocity,
	}
}

func ParseLiving(event []byte) *Living {
	str := Living{}
	str.FromBytes(event)
	return &str
}

func (str *Living) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeLiving)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Name)
	buffer = serialization.Append(buffer, str.Velocity)

	return buffer
}

func (str *Living) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Name)

	n += serialization.Read(data[n:], &str.Velocity)

	return n
}

func (str *Living) Type() int16 {
	return TypeLiving
}

func (str *Living) String() string {
	return fmt.Sprintf("Living: { Name: %v, Velocity: %v,  }", str.Name, str.Velocity)
}

func NewColor(R uint8, G uint8, B uint8, A uint8) *Color {
	return &Color{
		R: R,
		G: G,
		B: B,
		A: A,
	}
}

func ParseColor(event []byte) *Color {
	str := Color{}
	str.FromBytes(event)
	return &str
}

func (str *Color) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeColor)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.R)
	buffer = serialization.Append(buffer, str.G)
	buffer = serialization.Append(buffer, str.B)
	buffer = serialization.Append(buffer, str.A)

	return buffer
}

func (str *Color) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.R)

	n += serialization.Read(data[n:], &str.G)

	n += serialization.Read(data[n:], &str.B)

	n += serialization.Read(data[n:], &str.A)

	return n
}

func (str *Color) Type() int16 {
	return TypeColor
}

func (str *Color) String() string {
	return fmt.Sprintf("Color: { R: %v, G: %v, B: %v, A: %v,  }", str.R, str.G, str.B, str.A)
}

func NewCircle(Radius float32, Color *Color) *Circle {
	return &Circle{
		Radius: Radius,
		Color:  Color,
	}
}

func ParseCircle(event []byte) *Circle {
	str := Circle{}
	str.FromBytes(event)
	return &str
}

func (str *Circle) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeCircle)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.Radius)
	buffer = serialization.Append(buffer, str.Color)

	return buffer
}

func (str *Circle) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.Radius)

	str.Color = &Color{}
	n += serialization.Read(data[n:], str.Color)

	return n
}

func (str *Circle) Type() int16 {
	return TypeCircle
}

func (str *Circle) String() string {
	return fmt.Sprintf("Circle: { Radius: %v, Color: %v,  }", str.Radius, str.Color)
}

func NewMoveTo(X int32, Y int32) *MoveTo {
	return &MoveTo{
		X: X,
		Y: Y,
	}
}

func ParseMoveTo(event []byte) *MoveTo {
	str := MoveTo{}
	str.FromBytes(event)
	return &str
}

func (str *MoveTo) ToBytes(eventId int16) []byte {
	buffer := make([]byte, 0)
	buffer = serialization.Append(buffer, TypeMoveTo)
	buffer = serialization.Append(buffer, eventId)
	buffer = serialization.Append(buffer, str.X)
	buffer = serialization.Append(buffer, str.Y)

	return buffer
}

func (str *MoveTo) FromBytes(data []byte) int16 {
	var n int16 = 4
	n += serialization.Read(data[n:], &str.X)

	n += serialization.Read(data[n:], &str.Y)

	return n
}

func (str *MoveTo) Type() int16 {
	return TypeMoveTo
}

func (str *MoveTo) String() string {
	return fmt.Sprintf("MoveTo: { X: %v, Y: %v,  }", str.X, str.Y)
}
