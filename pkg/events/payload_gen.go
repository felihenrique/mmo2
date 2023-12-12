package events

import "encoding/binary"

func (str *Move) ToBytes() []byte {
	buffer := make([]byte, 4)
	var size int16 = 4
	WriteEventType(buffer, TypeMove)
	/////////
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dx))
	size += 4
	//////////
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(str.Dy))
	size += 4
	///////////
	WriteEventSize(buffer, size)
	return buffer
}

func (str *Move) FromBytes(data []byte) {
	position := 4
	///////////
	str.Dx = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
	///////////
	str.Dy = int32(binary.BigEndian.Uint32(data[position:]))
	position += 4
}
