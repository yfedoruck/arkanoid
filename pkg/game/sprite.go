package game

import (
	"github.com/faiface/pixel"
)

const (
	SpriteWidth = 180
	SpriteHeight = 136
)

type BasicPack struct {
	Image
}

func NewBasicPack() *BasicPack {
	var s = &BasicPack{}
	s.load("BasicArkanoidPack.png")
	return s
}

func (r BasicPack) Ball() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 32, 24, 33+24))
}
func (r BasicPack) Board() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 0, 160, 25))
}

const (
	Orange = iota
	Pink
	Red
	Green
	Blue
)

func (r BasicPack) Brick(color int) *pixel.Sprite {
	var rec = pixel.R(0,0,0,0)
	switch color {
	case Orange:
		rec = pixel.R(0, SpriteHeight - BrickHeight, BrickWidth, SpriteHeight)
	case Red:
		minX := BrickWidth + BrickGap
		rec = pixel.R(minX, SpriteHeight - BrickHeight, minX + BrickWidth, SpriteHeight)
	case Green:
		minX := 2*(BrickWidth + BrickGap)
		rec = pixel.R(minX, SpriteHeight - BrickHeight, minX + BrickWidth, SpriteHeight)
	case Blue:
		minX := 3*(BrickWidth + BrickGap)
		rec = pixel.R(minX, SpriteHeight - BrickHeight, minX + BrickWidth, SpriteHeight)
	case Pink:
		maxY := SpriteHeight - BrickHeight - BrickGap
		rec = pixel.R(0, maxY - BrickHeight, BrickWidth, maxY)

	}
	return pixel.NewSprite(r.picture, rec)
}
