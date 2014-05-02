package main

import (
	"fmt"
)

func ExampleTranscribe() {
	dna := "GATGGAACTTGACTACGTAAATT"
	fmt.Println(Transcribe(dna))
	// Output: GAUGGAACUUGACUACGUAAAUU
}
