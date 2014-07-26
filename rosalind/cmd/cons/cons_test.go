package main

import (
	"fmt"
	"log"
)

var seqs = []string{
	"ATCCAGCT",
	"GGGCAACT",
	"ATGGATCT",
	"AAGCAACC",
	"TTGGAACT",
	"ATGCCATT",
	"ATGGCACT",
}

func ExampleNewProfile() {
	profile, err := NewProfile(seqs, true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(profile)
	// Output:
	// A: 5 1 0 0 5 5 0 0
	// C: 0 0 1 4 2 0 6 1
	// G: 1 1 6 3 0 1 0 0
	// T: 1 5 0 0 0 1 1 6
}

func ExampleProfile_Cons() {
	profile, err := NewProfile(seqs, true)
	if err != nil {
		log.Fatalln(err)
	}
	cons := profile.Cons()
	fmt.Println(cons)
	// Output: ATGCAACT
}
