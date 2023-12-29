package shard

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/game/packets"
	"mmo2/pkg/event_utils"
	"mmo2/pkg/serialization"
	"time"
)

func (c *Client) ackRequest(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParseAckRequest(event)
	return data
}

func (c *Client) playerJoined(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParsePlayerJoined(event)
	ecs.MainWorld.NewEntityFromBytes(data.Entity)
	return data
}

func (c *Client) entityMoved(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParseEntityMoved(event)
	entity := ecs.MainWorld.GetEntity(data.EntityId)
	entity.Add(data.Move)
	fmt.Printf("[%d] Received move", time.Now().UnixMilli())
	return data
}

func (c *Client) requestError(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParseRequestError(event)
	fmt.Printf("Request error: %s \n", data.Message)
	return data
}

func (c *Client) joinShardResponse(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParseJoinShardResponse(event)
	entity := ecs.MainWorld.NewEntityFromBytes(data.PlayerEntity)
	entity.Add(ecs.NewPlayer())
	return data
}

func (c *Client) entityRemoved(event event_utils.Raw) serialization.ISerializable {
	data, _ := packets.ParseEntityRemoved(event)
	ecs.MainWorld.RemoveEntity(data.EntityId)
	return data
}
