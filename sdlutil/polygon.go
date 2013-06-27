package sdlutil

import (
	"fmt"
	"image/color"

	"sdl"
)

type polygon struct {
	ymin, ymax int
	edges      []*edge
	fill       func(ymin, ymax int, x1, x2 float32, scan int)
}

type edge struct {
	yUpper                int
	xIntersect, dxPerScan float32
	next                  *edge
}

func Poly(re *sdl.Display, texture *sdl.Texture, pts []sdl.Point, c color.RGBA) {
	var pl polygon
	re.SetDrawColor(c)
	re.SetTarget(texture)
	pl.fill = func(ymin, ymax int, x1, x2 float32, scan int) {
		re.DrawPoint(int(x1), scan)
		re.DrawPoint(int(x2), scan)
	}
	pl.scanFill(pts)
	re.SetTarget(nil)
}

func FillPoly(re *sdl.Display, texture *sdl.Texture, pts []sdl.Point, c color.RGBA) {
	var pl polygon
	re.SetDrawColor(c)
	re.SetTarget(texture)
	pl.fill = func(ymin, ymax int, x1, x2 float32, scan int) {
		re.DrawLine(int(x1), scan, int(x2), scan)
	}
	pl.scanFill(pts)
	re.SetTarget(nil)
}

func (pl *polygon) scanFill(pts []sdl.Point) {
	if len(pts) < 3 {
		panic(fmt.Errorf("%v points, not a polygon", len(pts)))
	}
	pl.ymin, pl.ymax = -1, -1
	for i := range pts {
		y := int(pts[i].Y)
		if y < pl.ymin || pl.ymin < 0 {
			pl.ymin = y
		}
		if y > pl.ymax || pl.ymax < 0 {
			pl.ymax = y + 1
		}
	}

	if pl.ymin < 0 || pl.ymax < 0 {
		panic(fmt.Errorf("invalid ymin/ymax bound: %v %v", pl.ymin, pl.ymax))
	}
	pl.edges = make([]*edge, pl.ymax)
	for i := range pl.edges {
		pl.edges[i] = new(edge)
	}
	pl.buildEdgeList(pts)
	active := new(edge)

	for scan := pl.ymin; scan < pl.ymax; scan++ {
		pl.buildActiveList(scan, active)
		if active.next != nil {
			pl.fillScan(scan, active)
			pl.updateActiveList(scan, active)
			pl.resortActiveList(active)
		}
	}
}

func (pl *polygon) fillScan(scan int, active *edge) {
	p1 := active.next
	for p1 != nil {
		p2 := p1.next
		pl.fill(pl.ymin, pl.ymax, p1.xIntersect, p2.xIntersect, scan)
		p1 = p2.next
	}
}

func (pl *polygon) resortActiveList(active *edge) {
	p := active.next
	active.next = nil
	for p != nil {
		q := p.next
		pl.insertEdge(active, p)
		p = q
	}
}

func (pl *polygon) deleteAfter(q *edge) {
	p := q.next
	q.next = p.next
}

func (pl *polygon) updateActiveList(scan int, active *edge) {
	q, p := active, active.next
	for p != nil {
		if scan >= int(p.yUpper) {
			p = p.next
			pl.deleteAfter(q)
		} else {
			p.xIntersect += p.dxPerScan
			q, p = p, p.next
		}
	}
}

func (pl *polygon) buildActiveList(scan int, active *edge) {
	p := pl.edges[scan].next
	for p != nil {
		q := p.next
		pl.insertEdge(active, p)
		p = q
	}
}

func (pl *polygon) buildEdgeList(pts []sdl.Point) {
	cnt := len(pts)

	yPrev := pts[cnt-2].Y
	v1 := pts[cnt-1]
	for i := 0; i < cnt; i++ {
		v2 := pts[i]
		if v1.Y != v2.Y {
			if v1.Y < v2.Y {
				pl.makeEdge(v1, v2, pl.yNext(i, pts))
			} else {
				pl.makeEdge(v2, v1, yPrev)
			}
		}
		yPrev = v1.Y
		v1 = v2
	}
}

func (pl *polygon) yNext(k int, pts []sdl.Point) sdl.Int {
	var i int

	cnt := len(pts)
	if k+1 <= cnt-1 {
		i = k + 1
	}
	for pts[k].Y == pts[i].Y {
		i = (i + 1) % cnt
	}
	return pts[i].Y
}

func (pl *polygon) makeEdge(lower, upper sdl.Point, yComp sdl.Int) {
	e := new(edge)
	e.dxPerScan = float32(upper.X-lower.X) / float32(upper.Y-lower.Y)
	e.xIntersect = float32(lower.X)
	e.yUpper = int(upper.Y)
	if upper.Y < yComp {
		e.yUpper--
	}
	pl.insertEdge(pl.edges[lower.Y], e)
}

func (pl *polygon) insertEdge(list, e *edge) {
	q := list
	p := q.next
	for p != nil {
		if e.xIntersect < p.xIntersect {
			p = nil
		} else {
			q, p = p, p.next
		}
	}
	e.next = q.next
	q.next = e
}
