package sdl

import (
	"strconv"
)

// String returns a string representation of p like "(3,4)".
func (p Point) String() string {
	return "(" + strconv.Itoa(int(p.X)) + "," + strconv.Itoa(int(p.Y)) + ")"
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Mul returns the vector p*k.
func (p Point) Mul(k int) Point {
	return Point{p.X * Int(k), p.Y * Int(k)}
}

// Div returns the vector p/k.
func (p Point) Div(k int) Point {
	return Point{p.X / Int(k), p.Y / Int(k)}
}

// In returns whether p is in r.
func (p Point) In(r Rect) bool {
	return r.X <= p.X && p.X < r.X+r.W &&
		r.Y <= p.Y && p.Y < r.Y+r.H
}

// Mod returns the point q in r such that p.X-q.X is a multiple of r's width
// and p.Y-q.Y is a multiple of r's height.
func (p Point) Mod(r Rect) Point {
	w, h := r.W, r.H
	p = p.Sub(Point{r.X, r.Y})
	p.X = p.X % w
	if p.X < 0 {
		p.X += w
	}
	p.Y = p.Y % h
	if p.Y < 0 {
		p.Y += h
	}
	return p.Add(Point{r.X, r.Y})
}

// Eq returns whether p and q are equal.
func (p Point) Eq(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

// ZP is the zero Point.
var ZP Point
