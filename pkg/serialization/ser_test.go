package serialization

import "testing"

func TestFloat32(t *testing.T) {
	var var1 float32 = 1.1112
	buffer := AppendFloat32([]byte{}, var1)
	if len(buffer) != 4 {
		panic("wrong")
	}
	var var2 float32 = 0
	n := ReadFloat32(buffer, &var2)
	if n != 4 || var2 != var1 {
		panic("wrong")
	}
}

func TestBool(t *testing.T) {
	var var1 bool = false
	buffer := AppendBool([]byte{}, var1)
	if len(buffer) != 1 {
		panic("wrong")
	}
	var var2 bool = true
	n := ReadBool(buffer, &var2)
	if n != 1 || var2 != var1 {
		panic("wrong")
	}
}

func TestInt8(t *testing.T) {
	var var1 int8 = 13
	buffer := AppendInt8([]byte{}, var1)
	if len(buffer) != 1 {
		panic("wrong")
	}
	var var2 int8 = 0
	n := ReadInt8(buffer, &var2)
	if n != 1 || var2 != var1 {
		panic("wrong")
	}
}

func TestInt16(t *testing.T) {
	var var1 int16 = 13
	buffer := AppendInt16([]byte{}, var1)
	if len(buffer) != 2 {
		panic("wrong")
	}
	var var2 int16 = 0
	n := ReadInt16(buffer, &var2)
	if n != 2 || var2 != var1 {
		panic("wrong")
	}
}

func TestInt32(t *testing.T) {
	var var1 int32 = 131
	buffer := AppendInt32([]byte{}, var1)
	if len(buffer) != 4 {
		panic("wrong")
	}
	var var2 int32 = 0
	n := ReadInt32(buffer, &var2)
	if n != 4 || var2 != var1 {
		panic("wrong")
	}
}

func TestString(t *testing.T) {
	var var1 string = "nossa, nossa"
	buffer := AppendString([]byte{}, var1)
	if len(buffer) != 14 {
		panic("wrong")
	}
	var var2 string = ""
	n := ReadString(buffer, &var2)
	if n != 14 || var2 != var1 {
		panic("wrong")
	}
}

func TestBoolSlice(t *testing.T) {
	var1 := []bool{true, false, false}
	buffer := AppendBoolSlice([]byte{}, var1)
	if len(buffer) != 5 {
		panic("wrong")
	}
	var var2 []bool
	n := ReadBoolSlice(buffer, &var2)
	if n != 5 || var2[0] != true || var2[1] != false || var2[2] != false || len(var2) != 3 {
		panic("wrong")
	}
}

func TestInt8Slice(t *testing.T) {
	var1 := []int8{12, 11, 78}
	buffer := AppendInt8Slice([]byte{}, var1)
	if len(buffer) != 5 {
		panic("wrong")
	}
	var var2 []int8
	n := ReadInt8Slice(buffer, &var2)
	if n != 5 || var2[0] != 12 || var2[1] != 11 || var2[2] != 78 || len(var2) != 3 {
		panic("wrong")
	}
}

func TestInt16Slice(t *testing.T) {
	var1 := []int16{12, 11, 78}
	buffer := AppendInt16Slice([]byte{}, var1)
	if len(buffer) != 8 {
		panic("wrong")
	}
	var var2 []int16
	n := ReadInt16Slice(buffer, &var2)
	if n != 8 || var2[0] != 12 || var2[1] != 11 || var2[2] != 78 || len(var2) != 3 {
		panic("wrong")
	}
}

func TestInt32Slice(t *testing.T) {
	var1 := []int32{12, 11, 78}
	buffer := AppendInt32Slice([]byte{}, var1)
	if len(buffer) != 14 {
		panic("wrong")
	}
	var var2 []int32
	n := ReadInt32Slice(buffer, &var2)
	if n != 14 || var2[0] != 12 || var2[1] != 11 || var2[2] != 78 || len(var2) != 3 {
		panic("wrong")
	}
}

func TestFloat32Slice(t *testing.T) {
	var1 := []float32{12.0, 11.1, 78.123}
	buffer := AppendFloat32Slice([]byte{}, var1)
	if len(buffer) != 14 {
		panic("wrong")
	}
	var var2 []float32
	n := ReadFloat32Slice(buffer, &var2)
	if n != 14 || var2[0] != 12.0 || var2[1] != 11.1 || var2[2] != 78.123 || len(var2) != 3 {
		panic("wrong")
	}
}
