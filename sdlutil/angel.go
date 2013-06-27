package sdlutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"sdl"
)

type angel struct {
	info struct {
		Face     string
		Size     int
		Bold     int
		Italic   int
		Charset  string
		Unicode  int
		StretchH int
		Smooth   int
		Aa       int
		Padding  []int
		Spacing  int
		Outline  int
	}
	common struct {
		LineHeight int
		Base       int
		ScaleW     int
		ScaleH     int
		Pages      int
	}
	pages []angelPage
	chars []angelChar

	re *sdl.Display
}

type angelPage struct {
	Id   int
	File string
}

type angelChar struct {
	Id               int
	X, Y             int
	Width, Height    int
	Xoffset, Yoffset int
	Xadvance         int
	Page             int
}

func (a *angel) Decode(r io.Reader) (*BMPFont, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		line, tag := a.getToken(line)
		if len(tag) == 0 {
			continue
		}

		switch tag {
		case "info":
			a.parse(line, &a.info)
		case "common":
			a.parse(line, &a.common)
		case "page":
			a.pages = append(a.pages, angelPage{})
			a.parse(line, &a.pages[len(a.pages)-1])
		case "char":
			a.chars = append(a.chars, angelChar{})
			a.parse(line, &a.chars[len(a.chars)-1])
		}
	}

	return a.makeFont()
}

func NewAngelFontFile(re *sdl.Display, name string) (*BMPFont, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	d := angel{re: re}
	return d.Decode(r)
}

func (a *angel) makeFont() (*BMPFont, error) {
	var textures []*sdl.Texture
	var page []int

	for i, p := range a.pages {
		texture, err := a.re.CreateTextureFromFile(p.File)
		if err != nil {
			return nil, err
		}
		textures = append(textures, texture)
		page = append(page, i)
	}

	bi := BMPFontInfo{
		Chars:   make(map[rune]BMPChar),
		Kern:    make(map[BMPKern]sdl.Point),
		Tabstop: tabstop,
		Width:   a.info.Size,
		Height:  a.common.LineHeight,
	}

	for _, c := range a.chars {
		var pg int
		var i int

		for i = 0; i < len(page); i++ {
			if c.Page == page[i] {
				pg = page[i]
				break
			}
		}
		if i == len(page) {
			return nil, fmt.Errorf("angelfont: page %v does not exist", c.Page)
		}

		bi.Chars[rune(c.Id)] = BMPChar{
			Texcoord: sdl.Rect{
				X: sdl.Int(c.X),
				Y: sdl.Int(c.Y),
				W: sdl.Int(c.Width),
				H: sdl.Int(c.Height),
			},
			Texture: textures[pg],
			Width:   c.Width,
			Height:  c.Height,
		}
	}

	return NewBMPFont(bi), nil
}

func (a *angel) parse(line string, p interface{}) {
	var key, value string
	v := reflect.ValueOf(p).Elem()

loop:
	for {
		line, key = a.getToken(line)
		if line == "" {
			break
		}

		if line[0] != '=' {
			line = line[1:]
			continue
		}

		line, value = a.getToken(line[1:])
		r, size := utf8.DecodeRuneInString(key)
		key = string(unicode.ToUpper(r)) + key[size:]

		field := v.FieldByName(key)
		if !field.IsValid() {
			continue
		}
		switch field.Kind() {
		case reflect.Int:
			n, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			field.SetInt(int64(n))
		case reflect.String:
			field.SetString(value)
		case reflect.Slice:
			snums := strings.Split(value, ",")
			var nums []int
			for _, sn := range snums {
				n, err := strconv.Atoi(sn)
				if err != nil {
					continue loop
				}
				nums = append(nums, n)
			}
			field.Set(reflect.ValueOf(nums))
		default:
			panic("unimplemented")
		}
	}
}

func (a *angel) getToken(line string) (string, string) {
	line = strings.Trim(line, " \r\t\n")

	if len(line) == 0 {
		return line, ""
	}

	if line[0] == '"' {
		line = line[1:]
		escape := 0
		for i, c := range line {
			if c == '\\' && escape == 0 {
				escape++
				continue
			}

			if c == '"' && escape == 0 {
				return line[i+1:], line[:i]
			}

			escape = 0
		}
	} else {
		for i, c := range line {
			if strings.ContainsRune(" \r\t\n=", c) {
				return line[i:], line[:i]
			}
		}
	}

	return "", line
}
