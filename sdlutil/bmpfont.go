package sdlutil

import (
	"fmt"
	"image/color"
	"sdl"
)

type BMPChar struct {
	Texcoord sdl.Rect
	Texture  *sdl.Texture
	Width    int
	Height   int
}

type BMPKern struct {
	Left, Right rune
}

type BMPFontInfo struct {
	Chars         map[rune]BMPChar
	Kern          map[BMPKern]sdl.Point
	Tabstop       int
	Width, Height int
}

type BMPFont struct {
	chars         map[rune]BMPChar
	kern          map[BMPKern]sdl.Point
	col           color.RGBA
	width, height int
	tabstop       int
}

func NewBMPFont(bi BMPFontInfo) *BMPFont {
	return &BMPFont{
		chars:   bi.Chars,
		kern:    bi.Kern,
		width:   bi.Width,
		height:  bi.Height,
		tabstop: bi.Tabstop,
		col:     White,
	}
}

func (b *BMPFont) SetColor(c color.RGBA) {
	b.col = c
}

func (b *BMPFont) Printf(re *sdl.Display, x, y int, format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	tile := ZT

	px, py := x, y
	for _, c := range text {
		cb, found := b.chars[c]
		if !found {
			cb, found = b.chars['?']
		}

		if c == ' ' || c == '\t' || !found {
			if c == '\t' {
				px += b.width * b.tabstop
			} else {
				px += b.width
			}
			continue
		}
		if c == '\n' {
			px = x
			py += b.height
			continue
		}

		pos := sdl.Rect{
			X: sdl.Int(px),
			Y: sdl.Int(py),
			W: sdl.Int(cb.Width),
			H: sdl.Int(cb.Height),
		}

		tile.Texture = cb.Texture
		tile.Texture.SetColorMod(b.col)
		tile.Draw(re, &pos, &cb.Texcoord)
		px += cb.Width
	}
}
