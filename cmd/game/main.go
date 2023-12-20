package main

import "mmo2/internal/game-client"

func main() {
	client := game.NewClient(game.ClientOptions{
		ShardAddress: "",
		ShardPort:    5555,
		Title:        "Best MMORPG",
	})
	client.Start()
}
