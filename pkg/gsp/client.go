package gsp

import (
	"errors"
	"fmt"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
	"net"
	"time"
)

type TcpClient struct {
	conn             net.Conn
	eventsChan       chan events.Raw
	disconnectedChan chan byte
	writer           *events.Writer
	connected        bool
}

func NewTcpClient() *TcpClient {
	client := TcpClient{}
	client.writer = events.NewWriter()
	client.disconnectedChan = make(chan byte)
	client.eventsChan = make(chan []byte)
	return &client
}

func (c *TcpClient) EventsChan() <-chan events.Raw {
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

func (c *TcpClient) SendRequest(event serialization.ISerializable) {
	eventBytes := event.ToBytes(serialization.IdGen.Next())
	c.writer.Append(eventBytes)
}

func (c *TcpClient) SendResponse(event events.Raw, response serialization.ISerializable) {
	eventId := events.GetEventId(event)
	bytes := response.ToBytes(eventId)
	c.writer.Append(bytes)
}

func (c *TcpClient) Close() {
	c.conn.Close()
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
	reader := events.NewReader()

	for c.connected {
		// READ
		c.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
		err := handleError(reader.FillFrom(c.conn))
		if err != nil && !errors.Is(err, ErrTimeout) {
			println(err.Error())
			break
		}
		for {
			rawEvent, err := reader.Next()
			if err != nil {
				break
			}
			c.eventsChan <- rawEvent
			reader.Pop()
		}
		// WRITE
		c.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 200))
		_, err = c.writer.Send(c.conn)
		err = handleError(err)
		if err != nil && !errors.Is(err, ErrTimeout) {
			println(err.Error())
			break
		}
	}
}
