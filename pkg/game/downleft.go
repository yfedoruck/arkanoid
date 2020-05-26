package game

type DownLeft struct {
	BlankBall
}

func NewDownLeft(bb *BlankBall) Ball {
	return &DownLeft{
		BlankBall: *bb,
	}
}

func (r *DownLeft) Move() Ball {
	if r.position.Y < (r.win.Bounds().Min.Y + r.radius + r.board.height) {
		if r.board.Area().X1 <= r.position.X && r.position.X <= r.board.Area().X2 {
			return NewUpLeft(CopyBlankBall(r))
		} else {
			r.Restart()
			return NewStopBall(CopyBlankBall(r))
		}
	}
	if r.position.X <= (r.win.Bounds().Min.X + r.radius) {
		return NewDownRight(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X -= r.DeltaX()
	return r
}

func (r *DownLeft) MoveLeft() {
}
func (r *DownLeft) MoveRight() {
}
