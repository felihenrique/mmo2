package shard

import (
	"mmo2/pkg/game"
	"mmo2/pkg/gsp"
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

		case <-disconChan:
			break main
		}
	}
	c.gspClient.Close()
}
