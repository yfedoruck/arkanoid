package game

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"time"
)

func Run() {
	screen := NewScreen()
	screen.Run()
}

func (r *Screen) Run() {
	var (
		win   = r.window
		delta = 0.0
		last  = time.Now()
		fps   = time.Tick(time.Second / 60)
	)
	r.board.OnStartPosition()
	r.ball.OnStartPosition()
	r.Level1()

	for !win.Closed() {
		var dt = time.Since(last).Seconds()
		last = time.Now()

		delta = dt * 500
		r.ball.SetDelta(delta)
		r.wall.SetDelta(delta)
		r.board.SetDelta(delta)
		r.background.Draw()

		if win.Pressed(pixelgl.KeyLeft) {
			r.board.MoveLeft(delta)
			r.ball.MoveLeft()
		}
		if win.Pressed(pixelgl.KeyRight) {
			r.board.MoveRight(delta)
			r.ball.MoveRight()
		}

		if win.Pressed(pixelgl.KeySpace) {
			r.ball.Start()
		}

		if r.wall.IsDestroyed() {
			r.wall.Clean()
			if r.NoMoreLevels() {
				win.Clear(colornames.Black)
				r.board.CleanSprite()
				r.ball.CleanSprite()
				r.text.Draw("You win!\nPress ESC to exit")
				r.playFinishOnce()
				r.listenExit()
			} else {
				r.NextLevel()
			}
		}

		if r.ball.IsStarted() {
			r.ball = r.ball.Move(r.wall)
		}
		if r.board.IsGun() || r.board.IsBig() {
			r.ball.Start()
		}

		r.board.Run(r.wall)
		r.ball.Draw()
		r.wall.Draw(delta)

		win.Update()
		<-fps
	}
}
