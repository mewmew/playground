// The replay tool downloads video lectures from play.kth.se.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mewkiz/pkg/httputil"
	"github.com/mewkiz/pkg/stringsutil"
	"github.com/pkg/errors"
)

func usage() {
	const use = `
Download video lectures from play.kth.se.

Usage:

	replay [OPTION]... URL

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// output specifies the output path.
	var output string
	flag.StringVar(&output, "o", "", "output path")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	url := flag.Arg(0)
	if err := replay(url, output); err != nil {
		log.Fatalf("%+v", err)
	}
}

// replay downloads the video lecture at the specified URL and stores it to the
// output path.
func replay(url, output string) error {
	title, wid, entryID, err := parseVideoPage(url)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(output) == 0 {
		output = fmt.Sprintf("%s.mp4", title)
	}
	playlist, err := downloadPlaylist(wid, entryID)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := downloadSegments(playlist, output); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// parseVideoPage extract the wid and entry ID from the HTML of the video page.
func parseVideoPage(url string) (title, wid, entryID string, err error) {
	page, err := httputil.GetString(url)
	if err != nil {
		return "", "", "", errors.WithStack(err)
	}
	title, err = part(page, "<title>", " - KTH Play</title>")
	if err != nil {
		return "", "", "", errors.WithStack(err)
	}
	wid, err = part(page, `'wid': '_`, `',`)
	if err != nil {
		return "", "", "", errors.WithStack(err)
	}
	entryID, err = part(page, `'entry_id' : '`, `',`)
	if err != nil {
		return "", "", "", errors.WithStack(err)
	}
	return title, wid, entryID, nil
}

// part returns the first occurrence of a string in s with the specified before
// and after patterns.
func part(s, before, after string) (string, error) {
	start := stringsutil.IndexAfter(s, before)
	if start == -1 {
		return "", errors.Errorf("unable to locate start %q in %q", before, s)
	}
	s = s[start:]
	end := strings.Index(s, after)
	if end == -1 {
		return "", errors.Errorf("unable to locate end %q in %q", after, s)
	}
	return s[:end], nil
}

// downloadPlaylist downloads the playlist of the given wid and entryID.
func downloadPlaylist(wid, entryID string) (string, error) {
	resolution := "0_63zp42a8"
	playlistURL := fmt.Sprintf("https://streaming.kaltura.nordu.net/hls/p/%s/sp/30800/serveFlavor/entryId/%s/v/2/ev/4/flavorId/%s/name/a.mp4/index.m3u8", wid, entryID, resolution)
	playlist, err := httputil.GetString(playlistURL)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return playlist, nil
}

// downloadSegments downloads the video segments of the given playlist.
func downloadSegments(playlist, output string) error {
	// Parse m3u playlist.
	s := bufio.NewScanner(strings.NewReader(playlist))
	var urls []string
	// #EXTM3U
	// #EXT-X-TARGETDURATION:10
	// #EXT-X-ALLOW-CACHE:YES
	// #EXT-X-PLAYLIST-TYPE:VOD
	// #EXT-X-VERSION:3
	// #EXT-X-MEDIA-SEQUENCE:1
	// #EXTINF:2.000,
	// seg-1-v1-a1.ts
	// #EXTINF:2.000,
	// seg-2-v1-a1.ts
	// ...
	// #EXTINF:6.025,
	// seg-65-v1-a1.ts
	// #EXT-X-ENDLIST
	for s.Scan() {
		seg := s.Text()
		if !strings.HasPrefix(seg, "seg-") {
			continue
		}
		url := "https://streaming.kaltura.nordu.net/hls/p/308/sp/30800/serveFlavor/entryId/0_91vao47k/v/2/ev/4/flavorId/0_lzx3ynr5/name/a.mp4/" + seg
		urls = append(urls, url)
	}
	if err := s.Err(); err != nil {
		return errors.WithStack(err)
	}

	var segNames []string
	for i, url := range urls {
		segName := segmentName(url)
		log.Printf("downloading segment %d of %d\n", i+1, len(urls))
		if err := downloadSegment(url, segName); err != nil {
			return errors.WithStack(err)
		}
		segNames = append(segNames, segName)
	}

	if err := merge(segNames, output); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// segmentName returns the name of the given video segment, based on its URL.
func segmentName(s string) string {
	pos := strings.Index(s, "a.mp4/seg-")
	if pos == -1 {
		panic(fmt.Errorf("unable to locate video segment name in %q", s))
	}
	name := s[pos+len("a.mp4/seg-"):]
	end := strings.IndexByte(name, '-')
	if end == -1 {
		panic(fmt.Errorf("unable to locate end of video segment name in %q", name))
	}
	return fmt.Sprintf("seg_%s.mp4", name[:end])
}

// downloadSegment downloads the given video segment and stores it to the
// specified output path.
func downloadSegment(url string, dst string) error {
	buf, err := httputil.Get(url)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := ioutil.WriteFile(dst, buf, 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// merge merges the given video segments into a single video file.
func merge(segNames []string, output string) error {
	data := &bytes.Buffer{}
	for _, segName := range segNames {
		fmt.Fprintf(data, "file '%s'\n", segName)
	}
	listfile := "segment_list.txt"
	if err := ioutil.WriteFile(listfile, data.Bytes(), 0644); err != nil {
		return errors.WithStack(err)
	}
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", listfile, "-c", "copy", output)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.WithStack(err)
	}
	if err := os.Remove(listfile); err != nil {
		return errors.WithStack(err)
	}
	for _, segName := range segNames {
		if err := os.Remove(segName); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
