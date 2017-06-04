package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		if err := reverse(path); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// reverse prints to stdout the reverse output of hexdump -C.
func reverse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	star := false
	var prev uint32
	var data []byte
	var zero [16]byte
	for s.Scan() {
		line := s.Text()
		var addr uint32
		var buf [16]byte
		if line == "*" {
			star = true
			continue
		}
		if _, err := fmt.Sscanf(line, "%08x", &addr); err != nil {
			return errors.WithStack(err)
		}
		line = line[len("00000000"):]
		n := 0
		for ; n < 16; n++ {
			line = strings.TrimSpace(line)
			if _, err := fmt.Sscanf(line, "%02x", &buf[n]); err != nil {
				break
			}
			if len(line) < len("00") {
				break
			}
			line = line[len("00"):]
		}
		if star {
			for i := prev; i < addr-16; i += 16 {
				data = append(data, zero[:]...)
			}
			star = false
		}
		prev = addr
		data = append(data, buf[:n]...)
	}
	fmt.Println(hex.Dump(data))
	if err := s.Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
