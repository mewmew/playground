// aquarium is a backup utility for grooveshark users.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mewmew/playground/archive/gs"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: aquarium USERNAME")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, `  Create a backup of testuser's song collection, favorites and playlists.`)
	fmt.Fprintln(os.Stderr, "    aquarium testuser")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := aquarium(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

func aquarium(username string) (err error) {
	gs.Verbose = true
	sess, err := gs.NewSession()
	if err != nil {
		return err
	}
	userID, err := sess.UserID(username)
	if err != nil {
		return err
	}

	now := time.Now()
	date := now.Format("2006-01-02")
	base := fmt.Sprintf("%s - %s", username, date)
	dir := fmt.Sprintf("%s/playlists", base)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Songs.
	err = songs(sess, userID, base)
	if err != nil {
		return err
	}

	// Favorites.
	err = favorites(sess, userID, base)
	if err != nil {
		return err
	}

	// Playlists.
	err = playlists(sess, userID, dir)
	if err != nil {
		return err
	}

	return nil
}

func writeSongs(w io.Writer, songs []*gs.Song) (err error) {
	for _, song := range songs {
		_, err = fmt.Fprintln(w, song)
		if err != nil {
			return err
		}
	}
	return nil
}

func songs(sess *gs.Session, userID int, dir string) (err error) {
	filePath := fmt.Sprintf("%s/songs.txt", dir)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	songs, err := sess.UserSongs(userID)
	if err != nil {
		return err
	}
	err = writeSongs(f, songs)
	if err != nil {
		return err
	}

	return nil
}

func favorites(sess *gs.Session, userID int, dir string) (err error) {
	filePath := fmt.Sprintf("%s/favorites.txt", dir)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	songs, err := sess.UserFavorites(userID)
	if err != nil {
		return err
	}
	err = writeSongs(f, songs)
	if err != nil {
		return err
	}

	return nil
}

func playlists(sess *gs.Session, userID int, dir string) (err error) {
	playlists, err := sess.UserPlaylists(userID)
	if err != nil {
		return err
	}

	for _, playlist := range playlists {
		name := strings.Replace(playlist.Name, "/", ",", -1)
		filePath := fmt.Sprintf("%s/%s.txt", dir, name)
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		err = writeSongs(f, playlist.Songs)
		if err != nil {
			return err
		}
	}

	return nil
}
