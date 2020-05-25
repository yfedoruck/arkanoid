package game

import (
	"github.com/faiface/pixel"
	"github.com/yfedoruck/billiards/pkg/env"
	"github.com/yfedoruck/billiards/pkg/fail"
	"image"
	"os"
	"path/filepath"
)

type Sprite struct {
	picture *pixel.PictureData
	ball    *pixel.Sprite
}

func NewSprite() *Sprite {
	var s = &Sprite{}
	s.loadPicture("BasicArkanoidPack.png")
	return s
}

func (r Sprite) Ball() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 32, 24, 33+24))
}
func (r Sprite) Board() *pixel.Sprite {
	return pixel.NewSprite(r.picture, pixel.R(0, 0, 160, 25))
}

func (r *Sprite) loadPicture(path string) {
	file, err := os.Open(env.BasePath() + filepath.FromSlash("/static/"+path))
	fail.Check(err)
	defer func() {
		var err = file.Close()
		fail.Check(err)
	}()
	img, _, err := image.Decode(file)
	fail.Check(err)
	r.picture = pixel.PictureDataFromImage(img)
}
