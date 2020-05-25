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
	if r.position.Y >= r.win.Bounds().Max.Y {
		return NewDownBall(CopyBlankBall(r))
	}
	r.position.Y += r.delta
	return r
}
