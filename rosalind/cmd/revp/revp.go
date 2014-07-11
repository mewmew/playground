package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewmew/playground/rosalind/rosa"
)

func main() {
	// Parse FASTA from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	// Locate the single DNA sequence present in the FASTA file.
	label, err := fas.Label(0)
	if err != nil {
		log.Fatalln(err)
	}
	dna := fas.Seqs[label]

	// Calculate the length and locating of each reverse palindrome within the
	// DNA sequence which has a length of between 4 and 12 nucleotides.
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

var (
	// baseComp is a map from each DNA base to its complement.
	baseComp = map[byte]byte{
		'A': 'T',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}
)

// IsRevPal returns true if the provided DNA sequence is a reverse palindrome,
// and false otherwise. A DNA sequence is a reverse palindrome if it is equal to
// its reverse complement.
func IsRevPal(dna string) bool {
	// The length of a reverse palindrome is always divisible by 2.
	if len(dna)%2 != 0 {
		return false
	}

	// This algorithm starts from both ends of the DNA sequence and successively
	// works towards the middle. It compares the start and end neucleotides to
	// verify that they are each others complement; and if not exits early.
	for i := 0; i < len(dna)/2; i++ {
		j := len(dna) - i - 1
		if baseComp[dna[i]] != dna[j] {
			return false
		}
	}
	return true
}
