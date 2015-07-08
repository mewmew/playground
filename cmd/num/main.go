//go:generate usagen num

// num is a tool which displays the binary, octal, decimal and hexadecimal
// representation of numbers.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func usage() {
	const use = `
Usage: num BIN|OCT|DEC|HEX

Examples:
  num 0b111101101
  num 0o755
  num 493
  num 0x1ED
`
	fmt.Fprint(os.Stderr, use[1:])
}

func main() {
	// Parse command line arguments.
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	s := flag.Arg(0)

	n, err := parseNum(s)
	if err != nil {
		log.Fatal(err)
	}
	const format = `
bin: 0b%b
oct: 0o%o
dec: %d
hex: 0x%X
`
	fmt.Printf(format[1:], n, n, n, n)
}

// parseNum parses the given string as a binary, octal, decimal or hexadecimal
// number based on its prefix; with "0b", "0o", no prefix, and "0x"
// respectively.
func parseNum(s string) (n uint64, err error) {
	var base int
	switch {
	case strings.HasPrefix(s, "0b"):
		// Binary
		base = 2
		s = s[len("0b"):]
	case strings.HasPrefix(s, "0o"):
		// Octal
		base = 8
		s = s[len("0o"):]
	case strings.HasPrefix(s, "0x"):
		// Hexadecimal
		base = 16
		s = s[len("0x"):]
	default:
		// Decimal
		base = 10
	}
	return strconv.ParseUint(s, base, 64)
}
