package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Get input from stdin.
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	dna := string(buf)

	// Count the nucleotide occurrences within the DNA sequence.
	fmt.Println(BaseCount(dna))
}

// BaseCount returns the respective number of times that the nucleotides 'A',
// 'C', 'G' and 'T' occurs in the provided DNA string.
func BaseCount(dna string) (a, c, g, t int) {
	for _, base := range dna {
		switch base {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		case 'T':
			t++
		}
	}
	return a, c, g, t
}
