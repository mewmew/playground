// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ref: http://golang.org/doc/play/sieve.go

package main

import "fmt"

func main() {
	primes := make(chan int)
	go sieve(primes)
	for i := 0; i < 1000; i++ {
		fmt.Println(<-primes)
	}
}

// sieve sends prime numbers on primes by creating a chain of channels, each of
// which filters previously recorded prime numbers.
func sieve(primes chan int) {
	c := make(chan int)
	go counter(c)
	for {
		prime := <-c
		primes <- prime
		newc := make(chan int)
		go filter(c, newc, prime)
		c = newc
	}
}

// counters sends incrementing integers on c, starting with the first prime (2).
func counter(c chan int) {
	i := 2
	for {
		c <- i
		i++
	}
}

// filter filters numbers from recv to send if they are evenly dividable by
// prime.
func filter(recv <-chan int, send chan<- int, prime int) {
	for {
		x := <-recv
		if x%prime != 0 {
			send <- x
		}
	}
}
