package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	win   *pixelgl.Window
	brick *Brick
	ball  *Ball
	wall  []*Brick
}

func NewWall(win *pixelgl.Window, ball *Ball) *Wall {
	w := &Wall{win: win, ball: ball}
	w.Build()
	return w
}

func (r *Wall) Build() {
	for i := 0.0; i < 4; i++ {
		brick := NewBrick(r.win)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2))
		r.wall = append(r.wall, brick)
	}
}

func (r Wall) Draw() {
	for _, brick := range r.wall {
		if brick.IsNotHit() {
			brick.Draw()
		}
	}
}
