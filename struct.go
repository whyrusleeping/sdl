package sdl

/*
#include <SDL.h>
#include "gosdl.h"
*/
import "C"

type TouchID C.SDL_TouchID
type FingerID C.SDL_FingerID
type Int C.int
type Joystick C.SDL_Joystick
type GameController C.SDL_GameController
type BlendMode int
type Keycode C.SDL_Keycode
type GLContext C.SDL_GLContext
type Window C.SDL_Window
type Renderer C.SDL_Renderer
type Texture C.SDL_Texture
type Surface C.SDL_Surface
type Flip int
type Display struct {
	*Window
	*Renderer
}

type DisplayMode struct {
	Format      uint32
	W, H        int
	RefreshRate int
}

type Point struct {
	X, Y Int
}

type Rect struct {
	X, Y, W, H Int
}

type AudioFormat uint16

type Audio struct {
	Freq     int
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	Size     uint32
	Callback func([]byte)
}

type Event interface{}

type FirstEvent struct{}
type QuitEvent struct{}

type KeyboardEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    bool
	Keysym
}

type Keysym struct {
	Mod      uint16
	ScanCode int32
	Sym      Keycode
	Unicode  uint32
}

type WindowEvent struct {
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	Data      [2]int
}

type MouseMotionEvent struct {
	Timestamp        uint32
	WindowID         uint32
	State            uint8
	X, Y, Xrel, Yrel int
}

type MouseButtonEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Button    uint8
	State     uint8
	X, Y      int
}

type MouseWheelEvent struct {
	Timestamp uint32
	WindowID  uint32
	X, Y      int
}

type TextEditingEvent struct {
	Timestamp uint32
	WindowID  uint32
	Text      [32]byte
	Start     int
	Length    int
}

type TextInputEvent struct {
	Timestamp uint32
	WindowID  uint32
	Text      [32]byte
}

type JoyAxisEvent struct {
	Timestamp uint32
	Which     int
	Axis      uint8
	Value     int16
}

type JoyBallEvent struct {
	Timestamp uint32
	Which     int
	Ball      uint8
	Xrel      int16
	Yrel      int16
}

type JoyHatEvent struct {
	Timestamp uint32
	Which     int
	Hat       uint8
	Value     uint8
}

type JoyButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int
	Button    uint8
	State     uint8
}

type ControllerAxisEvent struct {
	Timestamp uint32
	Which     int
	Axis      uint8
	Value     int16
}

type ControllerButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int
	Button    uint8
	State     uint8
}

type MessageBoxColor struct {
	R, G, B uint8
}

type MessageBoxColorScheme struct {
	Colors [MESSAGEBOX_COLOR_MAX]MessageBoxColor
}

type MessageBoxButtonData struct {
	Flags    uint32
	ButtonID int
	Text     string
}

type MessageBoxData struct {
	Flags       uint32
	Window      *Window
	Title       string
	Message     string
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

type Finger struct {
	ID             FingerID
	X, Y, Pressure float32
}
