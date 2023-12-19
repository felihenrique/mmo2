package main

import (
	"fmt"
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
	"mmo2/pkg/packets"
	"os"
)

func main() {
	world := game.NewWorld(10000)
	client := shard.NewClient(world)
	err := client.Connect("", 5555)
	if err != nil {
		panic(err)
	}
	for {
		var command string
		fmt.Fscan(os.Stdin, &command)
		if command == "join" {
			client.SendRequest(&packets.JoinShardRequest{
				Name:   "abc",
				Portal: 1,
			})
			println("sent request to join shard")
		} else if command == "entities" {
			for _, entity := range world.Entites() {
				println(entity.String())
			}
		}
	}
}
