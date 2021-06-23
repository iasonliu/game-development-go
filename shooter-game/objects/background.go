package objects

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iasonliu/game-development-go/shooter-game/assets"
	"github.com/iasonliu/game-development-go/shooter-game/utils"
)

type background struct {
	name string
}

func NewBackground(imgname string) Object {
	return &background{
		name: imgname,
	}
}

func (b *background) Tick(tick uint) {}

func (b *background) Draw(trgt *ebiten.Image) error {
	img, err := utils.GetImage(b.name, assets.Stall)
	if err != nil {
		return fmt.Errorf("drawing %s: %v", b.name, err)
	}
	trgtW, trgtH := trgt.Size()
	bgW, bgH := img.Size()
	x := int(math.Ceil(float64(trgtW) / float64(bgW)))
	y := int(math.Ceil(float64(trgtH) / float64(bgH)))
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			opts := &ebiten.DrawImageOptions{}
			tx := i * bgW
			ty := j * bgH
			opts.GeoM.Translate(float64(tx), float64(ty))
			trgt.DrawImage(img, opts)
		}
	}
	return nil
}

func (b *background) OnScreen() bool {
	return true
}
