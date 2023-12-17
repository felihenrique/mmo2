package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/packets"
)

func (c *Client) entityCreated(rawEvent events.Raw) {
	event := packets.EntityCreated{}
	events.Unserialize(rawEvent, &event)
}

func (c *Client) entityUpdated(rawEvent events.Raw) {
	event := packets.EntityUpdated{}
	events.Unserialize(rawEvent, &event)
}

func (c *Client) entityRemoved(rawEvent events.Raw) {
	event := packets.EntityRemoved{}
	events.Unserialize(rawEvent, &event)
}
