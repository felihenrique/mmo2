package game

import (
	assets_scene "mmo2/assets/scenes"
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
	options     ClientOptions
}

func NewClient(options ClientOptions) *Client {
	client := Client{}
	client.shardClient = shard.NewClient()
	client.shardClient.EnableTickChan()
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
	defer rl.CloseWindow()
	rl.InitWindow(800, 600, c.options.Title)
	rl.InitAudioDevice()
	scene.ChangeTo(assets_scene.MainMenu)
	tickChan := c.shardClient.TickChan()
	for !rl.WindowShouldClose() {
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		ecs.MainWorld.Update()
		scene.RenderGUI(c.shardClient)
		rl.EndDrawing()
		tickChan <- 1
	}
}
