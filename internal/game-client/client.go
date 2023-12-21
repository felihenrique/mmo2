package game

import (
	"fmt"
	"mmo2/assets"
	"mmo2/internal/shard-client"
	"mmo2/pkg/game"
	"time"

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
	var timeStep int64
	texture := rl.LoadTexture("assets/images/simple_rpg_gui.png")
	music := rl.LoadMusicStream("assets/music/main_title.mp3")
	logo := rl.LoadTexture("assets/images/logo.png")
	rl.PlayMusicStream(music)
	for !rl.WindowShouldClose() {
		now := time.Now()
		rl.UpdateMusicStream(music)
		<-tickChan
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.DrawTexturePro(
			logo,
			rl.NewRectangle(0, 0, 1024, 1024),
			rl.NewRectangle(400, 200, 400, 400),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
		rl.DrawTextureNPatch(
			texture, assets.Window1,
			rl.Rectangle{X: 100, Y: 100, Width: 300, Height: 200},
			rl.Vector2{X: 0, Y: 0}, 0, rl.White,
		)
		rl.DrawText(
			fmt.Sprintf("Frame %d", timeStep), 0, 0, 20, rl.LightGray,
		)
		rl.EndDrawing()
		tickChan <- 1
		timeStep = time.Since(now).Milliseconds()
	}
}
