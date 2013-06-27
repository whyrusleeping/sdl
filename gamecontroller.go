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

type GameControllerAxis int
type GameControllerButton int

const (
	CONTROLLER_AXIS_INVALID      GameControllerAxis = C.SDL_CONTROLLER_AXIS_INVALID
	CONTROLLER_AXIS_LEFTX                           = C.SDL_CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY                           = C.SDL_CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX                          = C.SDL_CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY                          = C.SDL_CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT                     = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT                    = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT
)

const (
	CONTROLLER_BUTTON_INVALID       GameControllerButton = C.SDL_CONTROLLER_BUTTON_INVALID
	CONTROLLER_BUTTON_A                                  = C.SDL_CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B                                  = C.SDL_CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X                                  = C.SDL_CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y                                  = C.SDL_CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK                               = C.SDL_CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE                              = C.SDL_CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START                              = C.SDL_CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK                          = C.SDL_CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK                         = C.SDL_CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER                       = C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER                      = C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP                            = C.SDL_CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN                          = C.SDL_CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT                          = C.SDL_CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT                         = C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX                                = C.SDL_CONTROLLER_BUTTON_MAX
)

func GameControllerAddMapping(s string) (int, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	v := int(C.SDL_GameControllerAddMapping(cs))
	if v < 0 {
		return v, errors.New(GetError())
	}
	return v, nil
}

func IsGameController(index int) bool {
	v := C.SDL_IsGameController(C.int(index))
	if v == C.SDL_FALSE {
		return false
	}
	return true
}

func GameControllerOpen(index int) (*GameController, error) {
	gc := (*GameController)(C.SDL_GameControllerOpen(C.int(index)))
	if gc == nil {
		return nil, errors.New(GetError())
	}
	return gc, nil
}

func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

func (gc *GameController) Attached() bool {
	v := C.SDL_GameControllerGetAttached((*C.SDL_GameController)(gc))
	if v == C.SDL_FALSE {
		return false
	}
	return true
}

func (gc *GameController) Close() {
	C.SDL_GameControllerClose((*C.SDL_GameController)(gc))
}

func (gc *GameController) Axis(axis GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis((*C.SDL_GameController)(gc), C.SDL_GameControllerAxis(axis)))
}

func (gc *GameController) Button(button GameControllerButton) uint8 {
	return uint8(C.SDL_GameControllerGetButton((*C.SDL_GameController)(gc), C.SDL_GameControllerButton(button)))
}
