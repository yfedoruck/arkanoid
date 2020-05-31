package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/fail"
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
	board := NewBoard(win)
	ball := NewStopBall(NewBlankBall(win, image, board))
	return &Screen{
		window: win,
		image:  image,
		wall:   NewWall(win, image),
		board:  board,
		ball:   ball,
	}
}

const (
	WinHeight = 768
	WinWidth = 714
)

func NewWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Arkanoid",
		Bounds: pixel.R(0, 0, WinWidth, WinHeight),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	fail.Check(err)

	win.SetSmooth(true)
	return win
}
