package sdlutil

import (
	"image"
	"image/color"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func transparent(img image.Image, c uint32) image.Image {
	p := image.NewNRGBA(img.Bounds())
	b := p.Bounds()

	cr, cg, cb := uint8(c&0xFF), uint8((c>>8)&0xFF), uint8((c>>16)&0xFF)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			col := img.At(x, y)
			r32, g32, b32, _ := col.RGBA()
			r, g, b := uint8(r32), uint8(g32), uint8(b32)
			if cr == r && cg == g && cb == b {
				p.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 0})
			} else {
				p.Set(x, y, col)
			}
		}
	}
	return p
}