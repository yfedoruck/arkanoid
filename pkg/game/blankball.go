package game

import (
	"github.com/faiface/beep"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BallRadius = 12
	BallScale  = 3
)

type Ball interface {
	Move(wall *Wall) Ball
	OnStartPosition()
	MoveLeft()
	MoveRight()
	Start()
	SetDelta(delta float64)
	IsStarted() bool
	Draw()
	Radius() float64
	Win() *pixelgl.Window
	Position() pixel.Vec
	Sprite() *pixel.Sprite
	CleanSprite()
	Delta() float64
	DeltaX() float64
	DeltaY() float64
	Board() *Board
	BeepHitBoard()
	WavHitBoard() *beep.Buffer
	WavHitBrick() *beep.Buffer
}

type BlankBall struct {
	radius      float64
	win         *pixelgl.Window
	position    pixel.Vec
	sprite      *pixel.Sprite
	started     bool
	delta       float64
	board       *Board
	wavHitBoard *beep.Buffer
	wavHitBrick *beep.Buffer
}

func NewBlankBall(win *pixelgl.Window, board *Board) *BlankBall {
	sp := SimpleSprite("ball.png")
	return &BlankBall{
		radius:      sp.Picture().Bounds().H() * BallScale / 2,
		win:         win,
		position:    pixel.ZV,
		sprite:      sp,
		started:     false,
		delta:       0.0,
		board:       board,
		wavHitBoard: LoadSound("ArkanoidSFX6.wav"),
		wavHitBrick: LoadSound("ArkanoidSFX7.wav"),
	}
}

func CopyBlankBall(b Ball) *BlankBall {
	return &BlankBall{
		radius:      b.Radius(),
		win:         b.Win(),
		position:    b.Position(),
		sprite:      b.Sprite(),
		started:     b.IsStarted(),
		delta:       b.Delta(),
		board:       b.Board(),
		wavHitBoard: b.WavHitBoard(),
		wavHitBrick: b.WavHitBrick(),
	}
}

func (r *BlankBall) OnStartPosition() {
	bp := r.board.StartPosition()
	r.position = pixel.V(bp.X, bp.Y+r.board.Height()/2+r.radius)
}

func (r *BlankBall) Delta() float64 {
	return r.delta
}
func (r *BlankBall) SetDelta(delta float64) {
	r.delta = delta
}

func (r *BlankBall) MoveLeft() {
	if r.IsStarted(){
		return
	}
	if r.position.X <= r.win.Bounds().Min.X+r.board.Width()/2+(BgBorderX*BgScale) {
		return
	}
	r.position.X -= r.delta
}

func (r *BlankBall) MoveRight() {
	if r.IsStarted(){
		return
	}
	if r.position.X >= r.win.Bounds().Max.X-r.board.Width()/2-(BgBorderX*BgScale) {
		return
	}
	r.position.X += r.delta
}

func (r *BlankBall) Start() {
	r.started = true
}

func (r *BlankBall) Stop() {
	r.started = false
}

func (r BlankBall) IsStarted() bool {
	return r.started
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
	return r.position.X >= (r.win.Bounds().Max.X - r.radius - (BgBorderX * BgScale))
}

func (r BlankBall) hitLeftBorder() bool {
	return r.position.X <= (r.win.Bounds().Min.X + r.radius + (BgBorderX * BgScale))
}

func (r BlankBall) hitCeil() bool {
	return r.position.Y >= (r.win.Bounds().Max.Y - r.radius - (BgBorderY * BgScale))
}

func (r BlankBall) hitBoard() bool {
	var isHit = r.board.Area().X1 <= r.position.X && r.position.X <= r.board.Area().X2
	if isHit {
		r.BeepHitBoard()
	}
	return isHit
}
func (r BlankBall) WavHitBrick() *beep.Buffer {
	return r.wavHitBrick
}

func (r BlankBall) WavHitBoard() *beep.Buffer {
	return r.wavHitBoard
}

func (r BlankBall) BeepHitBoard() {
	PlaySound(r.wavHitBoard)
}

func (r BlankBall) BeepHitBrick() {
	PlaySound(r.wavHitBrick)
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
		return true
	}

	return false
}

func (r BlankBall) hitBrickLeft(brick *Brick) bool {
	if r.isAfterBrick(brick) {
		return false
	}
	side := brick.Left()
	if (r.right() >= side.X) &&
		(side.Y1 <= r.bottom() && r.top() <= side.Y2) {
		return true
	}

	return false
}

func (r BlankBall) hitBrickRight(brick *Brick) bool {
	if r.isBeforeBrick(brick) {
		return false
	}
	side := brick.Right()
	if (r.left() <= side.X) &&
		(side.Y1 <= r.bottom() && r.top() <= side.Y2) {
		return true
	}

	return false
}

func (r BlankBall) hitBrick(brick *Brick) bool {
	return r.hitBrickRight(brick) ||
		r.hitBrickLeft(brick) ||
		r.hitBrickTop(brick) ||
		r.hitBrickBottom(brick)
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
	r.board.Simple()
	r.board.OnStartPosition()
}

func (r *BlankBall) CleanSprite() {
	r.sprite = TransparentPixel()
}
