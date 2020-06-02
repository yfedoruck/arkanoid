package game

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/jpeg"
	_ "image/png"
	"time"
)

func Run() {
	screen := NewScreen()
	screen.Run()
}

func (r *Screen) Run() {
	var (
		win   = r.Window()
		delta = 0.0
		last  = time.Now()
		fps   = time.Tick(time.Second / 60)
	)
	board := r.Board()
	board.OnStartPosition()
	r.ball.OnStartPosition()

	for !win.Closed() {
		var dt = time.Since(last).Seconds()
		last = time.Now()

		delta = dt * 500
		r.ball.SetDelta(delta)
		r.background.Draw()

		if win.Pressed(pixelgl.KeyLeft) {
			board.MoveLeft(delta)
			r.ball.MoveLeft()
		}
		if win.Pressed(pixelgl.KeyRight) {
			board.MoveRight(delta)
			r.ball.MoveRight()
		}

		if win.Pressed(pixelgl.KeySpace) {
			r.ball.Push()
		}

		if r.wall.IsDestroyed() {
			if r.NoMoreLevels() {
				win.Clear(colornames.Black)
			} else {
				r.NextLevel()
			}
		}

		if r.ball.IsPushed() {
			r.ball = r.ball.Move(r.wall)
		}

		board.Draw()
		r.ball.Draw()
		r.wall.Draw()

		win.Update()
		<-fps
	}
}
