package gsp

import (
	"fmt"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
	"net"
	"time"
)

type TcpPeer struct {
	conn              net.Conn
	writer            *events.Writer
	reader            *events.Reader
	connected         bool
	rateLimit         int
	lastCheck         time.Time
	lastSecondPackets int
	addr              string
}

func NewPeer(conn net.Conn) *TcpPeer {
	peer := TcpPeer{}
	peer.conn = conn
	peer.writer = events.NewWriter()
	peer.reader = events.NewReader()
	peer.addr = conn.RemoteAddr().String()
	peer.rateLimit = 30
	peer.lastCheck = time.Now()
	peer.lastSecondPackets = 0
	peer.connected = true
	return &peer
}

func (c *TcpPeer) updateRateLimit() {
	c.lastSecondPackets += 1
	if time.Since(c.lastCheck).Seconds() >= 1 {
		fmt.Printf("Peer %s current rate: %d \n", c.addr, c.lastSecondPackets)
		if c.lastSecondPackets >= c.rateLimit {
			fmt.Printf("Rate limit for %s reached. Closing connection \n", c.addr)
			c.connected = false
		}
		c.lastCheck = time.Now()
		c.lastSecondPackets = 0
	}
}

func (c *TcpPeer) readEvents() error {
	c.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
	return c.reader.FillFrom(c.conn)
}

func (c *TcpPeer) writeEvents() error {
	c.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 10))
	_, err := c.writer.Send(c.conn)
	return err
}

func (c *TcpPeer) SetRateLimit(limit int) {
	c.rateLimit = limit
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
