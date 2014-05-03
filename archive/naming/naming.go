// package naming specifies a few naming convensions by example.
package naming

import (
	"io"
)

type T struct{}

func New(r io.Reader) (t *T) {
	return t
}

func Open(path string) (t *T) {
	return t
}

func (t *T) ReadLine() (line string, err error) {
	return line, nil
}

func (t *T) ReadLines() (lines []string, err error) {
	return lines, nil
}

func ReadLines(r io.Reader) (lines []string, err error) {
	return lines, nil
}

func LoadLines(path string) (lines []string, err error) {
	return lines, nil
}
