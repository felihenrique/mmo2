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
	gui.SetGuiTexture(texture)
	window := gui.Window(800, 600, 100)
	panel := gui.NewWidget(assets.MainPanel, window, 30)
	cells := gui.NewGrid(panel, 2, 2)
	gui.NewWidget(assets.MainPanel, cells[0], 15)
	gui.NewWidget(assets.MainPanel, cells[1], 15)
	gui.NewWidget(assets.MainPanel, cells[2], 15)
	gui.NewWidget(assets.MainPanel, cells[3], 15)
	for !rl.WindowShouldClose() {
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		window.Render()
		rl.EndDrawing()
		tickChan <- 1
	}
}
