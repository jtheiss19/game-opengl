package engine

import (
	"errors"
	"image/png"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/sirupsen/logrus"
)

type Texture struct {
	textureID glObjectReference
	filePath  string
}

func NewTexture(textureFilePath string) *Texture {
	logrus.Info("creating new texture")
	newTexutre, err := loadTextureAlpha(textureFilePath)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return newTexutre
}

func (tex *Texture) Use() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(tex.textureID))
}

func loadTextureAlpha(filePath string) (*Texture, error) {
	logrus.Info("loading texture")
	infile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("can not open texture file: " + err.Error())
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		return nil, errors.New("can not decode texture file: " + err.Error())
	}

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}

	var texID uint32
	// create texutre
	gl.GenTextures(1, &texID)

	// set as the using texture gl.TEXTURE_2D
	gl.BindTexture(gl.TEXTURE_2D, texID)

	// configure the current gl.TEXTURE_2D texture settings
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	// put data into the buffer
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	// data optimization
	gl.GenerateMipmap(gl.TEXTURE_2D)

	// unbind
	gl.BindTexture(gl.TEXTURE_2D, 0)

	returnTexture := &Texture{
		textureID: glObjectReference(texID),
		filePath:  filePath,
	}
	return returnTexture, nil
}
