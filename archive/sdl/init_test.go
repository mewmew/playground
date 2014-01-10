package sdl

import (
	"testing"
)

func TestInit(t *testing.T) {
	Quit() // shouldn't be needed.

	// init audio
	err := Init(InitAudio)
	if err != nil {
		t.Errorf("failed to initialize audio: %s.", err)
	}
	if WasInit(InitAudio) != InitAudio {
		t.Errorf("audio wasn't initialized.")
	}
	Quit()

	// init video
	err = Init(InitVideo)
	if err != nil {
		t.Errorf("failed to initialize video: %s.", err)
	}
	if WasInit(InitVideo) != InitVideo {
		t.Errorf("video wasn't initialized.")
	}
	Quit()

	// init without parachute
	err = Init(InitNoParachute)
	if err != nil {
		t.Errorf("failed to initialize without parachute: %s.", err)
	}
	Quit()

	// init everything
	err = Init(InitEverything)
	if err != nil {
		t.Errorf("failed to initialize everything: %s.", err)
	}
	if WasInit(InitEverything) != InitEverything {
		t.Errorf("everything wasn't initialized.")
	}
	Quit()
}

func TestInitSubSystem(t *testing.T) {
	Quit() // shouldn't be needed.

	inited := InitFlag(0)

	// init audio
	err := InitSubSystem(InitAudio)
	if err != nil {
		t.Errorf("failed to initialize audio: %s.", err)
	}
	inited |= InitAudio
	if WasInit(inited) != inited {
		t.Errorf("audio wasn't initialized.")
	}

	// init video
	err = InitSubSystem(InitVideo)
	if err != nil {
		t.Errorf("failed to initialize video: %s.", err)
	}
	inited |= InitVideo
	if WasInit(inited) != inited {
		t.Errorf("video wasn't initialized.")
	}

	Quit()
}

func TestQuit(t *testing.T) {
	Quit() // shouldn't be needed.

	// check after quit audio
	err := Init(InitAudio)
	if err != nil {
		t.Errorf("failed to initialize audio: %s.", err)
	}
	if WasInit(InitAudio) != InitAudio {
		t.Errorf("audio wasn't initialized.")
	}
	Quit()
	if WasInit(InitAudio) != 0 {
		t.Errorf("audio is still initialized after quit.")
	}

	// check after quit video
	err = Init(InitVideo)
	if err != nil {
		t.Errorf("failed to initialize video: %s.", err)
	}
	if WasInit(InitVideo) != InitVideo {
		t.Errorf("video wasn't initialized.")
	}
	Quit()
	if WasInit(InitVideo) != 0 {
		t.Errorf("video is still initialized after quit.")
	}

	// check after quit everything
	err = Init(InitEverything)
	if err != nil {
		t.Errorf("failed to initialize everything: %s.", err)
	}
	if WasInit(InitEverything) != InitEverything {
		t.Errorf("everything wasn't initialized.")
	}
	Quit()
	if WasInit(InitEverything) != 0 {
		t.Errorf("something is still initialized after quit.")
	}
}

func TestQuitSubSystem(t *testing.T) {
	Quit() // shouldn't be needed.

	// init everything
	err := Init(InitEverything)
	if err != nil {
		t.Errorf("failed to initialize everything: %s.", err)
	}
	inited := InitEverything
	if WasInit(InitEverything) != inited {
		t.Errorf("everything wasn't initialized.")
	}

	// check after quit audio
	QuitSubSystem(InitAudio)
	inited &^= InitAudio
	if WasInit(InitEverything) != inited {
		t.Errorf("audio is still initialized.")
	}

	// check after quit video
	QuitSubSystem(InitVideo)
	inited &^= InitVideo
	if WasInit(InitEverything) != inited {
		t.Errorf("video is still initialized.")
	}

	Quit()
}

func TestWasInit(t *testing.T) {
	Quit() // shouldn't be needed.

	// check after init audio
	err := Init(InitAudio)
	if err != nil {
		t.Errorf("failed to initialize audio: %s.", err)
	}
	if WasInit(InitAudio) != InitAudio {
		t.Errorf("WasInit believes audio wasn't initialized.")
	}
	Quit()

	// check after init video
	err = Init(InitVideo)
	if err != nil {
		t.Errorf("failed to initialize video: %s.", err)
	}
	if WasInit(InitVideo) != InitVideo {
		t.Errorf("WasInit believes video wasn't initialized.")
	}
	Quit()

	// check after init everything
	err = Init(InitEverything)
	if err != nil {
		t.Errorf("failed to initialize everything: %s.", err)
	}
	if WasInit(InitEverything) != InitEverything {
		t.Errorf("WasInit believes everything wasn't initialized.")
	}
	Quit()
}
