package assets

import (
	"mmo2/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var MainPanel = gui.NewPanel(rl.NPatchInfo{
	Source: rl.Rectangle{
		X: 0, Y: 0, Width: 65, Height: 64,
	},
	Left: 20, Top: 20, Right: 20, Bottom: 20, Layout: rl.NPatchNinePatch,
})
