package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/fail"
	"golang.org/x/image/colornames"
)

type Screen struct {
	window *pixelgl.Window
}

func (r Screen) Window() *pixelgl.Window {
	return r.window
}

func NewScreen() *Screen {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	fail.Check(err)

	win.SetSmooth(true)
	win.Clear(colornames.Greenyellow)
	win.SetTitle("Arkanoid")

	return &Screen{window: win}
}