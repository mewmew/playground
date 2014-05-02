package main

import (
	"fmt"
)

func ExampleRevComp() {
	dna := "AAAACCCGGT"
	fmt.Println(RevComp(dna))
	// Output: ACCGGGTTTT
}
