package gsp

import (
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
	"net"
)

type IPeer interface {
	Close() error
	SendEvent(event serialization.ISerializable)
	SendResponse(event events.Raw, response serialization.ISerializable)
	SendBytes(data []byte)
	Addr() string
}

type TcpPeer struct {
	conn   net.Conn
	writer *events.Writer
	reader *events.Reader
	addr   string
}

func NewPeer(conn net.Conn) *TcpPeer {
	peer := TcpPeer{}
	peer.conn = conn
	peer.writer = events.NewWriter()
	peer.reader = events.NewReader()
	peer.addr = conn.RemoteAddr().String()
	return &peer
}

func (c *TcpPeer) Close() error {
	return c.conn.Close()
}

func (c *TcpPeer) SendEvent(event serialization.ISerializable) {
	eventBytes := event.ToBytes(serialization.IdGen.Next())
	c.writer.Append(eventBytes)
}

func (c *TcpPeer) SendResponse(event events.Raw, response serialization.ISerializable) {
	eventId := events.GetEventId(event)
	bytes := response.ToBytes(eventId)
	c.SendBytes(bytes)
}

func (c *TcpPeer) SendBytes(data []byte) {
	c.writer.Append(data)
}

func (c *TcpPeer) Addr() string {
	return c.addr
}
