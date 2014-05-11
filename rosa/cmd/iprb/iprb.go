package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get input from stdin.
	var k, m, n int
	_, err := fmt.Fscanf(os.Stdin, "%d %d %d", &k, &m, &n)
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the probability that two randomly selected mating organisms will
	// produce an individual possessing a dominant allele.
	prob := DominantProb(k, m, n)
	fmt.Printf("%.5f\n", prob)
}

// DominantProb returns the probability that two randomly selected mating
// organisms will produce an individual possessing a dominant allele. Any two
// organisms can mate from a population of k+m+n organisms: k individuals are
// homozygous dominant of a factor, m are heterozygous, and n are homozygous
// recessive.
func DominantProb(k, m, n int) (prob float64) {
	K, M, N := float64(k), float64(m), float64(n)
	total := K + M + N

	// The probability of producing an individual possessing a dominant allele is
	// calculated by subtracting the chance of a homozygous recessive offspring
	// from an initial probability of 100%.
	prob = 1

	// Dr+Dr: 25% recessive.
	prob -= 0.25 * M / total * (M - 1) / (total - 1)

	// Dr+rr: 50% recessive.
	prob -= 0.50 * M / total * N / (total - 1)

	// rr+Dr: 50% recessive.
	prob -= 0.50 * N / total * M / (total - 1)

	// rr+rr: 100% recessive.
	prob -= 1.00 * N / total * (N - 1) / (total - 1)

	return prob
}
