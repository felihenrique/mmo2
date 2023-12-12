package main

import (
	"mmo2/internal/shard"
)

func main() {
	server := shard.New("", 5555)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
