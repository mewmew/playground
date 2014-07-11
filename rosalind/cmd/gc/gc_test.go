package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mewmew/playground/rosalind/rosa"
)

const s = `>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT`

func ExampleMaxGC() {
	// Parse FASTA.
	fas, err := rosa.ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	label, gc := MaxGC(fas)
	fmt.Println(label)
	fmt.Printf("%.6f\n", gc)
	// Output:
	// Rosalind_0808
	// 60.919540
}

func ExampleGC() {
	gc := GC("CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC")
	fmt.Printf("%.6f\n", gc)
	// Output: 53.571429
}
