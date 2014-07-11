package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mewmew/playground/rosalind/rosa"
)

func main() {
	// Get input from stdin.
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	dna := string(buf)

	// Obtain the reverse complement of the provided DNA sequence.
	fmt.Println(rosa.RevComp(dna))
}
