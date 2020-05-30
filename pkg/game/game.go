package game

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/jpeg"
	_ "image/png"
	"time"
)

func Run() {
	var (
		screen = NewScreen()
		win    = screen.Window()
		delta  = 0.0
		last   = time.Now()
		fps    = time.Tick(time.Second / 60)
	)
	board := screen.Board()
	ball := screen.Ball()
	board.OnStartPosition()
	ball.OnStartPosition()

	wall := screen.Wall()
	bg := NewBackground(win)
	for !win.Closed() {
		var dt = time.Since(last).Seconds()
		last = time.Now()

		delta = dt * 500
		ball.SetDelta(delta)
		win.Clear(colornames.Firebrick)
		//win.Clear(colornames.Firebrick)
		bg.Draw()

		if win.Pressed(pixelgl.KeyLeft) {
			board.MoveLeft(delta)
			ball.MoveLeft()
		}
		if win.Pressed(pixelgl.KeyRight) {
			board.MoveRight(delta)
			ball.MoveRight()
		}

		if win.Pressed(pixelgl.KeySpace) {
			ball.Push()
		}
		if ball.IsPushed() {
			ball = ball.Move(wall)
		}

		board.Draw()
		ball.Draw()
		wall.Draw()

		win.Update()
		<-fps
	}
}
