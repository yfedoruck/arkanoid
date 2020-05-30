package game

import (
	"github.com/faiface/pixel"
)

type Wall struct {
	screen *Screen
	wall   []*Brick
}

func NewWall(screen *Screen) *Wall {
	var w = &Wall{}
	w.Build(screen)
	return w
}

func (r *Wall) Build(screen *Screen) {
	win := screen.Window()
	for i := 0.0; i < 15; i++ {
		brick := NewBrick(screen, Orange)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2))
		r.Add(brick)

		brick = NewBrick(screen, Green)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2+brick.height))
		r.Add(brick)

		brick = NewBrick(screen, Pink)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2+2*brick.height))
		r.Add(brick)

		brick = NewBrick(screen, Blue)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2+3*brick.height))
		r.Add(brick)

		brick = NewBrick(screen, Red)
		brick.MoveTo(pixel.V(win.Bounds().Min.X+brick.width*(i+1), win.Bounds().H()/2+4*brick.height))
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
