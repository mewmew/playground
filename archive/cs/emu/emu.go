// Package emu implements an emulator for the RISC dialect described in risc/op.
package emu

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/mewmew/playground/archive/cs/risc/op"
)

// Information about the emulator system.
const (
	// MemSize specifies the size of the memory in bytes.
	MemSize = 256
)

// PC is a system's program counter. It holds the address of the next
// instruction to be executed.
type PC uint8

// Inc increases the program counter, taking special precaution to limit integer
// overflows and underflows. To decrease the program counter provide a negative
// n value.
func (pc *PC) Inc(n int) (err error) {
	oldPC := *pc
	newPC := oldPC + PC(n)
	switch {
	case n < 0:
		// A negative n value should never make PC larger.
		if newPC > oldPC {
			return fmt.Errorf("PC.Inc: integer underflow occurred; old PC (%d), new PC (%d), n (%d)", oldPC, newPC, n)
		}
	case n > 0:
		// A positive n value should never make PC smaller.
		if newPC < oldPC {
			return fmt.Errorf("PC.Inc: integer overflow occurred; old PC (%d), new PC (%d), n (%d)", oldPC, newPC, n)
		}
	}
	*pc = newPC
	return nil
}

// A System capable of running the RISC dialect described in risc/op.
type System struct {
	// Program counter.
	PC PC
	// Registers r0 through r15.
	Regs [op.RegCount]uint8
	// Memory.
	Mem [MemSize]uint8
	// When running is true the system is executing instructions.
	running bool
}

// New allocates and returns a new system, initiating the memory with the
// contents read from r. The remaining memory, the program counter and all
// registers are set to 0.
//
// Remember to call sys.Start before executing instructions.
func New(r io.Reader) (sys *System, err error) {
	sys = new(System)
	_, err = r.Read(sys.Mem[:])
	if err != nil {
		return nil, err
	}
	return sys, nil
}

// FetchInst fetches the next instruction and increments the program counter.
func (sys *System) FetchInst() (buf uint16, err error) {
	if int(sys.PC)+op.InstSize > len(sys.Mem) {
		return 0, fmt.Errorf("System.FetchInst: instruction at PC (%d) is outside of Mem", sys.PC)
	}
	buf = binary.BigEndian.Uint16(sys.Mem[sys.PC:])
	err = sys.PC.Inc(op.InstSize)
	if err != nil {
		return 0, err
	}
	return buf, nil
}

// ErrHalted is returned when trying to execute an instruction while the system
// is halted.
var ErrHalted = errors.New("emu: system is halted")

// Step decodes and executes one instruction.
func (sys *System) Step() (err error) {
	if !sys.running {
		return ErrHalted
	}
	buf, err := sys.FetchInst()
	if err != nil {
		return err
	}
	inst, err := op.Decode(buf)
	if err != nil {
		return err
	}
	return sys.Exec(inst)
}

// Exec executes the provided instruction.
func (sys *System) Exec(inst interface{}) (err error) {
	if !sys.running {
		return ErrHalted
	}
	switch v := inst.(type) {
	case *op.Nop:
		sys.Nop()
	case *op.LoadMem:
		err = sys.LoadMem(v.Dst, v.Src)
		if err != nil {
			return err
		}
	case *op.LoadVal:
		err = sys.LoadVal(v.Dst, v.Src)
		if err != nil {
			return err
		}
	case *op.Store:
		err = sys.Store(v.Dst, v.Src)
		if err != nil {
			return err
		}
	case *op.Move:
		err = sys.Move(v.Dst, v.Src)
		if err != nil {
			return err
		}
	case *op.Add:
		err = sys.Add(v.Dst, v.Src1, v.Src2)
		if err != nil {
			return err
		}
	case *op.AddFloat:
		err = sys.AddFloat(v.Dst, v.Src1, v.Src2)
		if err != nil {
			return err
		}
	case *op.Or:
		err = sys.Or(v.Dst, v.Src1, v.Src2)
		if err != nil {
			return err
		}
	case *op.And:
		err = sys.And(v.Dst, v.Src1, v.Src2)
		if err != nil {
			return err
		}
	case *op.Xor:
		err = sys.Xor(v.Dst, v.Src1, v.Src2)
		if err != nil {
			return err
		}
	case *op.Ror:
		err = sys.Ror(v.Reg, v.X)
		if err != nil {
			return err
		}
	case *op.CmpBranch:
		err = sys.CmpBranch(v.Cmp, v.Addr)
		if err != nil {
			return err
		}
	case *op.Halt:
		err = sys.Halt()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("System.Step: instruction (%T) not handled", inst)
	}
	return nil
}

// Run starts the system and executes instructions until an error or ErrHalted.
// A successful call to Run return err == nil, not err == ErrHalted. Because Run
// is defined to execute instructions until ErrHalted, it does not treat
// ErrHalted as an error to report.
func (sys *System) Run() (err error) {
	sys.Start()
	for {
		err = sys.Step()
		if err != nil {
			if err == ErrHalted {
				return nil
			}
			return err
		}
	}
}

// Start starts the system, so it can execute instructions.
func (sys *System) Start() {
	sys.running = true
}

// Reset resets the system by clearing the memory and all registers. It will
// also halt the system.
//
// Remember to call sys.Start before executing instructions.
func (sys *System) Reset() {
	// Clear program counter.
	sys.PC = 0

	// Clear registers.
	for i := range sys.Regs {
		sys.Regs[i] = 0
	}

	// Clear memory.
	for i := range sys.Mem {
		sys.Mem[i] = 0
	}

	// Halt system.
	sys.running = false
}
