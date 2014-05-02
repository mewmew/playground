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

	fmt.Println(RevComp(dna))
}

// TODO(u): The limited scope of revc should make it possible for future
// compiler optimizations to remove redundant string allocations in RevComp.
// Only the final string value will be accessed from other parts of the code.

// RevComp returns the reverse complement of the provided DNA string. The bases
// are complemented as follows:
//    A: T
//    C: G
//    G: C
//    T: A
func RevComp(dna string) (revc string) {
	for i := len(dna) - 1; i >= 0; i-- {
		switch dna[i] {
		case 'A':
			revc += "T"
		case 'C':
			revc += "G"
		case 'G':
			revc += "C"
		case 'T':
			revc += "A"
		}
	}
	return revc
}
