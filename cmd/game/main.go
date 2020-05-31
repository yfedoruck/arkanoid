package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/arkanoid/pkg/game"
)

func main() {
	pixelgl.Run(game.Run)
}
