// ascii examines files and reports non-ascii characters.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		err := ascii(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func ascii(filePath string) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	lineNum := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		check(s.Text(), lineNum)
		lineNum++
	}
	err = s.Err()
	if err != nil {
		return err
	}

	return nil
}

func check(line string, lineNum int) {
	for col, r := range line {
		if r < 128 {
			if !unicode.IsSpace(rune(r)) && !unicode.IsPrint(rune(r)) {
				fmt.Printf("%d:%d - non-printable character 0x%02X.\n", lineNum, col, r)
			}
		} else {
			fmt.Printf("%d:%d - non-ascii character '%c'.\n", lineNum, col, r)
		}
	}
}
