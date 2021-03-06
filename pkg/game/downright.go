package game

type DownRight struct {
	BlankBall
}

func NewDownRight(bb *BlankBall) Ball {
	return &DownRight{
		BlankBall: *bb,
	}
}

func (r *DownRight) Move(wall *Wall) Ball {
	for _, brick := range wall.wall {
		if r.hitBrick(brick) {
			r.BeepHitBrick()
			brick.Delete()

			switch {
			case r.hitBrickTop(brick):
				return NewUpRight(CopyBlankBall(r))
			case r.hitBrickLeft(brick):
				return NewDownLeft(CopyBlankBall(r))
			}
		}
	}

	if r.crossBottomLine() {
		if r.hitBoard() {
			if r.board.IsSticky(){
				r.Stop()
			}
			return NewUpRight(CopyBlankBall(r))
		} else {
			r.Restart()
			return NewStartBall(CopyBlankBall(r))
		}
	}
	if r.hitRightBorder() {
		return NewDownLeft(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X += r.DeltaX()
	return r
}
