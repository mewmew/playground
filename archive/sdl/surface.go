package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

import (
	"encoding/binary"
	"image"
)

// A Surface is a collection of pixels.
type Surface struct {
	cSurface *C.SDL_Surface
}

// CreateSurface returns a newly allocated surface.
//
// Note: Free must be called when finished using the surface.
func CreateSurface(w, h int) (s *Surface, err error) {
	s = new(Surface)
	// SDL_PIXELFORMAT_RGBA8888
	rMask := C.Uint32(0xFF000000)
	gMask := C.Uint32(0x00FF0000)
	bMask := C.Uint32(0x0000FF00)
	aMask := C.Uint32(0x000000FF)
	if nativeByteOrder == binary.BigEndian {
		// SDL_PIXELFORMAT_ABGR8888
		rMask = 0x000000FF
		gMask = 0x0000FF00
		bMask = 0x00FF0000
		aMask = 0xFF000000
	}
	s.cSurface = C.SDL_CreateRGBSurface(0, C.int(w), C.int(h), 32, rMask, gMask, bMask, aMask)
	if s.cSurface == nil {
		return nil, getError()
	}
	return s, nil
}

// Free frees a surface.
func (s *Surface) Free() {
	C.SDL_FreeSurface(s.cSurface)
}

// Blit performs a fast blit from the source surface to the destination surface.
//
// srcRect represents the rectangle to be copied or image.ZR to copy the entire
// surface.
//
// dstPoint represents the destination position.
func (s *Surface) Blit(srcRect image.Rectangle, dst *Surface, dstPoint image.Point) (err error) {
	cSrcRect := cRect(srcRect)
	cDstRect := new(C.SDL_Rect)
	cDstRect.x = C.int(dstPoint.X)
	cDstRect.y = C.int(dstPoint.Y)
	if C.SDL_BlitSurface(s.cSurface, cSrcRect, dst.cSurface, cDstRect) != 0 {
		return getError()
	}
	return nil
}

// BlitScaled performs a scaled fast blit from the source surface to the
// destination surface.
//
// srcRect represents the rectangle to be copied or image.ZR to copy the entire
// surface.
//
// dstRect represents the rectangle to be copied into.
//
// If the srcRect and dstRect differs in size, the surface will be scaled to
// fit.
func (s *Surface) BlitScaled(srcRect image.Rectangle, dst *Surface, dstRect image.Rectangle) (err error) {
	cSrcRect := cRect(srcRect)
	cDstRect := cRect(dstRect)
	if C.SDL_BlitScaled(s.cSurface, cSrcRect, dst.cSurface, cDstRect) != 0 {
		return getError()
	}
	return nil
}

/*
   SDL_BlitScaled                [done] (SDL_UpperBlitScaled)
   SDL_BlitSurface               [done] (SDL_UpperBlit)
   SDL_ConvertPixels
   SDL_ConvertSurface
   SDL_ConvertSurfaceFormat
   SDL_CreateRGBSurface          [done]
   SDL_CreateRGBSurfaceFrom
   SDL_FillRect
   SDL_FillRects
   SDL_FreeSurface               [done]
   SDL_GetClipRect
   SDL_GetColorKey
   SDL_GetSurfaceAlphaMod
   SDL_GetSurfaceBlendMode
   SDL_GetSurfaceColorMod
   SDL_LoadBMP
   SDL_LoadBMP_RW
   SDL_LockSurface
   SDL_LowerBlit                 #skip#
   SDL_LowerBlitScaled           #skip#
   SDL_MUSTLOCK
   SDL_SaveBMP
   SDL_SaveBMP_RW
   SDL_SetClipRect
   SDL_SetColorKey
   SDL_SetSurfaceAlphaMod
   SDL_SetSurfaceBlendMode
   SDL_SetSurfaceColorMod
   SDL_SetSurfacePalette
   SDL_SetSurfaceRLE
   SDL_SoftStretch
   SDL_UnlockSurface
   SDL_UpperBlit                 #skip#
   SDL_UpperBlitScaled           #skip#
*/
