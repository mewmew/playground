package main

import (
	"fmt"
)

func ExampleBaseCount() {
	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	fmt.Println(BaseCount(dna))
	// Output: 20 12 17 21
}
