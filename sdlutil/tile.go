package sdlutil

import (
	"sdl"
)

var ZT = Tile{
	Texture: nil,
	ScaleX:  1,
	ScaleY:  1,
	Angle:   0,
	Center:  nil,
	Flip:    sdl.FLIP_NONE,
}

type Tile struct {
	Texture *sdl.Texture
	ScaleX  float64
	ScaleY  float64
	Angle   float64
	Center  *sdl.Point
	Flip    sdl.Flip
}

type TileSet struct {
	grid   [][][]tile
	tiles  []tile
	tw, th int
}

type tile struct {
	Tile
	id       int
	texcoord sdl.Rect
}

type tileset struct {
	TileSet
	prio    int
	visible bool
}

type TileMap struct {
	layers   []tileset
	textures []*sdl.Texture
}

func (t *Tile) Draw(re *sdl.Display, dst, src *sdl.Rect) {
	dst.W = sdl.Int(float64(dst.W) * t.ScaleX)
	dst.H = sdl.Int(float64(dst.H) * t.ScaleY)
	re.CopyEx(t.Texture, src, dst, t.Angle, t.Center, t.Flip)
}

func (ts *TileSet) Draw(re *sdl.Display, camera, viewport sdl.Rect) {
	var tile *tile
	var p, t sdl.Rect

	if viewport.X < 0 {
		camera.X -= viewport.X
		viewport.X = 0
	}
	if viewport.Y < 0 {
		camera.Y -= viewport.Y
		viewport.Y = 0
	}

	w := min(int(camera.W), int(viewport.W))
	h := min(int(camera.H), int(viewport.H))

	var tow, toh int
	tcy := int(viewport.Y)
	toy := int(viewport.Y) % ts.th

	for y := 0; y < h; {
		tcx := int(viewport.X)
		tox := int(viewport.X) % ts.tw

		for x := 0; x < w; {
			tx := tcx / ts.tw
			ty := tcy / ts.th
			tow = min(ts.tw-tox, w-x)
			toh = min(ts.th-toy, h-y)

			if ty < 0 || ty >= len(ts.grid) {
				goto incx
			}
			if tx < 0 || tx >= len(ts.grid[ty]) {
				goto incx
			}

			for i := range ts.grid[ty][tx] {
				tile = &ts.grid[ty][tx][i]
				if tile.Texture == nil {
					continue
				}

				tilew := min(tow, int(tile.texcoord.W))
				tileh := min(toh, int(tile.texcoord.H))
				t = sdl.Rect{
					X: tile.texcoord.X + sdl.Int(tox),
					Y: tile.texcoord.Y + sdl.Int(toy),
					W: sdl.Int(tilew),
					H: sdl.Int(tileh),
				}

				p = sdl.Rect{
					X: camera.X + sdl.Int(x),
					Y: camera.Y + sdl.Int(y),
					W: sdl.Int(tilew),
					H: sdl.Int(tileh),
				}
				tile.Draw(re, &p, &t)
			}

		incx:
			x += tow
			tcx += ts.tw
			tox = 0
		}

		y += toh
		tcy += ts.th
		toy = 0
	}
}

func (tm *TileMap) DrawLayer(layer int, re *sdl.Display, camera, viewport sdl.Rect) {
	tm.layers[layer].Draw(re, camera, viewport)
}

func (tm *TileMap) Draw(re *sdl.Display, camera, viewport sdl.Rect) {
	for i := range tm.layers {
		if !tm.layers[i].visible {
			continue
		}
		tm.DrawLayer(i, re, camera, viewport)
	}
}

func (tm *TileMap) Free() {
	for _, t := range tm.textures {
		t.Destroy()
	}
}

func (tm *TileMap) NumLayers() int {
	return len(tm.layers)
}
