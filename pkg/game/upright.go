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
	if r.position.Y >= (r.win.Bounds().Max.Y - r.radius) {
		return NewDownRight(CopyBlankBall(r))
	}
	if r.position.X >= (r.win.Bounds().Max.X - r.radius) {
		return NewUpLeft(CopyBlankBall(r))
	}
	r.position.Y += r.delta
	r.position.X += r.delta/5
	return r
}

func (r *UpRight) MoveLeft() {
}
func (r *UpRight) MoveRight() {
}