package sdl

// #include <SDL2/SDL.h>
import "C"

import (
	"encoding/binary"
	"errors"
	"image"
	"log"
	"unsafe"
)

// goRect converts a C SDL_Rect to a Go image.Rectangle.
func goRect(cRect C.SDL_Rect) (rect image.Rectangle) {
	x := int(cRect.x)
	y := int(cRect.y)
	w := int(cRect.w)
	h := int(cRect.h)
	return image.Rect(x, y, x+w, y+h)
}

// cRect converts a Go image.Rectangle to a C SDL_Rect.
func cRect(rect image.Rectangle) (cRect *C.SDL_Rect) {
	if rect == image.ZR {
		return nil
	}
	cRect = new(C.SDL_Rect)
	cRect.x = C.int(rect.Min.X)
	cRect.y = C.int(rect.Min.Y)
	cRect.w = C.int(rect.Max.X - rect.Min.X)
	cRect.h = C.int(rect.Max.Y - rect.Min.Y)
	return cRect
}

// nativeByteOrder corresponds to the native byte order of the system.
var nativeByteOrder binary.ByteOrder

// initNativeByteOrder determintes the native byte order of the system.
func initNativeByteOrder() (err error) {
	i := int32(0x01020304)
	p := (*byte)(unsafe.Pointer(&i))
	switch *p {
	case 0x01:
		nativeByteOrder = binary.BigEndian
		return nil
	case 0x04:
		nativeByteOrder = binary.LittleEndian
		return nil
	}
	return errors.New("sdl.initNativeByteOrder: unable to determine native byte order")
}

func init() {
	err := initNativeByteOrder()
	if err != nil {
		log.Fatalln(err)
	}
}
