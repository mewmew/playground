package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/mewmew/playground/rosalind/rosa"
)

func ExampleORFs() {
	dna := "AGCCATGTAGCTAACTCAGGTTACATGGGGATGACCCCGCGACTTGGATTAGAGTCTCTTTTGGAATAAGCCTGAATGATCCGAGTAGCATCTCAG"
	revc := rosa.RevComp(dna)

	// Locates each open reading frame (ORF) of the DNA-sequence and its reverse
	// complement. Use a map for the store unique proteins.
	orfs := ORFs(dna)
	orfs = append(orfs, ORFs(revc)...)
	uniq := make(map[string]bool)
	for _, orf := range orfs {
		rna := rosa.Trans(orf)
		prot, err := rosa.Prot(rna)
		if err != nil {
			log.Fatalln(err)
		}
		uniq[prot] = true
	}

	// Print sorted proteins.
	var prots []string
	for prot := range uniq {
		prots = append(prots, prot)
	}
	sort.Strings(prots)
	for _, prot := range prots {
		fmt.Println(prot)
	}
	// Output:
	// M
	// MGMTPRLGLESLLE
	// MLLGSFRLIPKETLIQVAGSSPCNLS
	// MTPRLGLESLLE
}
