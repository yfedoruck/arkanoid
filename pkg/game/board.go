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
	GunScale      = 0.25
)

type Board struct {
	width    float64
	height   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
	spSimple *pixel.Sprite
	spBig    *pixel.Sprite
	spGun    *pixel.Sprite
	sticky   bool
	shot     bool
	magazine []*Bullet
	scale    float64
	delta    float64
}

func NewBoard(win *pixelgl.Window) *Board {
	sp := SimpleSprite("racket.png")
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
		r.spBig = NewBasicPack().Board()
	}
	scale := BigBoardScale
	r.width = r.spBig.Frame().W() * scale
	r.height = r.spBig.Frame().H() * scale
	r.sprite = r.spBig
	r.scale = scale
	r.sticky = false
}

func (r *Board) GunBoard() {
	if r.spGun == nil {
		r.spGun = SimpleSprite("GunBoard.png")
	}
	scale := GunScale
	r.width = r.spGun.Frame().W() * scale
	r.height = r.spGun.Frame().H() * scale
	r.sprite = r.spGun
	r.scale = scale
	r.sticky = false
}

func (r *Board) Shot() {
	if r.scale != GunScale {
		return
	}
	r.shot = true
	r.addBullets()
}

func (r *Board) addBullets() {
	bullet1 := NewBullet(r.win, r.Top(), r.delta)
	bullet1.Left()

	bullet2 := NewBullet(r.win, r.Top(), r.delta)
	bullet2.Right()
	r.magazine = append(r.magazine, bullet1, bullet2)
}

func (r Board) IsShot() bool {
	return r.shot
}

func (r Board) Width() float64 {
	return r.width
}

func (r Board) Height() float64 {
	return r.height
}

func (r *Board) SetDelta(delta float64) {
	r.delta = delta
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
		X2: r.position.X + r.width/2,
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

func (r *Board) Draw(wall *Wall) {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, r.scale)
	r.sprite.Draw(r.win, mat.Moved(r.position))

	if r.IsShot() {
		r.Fire(wall)
	}
}

func (r *Board) Fire(wall *Wall) {
	mag := r.magazine[:0]
	for _, bullet := range r.magazine {
		if bullet.IsExploded() {
			continue
		}
		mag = append(mag, bullet)
		bullet.Fly(wall)
	}
	r.magazine = mag
}

func (r Board) Area() VecX {
	return VecX{
		r.position.X - r.width/2,
		r.position.X + r.width/2,
	}
}
