package main

import (
	"bytes"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	audioContext *audio.Context
	click        *audio.Player
)

const (
	screenWidth  = 320
	screenHeight = 240
	debouncer    = 100 * time.Millisecond
)

type Game struct {
	lastClickAt time.Time
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Now().Sub(g.lastClickAt) > debouncer {
		log.Printf("A pressed\n")
		click.Rewind()
		click.Play()
		g.lastClickAt = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Click A to play a sound")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	audioContext = audio.NewContext(44100)
	oggS, _ := vorbis.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(RagtimeSound))
	s := audio.NewInfiniteLoop(oggS, oggS.Length())

	background, _ := audio.NewPlayer(audioContext, s)

	click = audio.NewPlayerFromBytes(audioContext, ClickSound)
	background.Play()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World!!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
