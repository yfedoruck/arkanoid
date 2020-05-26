package game

type DownLeft struct {
	BlankBall
}

func NewDownLeft(bb *BlankBall) Ball {
	return &DownLeft{
		BlankBall: *bb,
	}
}

func (r *DownLeft) Move() Ball {
	if r.position.Y < (r.win.Bounds().Min.Y + r.radius) {
		return NewUpLeft(CopyBlankBall(r))
	}
	if r.position.X <= (r.win.Bounds().Min.X + r.radius) {
		return NewDownRight(CopyBlankBall(r))
	}
	r.position.Y -= r.DeltaY()
	r.position.X -= r.DeltaX()
	return r
}

func (r *DownLeft) MoveLeft() {
}
func (r *DownLeft) MoveRight() {
}