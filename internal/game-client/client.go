package game

import (
	"mmo2/game/ecs"
	"mmo2/game/scene"
	"mmo2/internal/shard-client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ClientOptions struct {
	ShardAddress string
	ShardPort    int
	Title        string
}

type Client struct {
	shardClient *shard.Client
	world       *ecs.World
	options     ClientOptions
}

func NewClient(options ClientOptions) *Client {
	client := Client{}
	client.world = ecs.NewWorld()
	client.shardClient = shard.NewClient(client.world)
	client.options = options
	return &client
}

func (c *Client) Start() error {
	err := c.shardClient.Connect(c.options.ShardAddress, c.options.ShardPort)
	if err != nil {
		return err
	}
	c.mainLoop()
	return nil
}

func (c *Client) mainLoop() {
	rl.InitWindow(800, 600, c.options.Title)
	rl.InitAudioDevice()
	scene := scene.MainMenu{
		ShardClient: c.shardClient,
	}
	scene.Init(c.world)
	defer rl.CloseWindow()
	tickChan := c.shardClient.TickChan()
	for !rl.WindowShouldClose() {
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		scene.Update(c.world, 0)
		rl.EndDrawing()
		tickChan <- 1
	}
}
