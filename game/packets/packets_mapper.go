package packets

import (
	"mmo2/pkg/serialization"
)

type unmarshal = func(data []byte) (serialization.ISerializable, int16)

var Mapper = []unmarshal{
	func(data []byte) (serialization.ISerializable, int16) {
		return ParsePing(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseAckRequest(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseRequestError(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseMoveRequest(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseJoinShardRequest(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseJoinShardResponse(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParsePlayerJoined(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseEntityMoved(data)
	},
	func(data []byte) (serialization.ISerializable, int16) {
		return ParseEntityRemoved(data)
	},
}
