package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/rhnvrm/lyric-api-go"
)

func main() {
	// Parse command line arguments.
	var (
		// Artist name.
		artist string
		// Song name.
		song string
	)
	const (
		defaultArtist = "Johnossi"
		defaultSong   = "What's the point"
	)
	flag.StringVar(&artist, "artist", defaultArtist, "artist name")
	flag.StringVar(&song, "song", defaultSong, "song name")
	flag.Parse()
	if artist == defaultArtist && song == defaultSong {
		args := flag.Args()
		switch flag.NArg() {
		case 1:
			parts := strings.Split(args[0], "-")
			artist = parts[0]
			song = parts[1]
		case 2:
			artist = args[0]
			song = args[1]
		}
	}

	// Search for lyric.
	l := lyrics.New()
	lyric, err := l.Search(artist, song)
	if err != nil {
		log.Fatalf("unable to locate lyrics for `%v - %v`", artist, song)
	}
	fmt.Println(lyric)
}
