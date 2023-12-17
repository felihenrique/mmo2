package serialization

type ISerializable interface {
	ToBytes() []byte
	FromBytes(data []byte) int16
	Type() int16
}
