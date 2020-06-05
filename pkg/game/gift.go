package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	GiftWidth  = 64.0
	GiftHeight = 32.0
	GiftScale = 3
)

type Gift struct {
	width    float64
	height   float64
	position pixel.Vec
	picture  *pixel.PictureData
	sprite   *pixel.Sprite
	spec     BrickSpec
}

func NewGift() *Gift {
	var img = LoadSprite("gift.png")
	return &Gift{
		position: pixel.ZV,
		picture:  pixel.PictureDataFromImage(img),
		width:    GiftWidth,
		height:   GiftHeight,
		spec:     SimpleBrick,
	}
}

func (r *Gift) Spec(spec BrickSpec) {
	switch spec {
	case GlueBrick:
		r.Blue()
	case GunBrick:
		r.Green()
	}
}

func (r *Gift) Blue() {
	r.sprite = pixel.NewSprite(r.picture, pixel.R(126, 60, r.picture.Bounds().Max.X, 70))
}

func (r *Gift) Green() {
	r.sprite = pixel.NewSprite(r.picture, pixel.R(126, 70, r.picture.Bounds().Max.X, 80))
}

func (r *Gift) SetSpec(spec BrickSpec) {
	r.spec = spec
}

func (r *Gift) Fall(delta float64) {
	r.position.Y -= delta
}

func (r Gift) Draw(win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, GiftScale)
	r.sprite.Draw(win, mat.Moved(r.position))
}

func (r *Gift) MoveTo(pos pixel.Vec) {
	r.position = pos
}
