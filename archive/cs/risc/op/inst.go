package op

import "fmt"

// Reg represents a register between 0 and 15.
type Reg uint8

func (reg Reg) String() string {
	return fmt.Sprintf("r%d", uint8(reg))
}

// Addr represents a memory address between 0 and 255.
type Addr uint8

func (addr Addr) String() string {
	return fmt.Sprintf("0x%02X", uint8(addr))
}

// Val represents an immediate value between 0 and 255.
type Val uint8

func (val Val) String() string {
	return fmt.Sprintf("$%d", uint8(val))
}

// The Nop instruction performs no operation.
type Nop struct {
	Code Code
}

func (inst *Nop) String() string {
	return inst.Code.String()
}

// The LoadMem instruction loads the contents of the src memory address into the
// dst register.
type LoadMem struct {
	Code Code
	Dst  Reg
	Src  Addr
}

func (inst *LoadMem) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// The LoadVal instruction loads the src immediate value into the dst register.
type LoadVal struct {
	Code Code
	Dst  Reg
	Src  Val
}

func (inst *LoadVal) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// The Store instruction stores the contents of the src register into the dst
// memory address.
type Store struct {
	Code Code
	Dst  Addr
	Src  Reg
}

func (inst *Store) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// The Move instruction moves the contents of the src register into the dst
// register.
type Move struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Move) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// The Add instruction adds the contents of the src1 and src2 registers, as
// though they represented values in two's complement notation, and store the
// result in the dst register.
type Add struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Add) String() string {
	return fmt.Sprintf("%-8s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// The AddFloat instruction adds the contents of the src1 and src2 registers, as
// though they represented values in floating-point notation, and store the
// result in the dst register.
type AddFloat struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *AddFloat) String() string {
	return fmt.Sprintf("%-8s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// The Or instruction performs a bitwise OR operation with the src1 and src2
// registers and store the result in the dst register.
type Or struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Or) String() string {
	return fmt.Sprintf("%-8s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// The And instruction performs a bitwise AND operation with the src1 and src2
// registers and store the result in the dst register.
type And struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *And) String() string {
	return fmt.Sprintf("%-8s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// The Xor instruction performs a bitwise XOR operation with the src1 and src2
// registers and store the result in the dst register.
type Xor struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Xor) String() string {
	return fmt.Sprintf("%-8s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// The Ror instruction rotates the bit pattern in the reg register x bits to the
// right. Each time a bit is rotated out of the low-order end it is placed at
// the high-order end.
type Ror struct {
	Code Code
	Reg  Reg
	X    Val
}

func (inst *Ror) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Reg, inst.X)
}

// The CmpBranch instruction jumps to the instruction located at the addr memory
// address if the contents of the cmp registers is equal to the contents of the
// 0 register. Otherwise, continue with the normal sequence of execution.
//
// This jump is "unconditional" when cmp == 0.
type CmpBranch struct {
	Code Code
	Cmp  Reg
	Addr Addr
}

func (inst *CmpBranch) String() string {
	return fmt.Sprintf("%-8s%s, %s", inst.Code, inst.Cmp, inst.Addr)
}

// The Halt instruction halts execution.
type Halt struct {
	Code Code
}

func (inst *Halt) String() string {
	return inst.Code.String()
}
