package gsp

import (
	"mmo2/pkg/event_utils"
	"mmo2/pkg/serialization"
)

type PeerEvent struct {
	Peer  IPeer
	Event event_utils.Raw
}

type IServer interface {
	Listening() bool
	NewConnectionChan() <-chan IPeer
	DisconnectedChan() <-chan IPeer
	EventsChan() <-chan PeerEvent
	Listen(host string, port int) error
	Close() error
}

type IPeer interface {
	Close() error
	SendResponse(event event_utils.Raw, response serialization.ISerializable)
	SendBytes(data []byte)
	Addr() string
}
