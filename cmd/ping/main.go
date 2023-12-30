package main

import (
	"fmt"
	"mmo2/game/packets"
	"mmo2/internal/shard-client"
	"mmo2/pkg/serialization"
	"time"
)

func main() {
	shardClient := shard.NewClient()
	// shardClient.Connect("159.203.96.193", 5555)
	shardClient.Connect("127.0.0.1", 5555)
	for {
		println("sending ping request")
		now := time.Now()
		shardClient.SendRequest(packets.NewPing(), func(response serialization.ISerializable) {
			ping := time.Since(now).Milliseconds()
			fmt.Printf("Ping: %d \n", ping)
		})
		time.Sleep(time.Second)
	}
}
