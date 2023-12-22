package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func calculateRect(parent *Widget) rl.Rectangle {
	return rl.NewRectangle(
		parent.destRect.X+float32(parent.padding),
		parent.destRect.Y+float32(parent.padding),
		parent.destRect.Width-float32(parent.padding)*2,
		parent.destRect.Height-float32(parent.padding)*2,
	)
}

func recursiveRender(w *Widget) {
	if w.renderer != nil {
		w.renderer.Render(w)
	}
	for _, child := range w.childs {
		recursiveRender(child)
	}
}

var guiTexture rl.Texture2D

func SetTexture(texture rl.Texture2D) {
	guiTexture = texture
}

type Renderer interface {
	Render(w *Widget)
}

func Window(renderer Renderer, x int, y int, width int, height int, padding int32) *Widget {
	w := Widget{
		renderer: renderer,
		parent:   nil,
		padding:  int32(padding),
		destRect: rl.NewRectangle(float32(x), float32(y), float32(width), float32(height)),
		childs:   make([]*Widget, 0),
	}
	return &w
}
