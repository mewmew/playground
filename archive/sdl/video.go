package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
//
// static SDL_Rect * makeRectArray(int size) {
//    return calloc(sizeof(SDL_Rect), size);
// }
//
// static void setArrayRect(SDL_Rect *cRects, SDL_Rect *rect, int index) {
//    memcpy(cRects + sizeof(SDL_Rect) * index, rect, sizeof(SDL_Rect));
// }
import "C"

import (
	"image"
	"unsafe"
)

// --- [ window ] --------------------------------------------------------------

// A Window represents a single graphics window.
type Window struct {
	cWin *C.SDL_Window
}

// WindowFlag is a bitfield of window flags.
type WindowFlag uint32

// Window flags.
const (
	WindowFullScreen   = WindowFlag(C.SDL_WINDOW_FULLSCREEN)
	WindowOpenGL       = WindowFlag(C.SDL_WINDOW_OPENGL)
	WindowShown        = WindowFlag(C.SDL_WINDOW_SHOWN)
	WindowHidden       = WindowFlag(C.SDL_WINDOW_HIDDEN)
	WindowBorderless   = WindowFlag(C.SDL_WINDOW_BORDERLESS)
	WindowResizeable   = WindowFlag(C.SDL_WINDOW_RESIZABLE)
	WindowMinimized    = WindowFlag(C.SDL_WINDOW_MINIMIZED)
	WindowMaximized    = WindowFlag(C.SDL_WINDOW_MAXIMIZED)
	WindowInputGrabbed = WindowFlag(C.SDL_WINDOW_INPUT_GRABBED)
	WindowInputFocus   = WindowFlag(C.SDL_WINDOW_INPUT_FOCUS)
	WindowMouseFocus   = WindowFlag(C.SDL_WINDOW_MOUSE_FOCUS)
	WindowForeign      = WindowFlag(C.SDL_WINDOW_FOREIGN)
)

// WindowPos specifies a window position (x, y) or a window position behavior.
type WindowPos int

// Window position flags.
const (
	WindowPosCentered  = WindowPos(C.SDL_WINDOWPOS_CENTERED)
	WindowPosUndefined = WindowPos(C.SDL_WINDOWPOS_UNDEFINED)
)

// CreateWindow creates a window with the specified position, dimensions, and
// flags.
//
// Note: Destroy must be called when finished using the window.
func CreateWindow(title string, x, y WindowPos, w, h int, flags WindowFlag) (win *Window, err error) {
	win = new(Window)
	win.cWin = C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if win.cWin == nil {
		return nil, getError()
	}
	return win, nil
}

// Destroy destroys the window.
func (win *Window) Destroy() {
	C.SDL_DestroyWindow(win.cWin)
}

// GetFlags returns the window flags.
func (win *Window) GetFlags() (flags WindowFlag) {
	return WindowFlag(C.SDL_GetWindowFlags(win.cWin))
}

// EnterFullScreen enters full screen mode for the window.
func (win *Window) EnterFullScreen() (err error) {
	if C.SDL_SetWindowFullscreen(win.cWin, C.SDL_TRUE) != 0 {
		return getError()
	}
	return nil
}

// LeaveFullScreen leaves full screen mode for the window.
func (win *Window) LeaveFullScreen() (err error) {
	if C.SDL_SetWindowFullscreen(win.cWin, C.SDL_FALSE) != 0 {
		return getError()
	}
	return nil
}

// Show shows the window.
func (win *Window) Show() {
	C.SDL_ShowWindow(win.cWin)
}

// Hide hides the window.
func (win *Window) Hide() {
	C.SDL_HideWindow(win.cWin)
}

// Minimize minimizes the window.
func (win *Window) Minimize() {
	C.SDL_MinimizeWindow(win.cWin)
}

/// ### todo ###
///   - doesn't seem to work.
///   - minimized windows are not raised after calling Raise().
/// ############

// Raise raises a window above other windows and set the input focus.
func (win *Window) Raise() {
	C.SDL_RaiseWindow(win.cWin)
}

// Maximize maximizes the window.
func (win *Window) Maximize() {
	C.SDL_MaximizeWindow(win.cWin)
}

/// ### todo ###
///   - doesn't seem to work.
///   - maximized and/or minimized windows are not restored after calling
///     Restore().
/// ############

// Restore restores the size and position of a minimized or maximized window.
func (win *Window) Restore() {
	C.SDL_RestoreWindow(win.cWin)
}

/// ### todo ###
///   - not yet tested.
/// ############

// GrabInputFocus grabs input focus for the window.
func (win *Window) GrabInputFocus() {
	C.SDL_SetWindowGrab(win.cWin, C.SDL_TRUE)
}

/// ### todo ###
///   - not yet tested.
/// ############

// ReleaseInputFocus releases input focus for the window.
func (win *Window) ReleaseInputFocus() {
	C.SDL_SetWindowGrab(win.cWin, C.SDL_FALSE)
}

/// ### todo ###
///   - not yet tested.
/// ############

// IsInputGrabbed returns true if input is grabbed by the window.
func (win *Window) IsInputGrabbed() bool {
	if C.SDL_GetWindowGrab(win.cWin) == C.SDL_TRUE {
		return true
	}
	return false
}

/// ### todo ###
///   - not yet tested.
/// ############

// Brightness returns the brightness (gamma correction) for the window.
func (win *Window) Brightness() (brightness float64) {
	return float64(C.SDL_GetWindowBrightness(win.cWin))
}

/// ### todo ###
///   - not yet tested.
/// ############

// SetBrightness sets the brightness (gamma correction) for the window.
func (win *Window) SetBrightness(brightness float64) (err error) {
	if C.SDL_SetWindowBrightness(win.cWin, C.float(brightness)) != 0 {
		return getError()
	}
	return nil
}

// Position returns the position of the window.
func (win *Window) Position() (x, y int) {
	var cX, cY C.int
	C.SDL_GetWindowPosition(win.cWin, &cX, &cY)
	return int(cX), int(cY)
}

// SetPosition set the position of the window.
func (win *Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition(win.cWin, C.int(x), C.int(y))
}

// Size returns the size of the window's client area.
func (win *Window) Size() (w, h int) {
	var cW, cH C.int
	C.SDL_GetWindowSize(win.cWin, &cW, &cH)
	return int(cW), int(cH)
}

// SetSize sets the size of the window's client area.
func (win *Window) SetSize(w, h int) {
	C.SDL_SetWindowSize(win.cWin, C.int(w), C.int(h))
}

// Title returns the title of the window.
func (win *Window) Title() (title string) {
	return C.GoString(C.SDL_GetWindowTitle(win.cWin))
}

// SetTitle sets the title of the window.
func (win *Window) SetTitle(title string) {
	C.SDL_SetWindowTitle(win.cWin, C.CString(title))
}

// Display returns the display index associated with the window.
func (win *Window) Display() (displayIndex int, err error) {
	displayIndex = int(C.SDL_GetWindowDisplayIndex(win.cWin))
	if displayIndex < 0 {
		return 0, getError()
	}
	return displayIndex, nil
}

// Surface returns the surface associated with the window.
func (win *Window) Surface() (surface *Surface, err error) {
	surface = new(Surface)
	surface.cSurface = C.SDL_GetWindowSurface(win.cWin)
	if surface.cSurface == nil {
		return nil, getError()
	}
	return surface, nil
}

/// ### todo ###
///   - not yet tested.
/// ############

// Update copies the window surface to the screen.
//
// Note: A Surface must be associated with the window before calling this
// function.
func (win *Window) Update() (err error) {
	if C.SDL_UpdateWindowSurface(win.cWin) != 0 {
		return getError()
	}
	return nil
}

/// ### todo ###
///   - not yet tested.
/// ############

// UpdateRects copies a number of rectangles on the window surface to the
// screen.
//
// Note: A Surface must be associated with the window before calling this
// function.
func (win *Window) UpdateRects(rects []image.Rectangle) (err error) {
	cRects := C.makeRectArray(C.int(len(rects)))
	defer C.SDL_free(unsafe.Pointer(cRects))
	for index, rect := range rects {
		cRect := cRect(rect)
		C.setArrayRect(cRects, cRect, C.int(index))
	}
	if C.SDL_UpdateWindowSurfaceRects(win.cWin, cRects, C.int(len(rects))) != 0 {
		return getError()
	}
	return nil
}

// --- [ display ] -------------------------------------------------------------

// DisplayBounds returns the desktop area represented by a display, with the
// primary display located at 0,0.
func DisplayBounds(displayIndex int) (rect image.Rectangle, err error) {
	var cRect C.SDL_Rect
	if C.SDL_GetDisplayBounds(C.int(displayIndex), &cRect) != 0 {
		return image.ZR, getError()
	}
	return goRect(cRect), nil
}

// DisplayCount returns the number of available video displays.
func DisplayCount() (n int, err error) {
	n = int(C.SDL_GetNumVideoDisplays())
	if n < 0 {
		return 0, getError()
	}
	return n, nil
}

// --- [ driver ] --------------------------------------------------------------

// CurrentVideoDriver returns the name of the currently initialized video
// driver.
func CurrentVideoDriver() (driver string) {
	cDriver := C.SDL_GetCurrentVideoDriver()
	if cDriver == nil {
		return ""
	}
	return C.GoString(cDriver)
}

// VideoDriver returns the name of a built-in video driver.
func VideoDriver(index int) (driver string) {
	cDriver := C.SDL_GetVideoDriver(C.int(index))
	if cDriver == nil {
		return ""
	}
	return C.GoString(cDriver)
}

// VideoDriverCount returns the number of video drivers compiled into SDL.
func VideoDriverCount() (n int, err error) {
	n = int(C.SDL_GetNumVideoDrivers())
	if n < 0 {
		return 0, getError()
	}
	return n, nil
}

/*
   SDL_CreateWindow                 [done]
   SDL_CreateWindowFrom             #skip#
   SDL_DestroyWindow                [done]
   SDL_DisableScreenSaver           #skip#
   SDL_EnableScreenSaver            #skip#
   SDL_GetClosestDisplayMode           _display mode_
   SDL_GetCurrentDisplayMode           _display mode_
   SDL_GetCurrentVideoDriver        [done]
   SDL_GetDesktopDisplayMode           _display mode_
   SDL_GetDisplayBounds             [done]
   SDL_GetDisplayMode                  _display mode_
   SDL_GetNumDisplayModes              _display mode_
   SDL_GetNumVideoDisplays          [done]
   SDL_GetNumVideoDrivers           [done]
   SDL_GetVideoDriver               [done]
   SDL_GetWindowBrightness          [done]
   SDL_GetWindowData                #skip#
   SDL_GetWindowDisplayIndex        [done]
   SDL_GetWindowDisplayMode            _display mode_
   SDL_GetWindowFlags               [done]
   SDL_GetWindowFromID              #skip#
   SDL_GetWindowGammaRamp              _gamma_
   SDL_GetWindowGrab                [done]
   SDL_GetWindowID                  #skip#
   SDL_GetWindowPixelFormat
   SDL_GetWindowPosition            [done]
   SDL_GetWindowSize                [done]
   SDL_GetWindowSurface             [done]
   SDL_GetWindowTitle               [done]
   SDL_GetWindowWMInfo
   SDL_GL_CreateContext                _OpenGL_
   SDL_GL_DeleteContext                _OpenGL_
   SDL_GL_ExtensionSupported           _OpenGL_
   SDL_GL_GetAttribute                 _OpenGL_
   SDL_GL_GetProcAddress               _OpenGL_
   SDL_GL_GetSwapInterval              _OpenGL_
   SDL_GL_LoadLibrary                  _OpenGL_
   SDL_GL_MakeCurrent                  _OpenGL_
   SDL_GL_SetAttribute                 _OpenGL_
   SDL_GL_SetSwapInterval              _OpenGL_
   SDL_GL_SwapWindow                   _OpenGL_
   SDL_GL_UnloadLibrary                _OpenGL_
   SDL_HideWindow                   [done]
   SDL_IsScreenSaverEnabled         #skip#
   SDL_MaximizeWindow               [done]
   SDL_MinimizeWindow               [done]
   SDL_RaiseWindow                  [done]
   SDL_RestoreWindow                [done]
   SDL_SetWindowBrightness          [done]
   SDL_SetWindowData                #skip#
   SDL_SetWindowDisplayMode            _display mode_
   SDL_SetWindowFullscreen          [done]
   SDL_SetWindowGammaRamp              _gamma_
   SDL_SetWindowGrab                [done]
   SDL_SetWindowIcon                   _window_
   SDL_SetWindowPosition            [done]
   SDL_SetWindowSize                [done]
   SDL_SetWindowTitle               [done]
   SDL_ShowWindow                   [done]
   SDL_UpdateWindowSurface          [done]
   SDL_UpdateWindowSurfaceRects     [done]
   SDL_VideoInit                    #skip#
   SDL_VideoQuit                    #skip#
*/
