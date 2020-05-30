package game

type StopBall struct {
	BlankBall
}

func NewStopBall(bb *BlankBall) Ball {
	return &StopBall{
		BlankBall: *bb,
	}
}

func (r *StopBall) Move(wall *Wall) Ball {
	return NewUpRight(CopyBlankBall(r))
}
