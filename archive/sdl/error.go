package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

import (
	"errors"
)

// getError returns the last error message.
func getError() (err error) {
	return errors.New(C.GoString(C.SDL_GetError()))
}
