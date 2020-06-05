package game

type StartBall struct {
	BlankBall
}

func NewStartBall(bb *BlankBall) Ball {
	return &StartBall{
		BlankBall: *bb,
	}
}

func (r *StartBall) Move(wall *Wall) Ball {
	return NewUpRight(CopyBlankBall(r))
}
