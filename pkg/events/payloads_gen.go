package events

const (
	TypeNone = int16(iota)
	TypeAck
	TypeMoveRequest
	TypeRotateRequest
	TypeJoinShardRequest
	TypeEntityCreated
	TypeEntityUpdated
)

func (str *Ack) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.EventId)

	return buffer
}

func (str *Ack) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.EventId)

	return n
}

func (str *Ack) evType() int16 {
	return TypeAck
}

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

func (str *RotateRequest) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.Quantity)

	return buffer
}

func (str *RotateRequest) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.Quantity)

	return n
}

func (str *RotateRequest) evType() int16 {
	return TypeRotateRequest
}

func (str *JoinShardRequest) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.Portal)

	return buffer
}

func (str *JoinShardRequest) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.Portal)

	return n
}

func (str *JoinShardRequest) evType() int16 {
	return TypeJoinShardRequest
}

func (str *EntityCreated) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.Data)

	return buffer
}

func (str *EntityCreated) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.Data)

	return n
}

func (str *EntityCreated) evType() int16 {
	return TypeEntityCreated
}

func (str *EntityUpdated) toBytes() []byte {
	buffer := make([]byte, 0)
	buffer = WriteBinary(buffer, str.Data)

	return buffer
}

func (str *EntityUpdated) fromBytes(data []byte) int16 {
	var n int16 = 0
	n += ReadBinary(data[n:], &str.Data)

	return n
}

func (str *EntityUpdated) evType() int16 {
	return TypeEntityUpdated
}
