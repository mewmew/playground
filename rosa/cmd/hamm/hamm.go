package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewkiz/pkg/bufioutil"
)

func main() {
	// Get input from stdin.
	br := bufioutil.NewReader(os.Stdin)
	a, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	b, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the Hamming distance between the two DNA sequences.
	n, err := HamDist(a, b)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

// HamDist calculates the Hamming distance between a and b.
func HamDist(a, b string) (n int, err error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("HamDist: length mismatch; len(a)=%d, len(b)=%d", len(a), len(b))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
		}
	}
	return n, nil
}
