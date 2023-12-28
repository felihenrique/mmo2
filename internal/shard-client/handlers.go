package shard

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/events"
	"mmo2/pkg/serialization"
)

func (c *Client) ackRequest(event events.Raw) serialization.ISerializable {
	return packets.ParseAckRequest(event)
}

func (c *Client) playerJoined(event events.Raw) serialization.ISerializable {
	data := packets.ParsePlayerJoined(event)
	ecs.MainWorld.NewEntityFrom(
		data.EntityId,
		data.Transform,
		data.Living,
		data.PlayerCircle,
	)
	return data
}

func (c *Client) entityMoved(event events.Raw) serialization.ISerializable {
	data := packets.ParseEntityMoved(event)
	entity := ecs.MainWorld.GetEntity(data.EntityId)
	entity.Add(
		data.Move,
	)
	return data
}

func (c *Client) requestError(event events.Raw) serialization.ISerializable {
	data := packets.ParseRequestError(event)
	fmt.Printf("Request error: %s \n", data.Message)
	return data
}

func (c *Client) joinShardResponse(event events.Raw) serialization.ISerializable {
	data := packets.ParseJoinShardResponse(event)
	ecs.MainWorld.NewEntityFrom(
		data.EntityId,
		data.Transform,
		data.Living,
		data.PlayerCircle,
		ecs.NewPlayer(),
	)

	return data
}

func (c *Client) entityRemoved(event events.Raw) serialization.ISerializable {
	println("AHHHHHHHHHHHHHHHHHHHHHHHHH")
	data := packets.ParseEntityRemoved(event)
	ecs.MainWorld.RemoveEntity(data.EntityId)
	return data
}
