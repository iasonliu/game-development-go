package utils

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iasonliu/game-development-go/shooter-game/assets"
)

func GetImage(name string, obj *assets.Object) (*ebiten.Image, error) {
	var rect image.Rectangle
	for _, img := range obj.Specs.Images {
		if img.Name == name {
			rect = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
			break
		}
	}
	img := obj.Image.SubImage(rect).(*ebiten.Image)
	return img, nil
}
