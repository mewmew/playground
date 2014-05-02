package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Get input from stdin.
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	dna := string(buf)

	fmt.Println(Transcribe(dna))
}

// Transcribe returns the transcribed RNA string based on the provided DNA
// string, where each 'T' nucleotide has been replaced by an 'U' nucleotide.
func Transcribe(dna string) (rna string) {
	return strings.Replace(dna, "T", "U", -1)
}
