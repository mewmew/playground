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
	url := fmt.Sprintf("https://cdnapisec.kaltura.com/p/%s/sp/%s00/playManifest/entryId/%s/flavorIds/1_qo1aqd11,1_i1aflt3v/format/applehttp/protocol/https/a.m3u8", wid, wid, entryID)
	playlists, err := httputil.GetString(url)
	if err != nil {
		return "", errors.WithStack(err)
	}
	var playlistURL string
	s := bufio.NewScanner(strings.NewReader(playlists))
	// Locate playlist URL with highest resolution.
	//
	// Example
	//
	//    playlist: #EXTM3U
	//    #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=2698240,RESOLUTION=1280x720
	//    https://cfvod.kaltura.com/hls/p/1813901/sp/181390100/serveFlavor/entryId/1_g2qcuvw5/v/1/flavorId/1_qo1aqd11/name/a.mp4/index.m3u8
	//    #EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=3524608,RESOLUTION=1920x1080
	//    https://cfvod.kaltura.com/hls/p/1813901/sp/181390100/serveFlavor/entryId/1_g2qcuvw5/v/1/flavorId/1_i1aflt3v/name/a.mp4/index.m3u8
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "https://") {
			continue
		}
		if len(playlistURL) == 0 {
			playlistURL = line
		} else {
			// 1_qo1aqd11 = 1280x720
			// 1_i1aflt3v = 1920x1080
			if strings.Contains(line, "1_i1aflt3v") {
				playlistURL = line
			}
		}
	}
	if err := s.Err(); err != nil {
		return "", errors.WithStack(err)
	}
	if len(playlistURL) == 0 {
		return "", errors.Errorf("unable to locate playlist URL in %q", playlists)
	}
	fmt.Println("playlistURL:", playlistURL)
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
	for s.Scan() {
		url := s.Text()
		if !strings.HasPrefix(url, "https://cfvod.kaltura.com/") {
			continue
		}
		urls = append(urls, url)
	}
	if err := s.Err(); err != nil {
		return errors.WithStack(err)
	}

	var segNames []string
	for i, url := range urls {
		segName := segmentName(url)
		log.Printf("downloading segment %d of %d\n", i, len(urls))
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
