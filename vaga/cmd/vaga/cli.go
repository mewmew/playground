package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/mewmew/playground/vaga/grid"
)

// cli plays a game of tic-tac-toe. The input is taken from the command line
// arguments.
func cli() (err error) {
	g := grid.NewGrid()
	for argNum, arg := range flag.Args() {
		// Parse marker position.
		pos, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		col := pos % g.Width()
		row := pos / g.Height()

		// Place marker.
		player := markA
		if argNum%2 != 0 {
			player = markB
		}
		err = g.Place(col, row, player)
		if err != nil {
			return err
		}

		// Print the grid.
		fmt.Println(g)

		// Check if there is a winner.
		markWin := g.Check()
		if markWin != grid.MarkNone {
			fmt.Printf("===> %q wins :)\n", markWin)
			return nil
		}
	}
	fmt.Println("no winner :(")
	return nil
}
