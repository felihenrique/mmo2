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

func GetEventId(data Raw) int16 {
	var requestId int16
	serialization.ReadInt16(data[2:], &requestId)
	return requestId
}

func getSize(data Raw) int16 {
	var size int16
	serialization.ReadInt16(data, &size)
	return size
}
