package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mewkiz/pkg/bufioutil"
)

func main() {
	// Get input from stdin.
	br := bufioutil.NewReader(os.Stdin)
	s, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	t, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	// Print all locations of the substring t in s with indicies starting at 1.
	locs := Subs(s, t)
	for i, loc := range locs {
		fmt.Print(loc + 1)
		if i != len(locs)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// Subs returns all locations where t occurs in s.
func Subs(s, t string) (locs []int) {
	for i := 0; i < len(s); {
		pos := strings.Index(s[i:], t)
		if pos == -1 {
			break
		}
		loc := i + pos
		locs = append(locs, loc)
		i = loc + 1
	}
	return locs
}
