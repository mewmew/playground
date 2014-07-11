package main

import (
	"fmt"
	"log"

	"github.com/mewmew/playground/rosalind/rosa"
)

func ExampleSplc() {
	// Splice the DNA, transcribe it into RNA and translate it into a protein.
	dna := "ATGGTCTACATAGCTGACAAACAGCACGTAGCAATCGGTCGAATCTCGAGAGGCATATGGTCACATGATCGGTCGAGCGTGTTTCAAAGTTTGCGCCTAG"
	introns := []string{"ATCGGTCGAA", "ATCGGTCGAGCGTGT"}
	rna := rosa.Trans(Splice(dna, introns))
	prot, err := rosa.Prot(rna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prot)
	// Output: MVYIADKQHVASREAYGHMFKVCA
}

func ExampleSplice() {
	dna := "ATGGTCTACATAGCTGACAAACAGCACGTAGCAATCGGTCGAATCTCGAGAGGCATATGGTCACATGATCGGTCGAGCGTGTTTCAAAGTTTGCGCCTAG"
	introns := []string{"ATCGGTCGAA", "ATCGGTCGAGCGTGT"}
	fmt.Println(Splice(dna, introns))
	// Output: ATGGTCTACATAGCTGACAAACAGCACGTAGCATCTCGAGAGGCATATGGTCACATGTTCAAAGTTTGCGCCTAG
}
