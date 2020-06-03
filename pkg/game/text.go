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
	basicTxt *text.Text
}

func NewText(win *pixelgl.Window) *Text {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	return &Text{
		window:   win,
		atlas:    atlas,
		basicTxt: text.New(pixel.V(100, 500), atlas),
	}
}

func (r Text) Draw(s string) {
	_, err := fmt.Fprintln(r.basicTxt, s)
	fail.Check(err)
	r.basicTxt.Draw(r.window, pixel.IM.Scaled(r.basicTxt.Orig, TextScale))
	r.basicTxt.Clear()
}
