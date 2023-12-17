package shard

import (
	"fmt"
	"mmo2/pkg/events"
	"mmo2/pkg/game"
	"mmo2/pkg/packets"
)

func (c *Client) entityCreated(rawEvent events.Raw) {
	event := packets.EntityCreated{}
	event.FromBytes(rawEvent)
	entity := c.world.NewEntity()
	entity.FromBytes(event.Entity)
}

func (c *Client) entityUpdated(rawEvent events.Raw) {
	event := packets.EntityUpdated{}
	event.FromBytes(rawEvent)
	entity := c.world.GetEntity(event.EntityId)
	if entity == nil {
		fmt.Printf("WRONG situation: updating entity with id %d", event.EntityId)
		return
	}
	for n := 0; n < len(event.Components); {
		component, readed := game.Read(event.Components)
		entity.Add(component)
		n += int(readed)
	}
}

func (c *Client) entityRemoved(rawEvent events.Raw) {
	event := packets.EntityRemoved{}
	event.FromBytes(rawEvent)
	c.world.RemoveEntity(event.EntityId)
}
