package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mewmew/playground/rosalind/rosa"
)

func main() {
	// Parse FASTA from stdin.
	fas, err := rosa.ParseFASTA(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var seqs []string
	for _, seq := range fas.Seqs {
		seqs = append(seqs, seq)
	}
	profile, err := NewProfile(seqs, true)
	if err != nil {
		log.Fatalln(err)
	}
	cons := profile.Cons()
	fmt.Println(cons)
}

// Profile records the number of times each base occurs in each position of a
// group of sequences.
//
// Some beautiful bit magic is used to distinguish the bases from each other. I
// first came in contact with this technique when reading Sonia's comment [1] at
// Rosalind. Lets take a look at the binary representation of each base before
// discussing the details.
//
//    'A' 01000001 (0x41)
//    'C' 01000011 (0x43)
//    'G' 01000111 (0x47)
//    'T' 01010100 (0x54)
//    'U' 01010101 (0x55)
//
// A closer look at the bit patterns reveals that the second and third bit can
// be used to uniquely distinguish the bases 'A', 'C', 'G' and 'T' from each
// other.
//
//    'A' .....00.
//    'C' .....01.
//    'G' .....11.
//    'T' .....10.
//    'U' .....10.
//
// Interestingly the second and third bit of 'T' and 'U' are the same, allowing
// the technique to be used for both DNA- and RNA-sequences!
//
// [1]: http://rosalind.info/problems/dna/solutions/#comment-222
type Profile struct {
	// p maps from position to base occurance in a group of sequences, where the
	// base is encoded using the bit magic described above.
	p [][4]int
	// The profile handles DNA-sequences if dna is set to true, and RNA-sequences
	// otherwise.
	dna bool
}

// NewProfile returns a profile which records the number of times each base
// occurs in each position of the provided sequences. seqs represent
// DNA-sequences if dna is set to true, and RNA-sequences otherwise.
func NewProfile(seqs []string, dna bool) (profile Profile, err error) {
	// Return an empty profile if no sequences have been provided.
	if len(seqs) == 0 {
		return Profile{}, errors.New("NewProfile: no sequences provided")
	}

	// Verify that each sequence have the same length.
	n := len(seqs[0])
	for i, seq := range seqs {
		if len(seq) != n {
			return Profile{}, fmt.Errorf("NewProfile: the length (%d) of sequence %d differs from the length (%d) of the first sequence", len(seq), i, n)
		}
	}

	profile = Profile{
		p:   make([][4]int, n),
		dna: dna,
	}
	for _, seq := range seqs {
		for i := 0; i < n; i++ {
			base := seq[i] >> 1 & 0x03
			profile.p[i][base]++
		}
	}

	return profile, nil
}

// Distinct bits of the various bases.
const (
	bitsA    = 0 // 00
	bitsC    = 1 // 01
	bitsG    = 3 // 11
	bitsTorU = 2 // 10
)

// Cons returns the consensus sequence (average sequence) based on the profile.
func (profile Profile) Cons() (cons string) {
	for _, pos := range profile.p {
		var max int
		var base byte

		// 'A' .....00.
		if v := pos[bitsA]; v > max {
			max = v
			base = 'A'
		}

		// 'C' .....01.
		if v := pos[bitsC]; v > max {
			max = v
			base = 'C'
		}

		// 'T' .....10.
		// 'U' .....10.
		if v := pos[bitsTorU]; v > max {
			max = v
			if profile.dna {
				base = 'T'
			} else {
				base = 'U'
			}
		}

		// 'G' .....11.
		if v := pos[bitsG]; v > max {
			max = v
			base = 'G'
		}

		cons += string(base)
	}
	return cons
}

func (profile Profile) String() string {
	buf := new(bytes.Buffer)

	// 'A'
	fmt.Fprint(buf, "A: ")
	for i, count := range profile.p {
		if i != 0 {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprint(buf, count[bitsA])
	}
	fmt.Fprintln(buf)

	// 'C'
	fmt.Fprint(buf, "C: ")
	for i, count := range profile.p {
		if i != 0 {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprint(buf, count[bitsC])
	}
	fmt.Fprintln(buf)

	// 'G'
	fmt.Fprint(buf, "G: ")
	for i, count := range profile.p {
		if i != 0 {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprint(buf, count[bitsG])
	}
	fmt.Fprintln(buf)

	// 'T' or 'U'
	if profile.dna {
		fmt.Fprint(buf, "T: ")
	} else {
		fmt.Fprint(buf, "U: ")
	}
	for i, count := range profile.p {
		if i != 0 {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprint(buf, count[bitsTorU])
	}

	return buf.String()
}
