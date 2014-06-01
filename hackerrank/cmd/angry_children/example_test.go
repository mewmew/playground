package main

import (
	"fmt"
)

func ExampleMinUnfair() {
	// Calculate the minimum unfairness.
	k := 3
	pkts := []int{10, 100, 300, 200, 1000, 20, 30}
	fmt.Println(MinUnfair(pkts, k))
	// Output: 20
}
