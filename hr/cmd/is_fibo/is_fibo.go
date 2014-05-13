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
	ns := make([]int, count)
	for i := 0; i < count; i++ {
		_, err = fmt.Scan(&ns[i])
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Check which input numbers that are fibonacci numbers.
	fs := NewFib()
	for _, n := range ns {
		if fs.IsFib(n) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}
	}
}

// Fib handles a dynamically generated fibonacci sequence of numbers.
type Fib struct {
	// seq represent a fibonacci sequence of numbers.
	seq []int
}

// NewFib returns a new dynamically generating fibonacci sequence.
func NewFib() (fib *Fib) {
	fib = &Fib{
		seq: []int{0, 1},
	}
	return fib
}

// Gen generates and stores the next number of the fibonacci sequence.
func (fib *Fib) Gen() {
	i := len(fib.seq)
	n := fib.seq[i-1] + fib.seq[i-2]
	fib.seq = append(fib.seq, n)
}

// Last returns the last number of the generated fibonacci sequence.
func (fib *Fib) Last() int {
	return fib.seq[len(fib.seq)-1]
}

// IsFib returns true if n is a fibonacci number, and false otherwise.
func (fib *Fib) IsFib(n int) bool {
	// Generate new numbers of the fibonacci sequence if needed.
	for fib.Last() < n {
		fib.Gen()
	}

	// Locate n using binary search.
	pos := sort.SearchInts(fib.seq, n)
	if pos < len(fib.seq) && fib.seq[pos] == n {
		return true
	}

	return false
}
