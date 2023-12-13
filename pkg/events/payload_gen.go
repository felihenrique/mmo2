package events

import "encoding/binary"

func (str *Move) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dx))
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dy))
	return buffer
}

func (str *Move) fromBytes(data []byte) {
	position := 0
	str.Dx = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	str.Dy = int32(binary.BigEndian.Uint32(data[position:]))
}

func (str *Move) evType() int16 {
	return TypeMove
}
