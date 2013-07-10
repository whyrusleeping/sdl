package sdl

/*
#include <SDL.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func PollEvent() Event {
	var ev C.SDL_Event
	if C.SDL_PollEvent(&ev) == 0 {
		return nil
	}

	return convertEvent(&ev)
}

func PumpEvents() {
	C.SDL_PumpEvents()
}

func WaitEvent() Event {
	var ev C.SDL_Event
	rc := C.SDL_WaitEvent(&ev)
	if rc == 0 {
		return nil
	}

	return convertEvent(&ev)
}

func FlushEvent(typ uint32) {
	C.SDL_FlushEvent(C.Uint32(typ))
}

func FlushEvents(minType, maxType uint32) {
	C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

func convertEvent(ev *C.SDL_Event) Event {
	typ := *((*uint32)(unsafe.Pointer(ev)))
	switch typ {
	case C.SDL_QUIT:
		return &QuitEvent{}

	case C.SDL_KEYDOWN, C.SDL_KEYUP:
		ev := (*C.SDL_KeyboardEvent)(unsafe.Pointer(ev))
		kb := &KeyboardEvent{
			Type:      uint32(ev._type),
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			State:     uint8(ev.state),
			Keysym: Keysym{
				ScanCode: int32(ev.keysym.scancode),
				Sym:      Keycode(ev.keysym.sym),
				Mod:      uint16(ev.keysym.mod),
			},
		}

		kb.Repeat = false
		if ev.repeat != 0 {
			kb.Repeat = true
		}

		return kb

	case C.SDL_WINDOWEVENT:
		ev := (*C.SDL_WindowEvent)(unsafe.Pointer(ev))
		return &WindowEvent{
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			Event:     uint8(ev.event),
			Data:      [2]int{int(ev.data1), int(ev.data2)},
		}

	case C.SDL_FIRSTEVENT:
		return &FirstEvent{}

	case C.SDL_MOUSEMOTION:
		ev := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(ev))
		return &MouseMotionEvent{
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			State:     uint8(ev.state),
			X:         int(ev.x),
			Y:         int(ev.y),
			Xrel:      int(ev.xrel),
			Yrel:      int(ev.yrel),
		}

	case C.SDL_MOUSEBUTTONDOWN, C.SDL_MOUSEBUTTONUP:
		ev := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(ev))
		return &MouseButtonEvent{
			Type:      uint32(ev._type),
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			Button:    uint8(ev.button),
			State:     uint8(ev.state),
			X:         int(ev.x),
			Y:         int(ev.y),
		}

	case C.SDL_MOUSEWHEEL:
		ev := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(ev))
		return &MouseWheelEvent{
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			X:         int(ev.x),
			Y:         int(ev.y),
		}

	case C.SDL_TEXTEDITING:
		ev := (*C.SDL_TextEditingEvent)(unsafe.Pointer(ev))
		t := &TextEditingEvent{
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
			Start:     int(ev.start),
			Length:    int(ev.length),
		}
		for i := range t.Text {
			t.Text[i] = byte(ev.text[i])
		}
		return t

	case C.SDL_TEXTINPUT:
		ev := (*C.SDL_TextInputEvent)(unsafe.Pointer(ev))
		t := &TextInputEvent{
			Timestamp: uint32(ev.timestamp),
			WindowID:  uint32(ev.windowID),
		}
		for i := range t.Text {
			t.Text[i] = byte(ev.text[i])
		}
		return t

	case C.SDL_JOYAXISMOTION:
		ev := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(ev))
		return &JoyAxisEvent{
			Timestamp: uint32(ev.timestamp),
			Which:     int(ev.which),
			Axis:      uint8(ev.axis),
			Value:     int16(ev.value),
		}

	case C.SDL_JOYBUTTONDOWN, C.SDL_JOYBUTTONUP:
		ev := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(ev))
		return &JoyButtonEvent{
			Type:      uint32(ev._type),
			Timestamp: uint32(ev.timestamp),
			Which:     int(ev.which),
			Button:    uint8(ev.button),
			State:     uint8(ev.state),
		}

	case C.SDL_CONTROLLERAXISMOTION:
		ev := (*C.SDL_ControllerAxisEvent)(unsafe.Pointer(ev))
		return &ControllerAxisEvent{
			Timestamp: uint32(ev.timestamp),
			Which:     int(ev.which),
			Axis:      uint8(ev.axis),
			Value:     int16(ev.value),
		}

	case C.SDL_CONTROLLERBUTTONDOWN, C.SDL_CONTROLLERBUTTONUP:
		ev := (*C.SDL_ControllerButtonEvent)(unsafe.Pointer(ev))
		return &ControllerButtonEvent{
			Type:      uint32(ev._type),
			Timestamp: uint32(ev.timestamp),
			Which:     int(ev.which),
			Button:    uint8(ev.button),
			State:     uint8(ev.state),
		}

	default:
		panic(fmt.Sprintf("sdl: unimplemented event: %v", typ))
	}

	panic("unreachable")
}
