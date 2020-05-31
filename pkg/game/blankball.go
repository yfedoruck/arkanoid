package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BallRadius = 12
	BallScale   = 3
)

type Ball interface {
	Move(wall *Wall) Ball
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
	Board() *Board
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

func NewBlankBall(win *pixelgl.Window, board *Board) *BlankBall {
	sp := BallSprite()
	return &BlankBall{
		radius:   BallSprite().Picture().Bounds().H()*BallScale/2,
		win:      win,
		position: pixel.ZV,
		sprite:   sp,
		pushed:   false,
		delta:    0.0,
		board:    board,
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
		board:    b.Board(),
	}
}

func (r *BlankBall) OnStartPosition() {
	bp := r.board.StartPosition()
	r.position = pixel.V(bp.X, bp.Y+r.board.Height()/2 + r.radius)
}

func (r *BlankBall) Delta() float64 {
	return r.delta
}
func (r *BlankBall) SetDelta(delta float64) {
	r.delta = delta
}

func (r *BlankBall) MoveLeft() {
	if r.position.X <= r.win.Bounds().Min.X+r.board.Width()/2+(BgBorderX*BgScale) {
		return
	}
	r.position.X -= r.delta
}

func (r *BlankBall) MoveRight() {
	if r.position.X >= r.win.Bounds().Max.X-r.board.Width()/2-(BgBorderX*BgScale) {
		return
	}
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
	mat := pixel.IM
	mat = mat.Scaled(pixel.ZV, BallScale)
	r.sprite.Draw(r.win, mat.Moved(r.position))
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

func (r *BlankBall) Board() *Board {
	return r.board
}

func (r BlankBall) hitRightBorder() bool {
	return r.position.X >= (r.win.Bounds().Max.X - r.radius - (BgBorderX*BgScale))
}

func (r BlankBall) hitLeftBorder() bool {
	return r.position.X <= (r.win.Bounds().Min.X + r.radius + (BgBorderX*BgScale))
}

func (r BlankBall) hitCeil() bool {
	return r.position.Y >= (r.win.Bounds().Max.Y - r.radius - (BgBorderY*BgScale))
}

func (r BlankBall) hitBoard() bool {
	return r.board.Area().X1 <= r.position.X && r.position.X <= r.board.Area().X2
}

func (r BlankBall) crossBottomLine() bool {
	return r.position.Y <= (r.win.Bounds().Min.Y + r.radius + r.board.Height())
}

func (r BlankBall) hitBrickBottom(brick *Brick) bool {
	if r.isAboveBrick(brick) {
		return false
	}
	side := brick.Bottom()
	if (r.top() >= side.Y) &&
		(side.X1 <= r.right() && r.left() <= side.X2) {
		fmt.Println("hitBrickBottom")
		brick.Delete()
		return true
	}

	return false
}

func (r BlankBall) isAboveBrick(brick *Brick) bool {
	return r.position.Y > brick.position.Y
}

func (r BlankBall) isUnderBrick(brick *Brick) bool {
	return r.position.Y < brick.position.Y
}

func (r BlankBall) isBeforeBrick(brick *Brick) bool {
	return r.position.X < brick.position.X
}
func (r BlankBall) isAfterBrick(brick *Brick) bool {
	return r.position.X > brick.position.X
}

func (r BlankBall) hitBrickTop(brick *Brick) bool {
	if r.isUnderBrick(brick) {
		return false
	}
	side := brick.Top()
	if (r.bottom() <= side.Y) &&
		(side.X1 <= r.right() && r.left() <= side.X2) {
		fmt.Println("hitBrickTop")
		brick.Delete()
		return true
	}

	return false
}

func (r BlankBall) hitBrickLeft(brick *Brick) bool {
	if r.isAfterBrick(brick) {
		//fmt.Println("isAfterBrick")
		return false
	}
	side := brick.Left()
	if (r.right() >= side.X) &&
		(side.Y1 <= r.bottom() && r.top() <= side.Y2) {
		fmt.Println("hitBrickLeft")
		brick.Delete()
		return true
	}

	return false
}

func (r BlankBall) hitBrickRight(brick *Brick) bool {
	if r.isBeforeBrick(brick) {
		//fmt.Println("isBeforeBrick")
		return false
	}
	side := brick.Right()
	if (r.left() <= side.X) &&
		(side.Y1 <= r.bottom() && r.top() <= side.Y2) {
		fmt.Println("hitBrickRight")
		brick.Delete()
		return true
	}

	return false
}

func (r BlankBall) top() float64 {
	return r.position.Y + r.radius
}
func (r BlankBall) bottom() float64 {
	return r.position.Y - r.radius
}
func (r BlankBall) right() float64 {
	return r.position.X + r.radius
}
func (r BlankBall) left() float64 {
	return r.position.X - r.radius
}

func (r *BlankBall) Restart() {
	r.Stop()
	r.position = pixel.ZV
	r.OnStartPosition()
	r.board.OnStartPosition()
}

func BallSprite() *pixel.Sprite {
	var picture = pixel.PictureDataFromImage(LoadSprite("ball.png"))
	return pixel.NewSprite(picture, picture.Bounds())
}
