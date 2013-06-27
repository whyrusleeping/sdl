#include <SDL.h>
#include "gosdl.h"
#include "_cgo_export.h"

SDL_Surface *sdl_loadbmp(const char *file) {
    return SDL_LoadBMP(file);
}

void set_audio_callback(SDL_AudioSpec *spec) {
    spec->callback = goAudioCallback;
}
