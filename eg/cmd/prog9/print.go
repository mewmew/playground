package main

import "container/list"
import dbg "fmt"
import "fmt"

// Print prints the left and the right side of each equation.
func Print(es *list.List) {
	for elem := es.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(Equation)
		e.Print()
	}
}

// Print prints the left and the right side of an equation.
func (e Equation) Print() {
	var i int
	var s string
	for shape, count := range e.Left {
		if count != 1 {
			s += fmt.Sprintf("%.f*", count)
		}
		s += fmt.Sprint(shape)
		if i < len(e.Left)-1 {
			s += " + "
		}
		i++
	}
	dbg.Printf("%-20s = ", s)
	s = ""
	for shape, count := range e.Right {
		if count != 1 {
			s += fmt.Sprintf("%.f*", count)
		}
		s += fmt.Sprint(shape)
		if i < len(e.Right)-1 {
			s += " + "
		}
		i++
	}
	dbg.Println(s)
}
