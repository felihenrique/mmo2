package assets

import (
	"mmo2/game/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var MainPanel = gui.NewPanel(rl.NPatchInfo{
	Source: rl.NewRectangle(0, 0, 65, 64),
	Left:   20, Top: 20, Right: 20, Bottom: 20, Layout: rl.NPatchNinePatch,
})

var Slot = gui.NewPanel(rl.NPatchInfo{
	Source: rl.NewRectangle(115, 36, 34, 34),
	Left:   5, Top: 5, Right: 5, Bottom: 5, Layout: rl.NPatchNinePatch,
})

var SlotWhite = gui.NewPanel(rl.NPatchInfo{
	Source: rl.NewRectangle(152, 36, 36, 34),
	Left:   10, Top: 10, Right: 10, Bottom: 10, Layout: rl.NPatchNinePatch,
})

var GreenButton = gui.NewPanel(rl.NPatchInfo{
	Source: rl.NewRectangle(263, 1, 70, 25),
	Left:   10, Top: 10, Right: 10, Bottom: 10, Layout: rl.NPatchNinePatch,
})
