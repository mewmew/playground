// flac2wav is a tool which converts FLAC files to WAV files.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"azul3d.org/audio.v1"
	_ "azul3d.org/audio/flac.dev"
	"github.com/davecheney/profile"
	"azul3d.org/audio/wav.v1"
	"github.com/mewkiz/pkg/osutil"
	"github.com/mewkiz/pkg/pathutil"
)

// flagForce specifies if file overwriting should be forced, when a WAV file of
// the same name already exists.
var flagForce bool

func init() {
	flag.BoolVar(&flagForce, "f", false, "Force overwrite.")
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	flag.Parse()
	for _, path := range flag.Args() {
		err := flac2wav(path)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// flac2wav converts the provided FLAC file to a WAV file.
func flac2wav(path string) error {
	// Open FLAC file.
	fr, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fr.Close()

	// Create FLAC decoder.
	dec, magic, err := audio.NewDecoder(fr)
	if err != nil {
		return err
	}
	fmt.Println("magic:", magic)
	conf := dec.Config()
	fmt.Println("conf:", conf)

	// Create WAV file.
	wavPath := pathutil.TrimExt(path) + ".wav"
	if !flagForce {
		exists, err := osutil.Exists(wavPath)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("the file %q exists already.", wavPath)
		}
	}
	fw, err := os.Create(wavPath)
	if err != nil {
		return err
	}
	defer fw.Close()

	// Create WAV encoder.
	enc, err := wav.NewEncoder(fw, conf)
	if err != nil {
		return err
	}
	defer enc.Close()

	// Encode WAV audio samples copied from the FLAC decoder.
	// TODO(u): Replace with audio.Copy as soon as that doesn't cause audio
	// sample conversions.
	buf := make(audio.PCM16Samples, (32*1024)/8)
	for {
		nr, er := dec.Read(buf)
		if nr > 0 {
			nw, ew := enc.Write(buf[0:nr])
			if ew != nil {
				return ew
			}
			if nr != nw {
				return audio.ErrShortWrite
			}
		}
		if er == audio.EOS {
			return nil
		}
		if er != nil {
			return er
		}
	}
}
