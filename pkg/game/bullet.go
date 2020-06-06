package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BulletScale = 0.5
)

type Bullet struct {
	width    float64
	height   float64
	delta    float64
	win      *pixelgl.Window
	sprite   *pixel.Sprite
	shot     bool
	scale    float64
	boardTop BrickSideX
	position pixel.Vec
}

func NewBullet(win *pixelgl.Window, top BrickSideX, delta float64) *Bullet {
	sp := SimpleSprite("bullet.png")
	scale := BulletScale
	bullet := &Bullet{
		width:    sp.Frame().W() * scale,
		height:   sp.Frame().H() * scale,
		win:      win,
		boardTop: top,
		sprite:   sp,
		scale:    scale,
		delta:    delta,
	}

	return bullet
}

func (r *Bullet) Left() {
	frame := r.sprite.Frame()
	r.position = pixel.V(r.boardTop.X1+frame.W()/4, r.boardTop.Y+frame.H()/4)
}

func (r *Bullet) Right() {
	frame := r.sprite.Frame()
	r.position = pixel.V(r.boardTop.X2-frame.W()/4, r.boardTop.Y+frame.H()/4)
}

func (r *Bullet) Shot() {
	r.position.Y += r.delta
}

func (r Bullet) Draw() {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, BulletScale)
	r.sprite.Draw(r.win, mat.Moved(r.position))
}
