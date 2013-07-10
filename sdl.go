package sdl

/*
#include <SDL.h>
#include "gosdl.h"
*/
import "C"

import (
	"errors"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"unsafe"

	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/go.image/tiff"
)

func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) < 0 {
		return errors.New(GetError())
	}

	return nil
}

func InitSubSystem(flags uint32) error {
	if C.SDL_InitSubSystem(C.Uint32(flags)) < 0 {
		return errors.New(GetError())
	}
	return nil
}

func QuitSubSystem(flags uint32) { C.SDL_QuitSubSystem(C.Uint32(flags)) }
func Quit()                      { C.SDL_Quit() }

func GetError() string { return C.GoString(C.SDL_GetError()) }

func NewDisplay(width, height int, wflags uint32) (*Display, error) {
	wr := &Display{}
	ok := C.SDL_CreateWindowAndRenderer(C.int(width), C.int(height), C.Uint32(wflags),
		(**C.SDL_Window)(unsafe.Pointer(&wr.Window)), (**C.SDL_Renderer)((unsafe.Pointer)(&wr.Renderer)))

	if ok < 0 {
		return nil, errors.New(GetError())
	}

	return wr, nil
}

func CreateRGBSurfaceFrom(pixels []byte, width, height, depth, pitch int, rmask, gmask, bmask, amask uint32) (*Surface, error) {
	s := C.SDL_CreateRGBSurfaceFrom(unsafe.Pointer(&pixels[0]), C.int(width), C.int(height), C.int(depth), C.int(pitch),
		C.Uint32(rmask), C.Uint32(gmask), C.Uint32(bmask), C.Uint32(amask))
	if s == nil {
		return nil, errors.New(GetError())
	}
	return (*Surface)(s), nil
}

func (s *Surface) Free() {
	C.SDL_FreeSurface((*C.SDL_Surface)(s))
}

func (w *Window) SetFullscreen(fullscreen bool) {
	f := C.Uint32(0)
	if fullscreen {
		f = 1
	}
	C.SDL_SetWindowFullscreen((*C.SDL_Window)(w), f)
}

func (w *Window) Maximize() {
	C.SDL_MaximizeWindow((*C.SDL_Window)(w))
}

func (w *Window) Minimize() {
	C.SDL_MinimizeWindow((*C.SDL_Window)(w))
}

func (w *Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition((*C.SDL_Window)(w), C.int(x), C.int(y))
}

func (w *Window) Position() (x, y int) {
	var cx, cy C.int
	C.SDL_GetWindowPosition((*C.SDL_Window)(w), &cx, &cy)
	return int(cx), int(cy)
}

func (w *Window) SetSize(x, y int) {
	C.SDL_SetWindowSize((*C.SDL_Window)(w), C.int(x), C.int(y))
}

func (w *Window) Size() (x, y int) {
	var cx, cy C.int
	C.SDL_GetWindowSize((*C.SDL_Window)(w), &cx, &cy)
	return int(cx), int(cy)
}

func (w *Window) SetMinimumSize(minw, minh int) {
	C.SDL_SetWindowMinimumSize((*C.SDL_Window)(w), C.int(minw), C.int(minh))
}

func (w *Window) MinimumSize() (mw, mh int) {
	var cw, ch C.int
	C.SDL_GetWindowMinimumSize((*C.SDL_Window)(w), &cw, &ch)
	return int(cw), int(ch)
}

func (w *Window) SetMaximumSize(maxw, maxh int) {
	C.SDL_SetWindowMaximumSize((*C.SDL_Window)(w), C.int(maxw), C.int(maxh))
}

func (w *Window) MaximumSize() (mw, mh int) {
	var cw, ch C.int
	C.SDL_GetWindowMaximumSize((*C.SDL_Window)(w), &cw, &ch)
	return int(cw), int(ch)
}

func (w *Window) Show() {
	C.SDL_ShowWindow((*C.SDL_Window)(w))
}

func (w *Window) Hide() {
	C.SDL_HideWindow((*C.SDL_Window)(w))
}

func (w *Window) Raise() {
	C.SDL_RaiseWindow((*C.SDL_Window)(w))
}

func (w *Window) Restore() {
	C.SDL_RestoreWindow((*C.SDL_Window)(w))
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow((*C.SDL_Window)(w))
}

func (w *Window) DisplayMode() DisplayMode {
	var mode C.SDL_DisplayMode
	C.SDL_GetWindowDisplayMode((*C.SDL_Window)(w), &mode)
	return DisplayMode{
		Format:      uint32(mode.format),
		W:           int(mode.w),
		H:           int(mode.h),
		RefreshRate: int(mode.refresh_rate),
	}
}

func (w *Window) CreateGLContext() GLContext {
	return GLContext(C.SDL_GL_CreateContext((*C.SDL_Window)(w)))
}

func (w *Window) MakeGLCurrent(context GLContext) {
	if context == nil {
		w.CreateGLContext()
	} else {
		C.SDL_GL_MakeCurrent((*C.SDL_Window)(w), C.SDL_GLContext(context))
	}
}

func (w *Window) SetGLSwapInterval(interval int) {
	C.SDL_GL_SetSwapInterval(C.int(interval))
}

func (w *Window) GLSwapInterval() int {
	return int(C.SDL_GL_GetSwapInterval())
}

func (w *Window) SetTitle(title string) {
	ctitle := C.CString(title)
	C.SDL_SetWindowTitle((*C.SDL_Window)(w), ctitle)
	C.free(unsafe.Pointer(ctitle))
}

func (w *Window) SetGrab(grabbed bool) {
	var g C.SDL_bool
	if grabbed {
		g = 1
	}
	C.SDL_SetWindowGrab((*C.SDL_Window)(w), g)
}

func (w *Window) Grabbed() bool {
	return C.SDL_GetWindowGrab((*C.SDL_Window)(w)) != 0
}

func (w *Window) UpdateSurface() {
	C.SDL_UpdateWindowSurface((*C.SDL_Window)(w))
}

func (re *Renderer) SetTarget(texture *Texture) error {
	if C.SDL_SetRenderTarget((*C.SDL_Renderer)(re), (*C.SDL_Texture)(texture)) < 0 {
		return errors.New(GetError())
	}
	return nil
}

func (re *Renderer) SetDrawColor(c color.RGBA) {
	C.SDL_SetRenderDrawColor((*C.SDL_Renderer)(re), C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B), C.Uint8(c.A))
}

func (re *Renderer) Clear() {
	C.SDL_RenderClear((*C.SDL_Renderer)(re))
}

func (re *Renderer) Present() {
	C.SDL_RenderPresent((*C.SDL_Renderer)(re))
}

func (re *Renderer) DrawLine(x1, y1, x2, y2 int) {
	C.SDL_RenderDrawLine((*C.SDL_Renderer)(re), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func (re *Renderer) DrawPoint(x, y int) {
	C.SDL_RenderDrawPoint((*C.SDL_Renderer)(re), C.int(x), C.int(y))
}

func (re *Renderer) DrawPoints(pts []Point) {
	C.SDL_RenderDrawPoints((*C.SDL_Renderer)(re), (*C.SDL_Point)(unsafe.Pointer(&pts[0])), C.int(len(pts)))
}

func (re *Renderer) DrawRect(r Rect) {
	cr := C.SDL_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}
	C.SDL_RenderDrawRect((*C.SDL_Renderer)(re), &cr)
}

func (re *Renderer) DrawRects(r []Rect) {
	C.SDL_RenderDrawRects((*C.SDL_Renderer)(re), (*C.SDL_Rect)(unsafe.Pointer(&r[0])), C.int(len(r)))
}

func (re *Renderer) FillRect(r Rect) {
	cr := C.SDL_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}
	C.SDL_RenderFillRect((*C.SDL_Renderer)(re), &cr)
}

func (re *Renderer) FillRects(r []Rect) {
	C.SDL_RenderFillRects((*C.SDL_Renderer)(re), (*C.SDL_Rect)(unsafe.Pointer(&r[0])), C.int(len(r)))
}

func (re *Renderer) SetScale(x, y float64) {
	C.SDL_RenderSetScale((*C.SDL_Renderer)(re), C.float(x), C.float(y))
}

func (re *Renderer) OutputSize() (w, h int, err error) {
	var cw, ch C.int
	v := C.SDL_GetRendererOutputSize((*C.SDL_Renderer)(re), &cw, &ch)
	w, h = int(cw), int(ch)
	if v < 0 {
		err = errors.New(GetError())
	}
	return
}

func (re *Renderer) SetViewport(r Rect) {
	cr := C.SDL_Rect{x: C.int(r.X), y: C.int(r.Y), w: C.int(r.W), h: C.int(r.H)}
	C.SDL_RenderSetViewport((*C.SDL_Renderer)(re), &cr)
}

func (re *Renderer) SetClipRect(r Rect) {
	cr := C.SDL_Rect{x: C.int(r.X), y: C.int(r.Y), w: C.int(r.W), h: C.int(r.H)}
	C.SDL_RenderSetClipRect((*C.SDL_Renderer)(re), &cr)
}

func (re *Renderer) Viewport() Rect {
	var cr C.SDL_Rect
	C.SDL_RenderGetViewport((*C.SDL_Renderer)(re), &cr)
	return Rect{X: Int(cr.x), Y: Int(cr.y), W: Int(cr.w), H: Int(cr.h)}
}

func (re *Renderer) ClipRect() Rect {
	var cr C.SDL_Rect
	C.SDL_RenderGetClipRect((*C.SDL_Renderer)(re), &cr)
	return Rect{X: Int(cr.x), Y: Int(cr.y), W: Int(cr.w), H: Int(cr.h)}
}

func (re *Renderer) CreateTexture(format uint32, access, w, h int) (*Texture, error) {
	t := (*Texture)(C.SDL_CreateTexture((*C.SDL_Renderer)(re), C.Uint32(format), C.int(access), C.int(w), C.int(h)))
	if t == nil {
		return nil, errors.New(GetError())
	}
	return t, nil
}

func (re *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	t := (*Texture)(C.SDL_CreateTextureFromSurface((*C.SDL_Renderer)(re), (*C.SDL_Surface)(surface)))
	if t == nil {
		return nil, errors.New(GetError())
	}
	return t, nil
}

func (re *Renderer) CreateTextureFromImage(img image.Image) (*Texture, error) {
	var s *Surface
	var err error

	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	pixels := C.malloc(C.size_t(w * h * 4))
	p := (*[math.MaxInt32]byte)(pixels)
	l := w * h * 4
	switch v := img.(type) {
	case *image.NRGBA:
		copy(p[:l], v.Pix)
		s, err = CreateRGBSurfaceFrom(p[:l], w, h, 32, w*4, 0xFF, 0xFF<<8, 0xFF<<16, 0xFF<<24)
	default:
		for i, y := 0, 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r32, g32, b32, a32 := img.At(x, y).RGBA()
				p[i], p[i+1], p[i+2], p[i+3] = uint8(r32), uint8(g32), uint8(b32), uint8(a32)
				i += 4
			}
		}
		s, err = CreateRGBSurfaceFrom(p[:l], w, h, 32, w*4, 0xFF, 0xFF<<8, 0xFF<<16, 0xFF<<24)
	}
	if err != nil {
		return nil, err
	}

	defer s.Free()
	return re.CreateTextureFromSurface(s)
}

func (re *Renderer) CreateTextureFromFile(name string) (*Texture, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return re.CreateTextureFromImage(img)
}

func (re *Renderer) Copy(texture *Texture, src, dst *Rect) int {
	var s, d *C.SDL_Rect

	if src != nil {
		s = &C.SDL_Rect{C.int(src.X), C.int(src.Y), C.int(src.W), C.int(src.H)}
	}

	if dst != nil {
		d = &C.SDL_Rect{C.int(dst.X), C.int(dst.Y), C.int(dst.W), C.int(dst.H)}
	}

	return int(C.SDL_RenderCopy((*C.SDL_Renderer)(re), (*C.SDL_Texture)(texture), s, d))
}

func (re *Renderer) CopyEx(texture *Texture, src, dst *Rect, angle float64, center *Point, flip Flip) int {
	var s, d *C.SDL_Rect
	var p *C.SDL_Point

	if src != nil {
		s = &C.SDL_Rect{C.int(src.X), C.int(src.Y), C.int(src.W), C.int(src.H)}
	}

	if dst != nil {
		d = &C.SDL_Rect{C.int(dst.X), C.int(dst.Y), C.int(dst.W), C.int(dst.H)}
	}

	if center != nil {
		p = &C.SDL_Point{C.int(center.X), C.int(center.Y)}
	}

	return int(C.SDL_RenderCopyEx((*C.SDL_Renderer)(re), (*C.SDL_Texture)(texture),
		s, d, C.double(angle), p, C.SDL_RendererFlip(flip)))
}

func (re *Renderer) DrawColor() color.RGBA {
	var r, g, b, a C.Uint8
	C.SDL_GetRenderDrawColor((*C.SDL_Renderer)(re), &r, &g, &b, &a)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func (re *Renderer) Destroy() {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(re))
}

func (t *Texture) SetColorMod(c color.RGBA) {
	C.SDL_SetTextureColorMod((*C.SDL_Texture)(t), C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B))
}

func (t *Texture) SetAlphaMode(alpha uint8) {
	C.SDL_SetTextureAlphaMod((*C.SDL_Texture)(t), C.Uint8(alpha))
}

func (t *Texture) SetBlendMode(mode BlendMode) int {
	return int(C.SDL_SetTextureBlendMode((*C.SDL_Texture)(t), C.SDL_BlendMode(mode)))
}

func (t *Texture) BlendMode() (BlendMode, error) {
	var mode C.SDL_BlendMode
	rc := C.SDL_GetTextureBlendMode((*C.SDL_Texture)(t), &mode)
	if rc < 0 {
		return BlendMode(0), errors.New(GetError())
	}
	return BlendMode(mode), nil
}

func (t *Texture) Update(rect *Rect, pixels []byte, pitch int) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	rc := C.SDL_UpdateTexture((*C.SDL_Texture)(t), cr, unsafe.Pointer(&pixels[0]), C.int(pitch))
	if rc < 0 {
		return errors.New(GetError())
	}
	return nil
}

func (t *Texture) Lock(rect Rect) ([]byte, int, error) {
	var pixels unsafe.Pointer
	var pitch C.int

	cr := C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	rc := C.SDL_LockTexture((*C.SDL_Texture)(t), &cr, &pixels, &pitch)
	if rc < 0 {
		return nil, 0, errors.New(GetError())
	}
	return nil, 0, nil
}

func (t *Texture) Query() (format uint32, access, w, h int) {
	var cformat C.Uint32
	var caccess, cw, ch C.int

	C.SDL_QueryTexture((*C.SDL_Texture)(t), &cformat, &caccess, &cw, &ch)
	return uint32(cformat), int(caccess), int(cw), int(ch)
}

func (t *Texture) Unlock() {
	C.SDL_UnlockTexture((*C.SDL_Texture)(t))
}

func (t *Texture) Destroy() {
	C.SDL_DestroyTexture((*C.SDL_Texture)(t))
}

func LoadBMP(fn string) (*Surface, error) {
	cfn := C.CString(fn)
	defer C.free(unsafe.Pointer(cfn))

	s := (*Surface)(C.sdl_loadbmp(cfn))
	if s == nil {
		return nil, errors.New(GetError())
	}

	return s, nil
}
