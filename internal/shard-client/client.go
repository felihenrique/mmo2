package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
	"mmo2/pkg/payloads"
)

type Client struct {
	gspClient *gsp.TcpClient
	world     *game.World
}

func NewClient(world *game.World) *Client {
	client := Client{}
	client.gspClient = gsp.NewTcpClient()
	client.world = world
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

func (c *Client) manageEvents() {
	eventsChan := c.gspClient.EventsChan()
	disconChan := c.gspClient.DisconnectedChan()
main:
	for {
		select {
		case eventBytes := <-eventsChan:
			evType := events.GetType(eventBytes)
			switch evType {
			case payloads.TypeEntityCreated:
				c.entityCreated(eventBytes)
			case payloads.TypeEntityUpdated:
				c.entityUpdated(eventBytes)
			case payloads.TypeEntityRemoved:
				c.entityRemoved(eventBytes)
			}
		case <-disconChan:
			break main
		}
	}
	c.gspClient.Close()
}
