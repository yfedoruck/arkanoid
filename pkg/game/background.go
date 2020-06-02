package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BgMaxX    = 232.0
	BgMinY    = 256.0
	BgWidth   = 230.0
	BgHeight  = 240.0
	BgBorderX = 8.0
	BgBorderY = 8.0
	BgScale   = 3.2
	BgGap     = 8.0
)

type Background struct {
	picture  *pixel.PictureData
	sprite   *pixel.Sprite
	win      *pixelgl.Window
	position pixel.Vec
	mat      pixel.Matrix
}

func NewBackground(win *pixelgl.Window) *Background {
	var img = LoadSprite("fields.png")
	return &Background{
		win:      win,
		position: pixel.V(BgWidth*BgScale/2, BgHeight*BgScale/2),
		mat:      pixel.IM.ScaledXY(pixel.ZV, pixel.V(BgScale, BgScale)),
		picture:  pixel.PictureDataFromImage(img),
	}
}

func (r *Background) Level1() {
	r.sprite = pixel.NewSprite(r.picture, pixel.R(0, BgMinY, BgMaxX, r.picture.Bounds().Max.Y))
}

func (r *Background) Level2() {
	minX := BgMaxX
	r.sprite = pixel.NewSprite(r.picture, pixel.R(minX, BgMinY, minX + BgMaxX, r.picture.Bounds().Max.Y))
}

func (r *Background) Draw() {
	r.sprite.Draw(r.win, r.mat.Moved(r.position))
}
