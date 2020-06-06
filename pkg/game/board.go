package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	//BoardWidth  = 160
	//BoardHeight = 25
	BoardScale    = 3.0
	BigBoardScale = 1.0
)

type Board struct {
	width    float64
	height   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
	spSimple *pixel.Sprite
	spBig    *pixel.Sprite
	sticky   bool
	scale    float64
}

func NewBoard(win *pixelgl.Window) *Board {
	sp := BoardSprite()
	scale := BoardScale
	return &Board{
		width:    sp.Frame().W() * scale,
		height:   sp.Frame().H() * scale,
		win:      win,
		position: pixel.ZV,
		sprite:   sp,
		spSimple: sp,
		scale:    scale,
	}
}

func (r *Board) Simple() {
	scale := BoardScale
	r.width = r.spSimple.Frame().W() * scale
	r.height = r.spSimple.Frame().H() * scale
	r.sprite = r.spSimple
	r.scale = scale
	r.sticky = false
}

func (r *Board) BigBoard() {
	if r.spBig == nil {
		r.spBig = NewImage().Board()
	}
	scale := BigBoardScale
	r.width = r.spBig.Frame().W() * scale
	r.height = r.spBig.Frame().H() * scale
	r.sprite = r.spBig
	r.scale = scale
	r.sticky = false
}

func (r Board) Width() float64 {
	return r.width
}

func (r Board) Height() float64 {
	return r.height
}

func (r *Board) Sticky() {
	r.sticky = true
}

func (r *Board) DelSticky() {
	r.sticky = false
}

func (r Board) IsSticky() bool {
	return r.sticky
}

func (r Board) TopCenter() pixel.Vec {
	return pixel.V(r.width/2+r.height, r.height)
}

func (r Board) Top() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.height/2,
		Y:  r.position.Y + r.height/2,
	}
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

func (r *Board) CleanSprite() {
	r.sprite = TransparentPixel()
}

func (r Board) Draw() {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, r.scale)
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
