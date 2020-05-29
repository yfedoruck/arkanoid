package game

type UpRight struct {
	BlankBall
}

func NewUpRight(bb *BlankBall) Ball {
	return &UpRight{
		BlankBall: *bb,
	}
}

func (r *UpRight) Move() Ball {
	if r.brick.IsNotHit() {
		if r.hitBrickLeft() {
			return NewUpLeft(CopyBlankBall(r))
		}

		if r.hitBrickBottom() {
			return NewDownRight(CopyBlankBall(r))
		}
	}

	if r.hitCeil() {
		return NewDownRight(CopyBlankBall(r))
	}

	if r.hitRightBorder() {
		return NewUpLeft(CopyBlankBall(r))
	}
	r.position.Y += r.DeltaY()
	r.position.X += r.DeltaX()
	return r
}

func (r *UpRight) MoveLeft() {
}
func (r *UpRight) MoveRight() {
}
