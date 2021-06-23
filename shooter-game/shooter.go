package shooter

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iasonliu/game-development-go/shooter-game/objects"
)

const (
	windowWidth  = 800
	windowHeight = 600
	maxUint      = ^uint(0)
)

type Game struct {
	tick    uint
	objects []objects.Object
	screen  *ebiten.Image
}

func (g *Game) Update() error {
	g.tick++
	if g.tick == maxUint {
		g.tick = 0
	}
	for _, obj := range g.objects {
		obj.Tick(g.tick)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, obj := range g.objects {
		if err := obj.Draw(screen); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{
		objects: []objects.Object{
			objects.NewBackground("bg_green.png"),
			objects.NewLevel1("water1.png", 4),
			objects.NewDesk("bg_wood.png"),
			objects.NewCurtains("curtain_straight.png", "curtain.png"),
		},
	}
	return g
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
