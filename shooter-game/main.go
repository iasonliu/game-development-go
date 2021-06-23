package main

import (
	"log"
	"math/rand"
	"time"

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

func main() {
	rand.Seed(time.Now().Unix())
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Shooter")
	g := &Game{
		objects: []objects.Object{
			objects.NewBackground("bg_green.png"),
			objects.NewLevel1("water1.png", 5),
			objects.NewDesk("bg_wood.png"),
			objects.NewCurtains("curtain_straight.png", "curtain.png"),
		},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
