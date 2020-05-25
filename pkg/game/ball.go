package game

import (
	"github.com/faiface/pixel"
)



type RedBall struct {
	BlankBall
	state    Ball
}

func (r *RedBall) OnStartPosition(board Board) {
	bp := board.StartPosition()
	r.position = pixel.V(bp.X, bp.Y+r.radius)
}

func (r *RedBall) SetDelta(delta float64) {
	r.delta = delta
}

func (r *RedBall) MoveLeft() {
	r.position.X -= r.delta
}

func (r *RedBall) MoveRight() {
	r.position.X += r.delta
}

func (r *RedBall) Push() {
	r.pushed = true
}

func (r RedBall) IsPushed() bool {
	return r.pushed
}

func (r RedBall) Move() Ball {
	//if r.position.Y <= r.win.Bounds().Max.Y {
	//	r.position.Y += r.delta
	//	return &UpBall{}
	//} else {
	//	return &DownBall{}
	//}
	return r.state.Move()
}

func (r *RedBall) MoveUp() {
	r.position.Y += r.delta
}

func (r *RedBall) MoveDown() {
	r.position.Y -= r.delta
}

func (r RedBall) Draw() {
	r.sprite.Draw(r.win, pixel.IM.Moved(r.position))
}

func (r RedBall) Radius() float64 {
	return r.radius
}
func (r RedBall) Diameter() float64 {
	return r.radius * 2
}
