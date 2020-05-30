package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BgMaxX    = 230.0
	BgMinY    = 256.0
	BgWidth   = 230.0
	BgHeight  = 240.0
	BgBorderX = 8.0
	BgBorderY = 7.0
	BgScale   = 3.2
)

type Background struct {
	picture  *pixel.PictureData
	sprite   *pixel.Sprite
	win      *pixelgl.Window
	position pixel.Vec
	mat      pixel.Matrix
}

func NewBackground(win *pixelgl.Window) *Background {
	var bg = &Background{
		win:      win,
		position: pixel.V(BgWidth*BgScale/2, BgHeight*BgScale/2),
		mat:      pixel.IM.ScaledXY(pixel.ZV, pixel.V(BgScale, BgScale)),
	}
	bg.load("fields.png")
	bg.Sprite()
	return bg
}

func (r *Background) Sprite() {
	r.sprite = pixel.NewSprite(r.picture, pixel.R(0, BgMinY, BgMaxX, r.picture.Bounds().Max.Y))
}

func (r *Background) Draw() {
	r.sprite.Draw(r.win, r.mat.Moved(r.position))
}

func (r *Background) load(path string) {
	var img = LoadSprite(path)
	r.picture = pixel.PictureDataFromImage(img)
}
