package payloads

import "testing"

func TestEntityCreated(t *testing.T) {
	data := EntityCreated{
		Entity: []byte{10, 167, 45},
	}
	bytes := data.ToBytes()
	data2 := EntityCreated{}
	data2.FromBytes(bytes)
	if data2.Entity[0] != 10 || data2.Entity[1] != 167 || data2.Entity[2] != 45 {
		panic("wrong")
	}
}
