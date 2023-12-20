package game

import (
	"fmt"
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
	"time"

	rlgui "github.com/gen2brain/raylib-go/raygui"
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
	defer rl.CloseWindow()
	tickChan := c.shardClient.TickChan()
	var timeStep int64
	var name string
	for !rl.WindowShouldClose() {
		now := time.Now()
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		mouseCell := rl.Vector2{}
		rlgui.Grid(rl.NewRectangle(50, 50, 700, 500), "", 30, 2, &mouseCell)
		rl.DrawText(
			"Enter your name: ", 288, 225, 20, rl.LightGray,
		)
		rlgui.TextBox(rl.NewRectangle(288, 250, 224, 32), &name, 10, true)
		if rlgui.Button(rl.NewRectangle(350, 350, 100, 40), "Join game") {
			fmt.Println("omg")
		}
		rl.DrawText(
			fmt.Sprintf("Frame %d", timeStep), 0, 0, 20, rl.LightGray,
		)
		rl.EndDrawing()
		tickChan <- 1
		timeStep = time.Since(now).Milliseconds()
	}
}
