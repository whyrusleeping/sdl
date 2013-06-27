package sdl

func (re *Renderer) DrawCircle(xc, yc, r int) {
	x := 0
	y := r
	p := 1 - r
	re.circlePlot(xc, yc, x, y)

	for x < y {
		x++
		if p < 0 {
			p += 2*x + 1
		} else {
			y--
			p += 2*(x-y) + 1
		}
		re.circlePlot(xc, yc, x, y)
	}
}

func (re *Renderer) circlePlot(xc, yc, x, y int) {
	re.DrawPoint(xc+x, yc+y)
	re.DrawPoint(xc-x, yc+y)
	re.DrawPoint(xc+x, yc-y)
	re.DrawPoint(xc-x, yc-y)
	re.DrawPoint(xc+y, yc+x)
	re.DrawPoint(xc-y, yc+x)
	re.DrawPoint(xc+y, yc-x)
	re.DrawPoint(xc-y, yc-x)
}

func round(x float64) int {
	return int(x + 0.5)
}

func (re *Renderer) DrawEllipse(xc, yc, rx, ry int) {
	rx2, ry2 := rx*rx, ry*ry
	fry := float64(ry)
	frx2, fry2 := float64(rx2), float64(ry2)
	x := 0
	y := ry
	px := 0
	py := 2 * rx2 * y

	re.ellipsePlot(xc, yc, x, y)
	p := round(fry2 - frx2*fry + (0.25 * frx2))
	for px < py {
		x++
		px += 2 * ry2
		if p < 0 {
			p += ry2 + px
		} else {
			y--
			py -= 2 * rx2
			p += ry2 + px - py
		}

		re.ellipsePlot(xc, yc, x, y)
	}

	fx := float64(x)
	fy := float64(y)

	p = round(fry2*(fx+0.5)*(fx+0.5) + frx2*(fy-1)*(fy-1) - frx2*fry2)
	for y > 0 {
		y--
		py -= 2 * rx2
		if p > 0 {
			p += rx2 - py
		} else {
			x++
			px += 2 * ry2
			p += rx2 - py + px
		}
		re.ellipsePlot(xc, yc, x, y)
	}
}

func (re *Renderer) ellipsePlot(xc, yc, x, y int) {
	re.DrawPoint(xc+x, yc+y)
	re.DrawPoint(xc-x, yc+y)
	re.DrawPoint(xc+x, yc-y)
	re.DrawPoint(xc-x, yc-y)
}
