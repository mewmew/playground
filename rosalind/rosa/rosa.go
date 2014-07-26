// Package rosa implements utility functions to facilitate problem solving on
// the bioinformatics platform Rosalind.
package rosa

import (
	"bytes"
	"fmt"
	"strings"
)

// Trans transcribes the provided DNA sequence into an RNA sequence where the
// nucleotide uracil is used in place of thymine.
func Trans(dna string) (rna string) {
	return strings.Replace(dna, "T", "U", -1)
}

// RevComp returns the reverse complement of the provided DNA sequence. The
// bases are complemented as follows:
//    A: T
//    C: G
//    G: C
//    T: A
func RevComp(dna string) (revc string) {
	buf := new(bytes.Buffer)
	for i := len(dna) - 1; i >= 0; i-- {
		switch dna[i] {
		case 'A':
			fmt.Fprint(buf, "T")
		case 'C':
			fmt.Fprint(buf, "G")
		case 'G':
			fmt.Fprint(buf, "C")
		case 'T':
			fmt.Fprint(buf, "A")
		}
	}
	return buf.String()
}

const (
	// Stop indicates the stop of amino acid translation.
	Stop byte = 0
)

// aminos is a map from codons to amino acids.
var aminos = map[string]byte{
	"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
	"UGC": 'C', "UGU": 'C',
	"GAC": 'D', "GAU": 'D',
	"GAA": 'E', "GAG": 'E', "UUC": 'F',
	"UUU": 'F',
	"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
	"CAC": 'H', "CAU": 'H',
	"AUA": 'I', "AUC": 'I', "AUU": 'I',
	"AAA": 'K', "AAG": 'K',
	"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L', "UUA": 'L', "UUG": 'L',
	"AUG": 'M',
	"AAC": 'N', "AAU": 'N',
	"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
	"CAA": 'Q', "CAG": 'Q',
	"AGA": 'R', "AGG": 'R', "CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
	"AGC": 'S', "AGU": 'S', "UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
	"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
	"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',
	"UGG": 'W',
	"UAC": 'Y', "UAU": 'Y',
	"UAA": Stop, "UAG": Stop, "UGA": Stop,
}

// Prot translates the provided RNA sequence to a protein which consists of a
// sequence of amino acids.
func Prot(rna string) (prot string, err error) {
	if len(rna)%3 != 0 {
		return "", fmt.Errorf("rosa.Prot: invalid RNA length; not divisible by 3")
	}

	buf := make([]byte, len(rna)/3)
	for i := 0; i < len(buf); i++ {
		j := i * 3
		codon := rna[j : j+3]
		amino, ok := aminos[codon]
		if !ok {
			return "", fmt.Errorf("rosa.Prot: invalid codon %q", codon)
		}

		// Break when a stop codon has been located.
		if amino == Stop {
			buf = buf[:i]
			break
		}

		buf[i] = amino
	}

	return string(buf), nil
}
