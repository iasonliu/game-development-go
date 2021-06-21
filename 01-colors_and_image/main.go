package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Game implements the ebiten.Game interface
type Game struct {
}

// Update() updates the game logic by 1 tick(60 ticks per second)
func (g *Game) Update() error {
	return nil
}

// Draw is optional, but suggested to maintain the logic of the Game Loop
// Draw() draws the screen based on the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with red
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})

	// Create a smaller green image (screen/2) and draw it on top of the above
	w, h := screen.Size()
	i1 := ebiten.NewImage(w/2, h/2)
	i1.Fill(color.RGBA{0, 0xff, 0, 0xff})
	screen.DrawImage(i1, nil)
	i1w, i1h := i1.Size()

	// Create an even smaller blue rectangle, a bit transparent. Then move and rotate it before
	// drawing it over the screen image
	i2 := ebiten.NewImage(w/3, h/3)
	i2.Fill(color.RGBA{0, 0, 0xff, 0x88})
	opts := &ebiten.DrawImageOptions{}
	// Translate (0, 0) is the top-left corner
	opts.GeoM.Translate(float64(i1w), float64(i1h))
	opts.GeoM.Rotate(0.5)
	opts.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(i2, opts)

	// Extra exercise: draw i2 over i1, not over screen
	// Tip: changing screen to i1 at line 34 isn't enough
	// because i1 has already been disposed
}

// Layout() gets the outside size(like the window size) and returns the game logical screen size
// can be fixed or can perform calculations to adapt the game to user's device size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World!!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
