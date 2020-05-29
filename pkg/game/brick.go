package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BrickWidth = 64
	BrickHeight = 32
)

type Brick struct {
	width    float64
	height   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
}

type BrickSideY struct {
	X, Y1, Y2 float64
}
type BrickSideX struct {
	X1, X2, Y float64
}

func NewBrick(win *pixelgl.Window) *Brick {
	return &Brick{
		win:      win,
		position: pixel.ZV,
		sprite:   NewSprite().Brick(),
		width:  BrickWidth,
		height:  BrickHeight,
	}
}

func (r Brick) Left() BrickSideY {
	return BrickSideY{
		X: r.position.X - r.width/2,
		Y1: r.position.Y + r.height/2,
		Y2: r.position.Y - r.height/2,
	}
}

func (r Brick) Right() BrickSideY {
	return BrickSideY{
		X: r.position.X + r.width/2,
		Y1: r.position.Y + r.height/2,
		Y2: r.position.Y - r.height/2,
	}
}

func (r Brick) Top() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.height/2,
		Y: r.position.Y + r.height/2,
	}
}
func (r Brick) Bottom() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.width/2,
		Y: r.position.Y - r.height/2,
	}
}


func (r *Brick) OnStartPosition() {
	r.position = pixel.V(r.win.Bounds().W()/2, r.win.Bounds().H()/2)
}

func (r Brick) Draw() {
	r.sprite.Draw(r.win, pixel.IM.Moved(r.position))
}