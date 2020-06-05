package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	win      *pixelgl.Window
	image    *Image
	wall     []*Brick
	giftPack []*Gift
}

func NewWall(win *pixelgl.Window, image *Image) *Wall {
	var w = &Wall{
		win:   win,
		image: image,
	}
	return w
}

func (r *Wall) level1() {
	dx := 8.0
	for i := 0.0; i < 10; i++ {
		brick := NewBrick(r.image, Orange)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2))
		r.Add(brick)

		brick = NewBrick(r.image, Green)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Pink)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+2*brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Blue)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+3*brick.height))
		r.Add(brick)

		brick = NewBrick(r.image, Red)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+4*brick.height))
		r.Add(brick)
	}
	r.SetGiftBricks()
}

func (r *Wall) SetGiftBricks() {
	for i, brick := range r.wall {
		switch {
		case i%4 == 0:
			brick.SetSpec(GlueBrick)
		case i%3 == 0:
			brick.SetSpec(GunBrick)
		}
	}
}

func (r *Wall) level2() {
	dx := 8.0
	var brick *Brick
	for i := 0.0; i < 10; i++ {
		brick = NewBrick(r.image, Green)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2))
		r.Add(brick)
	}
	for i := 0.0; i < 9; i++ {
		brick = NewBrick(r.image, Pink)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+brick.height))
		r.Add(brick)
	}
	for i := 0.0; i < 8; i++ {
		brick = NewBrick(r.image, Blue)
		brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(i+1)+dx, r.win.Bounds().H()/2+2*brick.height))
		r.Add(brick)
	}
}

func (r *Wall) levelTest1() {
	brick := NewBrick(r.image, Orange)
	brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(5)+8, r.win.Bounds().H()/2))
	r.Add(brick)
}
func (r *Wall) levelTest2() {
	brick := NewBrick(r.image, Green)
	brick.MoveTo(pixel.V(r.win.Bounds().Min.X+brick.width*(5)+8, r.win.Bounds().H()/2))
	r.Add(brick)
}

func (r Wall) IsDestroyed() bool {
	return 0 == len(r.wall)
}

func (r *Wall) Add(brick *Brick) {
	r.wall = append(r.wall, brick)
}

func (r *Wall) Draw(delta float64) {
	var idx []int
	for i, brick := range r.wall {
		if brick.IsNotHit() {
			brick.Draw(r.win)
		} else {
			if brick.HasGift() {
				gift := NewGift()
				gift.Spec(GlueBrick)
				gift.position = brick.position
				r.giftPack = append(r.giftPack, gift)
			}
			idx = append(idx, i)
		}
	}

	for _, gift := range r.giftPack {
		gift.Fall(delta/2)
		gift.Draw(r.win)
	}

	for _, i := range idx {
		r.DeleteBrick(i)
	}
}

func (r *Wall) DeleteBrick(i int) {
	r.wall[i] = r.wall[len(r.wall)-1]
	r.wall[len(r.wall)-1] = nil
	r.wall = r.wall[:len(r.wall)-1]
}
