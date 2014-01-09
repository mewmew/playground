package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mewmew/playground/archive/vaga/grid"
)

// markA specifies the marker for player A (the player that starts).
var markA = grid.MarkX

func init() {
	flag.Var(&markA, "s", "Start marker ('o', 'x').")
	flag.Usage = usage
}

// rawGrid is a representation of the grid and its cell positions.
var rawGrid = `
+-+-+-+
|0|1|2|
+-+-+-+
|3|4|5|
+-+-+-+
|6|7|8|
+-+-+-+
`

func usage() {
	fmt.Println("Usage: vaga [OPTION] [POS]...")
	fmt.Println("vaga uses a CLI if POS is present and a GUI otherwise.")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Below is a representation of the grid and its cell positions.")
	fmt.Println(rawGrid)
}

func main() {
	flag.Parse()
	err := vaga()
	if err != nil {
		log.Fatalln(err)
	}
}

// markB specifies the marker for player B (the second player).
var markB grid.Mark

func vaga() (err error) {
	// Parse start marker.
	switch markA {
	case grid.MarkO:
		markB = grid.MarkX
	case grid.MarkX:
		markB = grid.MarkO
	default:
		return fmt.Errorf("vaga: invalid start marker %q.", markA)
	}

	if flag.NArg() == 0 {
		err := gui()
		if err != nil {
			return err
		}
	} else {
		err := cli()
		if err != nil {
			return err
		}
	}
	return nil
}
