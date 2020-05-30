package game

type Pause struct {
	BlankBall
}

func NewPause(bb *BlankBall) Ball {
	return &Pause{
		BlankBall: *bb,
	}
}

func (r *Pause) Move(wall *Wall) Ball {
	return NewPause(CopyBlankBall(r))
}
