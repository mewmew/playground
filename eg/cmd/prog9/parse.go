package main

import "container/list"

///import dbg "fmt"
import "errors"
import "fmt"
import "strings"

import "github.com/mewkiz/pkg/stringsutil"

// Slot represents a slot in the game_grid.
type Slot struct {
	// Shape represents the shape of the slot.
	Shape Shape
	// Count represents the total number of shapes which are represented by the
	// slot.
	Count float64
}

// Shape represents the shape of a slot in the game_grid.
type Shape int

// The different shapes.
const (
	ShapeNumber = Shape(iota)
	ShapeCircle
	ShapeHeart
	ShapeSquare
	ShapeStar
	ShapeTriangle
)

// Equation represents the left and right side of an equation.
type Equation struct {
	// Left is a map from shape to count and represents the left side of the
	// equation.
	Left map[Shape]float64
	// Right is a map from shape to count and represents the right side of the
	// equation.
	Right map[Shape]float64
}

// NewEquation returns a newly initialized Equation object.
func NewEquation() (e Equation) {
	e = Equation{
		Left:  make(map[Shape]float64),
		Right: make(map[Shape]float64),
	}
	return e
}

// ParseGrid returns a slice of equations after parsing the #game_grid in the
// HTML page.
//
// Below is an illustration of a game_grid:
//
//    [◯][□][□][△][#]   1227
//    [♡][#][♡][#][◯]    707
//    [#][◯][#][☆][#]    469
//    [△][#][♡][△][#]    873
//    [#][☆][#][◯][♡]    662
//
//              1
//     8  7  6  1  5
//     5  5  6  4  1
//     4  2  9  9  4
//
// The game_grid results in the following equations:
//
//    ◯ + 2*□ + △    = 321 + 2*283 + 340     = 1227
//    ◯ + 2*♡        = 321 + 2*193           = 707
//    ◯ + ☆          = 321 + 148             = 469
//    ♡ + 2*△        = 193 + 2*340           = 873
//    ◯ + ♡ + ☆      = 321 + 193 + 148       = 662
//    ◯ + ♡ + △      = 321 + 193 + 340       = 854
//    ◯ + □ + ☆      = 321 + 283 + 148       = 752
//    2*♡ + □        = 2*193 + 283           = 669
//    ◯ + ☆ + 2*△    = 321 + 148 + 2*340     = 1149
//    ◯ + ♡          = 321 + 193             = 514
//
// The value of each shape is represented below:
//
//    ◯ = 321
//    ♡ = 193
//    □ = 283
//    ☆ = 148
//    △ = 340
//    # = 0
func ParseGrid(page string) (es *list.List, err error) {
	gridStart := stringsutil.IndexAfter(page, `<table id="game_grid">`)
	if gridStart == -1 {
		return nil, errors.New("unable to locate start table tag of game_grid.")
	}
	gridLen := strings.Index(page[gridStart:], "</table>")
	if gridLen == -1 {
		return nil, errors.New("unable to locate end table tag of game_grid.")
	}
	gridEnd := gridStart + gridLen
	grid := page[gridStart:gridEnd]
	err = dumpGrid(grid)
	if err != nil {
		return nil, err
	}
	//
	//    [◯][□][□][△][#]   #
	//    [♡][#][♡][#][◯]   #
	//    [#][◯][#][☆][#]   #
	//    [△][#][♡][△][#]   #
	//    [#][☆][#][◯][♡]   #
	//
	//     #  #  #  #  #
	//
	var lines [10][6]Slot
	var slot Slot
	// populate the first five lines
	for lineNum := 0; lineNum < 5; lineNum++ {
		for slotNum := 0; slotNum < 6; slotNum++ {
			slot, grid, err = GetSlot(grid)
			if err != nil {
				return nil, err
			}
			lines[lineNum][slotNum] = slot
		}
	}
	// populate the other five lines, using the slots from the first five lines.
	for lineNum := 5; lineNum < 10; lineNum++ {
		for slotNum := 0; slotNum < 5; slotNum++ {
			lines[lineNum][slotNum] = lines[slotNum][lineNum-5]
		}
		slot, grid, err = GetSlot(grid)
		if err != nil {
			return nil, err
		}
		lines[lineNum][5] = slot
	}
	es = list.New()
	for _, line := range lines {
		///dbg.Println("--- [ line ] ---")
		e := NewEquation()
		for slotNum := 0; slotNum < 5; slotNum++ {
			slot := line[slotNum]
			///dbg.Println("slot:", slot)
			if slot.Count != 0 {
				e.Left[slot.Shape] += slot.Count /// ### panic?
				///dbg.Println("shape:", slot.Shape)
				///dbg.Println("count:", e.Left[slot.Shape])
				///dbg.Println()
			}
		}
		slot := line[5]
		e.Right[slot.Shape] += slot.Count
		es.PushBack(e)
	}
	return es, nil
}

// GetSlot parses grid and returns a slot and the new grid position.
func GetSlot(grid string) (slot Slot, newGrid string, err error) {
	tdStart := stringsutil.IndexAfter(grid, "<td>")
	if tdStart == -1 {
		return Slot{}, "", errors.New("unable to locate start td tag in game_grid.")
	}
	tdLen := strings.Index(grid[tdStart:], "</td>")
	if tdLen == -1 {
		return Slot{}, "", errors.New("unable to locate end td tag in game_grid.")
	}
	tdEnd := tdStart + tdLen
	rawSlot := grid[tdStart:tdEnd]
	grid = grid[tdEnd:]
	slot, err = ParseSlot(rawSlot)
	if err != nil {
		return Slot{}, "", err
	}
	return slot, grid, nil
}

// ParseSlot parses the input and returns a slot corresponding to it.
func ParseSlot(rawSlot string) (slot Slot, err error) {
	switch rawSlot {
	case `<img src="shapes/circle.png">`:
		return Slot{Shape: ShapeCircle, Count: 1}, nil
	case `<img src="shapes/heart.png">`:
		return Slot{Shape: ShapeHeart, Count: 1}, nil
	case `<img src="shapes/square.png">`:
		return Slot{Shape: ShapeSquare, Count: 1}, nil
	case `<img src="shapes/star.png">`:
		return Slot{Shape: ShapeStar, Count: 1}, nil
	case `<img src="shapes/triangle.png">`:
		return Slot{Shape: ShapeTriangle, Count: 1}, nil
	case "&nbsp;":
		// zero
		return Slot{Shape: ShapeNumber, Count: 0}, nil
	}
	var x int
	n, err := fmt.Sscanf(rawSlot, "<strong>%d</strong>", &x)
	if err != nil || n != 1 {
		return Slot{}, fmt.Errorf("unable to parse slot ('%s').", rawSlot)
	}
	return Slot{Shape: ShapeNumber, Count: float64(x)}, nil
}

func (shape Shape) String() string {
	switch shape {
	case ShapeCircle:
		return "◯"
	case ShapeHeart:
		return "♡"
	case ShapeSquare:
		return "□"
	case ShapeStar:
		return "☆"
	case ShapeTriangle:
		return "△"
	case ShapeNumber:
		return "#"
	}
	return "invalid shape"
}
