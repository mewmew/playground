// The xeh tool reverses the output of hexdump -C (*.txt > *.bin).
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func usage() {
	const use = `
Reverse the output of hexdump -C.

Usage:

	xeh [OPTION]... FILE

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// output specifies the output path.
		output string
	)
	flag.StringVar(&output, "o", "", "output path")
	flag.Usage = usage
	flag.Parse()

	// Reverse the output of hexdump -C.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	for _, path := range flag.Args() {
		if err := reverse(w, path); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// reverse reverses the output of hexdump -C, writing to w.
func reverse(w io.Writer, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	star := false
	var prev uint64
	var data []byte
	var zero [16]byte
	for s.Scan() {
		line := s.Text()
		if line == "*" {
			star = true
			continue
		}
		if len(line) < len("00000000") {
			return errors.Errorf("invalid line %q, missing address prefix", line)
		}
		part := line[:len("00000000")]
		line = line[len("00000000"):]
		addr, err := strconv.ParseUint(part, 16, 32)
		if err != nil {
			return errors.WithStack(err)
		}
		if star {
			for i := prev + 16; i < addr; i += 16 {
				data = append(data, zero[:]...)
			}
			star = false
		}
		prev = addr
		pos := strings.IndexByte(line, '|')
		if pos != -1 {
			line = line[:pos]
		}
		for len(line) > 0 {
			line = strings.TrimSpace(line)
			pos := strings.IndexByte(line, ' ')
			if pos == -1 {
				pos = len(line)
			}
			part := line[:pos]
			line = line[pos:]
			b, err := strconv.ParseUint(part, 16, 8)
			if err != nil {
				return errors.WithStack(err)
			}
			data = append(data, byte(b))
		}
	}
	if err := s.Err(); err != nil {
		return errors.WithStack(err)
	}
	if _, err := w.Write(data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
