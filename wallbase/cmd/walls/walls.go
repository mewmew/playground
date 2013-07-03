/*

walls updates the desktop wallpaper at specified time intervals.

It performs a search for wallpapers on http://wallbase.cc/ based on the provided
search query. The wallpaper result order is random.

For persistent storage a custom download directory can be specified.

Installation:

	go get github.com/mewmew/wallbase/cmd/walls

Dependencies:

The `hsetroot` command is used to set the wallpaper.

Usage:

	walls [OPTION]... QUERY

Flags:

	-o (default="/tmp/wallbase")
		Output directory.
	-t (default="30m")
		Timeout interval between updates.

Examples:

1. Search for "nature waterfall" and update active wallpaper each 10s.

	wallbase -t 10s nature waterfall

2. Search for "nature" and store each wallpaper in "download/".

	wallbase -t 0s -o download/ nature

*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mewmew/playground/wallbase"
)

// flagTimeout is the timeout interval between wallpaper updates.
var flagTimeout string

// wallPath is the output directory in which all wallpapers are stored. The
// default is a none persistent directory.
var wallPath string

func init() {
	flag.StringVar(&wallPath, "o", "/tmp/wallbase", "Output directory.")
	flag.StringVar(&flagTimeout, "t", "30m", "Timeout interval between updates.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "walls [OPTION]... QUERY")
	fmt.Fprintln(os.Stderr, "Update the desktop wallpaper at specified time intervals.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, `  Search for "nature waterfall" and update active wallpaper each 10s.`)
	fmt.Fprintln(os.Stderr, "    walls -t 10s nature waterfall")
	fmt.Fprintln(os.Stderr, `  Search for "nature" and store each wallpaper in "download/".`)
	fmt.Fprintln(os.Stderr, "    walls -t 0s -o download/ nature")
}

func main() {
	flag.Parse()
	for {
		err := walls()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(10 * time.Second)
	}
}

// walls performs a search on wallbase.cc based on the provided search query,
// with random search result order, and updates the active wallpaper after the
// specified timeout interval.
func walls() (err error) {
	timeout, err := time.ParseDuration(flagTimeout)
	if err != nil {
		return err
	}
	err = os.MkdirAll(wallPath, 0755)
	if err != nil {
		return err
	}
	query := strings.Join(flag.Args(), " ")
	for {
		// Each call to search should return new wallpapers, since the search
		// result order is random.
		walls, err := wallbase.Search(query)
		if err != nil {
			return err
		}
		for _, wall := range walls {
			start := time.Now()
			err = update(wall)
			if err != nil {
				return err
			}
			elapsed := time.Since(start)
			time.Sleep(timeout - elapsed)
		}
	}
	return nil
}

// update downloads the provided wallpaper and updates the current wallpaper.
func update(wall *wallbase.Wall) (err error) {
	err = wall.Download()
	if err != nil {
		return err
	}
	imgPath := fmt.Sprintf("%s/%d.%s", wallPath, wall.Id, wall.Ext)
	err = ioutil.WriteFile(imgPath, wall.Buf, 0644)
	if err != nil {
		return err
	}
	cmd := exec.Command("hsetroot", "-fill", imgPath)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
