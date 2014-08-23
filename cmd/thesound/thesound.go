// thesound is a tool which extracts the original sound from videos. The raw
// audio stream is copied (not converted). It uses ffprobe to determine the
// audio codec, and ffmpeg to extract the sound.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/mewkiz/pkg/pathutil"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "thesound PATH...")
	fmt.Fprintln(os.Stderr, "Extract the original sound from videos.")
}

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		err := extract(path)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// extract extracts the original sound from the input video.
func extract(in string) error {
	// Probe to determine the audio codec.
	cmd := exec.Command("ffprobe", "-show_streams", "-select_streams", "a", in)
	buf, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("audio codec probe failed; %v", err)
	}
	re := regexp.MustCompile("codec_name=(.*)")
	matches := re.FindSubmatch(buf)
	if len(matches) < 2 {
		return fmt.Errorf("unable to locate codec_name")
	}
	codec := string(matches[1])

	// Copy the original sound from the input video.
	stderr := new(bytes.Buffer)
	out := fmt.Sprintf("%s.%s", pathutil.TrimExt(in), codec)
	cmd = exec.Command("ffmpeg", "-i", in, "-vn", "-y", "-acodec", "copy", out)
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		log.Println(stderr.String())
		return err
	}
	fmt.Printf("Created %q.\n", out)

	return nil
}
