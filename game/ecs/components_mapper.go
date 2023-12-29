package ecs

import (
	"mmo2/pkg/serialization"
)

type unmarshal = func(data []byte) (serialization.ISerializable, int16)

var Mapper = []unmarshal{
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseTransform(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParsePlayer(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseLiving(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseColor(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseCircle(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseMoveTo(data)
	},
}
