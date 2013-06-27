package sdl

/*
#include <SDL.h>
*/
import "C"

const (
	KEYDOWN = C.SDL_KEYDOWN
	KEYUP   = C.SDL_KEYUP
)

const (
	K_UNKNOWN            = C.SDLK_UNKNOWN
	K_RETURN             = C.SDLK_RETURN
	K_ESCAPE             = C.SDLK_ESCAPE
	K_BACKSPACE          = C.SDLK_BACKSPACE
	K_TAB                = C.SDLK_TAB
	K_SPACE              = C.SDLK_SPACE
	K_EXCLAIM            = C.SDLK_EXCLAIM
	K_QUOTEDBL           = C.SDLK_QUOTEDBL
	K_HASH               = C.SDLK_HASH
	K_PERCENT            = C.SDLK_PERCENT
	K_DOLLAR             = C.SDLK_DOLLAR
	K_AMPERSAND          = C.SDLK_AMPERSAND
	K_QUOTE              = C.SDLK_QUOTE
	K_LEFTPAREN          = C.SDLK_LEFTPAREN
	K_RIGHTPAREN         = C.SDLK_RIGHTPAREN
	K_ASTERISK           = C.SDLK_ASTERISK
	K_PLUS               = C.SDLK_PLUS
	K_COMMA              = C.SDLK_COMMA
	K_MINUS              = C.SDLK_MINUS
	K_PERIOD             = C.SDLK_PERIOD
	K_SLASH              = C.SDLK_SLASH
	K_0                  = C.SDLK_0
	K_1                  = C.SDLK_1
	K_2                  = C.SDLK_2
	K_3                  = C.SDLK_3
	K_4                  = C.SDLK_4
	K_5                  = C.SDLK_5
	K_6                  = C.SDLK_6
	K_7                  = C.SDLK_7
	K_8                  = C.SDLK_8
	K_9                  = C.SDLK_9
	K_COLON              = C.SDLK_COLON
	K_SEMICOLON          = C.SDLK_SEMICOLON
	K_LESS               = C.SDLK_LESS
	K_EQUALS             = C.SDLK_EQUALS
	K_GREATER            = C.SDLK_GREATER
	K_QUESTION           = C.SDLK_QUESTION
	K_AT                 = C.SDLK_AT
	K_LEFTBRACKET        = C.SDLK_LEFTBRACKET
	K_BACKSLASH          = C.SDLK_BACKSLASH
	K_RIGHTBRACKET       = C.SDLK_RIGHTBRACKET
	K_CARET              = C.SDLK_CARET
	K_UNDERSCORE         = C.SDLK_UNDERSCORE
	K_BACKQUOTE          = C.SDLK_BACKQUOTE
	K_a                  = C.SDLK_a
	K_b                  = C.SDLK_b
	K_c                  = C.SDLK_c
	K_d                  = C.SDLK_d
	K_e                  = C.SDLK_e
	K_f                  = C.SDLK_f
	K_g                  = C.SDLK_g
	K_h                  = C.SDLK_h
	K_i                  = C.SDLK_i
	K_j                  = C.SDLK_j
	K_k                  = C.SDLK_k
	K_l                  = C.SDLK_l
	K_m                  = C.SDLK_m
	K_n                  = C.SDLK_n
	K_o                  = C.SDLK_o
	K_p                  = C.SDLK_p
	K_q                  = C.SDLK_q
	K_r                  = C.SDLK_r
	K_s                  = C.SDLK_s
	K_t                  = C.SDLK_t
	K_u                  = C.SDLK_u
	K_v                  = C.SDLK_v
	K_w                  = C.SDLK_w
	K_x                  = C.SDLK_x
	K_y                  = C.SDLK_y
	K_z                  = C.SDLK_z
	K_CAPSLOCK           = C.SDLK_CAPSLOCK
	K_F1                 = C.SDLK_F1
	K_F2                 = C.SDLK_F2
	K_F3                 = C.SDLK_F3
	K_F4                 = C.SDLK_F4
	K_F5                 = C.SDLK_F5
	K_F6                 = C.SDLK_F6
	K_F7                 = C.SDLK_F7
	K_F8                 = C.SDLK_F8
	K_F9                 = C.SDLK_F9
	K_F10                = C.SDLK_F10
	K_F11                = C.SDLK_F11
	K_F12                = C.SDLK_F12
	K_PRINTSCREEN        = C.SDLK_PRINTSCREEN
	K_SCROLLLOCK         = C.SDLK_SCROLLLOCK
	K_PAUSE              = C.SDLK_PAUSE
	K_INSERT             = C.SDLK_INSERT
	K_HOME               = C.SDLK_HOME
	K_PAGEUP             = C.SDLK_PAGEUP
	K_DELETE             = C.SDLK_DELETE
	K_END                = C.SDLK_END
	K_PAGEDOWN           = C.SDLK_PAGEDOWN
	K_RIGHT              = C.SDLK_RIGHT
	K_LEFT               = C.SDLK_LEFT
	K_DOWN               = C.SDLK_DOWN
	K_UP                 = C.SDLK_UP
	K_NUMLOCKCLEAR       = C.SDLK_NUMLOCKCLEAR
	K_KP_DIVIDE          = C.SDLK_KP_DIVIDE
	K_KP_MULTIPLY        = C.SDLK_KP_MULTIPLY
	K_KP_MINUS           = C.SDLK_KP_MINUS
	K_KP_PLUS            = C.SDLK_KP_PLUS
	K_KP_ENTER           = C.SDLK_KP_ENTER
	K_KP_1               = C.SDLK_KP_1
	K_KP_2               = C.SDLK_KP_2
	K_KP_3               = C.SDLK_KP_3
	K_KP_4               = C.SDLK_KP_4
	K_KP_5               = C.SDLK_KP_5
	K_KP_6               = C.SDLK_KP_6
	K_KP_7               = C.SDLK_KP_7
	K_KP_8               = C.SDLK_KP_8
	K_KP_9               = C.SDLK_KP_9
	K_KP_0               = C.SDLK_KP_0
	K_KP_PERIOD          = C.SDLK_KP_PERIOD
	K_APPLICATION        = C.SDLK_APPLICATION
	K_POWER              = C.SDLK_POWER
	K_KP_EQUALS          = C.SDLK_KP_EQUALS
	K_F13                = C.SDLK_F13
	K_F14                = C.SDLK_F14
	K_F15                = C.SDLK_F15
	K_F16                = C.SDLK_F16
	K_F17                = C.SDLK_F17
	K_F18                = C.SDLK_F18
	K_F19                = C.SDLK_F19
	K_F20                = C.SDLK_F20
	K_F21                = C.SDLK_F21
	K_F22                = C.SDLK_F22
	K_F23                = C.SDLK_F23
	K_F24                = C.SDLK_F24
	K_EXECUTE            = C.SDLK_EXECUTE
	K_HELP               = C.SDLK_HELP
	K_MENU               = C.SDLK_MENU
	K_SELECT             = C.SDLK_SELECT
	K_STOP               = C.SDLK_STOP
	K_AGAIN              = C.SDLK_AGAIN
	K_UNDO               = C.SDLK_UNDO
	K_CUT                = C.SDLK_CUT
	K_COPY               = C.SDLK_COPY
	K_PASTE              = C.SDLK_PASTE
	K_FIND               = C.SDLK_FIND
	K_MUTE               = C.SDLK_MUTE
	K_VOLUMEUP           = C.SDLK_VOLUMEUP
	K_VOLUMEDOWN         = C.SDLK_VOLUMEDOWN
	K_KP_COMMA           = C.SDLK_KP_COMMA
	K_KP_EQUALSAS400     = C.SDLK_KP_EQUALSAS400
	K_ALTERASE           = C.SDLK_ALTERASE
	K_SYSREQ             = C.SDLK_SYSREQ
	K_CANCEL             = C.SDLK_CANCEL
	K_CLEAR              = C.SDLK_CLEAR
	K_PRIOR              = C.SDLK_PRIOR
	K_RETURN2            = C.SDLK_RETURN2
	K_SEPARATOR          = C.SDLK_SEPARATOR
	K_OUT                = C.SDLK_OUT
	K_OPER               = C.SDLK_OPER
	K_CLEARAGAIN         = C.SDLK_CLEARAGAIN
	K_CRSEL              = C.SDLK_CRSEL
	K_EXSEL              = C.SDLK_EXSEL
	K_KP_00              = C.SDLK_KP_00
	K_KP_000             = C.SDLK_KP_000
	K_THOUSANDSSEPARATOR = C.SDLK_THOUSANDSSEPARATOR
	K_DECIMALSEPARATOR   = C.SDLK_DECIMALSEPARATOR
	K_CURRENCYUNIT       = C.SDLK_CURRENCYUNIT
	K_CURRENCYSUBUNIT    = C.SDLK_CURRENCYSUBUNIT
	K_KP_LEFTPAREN       = C.SDLK_KP_LEFTPAREN
	K_KP_RIGHTPAREN      = C.SDLK_KP_RIGHTPAREN
	K_KP_LEFTBRACE       = C.SDLK_KP_LEFTBRACE
	K_KP_RIGHTBRACE      = C.SDLK_KP_RIGHTBRACE
	K_KP_TAB             = C.SDLK_KP_TAB
	K_KP_BACKSPACE       = C.SDLK_KP_BACKSPACE
	K_KP_A               = C.SDLK_KP_A
	K_KP_B               = C.SDLK_KP_B
	K_KP_C               = C.SDLK_KP_C
	K_KP_D               = C.SDLK_KP_D
	K_KP_E               = C.SDLK_KP_E
	K_KP_F               = C.SDLK_KP_F
	K_KP_XOR             = C.SDLK_KP_XOR
	K_KP_POWER           = C.SDLK_KP_POWER
	K_KP_PERCENT         = C.SDLK_KP_PERCENT
	K_KP_LESS            = C.SDLK_KP_LESS
	K_KP_GREATER         = C.SDLK_KP_GREATER
	K_KP_AMPERSAND       = C.SDLK_KP_AMPERSAND
	K_KP_DBLAMPERSAND    = C.SDLK_KP_DBLAMPERSAND
	K_KP_VERTICALBAR     = C.SDLK_KP_VERTICALBAR
	K_KP_DBLVERTICALBAR  = C.SDLK_KP_DBLVERTICALBAR
	K_KP_COLON           = C.SDLK_KP_COLON
	K_KP_HASH            = C.SDLK_KP_HASH
	K_KP_SPACE           = C.SDLK_KP_SPACE
	K_KP_AT              = C.SDLK_KP_AT
	K_KP_EXCLAM          = C.SDLK_KP_EXCLAM
	K_KP_MEMSTORE        = C.SDLK_KP_MEMSTORE
	K_KP_MEMRECALL       = C.SDLK_KP_MEMRECALL
	K_KP_MEMCLEAR        = C.SDLK_KP_MEMCLEAR
	K_KP_MEMADD          = C.SDLK_KP_MEMADD
	K_KP_MEMSUBTRACT     = C.SDLK_KP_MEMSUBTRACT
	K_KP_MEMMULTIPLY     = C.SDLK_KP_MEMMULTIPLY
	K_KP_MEMDIVIDE       = C.SDLK_KP_MEMDIVIDE
	K_KP_PLUSMINUS       = C.SDLK_KP_PLUSMINUS
	K_KP_CLEAR           = C.SDLK_KP_CLEAR
	K_KP_CLEARENTRY      = C.SDLK_KP_CLEARENTRY
	K_KP_BINARY          = C.SDLK_KP_BINARY
	K_KP_OCTAL           = C.SDLK_KP_OCTAL
	K_KP_DECIMAL         = C.SDLK_KP_DECIMAL
	K_KP_HEXADECIMAL     = C.SDLK_KP_HEXADECIMAL
	K_LCTRL              = C.SDLK_LCTRL
	K_LSHIFT             = C.SDLK_LSHIFT
	K_LALT               = C.SDLK_LALT
	K_LGUI               = C.SDLK_LGUI
	K_RCTRL              = C.SDLK_RCTRL
	K_RSHIFT             = C.SDLK_RSHIFT
	K_RALT               = C.SDLK_RALT
	K_RGUI               = C.SDLK_RGUI
	K_MODE               = C.SDLK_MODE
	K_AUDIONEXT          = C.SDLK_AUDIONEXT
	K_AUDIOPREV          = C.SDLK_AUDIOPREV
	K_AUDIOSTOP          = C.SDLK_AUDIOSTOP
	K_AUDIOPLAY          = C.SDLK_AUDIOPLAY
	K_AUDIOMUTE          = C.SDLK_AUDIOMUTE
	K_MEDIASELECT        = C.SDLK_MEDIASELECT
	K_WWW                = C.SDLK_WWW
	K_MAIL               = C.SDLK_MAIL
	K_CALCULATOR         = C.SDLK_CALCULATOR
	K_COMPUTER           = C.SDLK_COMPUTER
	K_AC_SEARCH          = C.SDLK_AC_SEARCH
	K_AC_HOME            = C.SDLK_AC_HOME
	K_AC_BACK            = C.SDLK_AC_BACK
	K_AC_FORWARD         = C.SDLK_AC_FORWARD
	K_AC_STOP            = C.SDLK_AC_STOP
	K_AC_REFRESH         = C.SDLK_AC_REFRESH
	K_AC_BOOKMARKS       = C.SDLK_AC_BOOKMARKS
	K_BRIGHTNESSDOWN     = C.SDLK_BRIGHTNESSDOWN
	K_BRIGHTNESSUP       = C.SDLK_BRIGHTNESSUP
	K_DISPLAYSWITCH      = C.SDLK_DISPLAYSWITCH
	K_KBDILLUMTOGGLE     = C.SDLK_KBDILLUMTOGGLE
	K_KBDILLUMDOWN       = C.SDLK_KBDILLUMDOWN
	K_KBDILLUMUP         = C.SDLK_KBDILLUMUP
	K_EJECT              = C.SDLK_EJECT
	K_SLEEP              = C.SDLK_SLEEP
)

const (
	KMOD_NONE     = C.KMOD_NONE
	KMOD_LSHIFT   = C.KMOD_LSHIFT
	KMOD_RSHIFT   = C.KMOD_RSHIFT
	KMOD_LCTRL    = C.KMOD_LCTRL
	KMOD_RCTRL    = C.KMOD_RCTRL
	KMOD_LALT     = C.KMOD_LALT
	KMOD_RALT     = C.KMOD_RALT
	KMOD_LGUI     = C.KMOD_LGUI
	KMOD_RGUI     = C.KMOD_RGUI
	KMOD_NUM      = C.KMOD_NUM
	KMOD_CAPS     = C.KMOD_CAPS
	KMOD_MODE     = C.KMOD_MODE
	KMOD_RESERVED = C.KMOD_RESERVED
)
