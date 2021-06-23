package main

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var coins *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}
	coins = ebiten.NewImageFromImage(img)
}

type framesSpec struct {
	Frames []frameSpec `json:"frames"`
}

type frameSpec struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Game struct {
	tick      uint64
	speed     float64
	frames    []frameSpec
	numFrames int
}

func (g *Game) Update() error {
	g.tick++
	return nil
}

func (g *Game) buildFrames(path string) error {
	j, _ := ioutil.ReadFile(path)
	fSpec := &framesSpec{}
	json.Unmarshal(j, fSpec)
	g.frames = fSpec.Frames
	g.numFrames = len(g.frames)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	frameNum := int(g.tick/uint64(g.speed)) % g.numFrames
	f := g.frames[frameNum]
	x, y := screen.Size()
	tx := x/2 - f.W/2
	ty := y/2 - f.H/2
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(tx), float64(ty))
	rect := image.Rect(f.X, f.Y, f.X+f.W, f.Y+f.H)
	subImg := coins.SubImage(rect).(*ebiten.Image)
	screen.DrawImage(subImg, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Coins!!!")
	g := &Game{
		speed: 60 / 10,
	}
	g.buildFrames("./coins.json")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
