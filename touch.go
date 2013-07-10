package sdl

/*
#include <SDL.h>
*/
import "C"

import (
	"unsafe"
)

func GetNumTouchDevices() int {
	return int(C.SDL_GetNumTouchDevices())
}

func GetTouchDevice(index int) TouchID {
	return TouchID(C.SDL_GetTouchDevice(C.int(index)))
}

func (id TouchID) NumTouchFingers() int {
	return int(C.SDL_GetNumTouchFingers(C.SDL_TouchID(id)))
}

func (id TouchID) TouchFinger(index int) *Finger {
	return (*Finger)(unsafe.Pointer(C.SDL_GetTouchFinger(C.SDL_TouchID(id), C.int(index))))
}
