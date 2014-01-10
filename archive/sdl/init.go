package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

// InitFlag is a bitfield of subsystem initialization flags.
type InitFlag uint32

// SDL subsystem init flags.
const (
	InitAudio       = InitFlag(C.SDL_INIT_AUDIO)
	InitVideo       = InitFlag(C.SDL_INIT_VIDEO)
	InitNoParachute = InitFlag(C.SDL_INIT_NOPARACHUTE) // Don't catch fatal signals.
	InitEverything  = InitAudio | InitVideo
)

///   InitTimer         = InitFlag(C.SDL_INIT_TIMER)
///   InitJoystick      = InitFlag(C.SDL_INIT_JOYSTICK)
///   InitHaptic        = InitFlag(C.SDL_INIT_HAPTIC)
///   InitEverything    = InitFlag(C.SDL_INIT_EVERYTHING)

// Init initializes the subsystems specified by flags.
//
// Note: Init must be called before using any other SDL function.
//
// Note: Quit must be called when finished using the SDL library.
//
// Note: Unless the InitNoParachute flag is set, it will install cleanup signal
// handlers for some commonly ignored fatal signals (like SIGSEGV).
func Init(flags InitFlag) (err error) {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return getError()
	}
	return nil
}

// InitSubSystem initializes specific SDL subsystems.
//
// Note: QuitSubSystem should be called when finished using the subsystem.
//
// Note: Init initializes assertions and crash protection. If you want to bypass
// those protections you can call InitSubSystem directly.
func InitSubSystem(flags InitFlag) (err error) {
	if C.SDL_InitSubSystem(C.Uint32(flags)) != 0 {
		return getError()
	}
	return nil
}

// Quit cleans up all initialized subsystems.
//
// Note: Quit should be called upon all exit conditions.
func Quit() {
	C.SDL_Quit()
}

// QuitSubSystem cleans up specific SDL subsystems.
func QuitSubSystem(flags InitFlag) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// WasInit returns a mask of the specified subsystems which have previously been
// initialized.
//
// Note: The return value doesn't include InitNoParachute.
func WasInit(flags InitFlag) (wasInitFlags InitFlag) {
	return InitFlag(C.SDL_WasInit(C.Uint32(flags)))
}
