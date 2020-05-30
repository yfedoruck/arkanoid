package game

type DownLeft struct {
	BlankBall
}

func NewDownLeft(bb *BlankBall) Ball {
	return &DownLeft{
		BlankBall: *bb,
	}
}

func (r *DownLeft) Move(wall *Wall) Ball {
	for _, brick := range wall.wall {
		if brick.IsNotHit() {
			if r.hitBrickTop(brick) {
				return NewUpLeft(CopyBlankBall(r))
			}

			if r.hitBrickRight(brick) {
				return NewDownRight(CopyBlankBall(r))
			}
		}
	}

	if r.crossBottomLine() {
		if r.hitBoard() {
			return NewUpLeft(CopyBlankBall(r))
		} else {
			r.Restart()
			return NewStopBall(CopyBlankBall(r))
		}
	}
	if r.hitLeftBorder() {
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
