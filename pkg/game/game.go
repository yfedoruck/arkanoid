package game

import (
	"github.com/faiface/pixel/pixelgl"
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
	bg := screen.Background()
	for !win.Closed() {
		var dt = time.Since(last).Seconds()
		last = time.Now()

		delta = dt * 500
		ball.SetDelta(delta)
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
