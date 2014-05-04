package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mewmew/playground/rosa"
)

func main() {
	// Get input from stdin.
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	dna := string(buf)

	// Transcribe the DNA sequence into an RNA sequence.
	fmt.Println(rosa.Trans(dna))
}
