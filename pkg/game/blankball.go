package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BallRadius = 12
)

type Ball interface {
	Move() Ball
	OnStartPosition()
	MoveLeft()
	MoveRight()
	Push()
	SetDelta(delta float64)
	IsPushed() bool
	Draw()
	Radius() float64
	Win() *pixelgl.Window
	Position() pixel.Vec
	Sprite() *pixel.Sprite
	Delta() float64
	DeltaX() float64
	DeltaY() float64
	Connect(b *Board)
	//Restart() Ball
}

type BlankBall struct {
	radius   float64
	win      *pixelgl.Window
	position pixel.Vec
	sprite   *pixel.Sprite
	pushed   bool
	delta    float64
	board    *Board
}

func NewBlankBall(win *pixelgl.Window) *BlankBall {
	return &BlankBall{
		radius:   BallRadius,
		win:      win,
		position: pixel.ZV,
		sprite:   NewSprite().Ball(),
		pushed:   false,
		delta:    0.0,
	}
}

func CopyBlankBall(b Ball) *BlankBall {
	return &BlankBall{
		radius:   b.Radius(),
		win:      b.Win(),
		position: b.Position(),
		sprite:   b.Sprite(),
		pushed:   b.IsPushed(),
		delta:    b.Delta(),
	}
}

func (r *BlankBall) OnStartPosition() {
	bp := r.board.StartPosition()
	r.position = pixel.V(bp.X, bp.Y+r.Diameter())
}

func (r *BlankBall) Delta() float64 {
	return r.delta
}
func (r *BlankBall) SetDelta(delta float64) {
	r.delta = delta
}

func (r *BlankBall) MoveLeft() {
	r.position.X -= r.delta
}

func (r *BlankBall) MoveRight() {
	r.position.X += r.delta
}

func (r *BlankBall) Push() {
	r.pushed = true
}

func (r *BlankBall) Stop() {
	r.pushed = false
}

func (r BlankBall) IsPushed() bool {
	return r.pushed
}

func (r *BlankBall) MoveUp() {
	r.position.Y += r.delta
}

func (r *BlankBall) MoveDown() {
	r.position.Y -= r.delta
}

func (r BlankBall) Draw() {
	r.sprite.Draw(r.win, pixel.IM.Moved(r.position))
}

func (r BlankBall) Radius() float64 {
	return r.radius
}
func (r BlankBall) Diameter() float64 {
	return r.radius * 2
}

func (r *BlankBall) Win() *pixelgl.Window {
	return r.win
}

func (r *BlankBall) Position() pixel.Vec {
	return r.position
}
func (r *BlankBall) Sprite() *pixel.Sprite {
	return r.sprite
}

func (r *BlankBall) DeltaX() float64 {
	return r.delta / 3
}

func (r *BlankBall) DeltaY() float64 {
	return r.delta
}

func (r *BlankBall) Connect(b *Board) {
	r.board = b
}

func (r *BlankBall) Restart() {
	r.Stop()
	r.position = pixel.ZV
	r.OnStartPosition()
	r.board.OnStartPosition()
}
