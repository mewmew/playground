package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mewmew/playground/rosa"
)

func main() {
	// Parse FASTA from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	// The longest sequence in the FASTA file is the DNA sequence and all other
	// sequences are introns.
	var dnaLabel string
	var max int
	for label, s := range fas {
		if len(s) > max {
			max = len(s)
			dnaLabel = label
		}
	}
	var dna string
	introns := make([]string, len(fas)-1)
	for label, s := range fas {
		if label == dnaLabel {
			dna = s
			continue
		}
		introns = append(introns, s)
	}

	// Splice the DNA, transcribe it into RNA and translate it into a protein.
	rna := rosa.Trans(Splice(dna, introns))
	prot, err := rosa.Prot(rna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prot)
}

// Splice splices the provided DNA sequence by removing any occurrence of the
// provided introns.
func Splice(dna string, introns []string) string {
	oldnew := make([]string, 2*len(introns))
	for i := range introns {
		oldnew[i*2] = introns[i]
	}
	r := strings.NewReplacer(oldnew...)
	return r.Replace(dna)
}
