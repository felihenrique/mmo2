package gsp

import (
	"io"
	"net"
)

type TcpPeer struct {
	conn net.Conn
	id   int64
}

func NewPeer(conn net.Conn, id int64) *TcpPeer {
	peer := TcpPeer{}
	peer.id = id
	peer.conn = conn
	return &peer
}

func (c *TcpPeer) Close() error {
	return c.conn.Close()
}

func (c *TcpPeer) Id() int64 {
	return c.id
}

func (c *TcpPeer) Writer() io.Writer {
	return c.conn
}

func (c *TcpPeer) Reader() io.Reader {
	return c.conn
}
