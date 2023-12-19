package serialization

type RawStruct = []byte

type ISerializable interface {
	ToBytes(eventId int16) []byte
	FromBytes(data []byte) int16
	Type() int16
	String() string
}
