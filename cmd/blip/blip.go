package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/mewpull/beep/flac"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: blip FILE\n")
		os.Exit(1)
	}
	fr, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("unable to open file; %v", err)
	}
	s, format, err := flac.Decode(fr)
	if err != nil {
		log.Fatalf("unable to decode file; %v", err)
	}
	defer s.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan struct{})
	f := func() {
		close(done)
	}
	notify := beep.Callback(f)
	speaker.Play(beep.Seq(s, notify))
	select {
	case <-done:
		break
	}
}
