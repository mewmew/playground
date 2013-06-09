// ascii examines files and reports non-ascii characters.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"unicode"
)

var flagVerbose bool

func init() {
	flag.BoolVar(&flagVerbose, "v", false, "Verbose.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: ascii PATH...")
	fmt.Fprintln(os.Stderr, "Report non-ascii characters in files.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Invoke ascii with one or more filenames or directories.")
}

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		if isDir(filePath) {
			err := checkDir(filePath)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err := checkFile(filePath)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// whilelist contains a list of all extensions believed to be plain text.
var whitelist = map[string]bool{
	".asm":  true,
	".css":  true,
	".c":    true,
	".go":   true,
	".html": true,
	".js":   true,
	".md":   true,
	".txt":  true,
}

func checkFile(filePath string) (err error) {
	ext := path.Ext(filePath)
	_, ok := whitelist[ext]
	if !ok {
		if flagVerbose {
			log.Printf("ignoring file %q with extension %q.\n", filePath, ext)
		}
		return nil
	}
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	lineNum := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		checkLine(s.Text(), lineNum)
		lineNum++
	}
	err = s.Err()
	if err != nil {
		return err
	}

	return nil
}

func checkLine(line string, lineNum int) {
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

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func checkDir(dir string) (err error) {
	err = filepath.Walk(dir, walk)
	if err != nil {
		return err
	}
	return nil
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.Mode().IsRegular() {
		err = checkFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}
