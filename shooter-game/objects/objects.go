package objects

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Tick(uint)
	Draw(*ebiten.Image) error
}
