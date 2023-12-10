package events

type IEventPacket interface {
	GetType() uint16
}

const (
	TypeMove = uint16(iota + 1)
)
