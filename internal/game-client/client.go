package game

import (
	"mmo2/assets"
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
	"mmo2/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ClientOptions struct {
	ShardAddress string
	ShardPort    int
	Title        string
}

type Client struct {
	shardClient *shard.Client
	world       *game.World
	options     ClientOptions
}

func NewClient(options ClientOptions) *Client {
	client := Client{}
	client.world = game.NewWorld()
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
	defer rl.CloseWindow()
	tickChan := c.shardClient.TickChan()
	texture := rl.LoadTexture("assets/images/simple_rpg_gui.png")
	gui.SetTexture(texture)
	container := gui.Window(assets.MainPanel, 0, 0, 800, 600, 100)
	window := gui.NewWidget(assets.MainPanel, container, 10)
	cells := gui.NewGrid(window, 15, 10, 0)
	for _, cell := range cells {
		gui.NewWidget(assets.SlotWhite, cell, 0)
	}
	for !rl.WindowShouldClose() {
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		window.Render()
		rl.EndDrawing()
		tickChan <- 1
	}
}
