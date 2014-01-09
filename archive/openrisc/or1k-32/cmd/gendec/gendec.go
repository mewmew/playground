// gendec generates the decoding logic for the Open RISC 1000 instruction sets.
// It is in no sense beautiful code, but gets the job done.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "gendec FILE")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := parseFile(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// parseFile parses the provided file, which has the following format:
//
//    000000NNNNNNNNNNNNNNNNNNNNNNNNNN l.j
//    000001NNNNNNNNNNNNNNNNNNNNNNNNNN l.jal
//    ...
func parseFile(filePath string) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Generate decoding logic as a switch statement.
	fmt.Println("switch {")
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return fmt.Errorf("invalid line %q, doesn't contain two parts", line)
		}
		// bits correspond to the bit representation of an instruction, for
		// instance: 000000NNNNNNNNNNNNNNNNNNNNNNNNNN
		bits := parts[0]
		if len(bits) != 32 {
			return fmt.Errorf("invalid bit string %q, doesnt' contain 32 bits", bits)
		}
		// mnemonic correspond to the mnemonic of an instruction, for instance:
		// l.j
		mnemonic := parts[1]
		parseInst(bits, mnemonic)
	}
	err = s.Err()
	if err != nil {
		return err
	}
	fmt.Println("}")

	return nil
}

// parseInst parses the bits of the provided instruction and generates decoding
// logic.
func parseInst(bits, mnemonic string) {
	if !strings.HasPrefix(mnemonic, "l.") {
		// Comment out cases that belong to other instruction sets than the Open
		// RISC Basic Instruction Set (ORBIS).
		fmt.Println("/*")
	}

	// In bits each '0' and '1' is part of the opcode mask and each '1' is part
	// of the actual opcode.
	var opMask, opCode uint
	// In bits each '-' is part of the padding mask.
	var padMask uint32
	for _, b := range bits {
		opMask <<= 1
		opCode <<= 1
		padMask <<= 1
		switch b {
		case '0':
			opMask |= 1
		case '1':
			opMask |= 1
			opCode |= 1
		case '-':
			padMask |= 1
		}
	}

	As := make([]*Offset, 0)
	Bs := make([]*Offset, 0)
	Ds := make([]*Offset, 0)
	Is := make([]*Offset, 0)
	Ks := make([]*Offset, 0)
	Ls := make([]*Offset, 0)
	Ns := make([]*Offset, 0)
	var A, B, D, I, K, L, N *Offset
	for i, b := range bits {
		switch b {
		case 'A':
			if A == nil {
				A = &Offset{start: 31 - i}
			}
		case 'B':
			if B == nil {
				B = &Offset{start: 31 - i}
			}
		case 'D':
			if D == nil {
				D = &Offset{start: 31 - i}
			}
		case 'I':
			if I == nil {
				I = &Offset{start: 31 - i}
			}
		case 'K':
			if K == nil {
				K = &Offset{start: 31 - i}
			}
		case 'L':
			if L == nil {
				L = &Offset{start: 31 - i}
			}
		case 'N':
			if N == nil {
				N = &Offset{start: 31 - i}
			}
		}
		if A != nil && (i == len(bits)-1 || b != 'A') {
			if i == len(bits)-1 {
				A.end = 0
			} else {
				A.end = 31 - (i - 1)
			}
			As = append(As, A)
			A = nil
		}
		if B != nil && (i == len(bits)-1 || b != 'B') {
			if i == len(bits)-1 {
				B.end = 0
			} else {
				B.end = 31 - (i - 1)
			}
			Bs = append(Bs, B)
			B = nil
		}
		if D != nil && (i == len(bits)-1 || b != 'D') {
			if i == len(bits)-1 {
				D.end = 0
			} else {
				D.end = 31 - (i - 1)
			}
			Ds = append(Ds, D)
			D = nil
		}
		if I != nil && (i == len(bits)-1 || b != 'I') {
			if i == len(bits)-1 {
				I.end = 0
			} else {
				I.end = 31 - (i - 1)
			}
			Is = append(Is, I)
			I = nil
		}
		if K != nil && (i == len(bits)-1 || b != 'K') {
			if i == len(bits)-1 {
				K.end = 0
			} else {
				K.end = 31 - (i - 1)
			}
			Ks = append(Ks, K)
			K = nil
		}
		if L != nil && (i == len(bits)-1 || b != 'L') {
			if i == len(bits)-1 {
				L.end = 0
			} else {
				L.end = 31 - (i - 1)
			}
			Ls = append(Ls, L)
			L = nil
		}
		if N != nil && (i == len(bits)-1 || b != 'N') {
			if i == len(bits)-1 {
				N.end = 0
			} else {
				N.end = 31 - (i - 1)
			}
			Ns = append(Ns, N)
			N = nil
		}
	}

	// Generate decoding logic as case statement.
	fmt.Println("//", mnemonic)
	fmt.Printf("case buf&0x%08X == 0x%08X:\n", opMask, opCode)
	fmt.Println("   //", bits)
	printPadding(padMask)
	printOperand("a", As)
	printOperand("b", Bs)
	printOperand("d", Ds)
	printOperand("i", Is)
	printOperand("k", Ks)
	printOperand("l", Ls)
	printOperand("n", Ns)
	if !strings.HasPrefix(mnemonic, "l.") {
		fmt.Println("*/")
	}
	fmt.Println()
}

// Offset contains the start and end offset in a bit pattern.
type Offset struct {
	start int
	end   int
}

// Mask returns the mask of off based on the it's start and end value.
func (off *Offset) Mask() (mask uint32) {
	for i := off.end; i <= off.start; i++ {
		mask <<= 1
		mask |= 1
	}
	mask <<= uint(off.end)
	return mask
}

const padFormat = `   if buf&0x%08X != 0 {
      return nil, errors.New("invalid padding")
   }
`

// printPadding prints padding check logic.
func printPadding(padMask uint32) {
	if padMask == 0 {
		return
	}
	fmt.Printf(padFormat, padMask)
}

// printOperand generates the decoding logic for the provided operand.
func printOperand(name string, xs []*Offset) {
	for i, x := range xs {
		op := ":="
		if i != 0 {
			op = "|="
		}
		var shift string
		if x.end != 0 {
			var sub int
			for _, y := range xs[i+1:] {
				sub += y.start - y.end
			}
			shift = fmt.Sprintf(" >> %d", x.end-sub)
		}
		fmt.Printf("   %s %s buf&0x%08X%s\n", name, op, x.Mask(), shift)
	}
}
