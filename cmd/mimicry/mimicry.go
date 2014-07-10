// The mimicry tool creates a git repository which mimics the image using a
// contribution history of carefully crafted commit dates. It expects a 51x7
// image with a transparent background.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/pathutil"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: mimicry IMAGE")
	fmt.Fprintln(os.Stderr, "Creates a git repository which mimics the image using a contribution history of carefully crafted commit dates. It expects a 51x7 image with a transparent background.")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := mimicry(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// Expected image dimensions.
const (
	Width  = 51
	Height = 7
)

// mimicry creates a git repository which mimics the image using a contribution
// history of carefully crafted commit dates. It expects a 51x7 image with a
// transparent background.
func mimicry(imgPath string) (err error) {
	img, err := imgutil.ReadFile(imgPath)
	if err != nil {
		return err
	}

	// Verify image dimensions.
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	if width != Width || height != Height {
		return fmt.Errorf("mimicry: invalid image dimensions; expected %dx%d, got %dx%d", Width, Height, width, height)
	}

	// Create an empty git repository.
	name := pathutil.FileName(imgPath)
	cmd := exec.Command("git", "init", name)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Forge a commit history based on the image.
	grid := new(Grid)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := img.At(x, y)
			if !isTrans(c) {
				// Create a commit with a date that represents the (x,y)-coordinate.
				date := coordDate(x, y)
				fmt.Println("date:", date)

				grid.Set(x, y)
				fmt.Println(grid)

				filePath := filepath.Join(name, "README")
				err = ioutil.WriteFile(filePath, grid.Bytes(), 0644)
				if err != nil {
					return err
				}

				cmd = exec.Command("git", "add", "README")
				cmd.Stderr = os.Stderr
				cmd.Dir = name
				err = cmd.Run()
				if err != nil {
					return err
				}

				message := fmt.Sprintf("readme: Add the cell at coordinate (%d, %d).", x, y)
				cmd = exec.Command("git", "commit", "-m", message, "--date", date.Format("2006-01-02 15:04:05"))
				cmd.Stderr = os.Stderr
				cmd.Dir = name
				err = cmd.Run()
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// isTrans returns true if the provided color is transparent.
func isTrans(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a == 0
}

func init() {
	// Each day of the last 52 weeks creates a rectangular grid of at least
	// 51x7 cells.

	// Locate the start of this grid which should be a Sunday almost or exactly
	// one year ago.
	start = time.Now().AddDate(-1, 0, 0)
	for {
		if start.Weekday() == time.Sunday {
			break
		}
		start = start.AddDate(0, 0, 1)
	}
}

// start represents a Sunday almost or exactly one year ago, which corresponds
// to the top left cell of the grid.
var start time.Time

// coordDate returns the date which corresponds to the provided
// (x,y)-coordinate.
func coordDate(x, y int) time.Time {
	days := x*Height + y
	return start.AddDate(0, 0, days)
}

// Grid represents a 2-dimensional grid of 51x7 cells.
type Grid [Width][Height]bool

// Set sets the provided coordinate to true.
func (grid *Grid) Set(x, y int) {
	grid[x][y] = true
}

func (grid *Grid) String() string {
	return string(grid.Bytes())
}

// Bytes returns a pretty-printed representation of grid as a slice of bytes.
func (grid *Grid) Bytes() []byte {
	buf := new(bytes.Buffer)
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if grid[x][y] {
				fmt.Fprint(buf, "#")
			} else {
				fmt.Fprint(buf, " ")
			}
		}
		fmt.Fprintln(buf)
	}
	return buf.Bytes()
}
