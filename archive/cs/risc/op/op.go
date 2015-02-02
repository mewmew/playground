// Package op provides access to the instruction set of a simple RISC
// architecture.
//
// ref: Appendix C: Computer science - An overview, 11th ed.
package op

import "fmt"

// Information about the architecture.
const (
	// InstSize specifies the size in bytes of an encoded instruction.
	InstSize = 2
	// RegCount specifies the number of general purpose registers on the system.
	RegCount = 16
	// RegSize specifies the bit size of the general purpose registers.
	RegSize = 8
)

// Code represents an op-code.
type Code uint8

// Op-codes.
const (
	// CodeNop performs no operation.
	//
	//    op-code: 0
	//    operand: 000
	CodeNop Code = iota
	// CodeLoadMem loads the contents of the src memory address into the dst
	// register.
	//
	//    op-code: 1
	//    operand: RXY
	//       R refers to the dst register.
	//       XY refers to the src memory address.
	CodeLoadMem
	// CodeLoadVal loads the src immediate value into the dst register.
	//
	//    op-code: 2
	//    operand: RXY
	//       R refers to the dst register.
	//       XY refers to the src immediate value.
	CodeLoadVal
	// CodeStore stores the contents of the src register into the dst memory
	// address.
	//
	//    op-code: 3
	//    operand: RXY
	//       R refers to the src register.
	//       XY refers to the dst memory address.
	CodeStore
	// CodeMove moves the contents of the src register into the dst register.
	//
	//    op-code: 4
	//    operand: 0RS
	//       R refers to the src register.
	//       S refers to the dst register.
	CodeMove
	// CodeAdd adds the contents of the src1 and src2 registers, as though they
	// represented values in two's complement notation, and store the result in
	// the dst register.
	//
	//    op-code: 5
	//    operand: RST
	//       R refers to the dst register.
	//       S refers to the src1 register.
	//       T refers to the src2 register.
	CodeAdd
	// CodeAddFloat adds the contents of the src1 and src2 registers, as though
	// they represented values in floating-point notation, and store the result
	// in the dst register.
	//
	//    op-code: 6
	//    operand: RST
	//       R refers to the dst register.
	//       S refers to the src1 register.
	//       T refers to the src2 register.
	CodeAddFloat
	// CodeOr performs a bitwise OR operation between the bit patterns in the
	// src1 and src2 registers and store the result in the dst register.
	//
	//    op-code: 7
	//    operand: RST
	//       R refers to the dst register.
	//       S refers to the src1 register.
	//       T refers to the src2 register.
	CodeOr
	// CodeAnd performs a bitwise AND operation between the bit patterns in the
	// src1 and src2 registers and store the result in the dst register.
	//
	//    op-code: 8
	//    operand: RST
	//       R refers to the dst register.
	//       S refers to the src1 register.
	//       T refers to the src2 register.
	CodeAnd
	// CodeXor performs a bitwise XOR operation between the bit patterns in the
	// src1 and src2 registers and store the result in the dst register.
	//
	//    op-code: 9
	//    operand: RST
	//       R refers to the dst register.
	//       S refers to the src1 register.
	//       T refers to the src2 register.
	CodeXor
	// CodeRor rotates the bit pattern in the reg register x bits to the right.
	// Each time a bit is rotated out of the low-order end it is placed at the
	// high-order end.
	//
	//    op-code: A
	//    operand: R0X
	//       R refers to the register.
	//       X refers to the immediate value x.
	CodeRor
	// CodeCmpBranch jumps to the instruction located at the addr memory address
	// if the contents of the cmp registers is equal to the contents of the 0
	// register. Otherwise, continue with the normal sequence of execution.
	//
	// This jump is "unconditional" when cmp == 0.
	//
	//    op-code: B
	//    operand: RXY
	//       R refers to the cmp register.
	//       XY refers to the memory address addr.
	CodeCmpBranch
	// CodeHalt halts execution.
	//
	//    op-code: C
	//    operand: 000
	CodeHalt
)

func (code Code) String() string {
	m := map[Code]string{
		CodeNop:       "NOP",  // no operation
		CodeLoadMem:   "LDR",  // load register
		CodeLoadVal:   "LDR",  // load register
		CodeStore:     "STR",  // store register
		CodeMove:      "MOV",  // move
		CodeAdd:       "ADD",  // add
		CodeAddFloat:  "FADD", // floating-point add
		CodeOr:        "OR",   // or
		CodeAnd:       "AND",  // and
		CodeXor:       "XOR",  // exclusive or
		CodeRor:       "ROR",  // rotate right
		CodeCmpBranch: "CBE",  // compare and branch equal
		CodeHalt:      "HLT",  // halt
	}
	s, ok := m[code]
	if ok {
		return s
	}
	return fmt.Sprintf("<invalid op-code: %d>", int(code))
}
