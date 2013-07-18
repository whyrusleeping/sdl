package sdl

/*
#include <SDL.h>
*/
import "C"

type AudioStatus int
type GLattr C.SDL_GLattr
type GLprofile C.SDL_GLprofile
type GLcontextFlag C.SDL_GLcontextFlag
type PowerState C.SDL_PowerState

const (
	POWERSTATE_UNKNOWN    PowerState = C.SDL_POWERSTATE_UNKNOWN
	POWERSTATE_ON_BATTERY            = C.SDL_POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY            = C.SDL_POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING              = C.SDL_POWERSTATE_CHARGING
	POWERSTATE_CHARGED               = C.SDL_POWERSTATE_CHARGED
)

const (
	GL_RED_SIZE                   GLattr = C.SDL_GL_RED_SIZE
	GL_GREEN_SIZE                        = C.SDL_GL_GREEN_SIZE
	GL_BLUE_SIZE                         = C.SDL_GL_BLUE_SIZE
	GL_ALPHA_SIZE                        = C.SDL_GL_ALPHA_SIZE
	GL_BUFFER_SIZE                       = C.SDL_GL_BUFFER_SIZE
	GL_DOUBLEBUFFER                      = C.SDL_GL_DOUBLEBUFFER
	GL_DEPTH_SIZE                        = C.SDL_GL_DEPTH_SIZE
	GL_STENCIL_SIZE                      = C.SDL_GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE                    = C.SDL_GL_ACCUM_RED_SIZE
	GL_ACCUM_BLUE_SIZE                   = C.SDL_GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE                  = C.SDL_GL_ACCUM_ALPHA_SIZE
	GL_STEREO                            = C.SDL_GL_STEREO
	GL_MULTISAMPLEBUFFERS                = C.SDL_GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES                = C.SDL_GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL                = C.SDL_GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING                  = C.SDL_GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION             = C.SDL_GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION             = C.SDL_GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL                       = C.SDL_GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS                     = C.SDL_GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK              = C.SDL_GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT        = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT
)

const (
	GL_CONTEXT_PROFILE_CORE          GLprofile = C.SDL_GL_CONTEXT_PROFILE_CORE
	GL_CONTEXT_PROFILE_COMPATIBILITY           = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY
	GL_CONTEXT_PROFILE_ES                      = C.SDL_GL_CONTEXT_PROFILE_ES
)

const (
	GL_CONTEXT_DEBUG_FLAG              GLcontextFlag = C.SDL_GL_CONTEXT_DEBUG_FLAG
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG               = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG
	GL_CONTEXT_ROBUST_ACCESS_FLAG                    = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG
	GL_CONTEXT_RESET_ISOLATION_FLAG                  = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG
)

const (
	ENABLE = C.SDL_ENABLE
	IGNORE = C.SDL_IGNORE
	QUERY  = C.SDL_QUERY
)

const (
	INIT_VIDEO      = C.SDL_INIT_VIDEO
	INIT_AUDIO      = C.SDL_INIT_AUDIO
	INIT_TIMER      = C.SDL_INIT_TIMER
	INIT_JOYSTICK   = C.SDL_INIT_JOYSTICK
	INIT_EVERYTHING = C.SDL_INIT_EVERYTHING
)

const (
	WINDOWEVENT_HIDDEN = C.SDL_WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED = C.SDL_WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED = C.SDL_WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED = C.SDL_WINDOWEVENT_RESIZED
	WINDOWEVENT_MINIMIZED = C.SDL_WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED = C.SDL_WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED = C.SDL_WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER = C.SDL_WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE = C.SDL_WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED = C.SDL_WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST = C.SDL_WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE = C.SDL_WINDOWEVENT_CLOSE
)

const (
	FIRSTEVENT = C.SDL_FIRSTEVENT

	QUIT = C.SDL_QUIT

	APP_TERMINATING         = C.SDL_APP_TERMINATING
	APP_LOWMEMORY           = C.SDL_APP_LOWMEMORY
	APP_WILLENTERBACKGROUND = C.SDL_APP_WILLENTERBACKGROUND
	APP_DIDENTERBACKGROUND  = C.SDL_APP_DIDENTERBACKGROUND
	APP_WILLENTERFOREGROUND = C.SDL_APP_WILLENTERFOREGROUND
	APP_DIDENTERFOREGROUND  = C.SDL_APP_DIDENTERFOREGROUND

	WINDOWEVENT = C.SDL_WINDOWEVENT
	SYSWMEVENT  = C.SDL_SYSWMEVENT

	KEYDOWN     = C.SDL_KEYDOWN
	KEYUP       = C.SDL_KEYUP
	TEXTEDITING = C.SDL_TEXTEDITING
	TEXTINPUT   = C.SDL_TEXTINPUT

	MOUSEMOTION     = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP
	MOUSEWHEEL      = C.SDL_MOUSEWHEEL

	JOYAXISMOTION    = C.SDL_JOYAXISMOTION
	JOYBALLMOTION    = C.SDL_JOYBALLMOTION
	JOYHATMOTION     = C.SDL_JOYHATMOTION
	JOYBUTTONDOWN    = C.SDL_JOYBUTTONDOWN
	JOYBUTTONUP      = C.SDL_JOYBUTTONUP
	JOYDEVICEADDED   = C.SDL_JOYDEVICEADDED
	JOYDEVICEREMOVED = C.SDL_JOYDEVICEREMOVED

	CONTROLLERAXISMOTION     = C.SDL_CONTROLLERAXISMOTION
	CONTROLLERBUTTONDOWN     = C.SDL_CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP       = C.SDL_CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED    = C.SDL_CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED  = C.SDL_CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED = C.SDL_CONTROLLERDEVICEREMAPPED

	FINGERDOWN   = C.SDL_FINGERDOWN
	FINGERUP     = C.SDL_FINGERUP
	FINGERMOTION = C.SDL_FINGERMOTION

	DOLLARGESTURE = C.SDL_DOLLARGESTURE
	DOLLARRECORD  = C.SDL_DOLLARRECORD
	MULTIGESTURE  = C.SDL_MULTIGESTURE

	CLIPBOARDUPDATE = C.SDL_CLIPBOARDUPDATE

	DROPFILE = C.SDL_DROPFILE

	USEREVENT = C.SDL_USEREVENT

	LASTEVENT = C.SDL_LASTEVENT
)

const (
	AUDIO_U8     AudioFormat = C.AUDIO_U8
	AUDIO_S8                 = C.AUDIO_S8
	AUDIO_U16LSB             = C.AUDIO_U16LSB
	AUDIO_S16LSB             = C.AUDIO_S16LSB
	AUDIO_U16MSB             = C.AUDIO_U16MSB
	AUDIO_S16MSB             = C.AUDIO_S16MSB
	AUDIO_U16                = C.AUDIO_U16
	AUDIO_S16                = C.AUDIO_S16
	AUDIO_S32LSB             = C.AUDIO_S32LSB
	AUDIO_S32MSB             = C.AUDIO_S32MSB
	AUDIO_S32                = C.AUDIO_S32
	AUDIO_F32LSB             = C.AUDIO_F32LSB
	AUDIO_F32MSB             = C.AUDIO_F32MSB
	AUDIO_F32                = C.AUDIO_F32
)

const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE    = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE  = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE       = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

const (
	AUDIO_STOPPED AudioStatus = C.SDL_AUDIO_STOPPED
	AUDIO_PAUSED  AudioStatus = C.SDL_AUDIO_PAUSED
	AUDIO_PLAYING AudioStatus = C.SDL_AUDIO_PLAYING
)

const (
	WINDOW_FULLSCREEN    = C.SDL_WINDOW_FULLSCREEN
	WINDOW_OPENGL        = C.SDL_WINDOW_OPENGL
	WINDOW_SHOWN         = C.SDL_WINDOW_SHOWN
	WINDOW_HIDDEN        = C.SDL_WINDOW_HIDDEN
	WINDOW_BORDERLESS    = C.SDL_WINDOW_BORDERLESS
	WINDOW_RESIZABLE     = C.SDL_WINDOW_RESIZABLE
	WINDOW_MINIMIZED     = C.SDL_WINDOW_MINIMIZED
	WINDOW_MAXIMIZED     = C.SDL_WINDOW_MAXIMIZED
	WINDOW_INPUT_GRABBED = C.SDL_WINDOW_INPUT_GRABBED
	WINDOW_INPUT_FOCUS   = C.SDL_WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS   = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_FOREIGN       = C.SDL_WINDOW_FOREIGN
)

const (
	RENDERER_SOFTWARE      = C.SDL_RENDERER_SOFTWARE
	RENDERER_ACCELERATED   = C.SDL_RENDERER_ACCELERATED
	RENDERER_PRESENTVSYNC  = C.SDL_RENDERER_PRESENTVSYNC
	RENDERER_TARGETTEXTURE = C.SDL_RENDERER_TARGETTEXTURE
)

const (
	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET
)

const (
	TEXTUREMODULATE_NONE  = C.SDL_TEXTUREMODULATE_NONE
	TEXTUREMODULATE_COLOR = C.SDL_TEXTUREMODULATE_COLOR
	TEXTUREMODULATE_ALPHA = C.SDL_TEXTUREMODULATE_ALPHA
)

const (
	BLENDMODE_NONE  = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND = C.SDL_BLENDMODE_BLEND
	BLENDMODE_ADD   = C.SDL_BLENDMODE_ADD
	BLENDMODE_MOD   = C.SDL_BLENDMODE_MOD
)

const (
	FLIP_NONE       Flip = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL      = C.SDL_FLIP_HORIZONTAL
	FLIP_VERTICAL        = C.SDL_FLIP_VERTICAL
)

const (
	PIXELFORMAT_UNKNOWN   = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_RGBA8888  = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ARGB8888  = C.SDL_PIXELFORMAT_ARGB8888
)

const (
	MESSAGEBOX_ERROR       = C.SDL_MESSAGEBOX_ERROR
	MESSAGEBOX_WARNING     = C.SDL_MESSAGEBOX_WARNING
	MESSAGEBOX_INFORMATION = C.SDL_MESSAGEBOX_INFORMATION
)

const (
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
)

const (
	MESSAGEBOX_COLOR_BACKGROUND        = C.SDL_MESSAGEBOX_COLOR_BACKGROUND
	MESSAGEBOX_COLOR_TEXT              = C.SDL_MESSAGEBOX_COLOR_TEXT
	MESSAGEBOX_COLOR_BUTTON_BORDER     = C.SDL_MESSAGEBOX_COLOR_BUTTON_BORDER
	MESSAGEBOX_COLOR_BUTTON_BACKGROUND = C.SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	MESSAGEBOX_COLOR_BUTTON_SELECTED   = C.SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED
	MESSAGEBOX_COLOR_MAX               = C.SDL_MESSAGEBOX_COLOR_MAX
)
