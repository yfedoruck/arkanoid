package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/fail"
	"golang.org/x/image/colornames"
)

type Screen struct {
	window *pixelgl.Window
	image  *Image
	wall   *Wall
	board  *Board
	ball   Ball
}

func (r Screen) Window() *pixelgl.Window {
	return r.window
}

func (r Screen) Wall() *Wall {
	return r.wall
}

func (r Screen) Board() *Board {
	return r.board
}

func (r Screen) Ball() Ball {
	return r.ball
}

func (r Screen) Image() *Image {
	return r.image
}

func NewScreen() *Screen {
	var (
		win   = NewWindow()
		image = NewImage()
	)
	board := NewBoard(win, image)
	ball := NewStopBall(NewBlankBall(win, image, board))
	return &Screen{
		window: win,
		image:  image,
		wall:   NewWall(win, image),
		board:  board,
		ball:   ball,
	}
}

func NewWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Arkanoid",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	fail.Check(err)

	win.SetSmooth(true)
	win.Clear(colornames.Greenyellow)
	return win
}
