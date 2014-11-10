package emu

import (
	"fmt"

	"github.com/mewmew/playground/archive/cs/float8"
	"github.com/mewmew/playground/archive/cs/risc/op"
)

// Nop performs no operation.
func (sys *System) Nop() {
}

// LoadMem loads the contents of the src memory address into the dst register.
func (sys *System) LoadMem(dst op.Reg, src op.Addr) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.LoadMem: invalid dst register %d.", dst)
	}
	sys.Regs[dst] = sys.Mem[src]
	return nil
}

// LoadVal loads the src immediate value into the dst register.
func (sys *System) LoadVal(dst op.Reg, src op.Val) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.LoadVal: invalid dst register %d.", dst)
	}
	sys.Regs[dst] = uint8(src)
	return nil
}

// Store stores the contents of the src register into the dst memory address.
func (sys *System) Store(dst op.Addr, src op.Reg) (err error) {
	if int(src) >= len(sys.Regs) {
		return fmt.Errorf("System.Store: invalid src register %d.", src)
	}
	sys.Mem[dst] = sys.Regs[src]
	return nil
}

// Move moves the contents of the src register into the dst register.
func (sys *System) Move(dst op.Reg, src op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.Move: invalid dst register %d.", dst)
	}
	if int(src) >= len(sys.Regs) {
		return fmt.Errorf("System.Move: invalid src register %d.", src)
	}
	sys.Regs[dst] = sys.Regs[src]
	return nil
}

// Add adds the contents of the src1 and src2 registers, as though they
// represented values in two's complement notation, and stores the result in the
// dst register.
func (sys *System) Add(dst op.Reg, src1 op.Reg, src2 op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.Add: invalid dst register %d.", dst)
	}
	if int(src1) >= len(sys.Regs) {
		return fmt.Errorf("System.Add: invalid src1 register %d.", src1)
	}
	if int(src2) >= len(sys.Regs) {
		return fmt.Errorf("System.Add: invalid src2 register %d.", src2)
	}
	sys.Regs[dst] = sys.Regs[src1] + sys.Regs[src2]
	return nil
}

// AddFloat adds the contents of the src1 and src2 registers, as though they
// represented values in floating-point notation, and stores the result in the
// dst register.
func (sys *System) AddFloat(dst op.Reg, src1 op.Reg, src2 op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.AddFloat: invalid dst register %d.", dst)
	}
	if int(src1) >= len(sys.Regs) {
		return fmt.Errorf("System.AddFloat: invalid src1 register %d.", src1)
	}
	if int(src2) >= len(sys.Regs) {
		return fmt.Errorf("System.AddFloat: invalid src2 register %d.", src2)
	}
	x := float8.Float8(sys.Regs[src1])
	y := float8.Float8(sys.Regs[src2])
	z, err := float8.Add(x, y)
	if err != nil {
		return err
	}
	sys.Regs[dst] = uint8(z)
	return nil
}

// Or performs a bitwise OR operation on the bit patterns in the src1 and src2
// registers and stores the result in the dst register.
func (sys *System) Or(dst op.Reg, src1 op.Reg, src2 op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.Or: invalid dst register %d.", dst)
	}
	if int(src1) >= len(sys.Regs) {
		return fmt.Errorf("System.Or: invalid src1 register %d.", src1)
	}
	if int(src2) >= len(sys.Regs) {
		return fmt.Errorf("System.Or: invalid src2 register %d.", src2)
	}
	sys.Regs[dst] = sys.Regs[src1] | sys.Regs[src2]
	return nil
}

// And performs a bitwise AND operation on the bit patterns in the src1 and src2
// registers and stores the result in the dst register.
func (sys *System) And(dst op.Reg, src1 op.Reg, src2 op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.And: invalid dst register %d.", dst)
	}
	if int(src1) >= len(sys.Regs) {
		return fmt.Errorf("System.And: invalid src1 register %d.", src1)
	}
	if int(src2) >= len(sys.Regs) {
		return fmt.Errorf("System.And: invalid src2 register %d.", src2)
	}
	sys.Regs[dst] = sys.Regs[src1] & sys.Regs[src2]
	return nil
}

// Xor performs a bitwise XOR operation on the bit patterns in the src1 and src2
// registers and stores the result in the dst register.
func (sys *System) Xor(dst op.Reg, src1 op.Reg, src2 op.Reg) (err error) {
	if int(dst) >= len(sys.Regs) {
		return fmt.Errorf("System.Xor: invalid dst register %d.", dst)
	}
	if int(src1) >= len(sys.Regs) {
		return fmt.Errorf("System.Xor: invalid src1 register %d.", src1)
	}
	if int(src2) >= len(sys.Regs) {
		return fmt.Errorf("System.Xor: invalid src2 register %d.", src2)
	}
	sys.Regs[dst] = sys.Regs[src1] ^ sys.Regs[src2]
	return nil
}

// Ror rotate the bit pattern in the reg register x bits to the right. Each time
// a bit is rotated out of the low-order end it is placed at the high-order end.
func (sys *System) Ror(reg op.Reg, x op.Val) (err error) {
	if int(reg) >= len(sys.Regs) {
		return fmt.Errorf("System.Ror: invalid register %d.", reg)
	}
	if x >= op.RegSize {
		return fmt.Errorf("System.Ror: invalid x (%d); above %d.", x, op.RegSize-1)
	}
	v := sys.Regs[reg]
	sys.Regs[reg] = (v >> x) | (v << (op.RegSize - x))
	return nil
}

// CmpBranch jumps to the instruction located at the addr memory address if the
// contents of the cmp registers is equal to the contents of the 0 register.
// Otherwise, continue with the normal sequence of execution.
//
// This jump is "unconditional" when cmp == 0.
func (sys *System) CmpBranch(cmp op.Reg, addr op.Addr) (err error) {
	if int(cmp) >= len(sys.Regs) {
		return fmt.Errorf("System.CmpBranch: invalid cmp register %d.", cmp)
	}
	if sys.Regs[cmp] == sys.Regs[0] {
		sys.PC = PC(addr)
	}
	return nil
}

// Halt halts execution of the system.
func (sys *System) Halt() (err error) {
	sys.running = false
	// Rewind the program counter, so it still points to the halt instruction.
	err = sys.PC.Inc(-op.InstSize)
	if err != nil {
		return err
	}
	return nil
}
