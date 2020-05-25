package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BoardWidth  = 160
	BoardHeight = 25
)

type Board struct {
	width    float64
	height   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
}

func NewBoard(win *pixelgl.Window) Board {
	return Board{
		width:    BoardWidth,
		height:   BoardHeight,
		win:      win,
		position: pixel.ZV,
		sprite:   NewSprite().Board(),
	}
}

func (r Board) Width() float64 {
	return r.width
}

func (r Board) Height() float64 {
	return r.height
}

func (r Board) TopCenter() pixel.Vec {
	return pixel.V(r.width/2+r.height, r.height)
}

func (r *Board) OnStartPosition() {
	r.position = pixel.V(r.win.Bounds().W()/2, r.height/2)
}

func (r *Board) StartPosition() pixel.Vec {
	return pixel.V(r.win.Bounds().W()/2, r.height/2)
}

func (r *Board) MoveLeft(delta float64) {
	r.position.X -= delta
}

func (r *Board) MoveRight(delta float64) {
	r.position.X += delta
}

func (r Board) Draw() {
	r.sprite.Draw(r.win, pixel.IM.Moved(r.position))
}
