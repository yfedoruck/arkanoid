package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	BrickWidth  = 64.0
	BrickHeight = 32.0
	BrickGap    = 8.0
)

const (
	SimpleBrick = iota
	GunBrick
	GlueBrick
)

type BrickSpec int

func (r BrickSpec) String() string {
	return [...]string{"simple", "gun", "glue"}[r]
}

type Brick struct {
	width    float64
	height   float64
	position pixel.Vec
	sprite   *pixel.Sprite
	live     bool
	spec     BrickSpec
}

type BrickSideY struct {
	X, Y1, Y2 float64
}
type BrickSideX struct {
	X1, X2, Y float64
}

func NewBrick(image *Image, color int) *Brick {
	return &Brick{
		position: pixel.ZV,
		sprite:   image.Brick(color),
		width:    BrickWidth,
		height:   BrickHeight,
		live:     true,
		spec:     SimpleBrick,
	}
}

func (r Brick) Left() BrickSideY {
	return BrickSideY{
		X:  r.position.X - r.width/2,
		Y1: r.position.Y - r.height/2,
		Y2: r.position.Y + r.height/2,
	}
}

func (r Brick) Right() BrickSideY {
	return BrickSideY{
		X:  r.position.X + r.width/2,
		Y1: r.position.Y - r.height/2,
		Y2: r.position.Y + r.height/2,
	}
}

func (r Brick) Top() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.height/2,
		Y:  r.position.Y + r.height/2,
	}
}
func (r Brick) Bottom() BrickSideX {
	return BrickSideX{
		X1: r.position.X - r.width/2,
		X2: r.position.X + r.width/2,
		Y:  r.position.Y - r.height/2,
	}
}

func (r *Brick) Delete() {
	r.live = false
}

func (r *Brick) SetSpec(spec BrickSpec) {
	r.spec = spec
}

func (r Brick) HasGift() bool {
	return r.spec != SimpleBrick
}

func (r Brick) IsNotHit() bool {
	return r.live
}

func (r Brick) Draw(win *pixelgl.Window) {
	r.sprite.Draw(win, pixel.IM.Moved(r.position))
}

func (r *Brick) MoveTo(pos pixel.Vec) {
	r.position = pos
}
