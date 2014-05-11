package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mewmew/playground/rosa"
)

func main() {
	// Get input from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	// The first sequence in the FASTA file is the DNA sequence and the second
	// sequence is the subsequence sep.
	dnaLabel, err := fas.Label(0)
	if err != nil {
		log.Fatalln(err)
	}
	sepLabel, err := fas.Label(1)
	if err != nil {
		log.Fatalln(err)
	}
	dna := fas.Seqs[dnaLabel]
	sep := fas.Seqs[sepLabel]

	// Print the location of each character in sep as a subsequence of dna.
	locs := SubSeq(dna, sep)
	for i, loc := range locs {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(loc + 1)
	}
	fmt.Println()
}

// SubSeq returns the first location of each character in sep as a subsequence
// of s, or nil if no match was found.
func SubSeq(s, sep string) (locs []int) {
	var i int
	for j := 0; j < len(sep); j++ {
		base := sep[j]
		pos := strings.IndexByte(s[i:], base)
		if pos == -1 {
			return nil
		}
		loc := i + pos
		locs = append(locs, loc)
		i += pos + 1
	}
	return locs
}
