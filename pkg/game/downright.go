package game

type DownRight struct {
	BlankBall
}

func NewDownRight(bb *BlankBall) Ball {
	return &DownRight{
		BlankBall: *bb,
	}
}

func (r *DownRight) Move() Ball {
	if r.position.Y <= (r.win.Bounds().Min.Y + r.radius) {
		return NewUpRight(CopyBlankBall(r))
	}
	if r.position.X > (r.win.Bounds().Max.X - r.radius) {
		return NewDownLeft(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X += r.DeltaX()
	return r
}

func (r *DownRight) MoveLeft() {
}
func (r *DownRight) MoveRight() {
}