package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	win   *pixelgl.Window
	image *Image
	wall  []*Brick
}

func NewWall(win *pixelgl.Window, image *Image) *Wall {
	var w = &Wall{
		win:   win,
		image: image,
	}
	w.Build()
	return w
}

func (r *Wall) Build() {
	for i := 0.0; i < 15; i++ {
		brick := NewBrick(r.image, Orange)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2))
		r.Add(brick)

		brick = NewBrick(r.image, Green)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2+brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Pink)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2+2*brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Blue)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2+3*brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Red)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1), r.win.Bounds().H()/2+4*brick.height))
		r.Add(brick)
	}
}

func (r *Wall) Add(brick *Brick) {
	r.wall = append(r.wall, brick)
}

func (r Wall) Draw() {
	for _, brick := range r.wall {
		if brick.IsNotHit() {
			brick.Draw(r.win)
		}
	}
}
