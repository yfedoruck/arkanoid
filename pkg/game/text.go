package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yfedoruck/arkanoid/pkg/fail"
	"golang.org/x/image/font/basicfont"
)

const TextScale = 3

type Text struct {
	window   *pixelgl.Window
	atlas    *text.Atlas
}

func NewText(win *pixelgl.Window) *Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	return &Text{
		window:   win,
		atlas:    atlas,
	}
}

func (r Text) Draw(s string) {
	basicTxt := text.New(pixel.V(100, 500), r.atlas)
	_, err := fmt.Fprintln(basicTxt, s)
	fail.Check(err)
	basicTxt.Draw(r.window, pixel.IM.Scaled(basicTxt.Orig, TextScale))
}
