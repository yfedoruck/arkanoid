package game

type Pause struct {
	BlankBall
}

func NewPause(bb *BlankBall) Ball {
	return &Pause{
		BlankBall: *bb,
	}
}

func (r *Pause) Move() Ball {
	return NewPause(CopyBlankBall(r))
}
