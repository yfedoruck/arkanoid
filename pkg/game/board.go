package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	//BoardWidth  = 160
	//BoardHeight = 25
	BoardScale   = 3
)

type Board struct {
	width    float64
	height   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
	picture  *pixel.PictureData
}

func NewBoard(win *pixelgl.Window) *Board {
	sp := BoardSprite()
	return &Board{
		width:    sp.Picture().Bounds().W()*BoardScale,
		height:   sp.Picture().Bounds().H()*BoardScale,
		win:      win,
		position: pixel.ZV,
		sprite:   sp,
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
	if r.position.X <= r.win.Bounds().Min.X+r.width/2+(BgBorderX*BgScale) {
		return
	}
	r.position.X -= delta
}

func (r *Board) MoveRight(delta float64) {
	if r.position.X >= r.win.Bounds().Max.X-r.width/2-(BgBorderX*BgScale) {
		return
	}
	r.position.X += delta
}

func (r Board) Draw() {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, BoardScale)
	r.sprite.Draw(r.win, mat.Moved(r.position))
}

func (r Board) Area() VecX {
	return VecX{
		r.position.X - r.width/2,
		r.position.X + r.width/2,
	}
}

func BoardSprite() *pixel.Sprite {
	var picture = pixel.PictureDataFromImage(LoadSprite("racket.png"))
	return pixel.NewSprite(picture, picture.Bounds())
}
