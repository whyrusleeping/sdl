package sdl

/*
#include <SDL.h>
#include "gosdl.h"
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

func (au *Audio) Lock() {
	C.SDL_LockAudio()
}

func (au *Audio) Unlock() {
	C.SDL_UnlockAudio()
}

//export goAudioCallback
func goAudioCallback(userdata unsafe.Pointer, stream *byte, length C.int) {
	var buf []byte

	sl := (*reflect.SliceHeader)((unsafe.Pointer(&buf)))
	sl.Cap = int(length)
	sl.Len = int(length)
	sl.Data = uintptr(unsafe.Pointer(stream))

	au := (*Audio)(userdata)
	au.Callback(buf)
}

func OpenAudio(desired Audio) (*Audio, error) {
	if desired.Callback == nil {
		panic("callback is nil")
	}
	au := &Audio{}
	o := C.SDL_AudioSpec{}
	d := C.SDL_AudioSpec{
		freq:     C.int(desired.Freq),
		format:   C.SDL_AudioFormat(desired.Format),
		channels: C.Uint8(desired.Channels),
		samples:  C.Uint16(desired.Samples),
		size:     C.Uint32(desired.Size),
		userdata: unsafe.Pointer(au),
	}

	C.set_audio_callback(&d)
	ok := C.SDL_OpenAudio(&d, &o)
	if ok < 0 {
		return nil, errors.New(GetError())
	}

	au.Freq = int(o.freq)
	au.Format = AudioFormat(o.format)
	au.Channels = uint8(o.channels)
	au.Silence = uint8(o.silence)
	au.Samples = uint16(o.samples)
	au.Size = uint32(o.size)
	au.Callback = desired.Callback
	return au, nil
}

func (au *Audio) Pause(pause int) {
	C.SDL_PauseAudio(C.int(pause))
}

func (au *Audio) Status() AudioStatus {
	return AudioStatus(C.SDL_GetAudioStatus())
}

func (au *Audio) Close() {
	C.SDL_CloseAudio()
}
