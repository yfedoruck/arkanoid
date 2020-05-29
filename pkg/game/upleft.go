package game

type UpLeft struct {
	BlankBall
}

func NewUpLeft(bb *BlankBall) Ball {
	return &UpLeft{
		BlankBall: *bb,
	}
}

func (r *UpLeft) Move() Ball {

	if r.hitCeil() || r.hitBrickBottom() {
		return NewDownLeft(CopyBlankBall(r))
	}

	if r.position.Y >= (r.win.Bounds().Max.Y - r.radius) {
		return NewDownLeft(CopyBlankBall(r))
	}
	if r.position.X < r.win.Bounds().Min.X + r.radius {
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