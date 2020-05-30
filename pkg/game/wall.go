package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	win   *pixelgl.Window
	wall  []*Brick
}

func NewWall(win *pixelgl.Window) *Wall {
	var w = &Wall{}
	w.Build(win)
	return w
}

func (r *Wall) Build(win *pixelgl.Window) {
	for i := 0.0; i < 10; i++ {
		brick := NewBrick(win)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2))
		r.Add(brick)
	}
}

func (r *Wall) Add(brick *Brick) {
	r.wall = append(r.wall, brick)
}

func (r Wall) Draw() {
	for _, brick := range r.wall {
		if brick.IsNotHit() {
			brick.Draw()
		}
	}
}
