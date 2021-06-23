package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	debouncer    = 100 * time.Millisecond
)

type Game struct {
	lastClickAt  time.Time // 0-value of time is 0001-01-01 00:00:00 +0000 UTC
	x, y         int
	currentColor int
}

var colors = []color.Color{
	color.RGBA{0, 0, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0xff, 0xff},
	color.RGBA{0, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Now().Sub(g.lastClickAt) > debouncer {
		log.Printf("A pressed")
		g.lastClickAt = time.Now()
		g.currentColor = (g.currentColor + 1) % len(colors)
	}
	g.x, g.y = ebiten.CursorPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("press A to switch background color. Mouse: %s", g.mousePosition())
	screen.Fill(colors[g.currentColor])
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) mousePosition() string {
	return fmt.Sprintf("(%d, %d)", g.x, g.y)
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
