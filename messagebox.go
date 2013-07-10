package sdl

/*
#include <SDL.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

func ShowSimpleMessageBox(flags uint32, title, message string, w *Window) error {
	ctitle := C.CString(title)
	cmessage := C.CString(message)
	v := C.SDL_ShowSimpleMessageBox(C.Uint32(flags), ctitle, cmessage, (*C.SDL_Window)(w))
	C.free(unsafe.Pointer(ctitle))
	C.free(unsafe.Pointer(cmessage))

	if v < 0 {
		return errors.New(GetError())
	}

	return nil
}
