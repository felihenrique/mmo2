package gsp

import (
	"fmt"
	"mmo2/pkg/errors"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
	"net"
	"time"
)

type TcpClient struct {
	conn             net.Conn
	eventsChan       chan event_utils.Raw
	disconnectedChan chan byte
	writer           *events.Writer
	reader           *events.Reader
	connected        bool
}

func NewTcpClient() *TcpClient {
	client := TcpClient{}
	client.writer = events.NewWriter()
	client.reader = events.NewReader()
	client.disconnectedChan = make(chan byte)
	client.eventsChan = make(chan []byte, 256)
	return &client
}

func (c *TcpClient) EventsChan() <-chan event_utils.Raw {
	return c.eventsChan
}

func (c *TcpClient) DisconnectedChan() <-chan byte {
	return c.disconnectedChan
}

func (c *TcpClient) Connect(host string, port int) error {
	dialer := net.Dialer{}
	conn, err := dialer.Dial("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	c.conn = conn
	c.connected = true
	go c.readEvents()
	return nil
}

func (c *TcpClient) SendBytes(data []byte) {
	c.writer.Append(data)
}

func (c *TcpClient) SendRequest(event serialization.ISerializable) int16 {
	id := serialization.IdGen.Next()
	eventBytes := event.ToBytes(id)
	c.writer.Append(eventBytes)
	return id
}

func (c *TcpClient) SendResponse(event event_utils.Raw, response serialization.ISerializable) {
	eventId := event_utils.GetEventId(event)
	bytes := response.ToBytes(eventId)
	c.writer.Append(bytes)
}

func (c *TcpClient) Close() {
	c.conn.Close()
}

func (c *TcpClient) handleDisconn(err error) {
	println(err.Error())
	c.disconnectedChan <- 1
	close(c.disconnectedChan)
	close(c.eventsChan)
}

func (c *TcpClient) readEvents() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("error in connection %s: \n", r)
		}
		c.connected = false
		c.conn.Close()
		c.disconnectedChan <- 1
	}()

	for c.connected {
		// READ
		c.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
		err := errors.Handle(c.reader.FillFrom(c.conn))
		if err != nil && !errors.Is(err, errors.ErrTimeout) {
			c.handleDisconn(err)
			break
		}
		for {
			rawEvent, err := c.reader.Next()
			if err != nil {
				break
			}
			c.eventsChan <- rawEvent
			c.reader.Pop()
		}
		// WRITE
		c.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 200))
		_, err = c.writer.Send(c.conn)
		err = errors.Handle(err)
		if err != nil && !errors.Is(err, errors.ErrTimeout) {
			c.handleDisconn(err)
			break
		}
	}
}
