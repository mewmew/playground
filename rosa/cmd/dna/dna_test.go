package main

import (
	"fmt"
)

func ExampleDNA() {
	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	fmt.Println(BaseCount(dna))
	// Output: 20 12 17 21
}
