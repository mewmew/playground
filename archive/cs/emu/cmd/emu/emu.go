// Command emu emulates a system capable of running the RISC dialect described
// in risc/op.
package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mewmew/playground/archive/cs/emu"
)

// flagHex is used for hexadecimal representation of instructions.
var flagHex string

func init() {
	flag.StringVar(&flagHex, "x", "", "Hexadecimal representation of instructions.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: emu [OPTION]... [FILE]")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "emu emulates a system capable of running the RISC dialect described in risc/op.")
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	var r io.Reader
	if len(flagHex) > 0 {
		buf, err := hex.DecodeString(flagHex)
		if err != nil {
			log.Fatalln(err)
		}
		r = bytes.NewReader(buf)
	} else if flag.NArg() == 1 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		r = f
	} else {
		flag.Usage()
		os.Exit(1)
	}
	err := emulate(r)
	if err != nil {
		log.Fatalln(err)
	}
}

func emulate(r io.Reader) (err error) {
	sys, err := emu.New(r)
	if err != nil {
		return err
	}
	fmt.Println(sys)
	sys.Start()
	for {
		err = sys.Step()
		if err != nil {
			if err == emu.ErrHalted {
				break
			}
			return err
		}
		fmt.Println(sys)
	}
	return nil
}
