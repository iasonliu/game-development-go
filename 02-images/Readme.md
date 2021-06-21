https://github.com/hajimehoshi/file2byteslice

To load the image

```
file2byteslice -input ./coin.png -output assets.go -package main -var coinImg
```
this command above will generate the assets.go file:
```
package main

var cionImg = []byte("...")
```

## Images form files

once the assets have been generated, the image can be created during initialization:
```
import _ "image/png"
var coin *ebiten.Image

func init(){
	img, _, _ := image.Decode(bytes.NewReader(coinImg))
	coin, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}
```

## Draw Images
The Draw function just moves the images to the center of the screen:

```
func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	cw, ch := coin.Size()
	sw, sh := screen.Size()
	// Move half of the screen size on the right/botton and
	// half of the image size on the left/top
	op.GeoM.Translate(float64(sw/2 - cw/2), float64(sh/2 - ch/2))
	screen.DrawImage(coin, op)
}
```

## using generate cmd

```
$ cat generate.go
//go:generate file2byteslice -input ./coin.png  -output assets.go -package main -var coinImg
package main

$ go generate .
```
