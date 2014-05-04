package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mewkiz/pkg/bufioutil"
	"github.com/mewmew/playground/rosa"
)

func main() {
	// Get input from stdin.
	br := bufioutil.NewReader(os.Stdin)
	rna, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	// Translate the RNA sequence into a protein.
	prot, err := rosa.Prot(rna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prot)
}
