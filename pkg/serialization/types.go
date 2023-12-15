package serialization

type ISerializable interface {
	ToBytes() []byte
	FromBytes(data []byte) int16
	EvType() int16
}
