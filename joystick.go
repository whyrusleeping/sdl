package sdl

/*
#include <SDL.h>
*/
import "C"

import (
	"errors"
)

func NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

func JoystickNameForIndex(index int) string {
	return C.GoString(C.SDL_JoystickNameForIndex(C.int(index)))
}

func JoystickOpen(index int) (*Joystick, error) {
	j := (*Joystick)(C.SDL_JoystickOpen(C.int(index)))
	if j == nil {
		return nil, errors.New(GetError())
	}
	return j, nil
}

func (j *Joystick) Name() string {
	return C.GoString(C.SDL_JoystickName((*C.SDL_Joystick)(j)))
}

func (j *Joystick) Attached() bool {
	v := C.SDL_JoystickGetAttached((*C.SDL_Joystick)(j))
	if v == C.SDL_FALSE {
		return false
	}
	return true
}

func (j *Joystick) ID() int {
	return int(C.SDL_JoystickInstanceID((*C.SDL_Joystick)(j)))
}

func (j *Joystick) NumHats() int {
	return int(C.SDL_JoystickNumHats((*C.SDL_Joystick)(j)))
}

func (j *Joystick) NumButtons() int {
	return int(C.SDL_JoystickNumButtons((*C.SDL_Joystick)(j)))
}
