package gsp

import (
	"mmo2/pkg/events"
	"net"
)

type TcpPeer struct {
	conn   net.Conn
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

func (c *TcpPeer) SendEvent(data []byte) {
	c.writer.Append(data)
}

func (c *TcpPeer) Addr() string {
	return c.addr
}
