package events

const (
	TypeAck = int16(iota + 1)
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeEntityCreated
	TypeEntityUpdated
)

func (str *MoveRequest) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.Dx)
	buffer = WriteBinary(buffer, str.Dy)
	return buffer
}

func (str *MoveRequest) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.Dx)
	n += ReadBinary(data[n:], &str.Dy)
	return n
}

func (str *MoveRequest) evType() int16 {
	return TypeMoveRequest
}
