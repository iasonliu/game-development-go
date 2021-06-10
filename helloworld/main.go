package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

// Game implements the ebiten.Game interface
type Game struct {
}

func (g *Game) Update() error {
	return nil
}

// Draw is optional, but suggested to maintain the logic of the Game Loop
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World!!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World!!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
