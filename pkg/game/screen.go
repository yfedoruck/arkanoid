package game

import (
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/arkanoid/pkg/fail"
)

type Screen struct {
	window     *pixelgl.Window
	image      *Image
	wall       *Wall
	board      *Board
	ball       Ball
	background *Background
	level      int
	text       *Text
	beepFinish bool
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
func (r Screen) Background() *Background {
	return r.background
}

func (r *Screen) NextLevel() {
	r.ball = NewStopBall(NewBlankBall(r.window, r.board))
	r.ball.OnStartPosition()
	r.board.OnStartPosition()
	r.level++
	r.Level()
}

func (r *Screen) Level() {
	switch r.level {
	case 1:
		r.Level1()
	case 2:
		r.Level2()
	}
}

func (r *Screen) Level1() {
	r.background.Level1()
	r.wall.level1()
}

func (r *Screen) Level2() {
	r.background.Level2()
	r.wall.level2()
}

func (r Screen) NoMoreLevels() bool {
	return r.level >= 2
}

func (r Screen) listenExit() {
	if r.window.Pressed(pixelgl.KeyEscape) {
		r.window.SetClosed(true)
	}
}

func (r *Screen) playFinishOnce() {
	if !r.beepFinish {
		buffer := LoadSound("ArkanoidSFX9.wav")
		melody := buffer.Streamer(0, buffer.Len())
		speaker.Play(melody)
		r.beepFinish = true
	}
}

func NewScreen() *Screen {
	var (
		win   = NewWindow()
		image = NewImage()
	)
	board := NewBoard(win)
	background := NewBackground(win)
	background.Level1()
	ball := NewStopBall(NewBlankBall(win, board))
	return &Screen{
		window:     win,
		image:      image,
		wall:       NewWall(win, image, board),
		board:      board,
		ball:       ball,
		background: background,
		level:      1,
		text:       NewText(win),
	}
}

const (
	WinHeight = 768
	WinWidth  = 714
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
