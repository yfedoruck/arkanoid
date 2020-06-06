package game

import (
	"github.com/faiface/pixel"
	"github.com/yfedoruck/arkanoid/pkg/env"
	"github.com/yfedoruck/arkanoid/pkg/fail"
	"image"
	"os"
	"path/filepath"
)

type Image struct {
	picture *pixel.PictureData
}

func NewImage(file string) *Image {
	var s = &Image{}
	s.load(file)
	return s
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

func TransparentPixel() *pixel.Sprite{
	var picture = pixel.PictureDataFromImage(LoadSprite("1x1.png"))
	return pixel.NewSprite(picture, picture.Bounds())
}

func SimpleSprite(file string) *pixel.Sprite {
	var picture = pixel.PictureDataFromImage(LoadSprite(file))
	return pixel.NewSprite(picture, picture.Bounds())
}
