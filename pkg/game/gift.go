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

func NewGift(spec BrickSpec) *Gift {
	var img = LoadSprite("gift.png")
	gift := &Gift{
		position: pixel.ZV,
		picture:  pixel.PictureDataFromImage(img),
		width:    GiftWidth,
		height:   GiftHeight,
		spec:     spec,
	}
	gift.Spec()
	return gift
}

func (r *Gift) Spec() {
	switch r.spec {
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

func (r *Gift) HitBoard(board *Board) bool {
	var giftBottom = r.Bottom()
	if (giftBottom.Y <= board.Top().Y) &&
		(giftBottom.X2 >= board.Top().X1 && giftBottom.X1 <= board.Top().X2) {
		return true
	}
	return false
}

func (r *Gift) FallAway(board *Board) bool {
	var giftBottom = r.Bottom()
	if (giftBottom.Y <= board.Top().Y) &&
		(giftBottom.X2 < board.Top().X1 || giftBottom.X1 > board.Top().X2) {
		return true
	}
	return false
}

func (r Gift) Draw(win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, GiftScale)
	r.sprite.Draw(win, mat.Moved(r.position))
}

func (r Gift) Bottom() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.width/2,
		Y:  r.position.Y - r.height/2,
	}
}
