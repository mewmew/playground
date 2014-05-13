package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	// Parse input.
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		log.Fatalln(err)
	}
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&cs[i])
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Calculate the minimum unfairness.
	min, err := MinUnfair(cs, k)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(min)
}

// MinUnfair calculates the minimum unfairness of distributing a packet of candy
// to each of k children.
func MinUnfair(cs []int, k int) (min int, err error) {
	if k < 1 {
		return 0, fmt.Errorf("MinUnfair: invalid k; number of children (%d) below 1", k)
	}
	if k > len(cs) {
		return 0, fmt.Errorf("MinUnfair: invalid k; number of children (%d) exceed the number of packages (%d)", k, len(cs))
	}

	sort.Ints(cs)
	for i := 0; i <= len(cs)-k; i++ {
		delta := cs[i+k-1] - cs[i]
		if i == 0 || delta < min {
			min = delta
		}
	}

	return min, nil
}
