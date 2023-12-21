package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func NewGrid(parent *Widget, columns int, rows int) []*Widget {
	cols := make([]*Widget, 0)
	innerRect := calculateRect(parent)
	cellWidth := innerRect.Width / float32(columns)
	cellHeight := innerRect.Height / float32(rows)
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			w := &Widget{
				renderer: nil,
				parent:   parent,
				destRect: rl.NewRectangle(
					innerRect.X+cellWidth*float32(i),
					innerRect.Y+cellHeight*float32(j),
					cellWidth,
					cellHeight,
				),
				childs: make([]*Widget, 0),
			}
			if parent.childs != nil {
				parent.childs = append(parent.childs, w)
			}
			cols = append(cols, w)
		}
	}
	return cols
}
