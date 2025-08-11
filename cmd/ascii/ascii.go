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

var (
	flagOnlyPlaintext bool
	flagQuiet         bool
)

func init() {
	flag.BoolVar(&flagOnlyPlaintext, "only-plaintext", false, "Only check file extensions with known plaintext.")
	flag.BoolVar(&flagQuiet, "q", false, "Suppress non-error log messages.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: ascii [OPTION]... PATH...")
	fmt.Fprintln(os.Stderr, "Report non-ascii characters in files.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Invoke ascii with one or more filenames or directories.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
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
	".c":    true,
	".css":  true,
	".go":   true,
	".html": true,
	".js":   true,
	".md":   true,
	".sml":  true,
	".txt":  true,
	".typ":  true,
}

func checkFile(filePath string) (err error) {
	ext := path.Ext(filePath)
	if flagOnlyPlaintext {
		_, ok := whitelist[ext]
		if !ok {
			if !flagQuiet {
				log.Printf("ignoring file %q with extension %q.\n", filePath, ext)
			}
			return nil
		}
	}
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	lineNum := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		err := checkLine(s.Text(), filePath, lineNum)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		lineNum++
	}
	err = s.Err()
	if err != nil {
		return err
	}

	return nil
}

func checkLine(line, filePath string, lineNum int) (err error) {
	for col, r := range line {
		if r < 128 {
			if !unicode.IsSpace(rune(r)) && !unicode.IsPrint(rune(r)) {
				return fmt.Errorf("%s (%d:%d) - non-printable character 0x%02X", filePath, lineNum, col, r)
			}
		} else {
			return fmt.Errorf("%s (%d:%d) - non-ascii character '%c'", filePath, lineNum, col, r)
		}
	}
	return nil
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
