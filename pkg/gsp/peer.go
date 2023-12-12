package gsp

import (
	"mmo2/pkg/events"
	"net"
)

type TcpPeer struct {
	conn   net.Conn
	writer *events.Writer
	id     int64
}

func NewPeer(conn net.Conn, id int64) *TcpPeer {
	peer := TcpPeer{}
	peer.id = id
	peer.conn = conn
	peer.writer = events.NewWriter()
	return &peer
}

func (c *TcpPeer) Close() error {
	return c.conn.Close()
}

func (c *TcpPeer) Id() int64 {
	return c.id
}

func (c *TcpPeer) SendEvent(data []byte) {
	c.writer.Append(data)
}
