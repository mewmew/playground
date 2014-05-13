package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	// Parse input.
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		log.Fatalln(err)
	}
	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		log.Fatalln(err)
	}
	pkts := make([]int, count)
	for i := 0; i < count; i++ {
		_, err = fmt.Scan(&pkts[i])
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Calculate the minimum unfairness.
	min, err := MinUnfairness(pkts, k)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(min)
}

// MinUnfairness calculates the minimum unfairness of distributing a packet of
// candy to each of k children.
func MinUnfairness(pkts []int, k int) (min int, err error) {
	sort.Ints(pkts)
	if k > len(pkts) {
		return 0, fmt.Errorf("MinUnfairness: invalid k; number of children (%d) exceed the number of packages (%d)", k, len(pkts))
	}
	for i, a := range pkts {
		j := i + k - 1
		if j >= len(pkts) {
			break
		}
		b := pkts[j]
		delta := abs(a - b)
		if i == 0 || delta < min {
			min = delta
		}
	}
	return min, nil
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
