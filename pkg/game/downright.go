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
	if r.position.Y <= (r.win.Bounds().Min.Y + r.radius + r.board.height) {
		//fmt.Println(r.board.Area())
		//fmt.Println(r.win.Bounds().Min.Y + r.radius + r.board.height)
		//fmt.Println("ball position", r.position.Y)
		//fmt.Println("ball radius", r.radius)
		//fmt.Println("board height", r.board.height)

		if r.board.Area().X1 <= r.position.X && r.position.X <= r.board.Area().X2 {
			return NewUpRight(CopyBlankBall(r))
		} else {
			return NewPause(CopyBlankBall(r))
		}
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