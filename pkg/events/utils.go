package events

import (
	"encoding/binary"
	"mmo2/pkg/ds"
	"mmo2/pkg/serialization"
)

type RawEvent = []byte

func GetSize(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data))
}

func GetID(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data[2:]))
}

func GetType(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data[4:]))
}

func Unserialize(data RawEvent, container serialization.ISerializable) {
	container.FromBytes(data[6:])
}

var idGen ds.SequentialID

func Serialize(event serialization.ISerializable) RawEvent {
	headers := make([]byte, 6)
	eventBytes := event.ToBytes()
	binary.BigEndian.PutUint16(headers, uint16(len(eventBytes)+len(headers)))
	binary.BigEndian.PutUint16(headers[2:], uint16(idGen.Next()))
	binary.BigEndian.PutUint16(headers[4:], uint16(event.Type()))
	return append(headers, eventBytes...)
}
