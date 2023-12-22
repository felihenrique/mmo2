package shard

import (
	"fmt"
	"mmo2/game/packets"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
)

func (c *Client) joinShardResponse(event events.Raw) serialization.ISerializable {
	join := packets.ParseJoinShardResponse(event)
	c.world.NewEntityFrom(
		join.EntityId,
		[]serialization.ISerializable{
			join.Position,
			join.Movable,
			join.Name,
		},
	)
	return join
}

func (c *Client) playerJoined(event events.Raw) serialization.ISerializable {
	data := packets.ParsePlayerJoined(event)
	c.world.NewEntityFrom(
		data.EntityId,
		[]serialization.ISerializable{
			data.Position,
			data.Name,
			data.Movable,
		},
	)
	return data
}

func (c *Client) entityMoved(event events.Raw) serialization.ISerializable {
	data := packets.ParseEntityMoved(event)
	entity := c.world.GetEntity(data.EntityId)
	entity.Add(data.NewPosition)
	return data
}

func (c *Client) requestError(event events.Raw) serialization.ISerializable {
	data := packets.ParseRequestError(event)
	fmt.Printf("Request error: %s \n", data.Message)
	return data
}
