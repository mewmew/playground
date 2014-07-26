package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mewmew/playground/rosalind/rosa"
)

func main() {
	// Parse FASTA from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	label, err := fas.Label(0)
	if err != nil {
		log.Fatalln(err)
	}
	dna := fas.Seqs[label]
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
}

// startSeq specifies the DNA-sequence of the start codon.
const startSeq = "ATG"

// stopSeqs specifies the DNA-sequences of the stop codons.
var stopSeqs = []string{"TAG", "TGA", "TAA"}

// ORFs locates each open reading frame (ORF) of dna. An ORF starts from the
// start codon (ATG) and ends by a stop codon (TAG, TGA, TAA), without any other
// stop codons in between.
func ORFs(dna string) (orfs []string) {
	const codonSeqLen = 3
loop:
	for {
		pos := strings.Index(dna, startSeq)
		if pos == -1 {
			break
		}
		s := dna[pos:]
		dna = dna[pos+codonSeqLen:]
		for i := 2 * codonSeqLen; i <= len(s); i += codonSeqLen {
			orf := s[:i]
			for _, stopSeq := range stopSeqs {
				if strings.HasSuffix(orf, stopSeq) {
					orfs = append(orfs, orf)
					continue loop
				}
			}
		}
	}
	return orfs
}
