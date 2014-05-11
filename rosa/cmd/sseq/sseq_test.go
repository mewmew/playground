package main

import (
	"fmt"
)

func ExampleBaseCount() {
	// Print the location of each character in sep as a subsequence of dna.
	dna := "ACGTACGTGACG"
	sep := "GTA"
	locs := SubSeq(dna, sep)
	for i, loc := range locs {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(loc + 1)
	}
	fmt.Println()
	// Output: 3 4 5
}
