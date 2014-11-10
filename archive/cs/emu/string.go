package emu

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/mewmew/playground/archive/cs/risc/op"
)

// String returns pretty-printed memory and register information about the
// system.
func (sys *System) String() string {
	// TODO(u): Clean up this mess :)
	b := new(bytes.Buffer)
	fmt.Fprintln(b, "=== [ system information ] ===")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "--- [ registers ] ---")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "PC  = %02X (%d)\n", sys.PC, sys.PC)
	for i := range sys.Regs {
		reg := sys.Regs[i]
		fmt.Fprintf(b, "r%-2d = %02X (%d)\n", i, reg, reg)
	}
	fmt.Fprintln(b)
	fmt.Fprintln(b, "--- [ memory ] ---")
	fmt.Fprintln(b)
	fmt.Fprintln(b, hex.Dump(sys.Mem[:]))
	fmt.Fprintln(b)
	fmt.Fprintln(b, "--- [ assembly ] ---")
	fmt.Fprintln(b)

	var prevNop, dots bool
	for i := 0; i < len(sys.Mem); i += op.InstSize {
		buf := binary.BigEndian.Uint16(sys.Mem[i:])
		inst, err := op.Decode(buf)
		if err != nil {
			fmt.Fprintf(b, "0x%02X: <invalid assembly> 0x%04X\n", i, buf)
			prevNop = false
			dots = false
			continue
		}
		_, ok := inst.(*op.Nop)
		if ok {
			if prevNop {
				if !dots {
					fmt.Fprintf(b, "0x%02X: %s\n", i, "...")
					dots = true
				}
			} else {
				fmt.Fprintf(b, "0x%02X: %s\n", i, inst)
				prevNop = true
			}
		} else {
			fmt.Fprintf(b, "0x%02X: %s\n", i, inst)
			prevNop = false
			dots = false
		}
	}

	return string(b.Bytes())
}
