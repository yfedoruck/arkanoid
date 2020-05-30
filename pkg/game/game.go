package game

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/jpeg"
	_ "image/png"
	"time"
)

func Run() {
	//sprite := NewSprite()
	//var boardSprite = sprite.Board()

	//var center = win.Bounds().Center()
	var (
		screen = NewScreen()
		win    = screen.Window()
		delta  = 0.0
		last   = time.Now()
		fps    = time.Tick(time.Second / 60)
		bb     = NewBlankBall(win)
	)
	//var center = win.Bounds().Center()
	board := NewBoard(screen)
	ball := NewStopBall(bb)
	ball.Connect(&board)
	board.OnStartPosition()
	ball.OnStartPosition()

	brick := NewBrick(win)
	//wall := NewWall(win, &ball)
	brick.OnStartPosition()
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
			ball = ball.Move(brick)
			ball.Connect(&board)
			//ball.SetBrick(brick)
			//fmt.Printf("%T", ball)
			//return
		}

		board.Draw()
		ball.Draw()
		//wall.Draw()
		if brick.IsNotHit() {
			brick.Draw()
		}

		//boardSprite.Draw(win, pixel.IM.Moved(vec))
		//ballSprite.Draw(win, pixel.IM.Moved(board.StartBallPosition(ball)))

		win.Update()
		<-fps
	}
}
