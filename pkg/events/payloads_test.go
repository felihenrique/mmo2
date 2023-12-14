package events

import "testing"

func TestEntityCreated(t *testing.T) {
	data := EntityCreated{
		Data: []byte{10, 167, 45},
	}
	bytes := data.toBytes()
	data2 := EntityCreated{}
	data2.fromBytes(bytes)
	if data2.Data[0] != 10 || data2.Data[1] != 167 || data2.Data[2] != 45 {
		panic("wrong")
	}
}
