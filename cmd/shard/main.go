package main

import (
	"fmt"
	"mmo2/game/ecs"
	"mmo2/internal/shard-server"
	"os"
)

func main() {
	server := shard.New()
	err := server.Start("", 5555)
	if err != nil {
		panic(err)
	}
	for {
		var command string
		fmt.Fscan(os.Stdin, &command)
		if command == "entities" {
			print("Entities: {")
			for _, entity := range ecs.MainWorld.Entites() {
				println(entity.String())
			}
			print(" } \n")
		}
	}
}
