package main

import (
	"fmt"
)

func ExampleIsFib() {
	// Check which numbers that are fibonacci numbers.
	ns := []int{5, 7, 8}
	fib := NewFib()
	for _, n := range ns {
		if fib.IsFib(n) {
			fmt.Println("IsFib")
		} else {
			fmt.Println("IsNotFib")
		}
	}
	// Output: IsFib
	// IsNotFib
	// IsFib
}
