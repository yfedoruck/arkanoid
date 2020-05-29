package game

type DownRight struct {
	BlankBall
}

func NewDownRight(bb *BlankBall) Ball {
	return &DownRight{
		BlankBall: *bb,
	}
}

func (r *DownRight) Move() Ball {

	if r.brick.IsNotHit() {
		if r.hitBrickTop() {
			return NewUpRight(CopyBlankBall(r))
		}

		if r.hitBrickLeft() {
			return NewDownLeft(CopyBlankBall(r))
		}
	}

	if r.crossBottomLine() {
		if r.hitBoard() {
			return NewUpRight(CopyBlankBall(r))
		} else {
			r.Restart()
			return NewStopBall(CopyBlankBall(r))
		}
	}
	if r.hitRightBorder() {
		return NewDownLeft(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X += r.DeltaX()
	return r
}

func (r *DownRight) MoveLeft() {
}
func (r *DownRight) MoveRight() {
}