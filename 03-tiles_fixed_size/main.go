package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	// size in pixels, square img
	imgSize = 16
	// number of frames in the spreadsheet
	numFrames = 8
)

var coins *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}
	coins = ebiten.NewImageFromImage(img)
}

type Game struct {
	tick  uint64
	speed float64
}

func (g *Game) Update() error {
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(sw/2-1), float64(sh/2-1))
	frameNum := uint64(g.tick/uint64(g.speed)) % numFrames
	// move right in the spreadsheet
	frameX := int(frameNum * imgSize)
	rect := image.Rect(frameX, 0, frameX+imgSize, imgSize)
	subImg := coins.SubImage(rect)

	screen.DrawImage(subImg.(*ebiten.Image), opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Coins!!!")
	g := &Game{
		speed: 60 / 6,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
