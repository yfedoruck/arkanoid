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
	if r.hitBrickTop() {
		return NewUpLeft(CopyBlankBall(r))
	}

	if r.crossBottomLine() {
		if r.hitBoard() {
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
