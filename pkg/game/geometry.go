package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func startBallPosition(win *pixelgl.Window) pixel.Vec {

	var v = pixel.V(win.Bounds().W()/2, win.Bounds().H()*10/100)
	return v
}
