package objects

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iasonliu/game-development-go/shooter-game/assets"
	"github.com/iasonliu/game-development-go/shooter-game/utils"
)

const (
	duckName        = "duck_outline_target_white.png"
	stickName       = "stick_woodFixed_outline.png"
	ducksXSpeed     = 1.8 // horizontal speed
	ducksYSpeed     = 0.6 // vertical speed
	ducksMaxOffsetY = 16  // max vertical movement for animation
)

type duck struct {
	duckImg        *ebiten.Image
	stickImg       *ebiten.Image
	duckW          int
	duckH          int
	offsetX        float64   // horizontal position
	offsetY        float64   // vertical position
	initialOffsetY float64   // initial vertical position. set by the caller
	yDirection     direction // whether the vertical animation is going up or down
	onScreen       bool      // true when the image is visible in the screen
}

// newDuck generates a new duck with an initial vertical position
func newDuck(initialOffsetY int) *duck {
	duckImg, err := utils.GetImage(duckName, assets.Objects)
	stickImg, err := utils.GetImage(stickName, assets.Objects)
	if err != nil {
		log.Fatalf("drawing %s: %v", duckName, err)
	}

	w, h := duckImg.Size()

	return &duck{
		duckImg:        duckImg,
		stickImg:       stickImg,
		duckW:          w,
		duckH:          h,
		initialOffsetY: float64(initialOffsetY),
		offsetX:        float64(-w),
		offsetY:        0,
		yDirection:     down,
		onScreen:       true,
	}
}

func (d *duck) shoot(clickX, clickY int) bool {
	x := int(d.offsetX)
	y := int(d.offsetY + d.initialOffsetY)

	if clickX >= x && clickY <= x+d.duckW && clickY >= y && clickY <= y+d.duckH {
		d.onScreen = false
		return true
	}
	return false
}

func (d *duck) Tick(_ uint) {
	// horizontal movement
	d.offsetX = d.offsetX + ducksXSpeed

	// calculate the vertical direction and offset (for animation)
	if ducksMaxOffsetY-math.Abs(d.offsetY) < 0 {
		d.yDirection = d.yDirection.invert()
	}
	d.offsetY = d.offsetY + float64(d.yDirection)*ducksYSpeed
}

func (d *duck) Draw(trgt *ebiten.Image) error {
	// when the duck is over the screen size, it's no more visible
	screenW := 800
	if d.offsetX > float64(screenW) {
		d.onScreen = false
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.offsetX, d.offsetY+d.initialOffsetY)
	trgt.DrawImage(d.duckImg, op)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(d.offsetX+float64(d.duckW/4+10), d.offsetY+float64(d.duckH)+d.initialOffsetY)
	trgt.DrawImage(d.stickImg, opt)
	return nil
}

func (d *duck) OnScreen() bool {
	return d.onScreen
}
