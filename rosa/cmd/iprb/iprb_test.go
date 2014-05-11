package main

import (
	"fmt"
)

func ExampleBaseCount() {
	// Calculate the probability that two randomly selected mating organisms will
	// produce an individual possessing a dominant allele.
	k, m, n := 2, 2, 2
	prob := DominantProb(k, m, n)
	fmt.Printf("%.5f\n", prob)
	// Output: 0.78333
}
