package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

import (
	"unsafe"
)

// GetClipboardText returns the clipboard text.
//
// Note: A window must be created before calling this function.
func GetClipboardText() (text string, err error) {
	s := C.SDL_GetClipboardText()
	if s == nil {
		return "", getError()
	}
	text = C.GoString(s)
	C.SDL_free(unsafe.Pointer(s))
	return text, nil
}

// HasClipboardText returns true if the clipboard contains a non-empty text
// string, and false otherwise.
//
// Note: A window must be created before calling this function.
func HasClipboardText() bool {
	if C.SDL_HasClipboardText() == C.SDL_TRUE {
		return true
	}
	return false
}

// SetClipboardText sets the clipboard text to the specified string.
//
// Note: A window must be created before calling this function.
func SetClipboardText(text string) (err error) {
	if C.SDL_SetClipboardText(C.CString(text)) != 0 {
		return getError()
	}
	return nil
}
