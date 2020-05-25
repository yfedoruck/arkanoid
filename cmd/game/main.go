package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/yfedoruck/billiards/pkg/game"
)

func main() {
	pixelgl.Run(game.Run)
}
