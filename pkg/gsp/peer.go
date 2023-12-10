package gsp

import (
	"io"
	"net"
)

type TcpPeer struct {
	conn net.Conn
	id   int32
}

func (c *TcpPeer) Close() error {
	return c.conn.Close()
}

func (c TcpPeer) Id() int32 {
	return c.id
}

func (c TcpPeer) Writer() io.Writer {
	return c.conn
}

func (c TcpPeer) Reader() io.Reader {
	return c.conn
}
