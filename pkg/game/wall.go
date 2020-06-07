package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Wall struct {
	win      *pixelgl.Window
	image    *BasicPack
	wall     []*Brick
	giftPack []*Gift
	board    *Board
	delta    float64
}

func NewWall(win *pixelgl.Window, image *BasicPack, board *Board) *Wall {
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
		case i%2 == 0:
			brick.SetSpec(GunBoard)
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
	r.SetGiftBricks()
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

func (r *Wall) AddGiftFrom(brick *Brick) {
	gift := NewGift(brick.spec)
	gift.position = brick.position
	r.giftPack = append(r.giftPack, gift)
}

func (r *Wall) Draw(delta float64) {
	r.DrawBricks()
	r.DrawGifts(delta)
}

func (r *Wall) DrawBricks() {
	var wall = r.wall[:0]
	for _, brick := range r.wall {
		if brick.IsNotHit() {
			brick.Draw(r.win)
			wall = append(wall, brick)
		} else {
			if brick.HasGift() {
				r.AddGiftFrom(brick)
			}
		}
	}
	r.wall = wall
}

func (r *Wall) DrawGifts(delta float64) {
	var giftPack = r.giftPack[:0]
	for _, gift := range r.giftPack {
		gift.Fall(delta / 2)
		switch {
		case gift.HitBoard(r.board):
			r.UseGift(gift.spec)
		case gift.FallAway(r.board):
		default:
			giftPack = append(giftPack, gift)

		}
		gift.Draw(r.win)
	}
	r.giftPack = giftPack
}

func (r *Wall) SetDelta(delta float64) {
	r.delta = delta
}
func (r *Wall) Clean() {
	r.board.CleanMagazine()
	r.board.StopFire()
	r.CleanGifts()
}

func (r *Wall) CleanGifts() {
	r.giftPack = r.giftPack[:0]
}

func (r *Wall) UseGift(spec BrickSpec) {
	switch spec {
	case GlueBoard:
		r.StickyBoard()
	case BigBoard:
		r.BigBoard()
	case GunBoard:
		r.GunBoard()
	}
}

func (r *Wall) StickyBoard() {
	r.board.Simple()
	r.board.Sticky()
}

func (r *Wall) BigBoard() {
	r.board.BigBoard()
}

func (r *Wall) GunBoard() {
	r.board.CleanMagazine()
	r.board.StopFire()
	r.board.GunBoard()
}
