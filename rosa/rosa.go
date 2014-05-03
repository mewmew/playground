// Package rosa implements utility functions to facilitate problem solving on
// the bioinformatics platform Rosalind.
package rosa

import (
	"bufio"
	"errors"
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
