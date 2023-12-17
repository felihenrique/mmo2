package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/payloads"
)

func (c *Client) entityCreated(rawEvent events.RawEvent) {
	event := payloads.EntityCreated{}
	events.Unserialize(rawEvent, &event)
}

func (c *Client) entityUpdated(rawEvent events.RawEvent) {
	event := payloads.EntityUpdated{}
	events.Unserialize(rawEvent, &event)
}

func (c *Client) entityRemoved(rawEvent events.RawEvent) {
	event := payloads.EntityRemoved{}
	events.Unserialize(rawEvent, &event)
}
