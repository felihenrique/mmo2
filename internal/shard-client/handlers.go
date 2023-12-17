package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/packets"
)

func (c *Client) entityCreated(rawEvent events.Raw) {
	event := packets.EntityCreated{}
	event.FromBytes(rawEvent)
}

func (c *Client) entityUpdated(rawEvent events.Raw) {
	event := packets.EntityUpdated{}
	event.FromBytes(rawEvent)
}

func (c *Client) entityRemoved(rawEvent events.Raw) {
	event := packets.EntityRemoved{}
	event.FromBytes(rawEvent)
}
