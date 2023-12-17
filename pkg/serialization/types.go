package serialization

type RawStruct = []byte

type ISerializable interface {
	ToBytes() []byte
	FromBytes(data []byte) int16
	Type() int16
}
