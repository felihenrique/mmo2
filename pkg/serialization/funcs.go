package serialization

import (
	"fmt"
	"mmo2/pkg/ds"
)

var IdGen ds.SequentialID

func Append(buffer []byte, data any) []byte {

	switch data := data.(type) {
	case byte:
		return append(buffer, data)
	case string:
		return AppendString(buffer, data)
	case bool:
		return AppendBool(buffer, data)
	case int8:
		return AppendInt8(buffer, data)
	case int16:
		return AppendInt16(buffer, data)
	case int32:
		return AppendInt32(buffer, data)
	case uint32:
		return AppendUint32(buffer, data)
	case float32:
		return AppendFloat32(buffer, data)
	case []byte:
		buffer = AppendInt16(buffer, int16(len(data)))
		return append(buffer, data...)
	case []bool:
		return AppendBoolSlice(buffer, data)
	case []int8:
		return AppendInt8Slice(buffer, data)
	case []int16:
		return AppendInt16Slice(buffer, data)
	case []int32:
		return AppendInt32Slice(buffer, data)
	case []float32:
		return AppendFloat32Slice(buffer, data)
	case ISerializable:
		return append(buffer, data.ToBytes(0)...)
	default:
		panic(fmt.Sprintf("wrong data type %s", data))
	}
}

func Read(buffer []byte, data any) int16 {
	switch data := data.(type) {
	case *byte:
		*data = buffer[0]
		return 1
	case *string:
		return ReadString(buffer, data)
	case *bool:
		return ReadBool(buffer, data)
	case *int8:
		return ReadInt8(buffer, data)
	case *int16:
		return ReadInt16(buffer, data)
	case *int32:
		return ReadInt32(buffer, data)
	case *uint32:
		return ReadUint32(buffer, data)
	case *float32:
		return ReadFloat32(buffer, data)
	case *[]byte:
		var size int16
		n := ReadInt16(buffer, &size)
		*data = buffer[:n]
		return n + 2
	case *[]bool:
		return ReadBoolSlice(buffer, data)
	case *[]int8:
		return ReadInt8Slice(buffer, data)
	case *[]int16:
		return ReadInt16Slice(buffer, data)
	case *[]int32:
		return ReadInt32Slice(buffer, data)
	case *[]float32:
		return ReadFloat32Slice(buffer, data)
	case ISerializable:
		return data.FromBytes(buffer)
	default:
		panic(fmt.Sprintf("wrong data type %s", data))
	}
}
