package events

import "encoding/binary"

type IEvent interface {
	ToBytes() []byte
	FromBytes(data []byte)
}

func GetEventSize(data []byte) int16 {
	return int16(binary.BigEndian.Uint16(data))
}

func GetEventType(data []byte) int16 {
	return int16(binary.BigEndian.Uint16(data[2:]))
}

func GetEventId(data []byte) int16 {
	return int16(binary.BigEndian.Uint16(data[4:]))
}

func GetEventPayload(data []byte) []byte {
	return data[6:]
}

func WriteEventSize(data []byte, size int16) {
	binary.BigEndian.PutUint16(data, uint16(size))
}

func WriteEventType(data []byte, eventType int16) {
	binary.BigEndian.PutUint16(data[2:], uint16(eventType))
}

func WriteEventId(data []byte, eventId int16) {
	binary.BigEndian.PutUint16(data[4:], uint16(eventId))
}
