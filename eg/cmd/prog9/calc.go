package main

import "container/list"
import dbg "fmt"
import "os"

// Calc attempts to locate and solve the easiest equation.
func Calc(es *list.List) {
	for elem := es.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(Equation)
		if len(e.Left) == 1 {
			Propagate(es, elem)
			return
		}
	}
	os.Exit(1)
}

// Propagate will propagate the value of the shape in killElem, remove the
// equation from the list of equations.
func Propagate(es *list.List, killElem *list.Element) {
	e := es.Remove(killElem).(Equation)
	var killSlot Slot
	for shape, count := range e.Left {
		killSlot = Slot{Shape: shape, Count: count}
	}
	var killValue float64
	for _, value := range e.Right {
		killValue = value / killSlot.Count
	}
	dbg.Print("found:", killSlot.Shape)
	dbg.Print(" = ")
	dbg.Printf("%.f\n", killValue)
	Found(killSlot.Shape, int(killValue))
	for elem := es.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(Equation)
		for shape, count := range e.Left {
			if shape == killSlot.Shape {
				e.Right[ShapeNumber] -= count * killValue
				delete(e.Left, shape)
				// kill equation
				if len(e.Left) < 1 {
					es.Remove(elem)
				}
			}
		}
	}
}

var circleValue, heartValue, squareValue, starValue, triangleValue int

// Found stores the located shape's value.
func Found(shape Shape, value int) {
	switch shape {
	case ShapeCircle:
		circleValue = value
	case ShapeHeart:
		heartValue = value
	case ShapeSquare:
		squareValue = value
	case ShapeStar:
		starValue = value
	case ShapeTriangle:
		triangleValue = value
	}
}
