package gui

import rl "github.com/gen2brain/raylib-go/raylib"

type Widget struct {
	renderer Renderer
	parent   *Widget
	destRect rl.Rectangle
	padding  int32
	childs   []*Widget
}

func (w *Widget) Render() {
	recursiveRender(w)
}

func NewWidget(renderer Renderer, parent *Widget, padding int32) *Widget {
	widget := &Widget{
		renderer: renderer,
		parent:   parent,
		destRect: calculateRect(parent),
		padding:  padding,
		childs:   make([]*Widget, 0),
	}
	parent.childs = append(parent.childs, widget)
	return widget
}
