package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/fail"
	"golang.org/x/image/colornames"
	_ "image/jpeg"
	_ "image/png"
	"time"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	fail.Check(err)

	win.SetSmooth(true)

	//sprite := NewSprite()
	//var boardSprite = sprite.Board()

	win.Clear(colornames.Greenyellow)

	//var center = win.Bounds().Center()
	var (
		delta = 0.0
		last  = time.Now()
		fps   = time.Tick(time.Second / 60)
		bb = NewBlankBall(win)
	)
	//var center = win.Bounds().Center()
	board := NewBoard(win)
	//ball := NewRedBall(bb)
	ball := NewStopBall(bb)
	board.OnStartPosition()
	ball.OnStartPosition(board)
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
			ball = ball.Move()
			//fmt.Printf("%T", ball)
			//return
		}

		board.Draw()
		ball.Draw()

		//boardSprite.Draw(win, pixel.IM.Moved(vec))
		//ballSprite.Draw(win, pixel.IM.Moved(board.StartBallPosition(ball)))

		win.Update()
		<-fps
	}
}
