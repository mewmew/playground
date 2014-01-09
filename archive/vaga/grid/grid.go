// Package grid implements a grid with markers for tic-tac-toe.
package grid

import (
	"fmt"
)

// Mark specifies a marker on the grid.
type Mark int

// The various markers.
const (
	MarkNone Mark = iota
	MarkO
	MarkX
)

func (mark Mark) String() string {
	m := map[Mark]string{
		MarkNone: "-",
		MarkO:    "o",
		MarkX:    "x",
	}
	return m[mark]
}

// Set sets the marker based on the provided flag value. Mark satisfies the
// flag.Value interface.
func (mark *Mark) Set(s string) (err error) {
	m := map[string]Mark{
		"-": MarkNone,
		"o": MarkO,
		"x": MarkX,
	}
	var ok bool
	x, ok := m[s]
	if !ok {
		return fmt.Errorf("Mark.Set: unable to parse mark %q", s)
	}
	*mark = x
	return nil
}

// Width and height of the grid.
const (
	Width  = 3
	Height = 3
)

// Grid contains a grid. Each cell can hold a marker.
type Grid [Width][Height]Mark

// NewGrid returns a new Grid.
func NewGrid() (g *Grid) {
	return new(Grid)
}

// Width returns the grid's width, i.e. the number of columns.
func (g *Grid) Width() int {
	return len(g)
}

// Height returns the grid's height, i.e. the number of rows.
func (g *Grid) Height() int {
	return len(g[0])
}

func (g *Grid) String() (s string) {
	for row := 0; row < g.Height(); row++ {
		for col := 0; col < g.Width(); col++ {
			s += g[col][row].String()
		}
		s += "\n"
	}
	return s
}

// Place places the marker at the specified grid cell.
func (g *Grid) Place(col, row int, mark Mark) (err error) {
	if col < 0 {
		return fmt.Errorf("Grid.Place: negative col (%d)", col)
	}
	if col >= g.Width() {
		return fmt.Errorf("Grid.Place: col (%d) above max (%d)", col, g.Width())
	}
	if row < 0 {
		return fmt.Errorf("Grid.Place: negative row (%d)", row)
	}
	if row >= g.Height() {
		return fmt.Errorf("Grid.Place: row (%d) above max (%d)", row, g.Height())
	}
	if g[col][row] != MarkNone {
		return fmt.Errorf("Grid.Place: grid cell (%d, %d) not empty", col, row)
	}
	g[col][row] = mark
	return nil
}

// Check checks if there is a winner and if so, returns it's marker. MarkNone is
// returned when no winner has been located.
func (g *Grid) Check() (mark Mark) {
	// Check vertical lines.
	for col := 0; col < g.Width(); col++ {
		mark = g.FullVert(col)
		if mark != MarkNone {
			return mark
		}
	}
	// Check horizontal lines.
	for row := 0; row < g.Height(); row++ {
		mark = g.FullHori(row)
		if mark != MarkNone {
			return mark
		}
	}
	// Check diagonal lines.
	mark = g.FullDiag()
	if mark != MarkNone {
		return mark
	}
	return MarkNone
}

// FullVert checks if any marker has a full vertical line and if so, returns
// that marker. MarkNone is returned when no winner has been located.
func (g *Grid) FullVert(col int) (mark Mark) {
	mark = g[col][0]
	for row := 1; row < g.Height(); row++ {
		if mark != g[col][row] {
			return MarkNone
		}
	}
	return mark
}

// FullHori checks if any marker has a full horizontal line and if so, returns
// that marker. MarkNone is returned when no winner has been located.
func (g *Grid) FullHori(row int) (mark Mark) {
	mark = g[0][row]
	for col := 1; col < g.Width(); col++ {
		if mark != g[col][row] {
			return MarkNone
		}
	}
	return mark
}

// FullDiag checks if any marker has a full diagonal line and if so, returns
// that marker. MarkNone is returned when no winner has been located.
func (g *Grid) FullDiag() (mark Mark) {
	mark = g.fullDiagTopLeft()
	if mark != MarkNone {
		return mark
	}
	return g.fullDiagBottomLeft()
}

// fullDiagTopLeft checks if any marker has a full diagonal line from top left
// to bottom right and if so, returns that marker. MarkNone is returned when no
// winner has been located.
func (g *Grid) fullDiagTopLeft() (mark Mark) {
	mark = g[0][0]
	row := 1
	for col := 1; col < g.Width(); col++ {
		if mark != g[col][row] {
			return MarkNone
		}
		row++
	}
	return mark
}

// fullDiagBottomLeft checks if any marker has a full diagonal line from bottom
// left to top right and if so, returns that marker. MarkNone is returned when
// no winner has been located.
func (g *Grid) fullDiagBottomLeft() (mark Mark) {
	mark = g[0][g.Height()-1]
	row := g.Height() - 2
	for col := 1; col < g.Width(); col++ {
		if mark != g[col][row] {
			return MarkNone
		}
		row--
	}
	return mark
}

// Clear clears the grid of all markers.
func (g *Grid) Clear() {
	for col := 0; col < g.Width(); col++ {
		for row := 0; row < g.Height(); row++ {
			g[col][row] = MarkNone
		}
	}
}
