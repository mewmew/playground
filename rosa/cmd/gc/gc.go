package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Parse FASTA from stdin.
	f, err := ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	label, gc := MaxGC(f)
	fmt.Println(label)
	fmt.Printf("%.6f\n", gc)
}

// MaxGC returns the label and GC-content of the DNA sequence with the highest
// GC-content in f.
func MaxGC(f fasta) (maxLabel string, maxGC float64) {
	for label, dna := range f {
		gc := GC(dna)
		if gc > maxGC {
			maxGC = gc
			maxLabel = label
		}
	}
	return maxLabel, maxGC
}

// fasta is a map from FASTA label names to DNA sequences.
type fasta map[string]string

// ParseFASTA reads data from r and parses it according to the FASTA file
// format.
func ParseFASTA(r io.Reader) (f fasta, err error) {
	s := bufio.NewScanner(r)
	var label string
	f = make(fasta)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, ">") {
			label = line[1:]
			continue
		}
		if len(label) == 0 {
			return nil, errors.New("parse: invalid label; zero length")
		}
		f[label] += strings.Replace(line, "\n", "", -1)
	}
	return f, s.Err()
}

// GC returns the percentage of 'G' and 'C' nucleotides in a given DNA sequence.
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
