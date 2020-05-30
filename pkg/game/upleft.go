package game

type UpLeft struct {
	BlankBall
}

func NewUpLeft(bb *BlankBall) Ball {
	return &UpLeft{
		BlankBall: *bb,
	}
}

func (r *UpLeft) Move(brick *Brick) Ball {

	if brick.IsNotHit() {
		if r.hitBrickRight(brick) {
			return NewUpRight(CopyBlankBall(r))
		}

		if r.hitBrickBottom(brick) {
			return NewDownLeft(CopyBlankBall(r))
		}
	}

	if r.hitCeil() {
		return NewDownLeft(CopyBlankBall(r))
	}

	if r.hitLeftBorder() {
		return NewUpRight(CopyBlankBall(r))
	}
	r.position.Y += r.DeltaY()
	r.position.X -= r.DeltaX()
	return r
}

func (r *UpLeft) MoveLeft() {
}
func (r *UpLeft) MoveRight() {
}