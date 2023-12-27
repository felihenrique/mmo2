package shard

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
)

func (c *Client) joinShardResponse(event events.Raw) serialization.ISerializable {
	join := packets.ParseJoinShardResponse(event)
	ecs.MainWorld.NewEntityFrom(
		join.EntityId,
		[]ecs.IComponent{
			join.Transform,
			join.Living,
			join.PlayerCircle,
		},
	)
	return join
}

func (c *Client) playerJoined(event events.Raw) serialization.ISerializable {
	data := packets.ParsePlayerJoined(event)
	ecs.MainWorld.NewEntityFrom(
		data.EntityId,
		[]ecs.IComponent{
			data.Transform,
			data.Living,
			data.PlayerCircle,
		},
	)
	return data
}

func (c *Client) entityMoved(event events.Raw) serialization.ISerializable {
	data := packets.ParseEntityMoved(event)
	entity := ecs.MainWorld.GetEntity(data.EntityId)
	entity.Add(data.NewPosition)
	return data
}

func (c *Client) requestError(event events.Raw) serialization.ISerializable {
	data := packets.ParseRequestError(event)
	fmt.Printf("Request error: %s \n", data.Message)
	return data
}
