package shard

import (
	"fmt"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/packets"
	"mmo2/pkg/serialization"
)

type Client struct {
	gspClient *gsp.TcpClient
	world     *game.World
	handlers  map[int16]EventHandler
}

func NewClient(world *game.World) *Client {
	client := Client{}
	client.gspClient = gsp.NewTcpClient()
	client.world = world
	client.handlers = make(map[int16]EventHandler)
	client.handlers[packets.TypeJoinShardResponse] = client.joinShardResponse
	client.handlers[packets.TypePlayerJoined] = client.playerJoined
	return &client
}

func (c *Client) Connect(host string, port int) error {
	err := c.gspClient.Connect(host, port)
	if err != nil {
		return err
	}
	go c.manageEvents()
	return nil
}

func (c *Client) SendRequest(event serialization.ISerializable) {
	c.gspClient.SendRequest(event)
}

func (c *Client) handleEvent(event events.Raw) {
	evType := events.GetType(event)
	handler := c.handlers[evType]
	if handler == nil {
		fmt.Printf("Handler for %d not found \n", evType)
		return
	}
	handler(event)
}

func (c *Client) manageEvents() {
	eventsChan := c.gspClient.EventsChan()
	disconChan := c.gspClient.DisconnectedChan()
main:
	for {
		select {
		case eventBytes := <-eventsChan:
			c.handleEvent(eventBytes)
		case <-disconChan:
			break main
		}
	}
	c.gspClient.Close()
}
