package main

import (
	"fmt"
)

func ExampleSubs() {
	s := "GATATATGCATATACTT"
	t := "ATAT"

	// Print all locations with indicies starting at 1.
	locs := Subs(s, t)
	for i, loc := range locs {
		fmt.Print(loc + 1)
		if i != len(locs)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	// Output: 2 4 10
}
