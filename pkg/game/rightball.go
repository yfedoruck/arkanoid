package game

type RightBall struct {
	BlankBall
}

func NewRightBall(bb *BlankBall) Ball {
	return &RightBall{
		BlankBall: *bb,
	}
}

func (r *RightBall) Move() Ball {
	if r.position.X >= (r.win.Bounds().Max.X - r.radius) {
		return NewLeftBall(CopyBlankBall(r))
	}
	r.position.X += r.delta
	return r
}

func (r *RightBall) MoveLeft() {
}
func (r *RightBall) MoveRight() {
}