package game

type UpBall struct {
	BlankBall
}

func NewUpBall(bb *BlankBall) Ball {
	return &UpBall{
		BlankBall: *bb,
	}
}

func (r *UpBall) Move() Ball {
	if r.position.Y >= (r.win.Bounds().Max.Y - r.radius) {
		return NewDownBall(CopyBlankBall(r))
	}
	r.position.Y += r.delta
	r.position.X += r.delta/5
	return r
}

func (r *UpBall) MoveLeft() {
}
func (r *UpBall) MoveRight() {
}