package main

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 320
	screenHeight = 240
	sampleText   = `hello, Gophers!!!`
	dpi          = 72
	fontSize     = 36
)

var mplusNormalFont font.Face

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// calculate the rectangle containing the text
	bounds := text.BoundString(mplusNormalFont, sampleText)
	// write moving the text down by its height
	text.Draw(screen, sampleText, mplusNormalFont, 10, bounds.Dy(), color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func init() {
	tt, err := truetype.Parse(Font)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World!!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
