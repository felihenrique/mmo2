package serialization

import (
	"encoding/binary"
	"math"
)

func ReadString(buffer []byte, data *string) int16 {
	strSize := int16(binary.BigEndian.Uint16(buffer))
	*data = string(buffer[2 : strSize+2])
	return int16(2 + strSize)
}

func ReadBool(buffer []byte, data *bool) int16 {
	*data = buffer[0] != 0
	return 1
}

func ReadInt8(buffer []byte, data *int8) int16 {
	*data = int8(buffer[0])
	return 1
}

func ReadInt16(buffer []byte, data *int16) int16 {
	*data = int16(binary.BigEndian.Uint16(buffer))
	return 2
}

func ReadInt32(buffer []byte, data *int32) int16 {
	*data = int32(binary.BigEndian.Uint32(buffer))
	return 4
}

func ReadFloat32(buffer []byte, data *float32) int16 {
	*data = math.Float32frombits(binary.BigEndian.Uint32(buffer))
	return 4
}

func ReadBoolSlice(buffer []byte, data *[]bool) int16 {
	size := int16(binary.BigEndian.Uint16(buffer))
	tempData := make([]bool, size)
	for i := int16(0); i < size; i++ {
		tempData[i] = buffer[i+2] != 0
	}
	*data = tempData
	return 2 + int16(len(tempData))
}

func ReadInt8Slice(buffer []byte, data *[]int8) int16 {
	size := int16(binary.BigEndian.Uint16(buffer))
	tempData := make([]int8, size)
	for i := int16(0); i < size; i++ {
		tempData[i] = int8(buffer[i+2])
	}
	*data = tempData
	return 2 + int16(len(tempData))
}

func ReadInt16Slice(buffer []byte, data *[]int16) int16 {
	size := int16(binary.BigEndian.Uint16(buffer))
	tempData := make([]int16, size)
	for i := int16(0); i < size; i++ {
		tempData[i] = int16(binary.BigEndian.Uint16(buffer[2*i+2:]))
	}
	*data = tempData
	return 2 + int16(2*len(tempData))
}

func ReadInt32Slice(buffer []byte, data *[]int32) int16 {
	size := int16(binary.BigEndian.Uint16(buffer))
	tempData := make([]int32, size)
	for i := int16(0); i < size; i++ {
		tempData[i] = int32(binary.BigEndian.Uint32(buffer[4*i+2:]))
	}
	*data = tempData
	return 2 + int16(4*len(tempData))
}

func ReadFloat32Slice(buffer []byte, data *[]float32) int16 {
	size := int16(binary.BigEndian.Uint16(buffer))
	tempData := make([]float32, size)
	for i := int16(0); i < size; i++ {
		tempData[i] = math.Float32frombits(binary.BigEndian.Uint32(buffer[4*i+2:]))
	}
	*data = tempData
	return 2 + int16(4*len(tempData))
}
