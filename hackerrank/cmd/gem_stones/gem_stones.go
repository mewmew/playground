package main

import (
	"fmt"
	"log"
)

func main() {
	// Parse input.
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}
	stones := make([]string, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&stones[i])
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Calculate the number of gem-elements.
	fmt.Println(GemElems(stones))
}

// GemElems returns the number of elements that occur in each stone, i.e. the
// number of gem-elements.
func GemElems(stones []string) int {
	gems := Elems(stones[0])
	for _, stone := range stones[1:] {
		elems := Elems(stone)
		// If an element of the first stone didn't occur in the current stone it
		// cannot be a gem-element.
		for elem := range gems {
			if !elems[elem] {
				delete(gems, elem)
			}
		}
	}
	return len(gems)
}

// Elems returns a map of elements contained within the stone.
func Elems(stone string) (elems map[byte]bool) {
	elems = make(map[byte]bool)
	for i := 0; i < len(stone); i++ {
		elems[stone[i]] = true
	}
	return elems
}
