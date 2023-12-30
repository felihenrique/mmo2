package shard

import (
	"fmt"
	"mmo2/game/packets"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/gsp"
	"mmo2/pkg/serialization"
	"time"
)

type Client struct {
	gspClient       *gsp.TcpClient
	handlers        map[int16]EventHandler
	callbacks       map[int16]ResponseHandler
	disconnChan     chan byte
	tickChan        chan byte
	tickChanEnabled bool
}

func NewClient() *Client {
	client := Client{}
	client.gspClient = gsp.NewTcpClient()
	client.disconnChan = make(chan byte)
	client.tickChan = make(chan byte)
	client.handlers = make(map[int16]EventHandler)
	client.callbacks = make(map[int16]ResponseHandler)
	client.handlers[packets.TypePlayerJoined] = client.playerJoined
	client.handlers[packets.TypeEntityMoved] = client.entityMoved
	client.handlers[packets.TypeRequestError] = client.requestError
	client.handlers[packets.TypeAckRequest] = client.ackRequest
	client.handlers[packets.TypeJoinShardResponse] = client.joinShardResponse
	client.handlers[packets.TypeEntityRemoved] = client.entityRemoved
	return &client
}

func (c *Client) EnableTickChan() {
	c.tickChanEnabled = true
}

func (c *Client) Connect(host string, port int) error {
	err := c.gspClient.Connect(host, port)
	if err != nil {
		return err
	}
	go c.manageEvents()
	return nil
}

func (c *Client) TickChan() chan byte {
	return c.tickChan
}

func (c *Client) SendRequest(event serialization.ISerializable, callback ResponseHandler) {
	id := c.gspClient.SendRequest(event)
	c.callbacks[id] = callback
}

func (c *Client) handleEvent(event event_utils.Raw) {
	evType := event_utils.GetType(event)
	handler := c.handlers[evType]
	if handler == nil {
		fmt.Printf("Handler for %d not found \n", evType)
		return
	}
	response := handler(event)
	id := event_utils.GetEventId(event)
	if id == 0 { // SERVER EVENT
		return
	}
	callback := c.callbacks[id]
	if callback == nil {
		return
	}
	callback(response)
	delete(c.callbacks, id)
}

func (c *Client) manageEvents() {
	eventsChan := c.gspClient.EventsChan()
	disconChan := c.gspClient.DisconnectedChan()
	ticker := time.NewTicker(time.Second / 60)
main:
	for {
		select {
		case eventBytes := <-eventsChan:
			c.handleEvent(eventBytes)
		case <-disconChan:
			c.disconnChan <- 1
			break main
		case <-ticker.C:
			if c.tickChanEnabled {
				c.tickChan <- 1
				<-c.tickChan
			}
		case event := <-SendEventsChan:
			c.SendRequest(event, nil)
		}
	}
	c.gspClient.Close()
}
