package shard

import (
	"mmo2/pkg/events"
	"mmo2/pkg/packets"
	"mmo2/pkg/serialization"
)

func (c *Client) joinShardResponse(event events.Raw) {
	join := packets.JoinShardResponse{}
	join.FromBytes(event)
	c.world.NewEntityFrom(join.EntityId, []serialization.ISerializable{&join.Position})
}

func (c *Client) playerJoined(event events.Raw) {
	data := packets.PlayerJoined{}
	data.FromBytes(event)
	c.world.NewEntityFrom(data.EntityId, []serialization.ISerializable{&data.Position})
}
