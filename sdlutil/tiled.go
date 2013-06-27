package sdlutil

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"sdl"
)

// flags for tile ids
const (
	flipHorizontal   = 1 << 31
	flipVertical     = 1 << 30
	flipAntiDiagonal = 1 << 29

	tmxMask = (flipHorizontal | flipVertical | flipAntiDiagonal)
)

type tiledxml struct {
	XMLName     xml.Name `xml:"map"`
	Version     string   `xml:"version,attr"`
	Orientation string   `xml:"orientation,attr"`
	Width       int      `xml:"width,attr"`
	Height      int      `xml:"height,attr"`
	Tilewidth   int      `xml:"tilewidth,attr"`
	Tileheight  int      `xml:"tileheight,attr"`
	Properties  []struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	} `xml:"properties>property"`
	Tileset []struct {
		Firstgid   int    `xml:"firstgid,attr"`
		Name       string `xml:"name,attr"`
		Tilewidth  int    `xml:"tilewidth,attr"`
		Tileheight int    `xml:"tileheight,attr"`
		Margin     int    `xml:"margin,attr"`
		Spacing    int    `xml:"spacing,attr"`
		Image      struct {
			Source string `xml:"source,attr"`
			Trans  string `xml:"trans,attr"`
			Width  int    `xml:"width,attr"`
			Height int    `xml:"height,attr"`
		} `xml:"image"`
	} `xml:"tileset"`
	Layer []struct {
		Name    string `xml:"name,attr"`
		Width   int    `xml:"width,attr"`
		Height  int    `xml:"height,attr"`
		Visible *int   `xml:"visible,attr"`
		Data    struct {
			Encoding    string `xml:"encoding,attr"`
			Compression string `xml:"compression,attr"`
			Tile        []struct {
				Gid int `xml:"gid,attr"`
			} `xml:"tile"`
			Chardata string `xml:",chardata"`
		} `xml:"data"`
	} `xml:"layer"`
}

type tiled struct {
	tiledxml

	re       *sdl.Display
	loader   TextureLoader
	textures []*sdl.Texture
	tilesets []tileset
	tiles    []tile
}

func (ti *tiled) Decode(re *sdl.Display, r io.Reader) (*TileMap, error) {
	ti.re = re

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(b, &ti.tiledxml)
	if err != nil {
		return nil, err
	}

	ti.sortTileSet()

	err = ti.buildTextures()
	if err != nil {
		return nil, err
	}

	err = ti.buildTiles()
	if err != nil {
		return nil, err
	}

	tm := &TileMap{
		layers:   make([]tileset, len(ti.tilesets)),
		textures: ti.textures,
	}
	for i := range tm.layers {
		tm.layers[i] = ti.tilesets[i]
		tm.layers[i].prio = i
		tm.layers[i].visible = true
	}

	return tm, nil
}

func (ti *tiled) sortTileSet() {
	set := ti.Tileset[:]

	// sort in ascending order
	for i := 1; i < len(set); i++ {
		j := i
		for j > 0 && set[j].Firstgid < set[j-1].Firstgid {
			set[j], set[j-1] = set[j-1], set[j]
			j--
		}
	}
}

func (ti *tiled) buildTextures() error {
	var err error

	set := ti.Tileset[:]
	for i := range set {
		p := &set[i]

		c := int64(-1)
		if p.Image.Trans != "" {
			c, err = strconv.ParseInt(p.Image.Trans, 16, 64)
			if err != nil {
				return err
			}
		}
		t, err := ti.loader.LoadTexture(ti.re, p.Image.Source, TextureLoaderInfo{Transparent: c})
		if err != nil {
			return err
		}

		ti.textures = append(ti.textures, t)
	}
	return nil
}

func (ti *tiled) buildTiles() error {
	set := ti.Tileset[:]
	t := tile{Tile: ZT}

	for i := range ti.Layer {
		l := &ti.Layer[i]

		if l.Width <= 0 {
			l.Width = ti.Width
		}
		if l.Height <= 0 {
			l.Height = ti.Height
		}
		if l.Width <= 0 || l.Height <= 0 {
			return fmt.Errorf("tiled: invalid dimensions from layer %v", i)
		}

		tiles, err := ti.getTiles(i)
		if err != nil {
			return err
		}

		ti.tiles = make([]tile, 256)
		for _, v := range tiles {
			var j int
			var tx, ty int

			tid := v &^ tmxMask
			if tid < 0 {
				return fmt.Errorf("tiled: tile id less than 0: %v", tid)
			}
			if tid >= len(ti.tiles) {
				newtiles := make([]tile, tid*2)
				copy(newtiles, ti.tiles)
				ti.tiles = newtiles
			}

			gid, j := ti.tileGID(tid)
			if gid < 0 {
				continue
			}
			if tid < len(ti.tiles) && ti.tiles[tid].Texture != nil {
				continue
			}

			id := tid - gid
			p := &set[j]
			if p.Tilewidth <= 0 || p.Tileheight <= 0 {
				return fmt.Errorf("tiled: invalid tile dimension: %vx%v",
					p.Tilewidth, p.Tileheight)

			}

			row := p.Image.Width / (p.Tilewidth + p.Margin)
			if row != 0 {
				tx = id % row
				ty = id / row
			} else {
				tx = 0
				ty = 0
			}
			t.Texture = ti.textures[j]
			t.texcoord = sdl.Rect{
				X: sdl.Int(tx*p.Tilewidth + ((tx + 1) * p.Margin)),
				Y: sdl.Int(ty*p.Tileheight + ((ty + 1) * p.Spacing)),
				W: sdl.Int(p.Tilewidth),
				H: sdl.Int(p.Tileheight),
			}

			ti.tiles[tid] = t
		}

		ts := tileset{}
		grid := make([][]tile, l.Width*l.Height)
		ts.grid = make([][][]tile, l.Height)
		ts.tw, ts.th = ti.Tilewidth, ti.Tileheight
		for j := range ts.grid {
			ts.grid[j] = grid[j*l.Width : (j+1)*l.Width]
		}

		for j, y := 0, 0; y < l.Height; y++ {
			for x := 0; x < l.Width; x++ {
				tid := tiles[j] &^ tmxMask
				t := ti.tiles[tid]
				f := tiles[j] & tmxMask

				if f&flipHorizontal != 0 {
					t.Flip |= sdl.FLIP_HORIZONTAL
				}

				if f&flipVertical != 0 {
					t.Flip |= sdl.FLIP_VERTICAL
				}

				// yeah ok, whatever
				if f&flipAntiDiagonal != 0 {
				}

				tx := int(t.texcoord.X)
				ty := int(t.texcoord.Y)

				tpy := ty
				iy := y

				for th := int(t.texcoord.H); th > 0; th -= ti.Tileheight {
					if iy >= len(ts.grid) {
						break
					}

					tileh := min(th, ti.Tileheight)
					tpx := tx
					ix := x
					for tw := int(t.texcoord.W); tw > 0; tw -= ti.Tilewidth {
						if ix >= len(ts.grid[iy]) {
							break
						}

						tilew := min(tw, ti.Tilewidth)
						t.id = tid

						t.texcoord.X = sdl.Int(tpx)
						t.texcoord.Y = sdl.Int(tpy)
						t.texcoord.W = sdl.Int(tilew)
						t.texcoord.H = sdl.Int(tileh)

						tpx += tilew

						ts.grid[iy][ix] = append(ts.grid[iy][ix], t)
						ix++
					}
					iy++
					tpy += tileh
				}

				j++
			}
		}
		ti.tilesets = append(ti.tilesets, ts)
	}

	for i := range ti.tilesets {
		ti.tilesets[i].tiles = ti.tiles
	}

	return nil
}

func (ti *tiled) getTiles(i int) ([]int, error) {
	var tiles []int

	l := &ti.Layer[i]
	d := &l.Data
	switch d.Encoding {
	case "base64":
		var cr io.Reader
		var buf []byte

		d.Chardata = strings.Trim(d.Chardata, " \n")
		buf, err := base64.StdEncoding.DecodeString(d.Chardata)
		if err != nil {
			return nil, err
		}

		br := bufio.NewReader(bytes.NewBuffer(buf))
		switch d.Compression {
		case "gzip":
			cr, err = gzip.NewReader(br)
		case "zlib":
			cr, err = zlib.NewReader(br)
		case "":
			cr = br
		default:
			return nil, errors.New("tiled: unknown tile compression format: " + d.Compression)
		}
		if err != nil {
			return nil, err
		}

		var v uint32
		for i := 0; i < l.Width*l.Height; i++ {
			err := binary.Read(cr, binary.LittleEndian, &v)
			if err != nil {
				return nil, err
			}
			tiles = append(tiles, int(v))
		}

	case "csv":
		d.Chardata = strings.Map(func(r rune) rune {
			if strings.ContainsRune(" \t\r\n", r) {
				return -1
			}
			return r
		}, d.Chardata)

		r := csv.NewReader(bytes.NewBufferString(d.Chardata))
		sp, err := r.Read()
		if err != nil {
			return nil, err
		}
		for i := range sp {
			v, err := strconv.Atoi(sp[i])
			if err != nil {
				return nil, err
			}
			tiles = append(tiles, v)
		}

	case "":
		for _, tile := range d.Tile {
			tiles = append(tiles, tile.Gid)
		}

	default:
		return nil, errors.New("tiled: unknown tile encoding: " + d.Encoding)
	}

	if len(tiles) != l.Width*l.Height {
		return nil, fmt.Errorf("tiled: layer '%v': mismatched tile dimensions: %v (%v x %v) = %v",
			l.Name, len(tiles), l.Width, l.Height, l.Width*l.Height)
	}

	return tiles, nil
}

// finds the first gid set that the tile falls in
// and the index to set
func (ti *tiled) tileGID(tile int) (gid, index int) {
	set := ti.Tileset[:]
	for i := len(set) - 1; i >= 0; i-- {
		if set[i].Firstgid <= tile {
			return set[i].Firstgid, i
		}
	}
	return -1, -1
}

func NewTiledFile(re *sdl.Display, name string, l TextureLoader) (*TileMap, error) {
	var ti tiled

	ti.loader = l
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ti.Decode(re, r)
}
