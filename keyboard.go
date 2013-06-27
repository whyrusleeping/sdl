package sdl

/*
#include <SDL.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

func GetKeyboardState() []uint8 {
	var numkeys C.int
	var keystate []uint8

	state := C.SDL_GetKeyboardState(&numkeys)
	sl := (*reflect.SliceHeader)((unsafe.Pointer(&keystate)))
	sl.Cap = int(numkeys)
	sl.Len = int(numkeys)
	sl.Data = uintptr(unsafe.Pointer(state))
	return keystate
}
