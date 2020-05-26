package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/fail"
	"golang.org/x/image/colornames"
)

func NewWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	fail.Check(err)

	win.SetSmooth(true)
	win.Clear(colornames.Greenyellow)

	return win
}