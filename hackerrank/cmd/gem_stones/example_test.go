package main

import (
	"fmt"
)

func ExampleGemElems() {
	// Calculate the number of gem-elements.
	stones := []string{"abcdde", "baccd", "eeabg"}
	fmt.Println(GemElems(stones))
	// Output: 2
}
