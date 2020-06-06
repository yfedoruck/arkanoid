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
	board    *Board
}

func NewWall(win *pixelgl.Window, image *Image, board *Board) *Wall {
	var w = &Wall{
		win:   win,
		image: image,
		board: board,
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
			brick.SetSpec(GlueBoard)
		case i%3 == 0:
			brick.SetSpec(BigBoard)
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
				gift := NewGift(brick.spec)
				gift.position = brick.position
				r.giftPack = append(r.giftPack, gift)
			}
			idx = append(idx, i)
		}
	}

	var delGift []int
	for i, gift := range r.giftPack {
		gift.Fall(delta / 2)
		switch {
		case gift.HitBoard(r.board):
			r.UseGift(gift.spec)
			delGift = append(delGift, i)
		case gift.FallAway(r.board):
			delGift = append(delGift, i)
		}
		gift.Draw(r.win)
	}

	for _, i := range idx {
		r.DeleteBrick(i)
	}

	for _, i := range delGift {
		r.DeleteGift(i)
	}
}

func (r *Wall) DeleteBrick(i int) {
	r.wall[i] = r.wall[len(r.wall)-1]
	r.wall[len(r.wall)-1] = nil
	r.wall = r.wall[:len(r.wall)-1]
}

func (r *Wall) DeleteGift(i int) {
	r.giftPack[i] = r.giftPack[len(r.giftPack)-1]
	r.giftPack[len(r.giftPack)-1] = nil
	r.giftPack = r.giftPack[:len(r.giftPack)-1]
}

func (r *Wall) UseGift(spec BrickSpec) {
	switch spec {
	case GlueBoard:
		r.StickyBoard()
	case BigBoard:
		r.BigBoard()
	}
}

func (r *Wall) StickyBoard() {
	r.board.Simple()
	r.board.Sticky()
}

func (r *Wall) BigBoard() {
	r.board.BigBoard()
}
