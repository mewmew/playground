package main

import (
	"fmt"
	"log"
)

func ExampleProt() {
	rna := "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"
	prot, err := Prot(rna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prot)
	// Output: MAMAPRTEINSTRING
}
