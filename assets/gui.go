package assets

import rl "github.com/gen2brain/raylib-go/raylib"

var Window1 = rl.NPatchInfo{
	Source: rl.Rectangle{
		X: 0, Y: 0, Width: 65, Height: 64,
	},
	Left: 20, Top: 20, Right: 20, Bottom: 20, Layout: rl.NPatchNinePatch,
}
