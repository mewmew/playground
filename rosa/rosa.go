// Package rosa implements utility functions to facilitate problem solving on
// the bioinformatics platform Rosalind.
package rosa

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// TODO(u): The limited scope of revc should make it possible for future
// compiler optimizations to remove redundant string allocations in RevComp.
// Only the final string value will be accessed from other parts of the code.

// RevComp returns the reverse complement of the provided DNA sequence. The
// bases are complemented as follows:
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

// FASTA is a map from FASTA label names to DNA sequences.
type FASTA map[string]string

// ParseFASTA reads data from r and parses it according to the FASTA file
// format.
func ParseFASTA(r io.Reader) (fas FASTA, err error) {
	s := bufio.NewScanner(r)
	var label string
	fas = make(FASTA)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, ">") {
			label = line[1:]
			continue
		}
		if len(label) == 0 {
			return nil, errors.New("parse: invalid label; zero length")
		}
		fas[label] += strings.Replace(line, "\n", "", -1)
	}
	return fas, s.Err()
}

const (
	// Stop indicates the stop of amino acid translation.
	Stop = 0
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
		return "", fmt.Errorf("Prot: invalid RNA length; not divisible by 3")
	}

	buf := make([]byte, len(rna)/3)
	for i := 0; i < len(buf); i++ {
		j := i * 3
		codon := rna[j : j+3]
		amino, ok := aminos[codon]
		if !ok {
			return "", fmt.Errorf("Prot: invalid codon %q", codon)
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
