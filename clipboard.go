package sdl

/*
#include <SDL.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func SetClipboardText(text string) {
	s := C.CString(text)
	C.SDL_SetClipboardText(s)
	C.free(unsafe.Pointer(s))
}

func GetClipboardText() string {
	s := C.SDL_GetClipboardText()
	return C.GoString(s)
}

func HasClipboardText() bool {
	v := C.SDL_HasClipboardText()
	if v == 0 {
		return false
	}
	return true
}
