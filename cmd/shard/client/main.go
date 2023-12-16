package main

import (
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
)

func main() {
	world := game.NewWorld(10000)
	client := shard.NewClient(world)
	err := client.Connect("", 5555)
	if err != nil {
		panic(err)
	}
}
