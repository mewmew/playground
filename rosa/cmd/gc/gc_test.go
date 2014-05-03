package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
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
	f, err := ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	label, gc := MaxGC(f)
	fmt.Println(label)
	fmt.Printf("%.6f\n", gc)
	// Output:
	// Rosalind_0808
	// 60.919540
}

func ExampleParseFASTA() {
	// Parse FASTA.
	f, err := ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	// Sort keys.
	var labels []string
	for label := range f {
		labels = append(labels, label)
	}
	sort.Strings(labels)

	for _, label := range labels {
		dna := f[label]
		fmt.Println(label)
		fmt.Println(dna)
	}
	// Output:
	// Rosalind_0808
	// CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGACTGGGAACCTGCGGGCAGTAGGTGGAAT
	// Rosalind_5959
	// CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC
	// Rosalind_6404
	// CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCCTCCCACTAATAATTCTGAGG
}

func ExampleGC() {
	gc := GC("CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC")
	fmt.Printf("%.6f\n", gc)
	// Output: 53.571429
}
