package wrapper

import (
	"image"
	"log"
	"os"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Texture struct {
	image *ebiten.Image
}

func NewTexture(filepath string) *Texture {
	infile, err := os.Open(filepath)
	if err != nil {
		logrus.Error(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		log.Fatal(err)
	}

	origEbitenImage := ebiten.NewImageFromImage(img)

	w, h := origEbitenImage.Size()
	ebitenImage := ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	ebitenImage.DrawImage(origEbitenImage, op)

	return &Texture{
		image: ebitenImage,
	}
}

func (tex *Texture) Draw(screen *ebiten.Image, position, scale mgl32.Vec2, rotation float64) {
	drawOptions := getImageDrawOptions(position, scale, rotation)
	screen.DrawImage(tex.image, &drawOptions)
}

func getImageDrawOptions(position, scale mgl32.Vec2, rotation float64) ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(float64(position.X()), float64(position.Y()))
	return op
}
