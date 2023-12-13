package events

import "encoding/binary"

type RawEvent = []byte

type ISerializable interface {
	toBytes() []byte
	fromBytes(data []byte)
	evType() int16
}

func GetSize(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data))
}

func GetId(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data[2:]))
}

func GetType(data RawEvent) int16 {
	return int16(binary.BigEndian.Uint16(data[4:]))
}

func Unserialize[T ISerializable](data RawEvent, container T) {
	container.fromBytes(data[6:])
}

func Serialize(event ISerializable, id int16) RawEvent {
	headers := make([]byte, 6)
	eventBytes := event.toBytes()
	binary.BigEndian.PutUint16(headers, uint16(len(eventBytes)+len(headers)))
	binary.BigEndian.PutUint16(headers[2:], uint16(id))
	binary.BigEndian.PutUint16(headers[4:], uint16(event.evType()))
	return append(headers, eventBytes...)
}
