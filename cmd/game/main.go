package main

import "mmo2/internal/game-client"

func main() {
	client := game.NewClient(game.ClientOptions{
		ShardAddress: "159.203.96.193",
		// ShardAddress: "",
		ShardPort: 5555,
		Title:     "Best MMORPG",
	})
	err := client.Start()
	if err != nil {
		println(err.Error())
	}
}
