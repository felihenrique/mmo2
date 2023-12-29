package gsp

import (
	"mmo2/pkg/event_utils"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
	"net"
)

type IPeer interface {
	Close() error
	SendResponse(event event_utils.Raw, response serialization.ISerializable)
	SendBytes(data []byte)
	Addr() string
}

type TcpPeer struct {
	conn      net.Conn
	writer    *events.Writer
	reader    *events.Reader
	connected bool
	addr      string
}

func NewPeer(conn net.Conn) *TcpPeer {
	peer := TcpPeer{}
	peer.conn = conn
	peer.writer = events.NewWriter()
	peer.reader = events.NewReader()
	peer.addr = conn.RemoteAddr().String()
	peer.connected = true
	return &peer
}

func (c *TcpPeer) Close() error {
	c.connected = false
	return c.conn.Close()
}

func (c *TcpPeer) SendResponse(event event_utils.Raw, response serialization.ISerializable) {
	eventId := event_utils.GetEventId(event)
	bytes := response.ToBytes(eventId)
	c.SendBytes(bytes)
}

func (c *TcpPeer) SendBytes(data []byte) {
	c.writer.Append(data)
}

func (c *TcpPeer) Addr() string {
	return c.addr
}
