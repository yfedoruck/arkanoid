package game

type DownBall struct {
	BlankBall
}

func NewDownBall(bb *BlankBall) Ball {
	return &DownBall{
		BlankBall: *bb,
	}
}

func (r *DownBall) Move() Ball {
	if r.position.Y <= (BoardHeight + r.radius) {
		return NewUpBall(CopyBlankBall(r))
	}
	r.position.Y -= r.delta
	return r
}
func (r *DownBall) MoveLeft() {
}
func (r *DownBall) MoveRight() {
}
