package events

import (
	"mmo2/pkg/serialization"
)

type Raw = []byte

func GetType(data Raw) int16 {
	var evType int16
	serialization.ReadInt16(data, &evType)
	return evType
}
