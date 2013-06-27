package sdlutil

import (
	"fmt"
	"image/color"
	"sdl"
)

type XFont struct {
	pix     []byte
	col     color.RGBA
	w, h    int
	rw      int
	tabstop int
}

type XFontInfo struct {
	Pix           []byte
	Width, Height int
	Rwidth        int
	Tabstop       int
}

var (
	XFont16 = XFontInfo{xfont16, 9, 16, 8, tabstop}
	XFont32 = XFontInfo{xfont32, 18, 32, 16, tabstop}
)

func NewXFont(xf XFontInfo) *XFont {
	return &XFont{
		pix:     xf.Pix,
		col:     White,
		w:       xf.Width,
		h:       xf.Height,
		rw:      xf.Rwidth,
		tabstop: xf.Tabstop,
	}
}

func (ft *XFont) SetColor(c color.RGBA) {
	ft.col = c
}

func (ft *XFont) Printf(re *sdl.Display, x, y int, format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	re.SetDrawColor(ft.col)

	fontwidth := (ft.rw + 7) / 8
	for _, c := range text {
		if c == '\n' {
			y += ft.h
			continue
		} else if c == '\t' {
			y += ft.w * ft.tabstop
			continue
		}
		l := int(c) * ft.h * fontwidth
		if l > len(ft.pix) {
			c = '?'
		}
		p := ft.pix[int(c)*ft.h*fontwidth:]
		for py := 0; py < ft.h; py++ {
			px := 0
			for i := 0; i < fontwidth; i++ {
				bit := p[i]
				for i := uint(0); i < 8; i++ {
					if bit&(1<<(7-i)) != 0 {
						re.DrawPoint(x+px, y+py)
					}
					px++
				}
			}
			p = p[fontwidth:]
		}
		x += ft.w
	}
}
