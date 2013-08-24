package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("unixdict.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var cie, ie int
	var cei, ei int
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "cie") {
			cie++
		} else if strings.Contains(line, "ie") {
			ie++
		}
		if strings.Contains(line, "cei") {
			cei++
		} else if strings.Contains(line, "ei") {
			ei++
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatalln(err)
	}

	check(ie, cie, "I before E when not preceded by C")
	check(cei, ei, "I before E when not preceded by C")
}

// check checks if a statement is plausible. Something is plausible if a is more
// than two times b.
func check(a, b int, s string) {
	if a > b*2 {
		fmt.Printf("%q is plausible.\n", s)
	} else {
		fmt.Printf("%q is implausible.\n", s)
	}
}
