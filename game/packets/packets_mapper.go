package packets

import (
	"mmo2/pkg/serialization"
)

type unmarshal = func(data []byte) serialization.ISerializable

var Mapper = []unmarshal{
	func(data []byte) serialization.ISerializable {
		return ParseAckRequest(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseRequestError(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseMoveRequest(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseJoinShardRequest(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseJoinShardResponse(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParsePlayerJoined(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseEntityMoved(data)
	},
	func(data []byte) serialization.ISerializable {
		return ParseEntityRemoved(data)
	},
}
