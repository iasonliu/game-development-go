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
type Game struct{}

// Update() updates the game logic by 1 tick(60 ticks per second)
func (g *Game) Update() error {
	return nil
}

// Draw is optional, but suggested to maintain the logic of the Game Loop
// Draw() draws the screen based on the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World!!")
}

// Layout() gets the outside size(like the window size) and returns the game logical screen size
// can be fixed or can perform calculations to adapt the game to user's device size
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
