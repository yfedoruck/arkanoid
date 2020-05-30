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
		bb     = NewBlankBall(screen)
	)
	board := NewBoard(screen)
	ball := NewStopBall(bb)
	ball.Connect(&board)
	board.OnStartPosition()
	ball.OnStartPosition()

	wall := NewWall(screen)
	for !win.Closed() {
		var dt = time.Since(last).Seconds()
		last = time.Now()

		delta = dt * 500
		ball.SetDelta(delta)
		win.Clear(colornames.Firebrick)

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
			ball.Connect(&board)
		}

		board.Draw()
		ball.Draw()
		wall.Draw()

		win.Update()
		<-fps
	}
}
