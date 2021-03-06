package game

type DownLeft struct {
	BlankBall
}

func NewDownLeft(bb *BlankBall) Ball {
	return &DownLeft{
		BlankBall: *bb,
	}
}

func (r *DownLeft) Move(wall *Wall) Ball {
	for _, brick := range wall.wall {
		if r.hitBrick(brick) {
			r.BeepHitBrick()
			brick.Delete()

			switch {
			case r.hitBrickTop(brick):
				return NewUpLeft(CopyBlankBall(r))
			case r.hitBrickRight(brick):
				return NewDownRight(CopyBlankBall(r))
			}
		}
	}

	if r.crossBottomLine() {
		if r.hitBoard() {
			if r.board.IsSticky(){
				r.Stop()
			}
			return NewUpLeft(CopyBlankBall(r))
		} else {
			r.Restart()
			return NewStartBall(CopyBlankBall(r))
		}
	}
	if r.hitLeftBorder() {
		return NewDownRight(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X -= r.DeltaX()
	return r
}
