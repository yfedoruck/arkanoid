package game

import (
	"github.com/faiface/pixel"
	"github.com/yfedoruck/billiards/pkg/env"
	"github.com/yfedoruck/billiards/pkg/fail"
	"image"
	"os"
	"path/filepath"
)

const (
	SpriteWidth = 180
	SpriteHeight = 136
)

type Image struct {
	picture *pixel.PictureData
}

func NewImage() *Image {
	var s = &Image{}
	s.load("BasicArkanoidPack.png")
	return s
}

func (r Image) Ball() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 32, 24, 33+24))
}
func (r Image) Board() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 0, 160, 25))
}

const (
	Orange = iota
	Pink
	Red
	Green
	Blue
)

func (r Image) Brick(color int) *pixel.Sprite {
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

func (r *Image) load(path string) {
	var img = LoadSprite(path)
	r.picture = pixel.PictureDataFromImage(img)
}

func LoadSprite(path string) image.Image {
	file, err := os.Open(env.BasePath() + filepath.FromSlash("/static/"+path))
	fail.Check(err)
	defer func() {
		var err = file.Close()
		fail.Check(err)
	}()
	img, _, err := image.Decode(file)
	fail.Check(err)

	return img
}