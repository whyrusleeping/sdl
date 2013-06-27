package sdlutil

import (
	"fmt"
	"image"
	"os"

	"sdl"
)

type TextureLoader interface {
	LoadTexture(re *sdl.Display, name string, info TextureLoaderInfo) (*sdl.Texture, error)
}

type TextureLoaderInfo struct {
	Transparent int64
}

type defaultLoader struct{}

var DefaultLoader *defaultLoader

func (l *defaultLoader) LoadTexture(re *sdl.Display, name string, info TextureLoaderInfo) (*sdl.Texture, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", name, err)
	}

	if info.Transparent >= 0 {
		img = transparent(img, uint32(info.Transparent))
	}
	t, err := re.CreateTextureFromImage(img)
	if err != nil {
		return nil, err
	}
	return t, nil
}
