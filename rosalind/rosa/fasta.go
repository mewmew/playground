package rosa

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// FASTA handles labeled DNA sequences.
type FASTA struct {
	// Seqs is a map from FASTA label names to DNA sequences.
	Seqs map[string]string
	// labels maps from index to label; it keeps track of the order in which
	// labels occur in the FASTA file.
	labels []string
	// indices maps from label to index.
	indices map[string]int
}

// ParseFASTA reads data from r and parses it according to the FASTA file
// format.
func ParseFASTA(r io.Reader) (fas *FASTA, err error) {
	fas = &FASTA{
		Seqs:    make(map[string]string),
		indices: make(map[string]int),
	}
	s := bufio.NewScanner(r)
	var label string
	index := 0
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, ">") {
			label = line[1:]
			fas.labels = append(fas.labels, label)
			fas.indices[label] = index
			index++
			continue
		}
		if len(label) == 0 {
			return nil, errors.New("rosa.ParseFASTA: invalid label; zero length")
		}
		fas.Seqs[label] += strings.Replace(line, "\n", "", -1)
	}
	return fas, s.Err()
}

// Label returns the nth label of the FASTA file.
func (fas *FASTA) Label(n int) (label string, err error) {
	if n >= len(fas.labels) {
		return "", fmt.Errorf("rosa.FASTA.Label: invalid index %d; out of bounds for %d-element slice", n, len(fas.labels))
	}
	return fas.labels[n], nil
}

// Index returns the index of the provided label.
func (fas *FASTA) Index(label string) (n int, err error) {
	n, ok := fas.indices[label]
	if !ok {
		return 0, fmt.Errorf("rosa.FASTA.Index: unable to locate index of label %q", label)
	}
	return n, nil
}
