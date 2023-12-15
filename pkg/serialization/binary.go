package serialization

import (
	"encoding/binary"
	"fmt"
	"math"
)

func WriteBinary(buffer []byte, data any) []byte {
	switch v := data.(type) {
	case bool:
		if v {
			return append(buffer, 1)
		}
		return append(buffer, 0)
	case int8:
		return append(buffer, byte(v))
	case uint8:
		return append(buffer, byte(v))
	case int16:
		return binary.BigEndian.AppendUint16(buffer, uint16(v))
	case uint16:
		return binary.BigEndian.AppendUint16(buffer, uint16(v))
	case int32:
		return binary.BigEndian.AppendUint32(buffer, uint32(v))
	case uint32:
		return binary.BigEndian.AppendUint32(buffer, uint32(v))
	case int64:
		return binary.BigEndian.AppendUint64(buffer, uint64(v))
	case uint64:
		return binary.BigEndian.AppendUint64(buffer, uint64(v))
	case float32:
		return binary.BigEndian.AppendUint32(buffer, math.Float32bits(v))
	case float64:
		return binary.BigEndian.AppendUint64(buffer, math.Float64bits(v))
	case []bool:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			if x {
				buffer = append(buffer, 1)
			} else {
				buffer = append(buffer, 0)
			}
		}
		return buffer
	case []int8:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			buffer = append(buffer, byte(x))
		}
	case []uint8:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		return append(buffer, v...)
	case []int16:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint16(buffer, uint16(x))
		}
	case []uint16:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint16(buffer, x)
		}
	case []int32:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint32(buffer, uint32(x))
		}
	case []uint32:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint32(buffer, x)
		}
	case []int64:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint64(buffer, uint64(x))
		}
	case []uint64:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint64(buffer, x)
		}
	case []float32:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint32(buffer, math.Float32bits(x))
		}
	case []float64:
		buffer = binary.BigEndian.AppendUint16(buffer, uint16(len(v)))
		for _, x := range v {
			return binary.BigEndian.AppendUint64(buffer, math.Float64bits(x))
		}
	default:
		panic(fmt.Sprintf("wrong data type %s", v))
	}
	return buffer
}

func ReadBinary(buffer []byte, data any) int16 {
	switch data := data.(type) {
	case *bool:
		*data = buffer[0] != 0
		return 1
	case *int8:
		*data = int8(buffer[0])
		return 1
	case *uint8:
		*data = buffer[0]
		return 1
	case *int16:
		*data = int16(binary.BigEndian.Uint16(buffer))
		return 2
	case *uint16:
		*data = binary.BigEndian.Uint16(buffer)
		return 2
	case *int32:
		*data = int32(binary.BigEndian.Uint32(buffer))
		return 4
	case *uint32:
		*data = binary.BigEndian.Uint32(buffer)
		return 4
	case *int64:
		*data = int64(binary.BigEndian.Uint64(buffer))
		return 8
	case *uint64:
		*data = binary.BigEndian.Uint64(buffer)
		return 8
	case *float32:
		*data = math.Float32frombits(binary.BigEndian.Uint32(buffer))
		return 4
	case *float64:
		*data = math.Float64frombits(binary.BigEndian.Uint64(buffer))
		return 8
	case *[]bool:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]bool, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = buffer[i+2] != 0
		}
		*data = tempData
		return 2 + int16(len(tempData))
	case *[]int8:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]int8, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = int8(buffer[i+2])
		}
		*data = tempData
		return 2 + int16(len(tempData))
	case *[]uint8:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]uint8, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = uint8(buffer[i+2])
		}
		*data = tempData
		return 2 + int16(len(tempData))
	case *[]int16:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]int16, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = int16(binary.BigEndian.Uint16(buffer[2*i+2:]))
		}
		*data = tempData
		return 2 + int16(2*len(tempData))
	case *[]uint16:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]uint16, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = binary.BigEndian.Uint16(buffer[2*i+2:])
		}
		*data = tempData
		return 2 + int16(2*len(tempData))
	case *[]int32:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]int32, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = int32(binary.BigEndian.Uint32(buffer[4*i+2:]))
		}
		*data = tempData
		return 2 + int16(4*len(tempData))
	case *[]uint32:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]uint32, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = binary.BigEndian.Uint32(buffer[4*i+2:])
		}
		*data = tempData
		return 2 + int16(4*len(tempData))
	case *[]int64:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]int64, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = int64(binary.BigEndian.Uint64(buffer[8*i+2:]))
		}
		*data = tempData
		return 2 + int16(8*len(tempData))
	case *[]uint64:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]uint64, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = binary.BigEndian.Uint64(buffer[8*i+2:])
		}
		*data = tempData
		return 2 + int16(8*len(tempData))
	case *[]float32:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]float32, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = math.Float32frombits(binary.BigEndian.Uint32(buffer[4*i+2:]))
		}
		*data = tempData
		return 2 + int16(4*len(tempData))
	case *[]float64:
		size := int16(binary.BigEndian.Uint16(buffer))
		tempData := make([]float64, size)
		for i := int16(0); i < size; i++ {
			tempData[i] = math.Float64frombits(binary.BigEndian.Uint64(buffer[8*i+2:]))
		}
		*data = tempData
		return 2 + int16(8*len(tempData))
	default:
		panic(fmt.Sprintf("wrong data type %s", data))
	}
}
