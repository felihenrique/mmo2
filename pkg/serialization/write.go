package serialization

import (
	"encoding/binary"
	"math"
)

func AppendString(buffer []byte, data string) []byte {
	temp := binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	return append(temp, []byte(data)...)
}

func AppendBool(buffer []byte, data bool) []byte {
	if data {
		return append(buffer, 1)
	}
	return append(buffer, 0)
}

func AppendInt8(buffer []byte, data int8) []byte {
	return append(buffer, byte(data))
}

func AppendInt16(buffer []byte, data int16) []byte {
	return binary.BigEndian.AppendUint16(buffer, uint16(data))
}

func AppendInt32(buffer []byte, data int32) []byte {
	return binary.BigEndian.AppendUint32(buffer, uint32(data))
}

func AppendUint32(buffer []byte, data uint32) []byte {
	return binary.BigEndian.AppendUint32(buffer, data)
}

func AppendFloat32(buffer []byte, data float32) []byte {
	return binary.BigEndian.AppendUint32(buffer, math.Float32bits(data))
}

func AppendBoolSlice(buffer []byte, data []bool) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	for _, x := range data {
		if x {
			buffer = append(buffer, 1)
		} else {
			buffer = append(buffer, 0)
		}
	}
	return buffer
}

func AppendInt8Slice(buffer []byte, data []int8) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	for _, x := range data {
		buffer = append(buffer, byte(x))
	}
	return buffer
}

func AppendInt16Slice(buffer []byte, data []int16) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	for _, x := range data {
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(x))
	}
	return buffer
}

func AppendInt32Slice(buffer []byte, data []int32) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	for _, x := range data {
		buffer = binary.BigEndian.AppendUint32(buffer, uint32(x))
	}
	return buffer
}

func AppendFloat32Slice(buffer []byte, data []float32) []byte {
	buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(data)))
	for _, x := range data {
		buffer = binary.BigEndian.AppendUint32(buffer, math.Float32bits(x))
	}
	return buffer
}
