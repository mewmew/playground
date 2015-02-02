// TODO(u): make disasm more fault tolerant. display "<invalid instruction>"
// instead of terminating execution.

// TODO(u): handle input from stdin.

// Command disasm disassembles instructions encoded in the RISC dialect
// described in risc/op.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mewmew/playground/archive/cs/risc/op"
)

// flagHex is used for hexadecimal representation of instructions.
var flagHex string

func init() {
	flag.StringVar(&flagHex, "x", "", "Hexadecimal representation of instructions.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: disasm [OPTION]... [FILE]")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "disasm disassembles instructions encoded in the RISC dialect described in risc/op.")
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	insts, err := disasm()
	if err != nil {
		log.Fatalln(err)
	}
	for _, inst := range insts {
		fmt.Println(inst)
	}
}

func disasm() (insts []interface{}, err error) {
	if len(flagHex) > 0 {
		return parseHex(flagHex)
	} else if flag.NArg() == 1 {
		return parseFile(flag.Arg(0))
	}
	flag.Usage()
	os.Exit(1)
	panic("unreachable")
}

// parseHex parses the provided hexadecimal string into a slice of instructions.
func parseHex(s string) (insts []interface{}, err error) {
	if len(s)%4 != 0 {
		return nil, fmt.Errorf("parseHex: string len (%d) not evenly dividable by 4", len(s))
	}

	insts = make([]interface{}, 0, len(s)/4)
	for i := 0; i < len(s); i += 4 {
		buf, err := strconv.ParseUint(s[i:i+4], 16, 16)
		if err != nil {
			return nil, err
		}
		inst, err := op.Decode(uint16(buf))
		if err != nil {
			return nil, err
		}
		insts = append(insts, inst)
	}

	return insts, nil
}

// parseFile parses the provided file into a slice of instructions.
func parseFile(filePath string) (insts []interface{}, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf uint16
	insts = make([]interface{}, 0)
	for {
		err = binary.Read(f, binary.BigEndian, &buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		inst, err := op.Decode(buf)
		if err != nil {
			return nil, err
		}
		insts = append(insts, inst)
	}

	return insts, nil
}
