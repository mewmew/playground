package main

import (
	"fmt"
	"log"
)

func ExampleHamDist() {
	// Calculate the Hamming distance between the two DNA sequences.
	dna1 := "GAGCCTACTAACGGGAT"
	dna2 := "CATCGTAATGACGGCCT"
	n, err := HamDist(dna1, dna2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
	// Output: 7
}
