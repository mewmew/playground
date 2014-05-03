package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewmew/playground/rosa"
)

func main() {
	// Parse FASTA from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var dna string
	for label := range fas {
		dna = fas[label]
		break
	}

	locs, ns := RevPal(dna)
	for i := range locs {
		loc := locs[i]
		n := ns[i]
		fmt.Println(loc+1, n)
	}
}

// RevPal returns the location of length of every reverse palindrome in the
// provided DNA sequence having a length between 4 and 12 nucleotides.
func RevPal(dna string) (locs, ns []int) {
	for loc := 0; loc < len(dna); loc++ {
		// The length of a reverse palindrome is always divisible by 2.
		for n := 4; n <= 12; n += 2 {
			end := loc + n
			if end > len(dna) {
				break
			}
			if IsRevPal(dna[loc:end]) {
				locs = append(locs, loc)
				ns = append(ns, n)
			}
		}
	}
	return locs, ns
}

// IsRevPal returns true if the provided DNA sequence is a reverse palindrome,
// and false otherwise. A DNA sequence is a reverse palindrome if it is equal to
// its reverse complement.
func IsRevPal(dna string) bool {
	return rosa.RevComp(dna) == dna
}
