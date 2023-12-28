package ecs

import (
	"mmo2/pkg/serialization"
)

type unmarshal = func(data []byte) serialization.ISerializable

var Mapper = []unmarshal{
	func(data []byte) serialization.ISerializable {
		return ParseTransform(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParsePlayer(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseLiving(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseColor(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseCircle(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseMove(data)
	},
}
