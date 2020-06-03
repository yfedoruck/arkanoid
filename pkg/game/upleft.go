package game

type UpLeft struct {
	BlankBall
}

func NewUpLeft(bb *BlankBall) Ball {
	return &UpLeft{
		BlankBall: *bb,
	}
}

func (r *UpLeft) Move(wall *Wall) Ball {
	for _, brick := range wall.wall {
		if r.hitBrick(brick) {
			r.BeepHitBrick()
			brick.Delete()

			switch {
			case r.hitBrickRight(brick):
				return NewUpRight(CopyBlankBall(r))
			case r.hitBrickBottom(brick):
				return NewDownLeft(CopyBlankBall(r))
			}
		}
	}

	if r.hitCeil() {
		return NewDownLeft(CopyBlankBall(r))
	}

	if r.hitLeftBorder() {
		return NewUpRight(CopyBlankBall(r))
	}
	r.position.Y += r.DeltaY()
	r.position.X -= r.DeltaX()
	return r
}

func (r *UpLeft) MoveLeft() {
}
func (r *UpLeft) MoveRight() {
}
