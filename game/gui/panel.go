package gui

import rl "github.com/gen2brain/raylib-go/raylib"

type Panel struct {
	nPatch rl.NPatchInfo
}

func NewPanel(kind rl.NPatchInfo) *Panel {
	panel := Panel{}
	panel.nPatch = kind
	return &panel
}

func (p *Panel) Render(w *Widget) {
	rl.DrawTextureNPatch(
		guiTexture,
		p.nPatch,
		w.destRect,
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}
