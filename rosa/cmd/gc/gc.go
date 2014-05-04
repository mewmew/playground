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

	// Locate the DNA sequence with the highest GC-content.
	label, gc := MaxGC(fas)
	fmt.Println(label)
	fmt.Printf("%.6f\n", gc)
}

// MaxGC returns the label and GC-content of the DNA sequence with the highest
// GC-content in fas.
func MaxGC(fas rosa.FASTA) (maxLabel string, maxGC float64) {
	for label, dna := range fas {
		gc := GC(dna)
		if gc > maxGC {
			maxGC = gc
			maxLabel = label
		}
	}
	return maxLabel, maxGC
}

// GC returns the percentage of the provided DNA sequence's bases that are
// either guanine or cytosine.
func GC(dna string) (gc float64) {
	for _, base := range dna {
		switch base {
		case 'C', 'G':
			gc++
		}
	}
	gc = 100 * gc / float64(len(dna))
	return gc
}
