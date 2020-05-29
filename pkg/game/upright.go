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
	if r.hitCeil() || r.hitBrickBottom() {
		return NewDownRight(CopyBlankBall(r))
	}
	if r.hitRightWall() {
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
