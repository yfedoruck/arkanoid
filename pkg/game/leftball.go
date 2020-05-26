package game

type LeftBall struct {
	BlankBall
}

func NewLeftBall(bb *BlankBall) Ball {
	return &LeftBall{
		BlankBall: *bb,
	}
}

func (r *LeftBall) Move() Ball {
	if r.position.X <= r.radius {
		return NewRightBall(CopyBlankBall(r))
	}
	r.position.X -= r.delta
	return r
}

func (r *LeftBall) MoveLeft() {
}
func (r *LeftBall) MoveRight() {
}