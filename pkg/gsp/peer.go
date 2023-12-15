package gsp

import (
	"mmo2/pkg/ds"
	"mmo2/pkg/events"
	"net"
)

type TcpPeer struct {
	conn   net.Conn
	idGen  ds.SequentialID
	writer *events.Writer
	addr   string
}

func NewPeer(conn net.Conn) *TcpPeer {
	peer := TcpPeer{}
	peer.conn = conn
	peer.writer = events.NewWriter()
	peer.addr = conn.RemoteAddr().String()
	return &peer
}

func (c *TcpPeer) Close() error {
	return c.conn.Close()
}

func (c *TcpPeer) SendEvent(event events.ISerializable) {
	packet := events.Serialize(event, c.idGen.Next())
	c.writer.Append(packet)
}

func (c *TcpPeer) AckEvent(eventId int16) {
	packet := events.Serialize(&events.Ack{EventId: eventId}, c.idGen.Next())
	c.writer.Append(packet)
}

func (c *TcpPeer) Addr() string {
	return c.addr
}
