package rosa

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func ExampleTrans() {
	dna := "GATGGAACTTGACTACGTAAATT"
	fmt.Println(Trans(dna))
	// Output: GAUGGAACUUGACUACGUAAAUU
}

func ExampleRevComp() {
	dna := "AAAACCCGGT"
	fmt.Println(RevComp(dna))
	// Output: ACCGGGTTTT
}

const s = `>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT`

func ExampleParseFASTA() {
	// Parse FASTA.
	fas, err := ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	// Sort keys.
	var labels []string
	for label := range fas.Seqs {
		labels = append(labels, label)
	}
	sort.Strings(labels)

	for _, label := range labels {
		dna := fas.Seqs[label]
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

func ExampleFASTA_Label() {
	// Parse FASTA.
	fas, err := ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	// Locate the second label.
	label, err := fas.Label(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(label)
	// Output:
	// Rosalind_5959
}

func ExampleFASTA_Index() {
	// Parse FASTA.
	fas, err := ParseFASTA(strings.NewReader(s))
	if err != nil {
		log.Fatalln(err)
	}

	// Locate the index of the label named "Rosalind_0808".
	i, err := fas.Index("Rosalind_0808")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
	// Output:
	// 2
}

func ExampleProt() {
	rna := "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"
	prot, err := Prot(rna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prot)
	// Output: MAMAPRTEINSTRING
}
