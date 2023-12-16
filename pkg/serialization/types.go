package serialization

type ISerializable interface {
	ToBytes() []byte
	FromBytes(data []byte) int16
	ID() int16
}
