// Package orbis provides access to the OpenRISC Basic Instruction Set (ORBIS32)
// [1].
//
// [1]: http://opencores.org/websvn,filedetails?repname=openrisc&path=%2Fopenrisc%2Ftrunk%2Fdocs%2Fopenrisc-arch-1.0-rev0.pdf
package orbis

import (
	"fmt"
)

// General architecture information.
const (
	// InstSize specifies the size in bytes of an encoded instruction.
	InstSize = 2
	// RegSize specifies the bit size of the general purpose registers.
	RegSize = 32
)

// Customizable architecture information.
var (
	// RegCount specifies the number of general purpose registers that are
	// available. The OpenRISC 1000 register set includes thirty-two or sixteen
	// 32-bit general purpose registers.
	RegCount = 32
)

// Code represents an opcode.
type Code int

// Opcodes.
const (
	// Name:
	//    Add
	// Format:
	//    l.add rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA are added to the contents
	//    of general-purpose register rB to form the result. The result is placed
	//    into general-purpose register rD.
	//
	//    The instruction will set the carry flag on unsigned overflow, and the
	//    overflow flag on signed overflow.
	CodeAdd Code = iota
	// Name:
	//    Add and Carry
	// Format:
	//    l.addc rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA are added to the contents
	//    of general-purpose register rB and carry SR[CY] to form the result. The
	//    result is placed into general-purpose register rD.
	//
	//    The instruction will set the carry flag on unsigned overflow, and the
	//    overflow flag on signed overflow.
	CodeAddc
	// Name:
	//    Add Immediate
	// Format:
	//    l.addi rD,rA,I
	// Description:
	//    The immediate value is sign-extended and added to the contents of
	//    general-purpose register rA to form the result. The result is placed
	//    into general-purpose register rD.
	//
	//    The instruction will set the carry flag on unsigned overflow, and the
	//    overflow flag on signed overflow.
	CodeAddi
	// Name:
	//    Add Immediate and Carry
	// Format:
	//    l.addic rD,rA,I
	// Description:
	//    The immediate value is sign-extended and added to the contents of
	//    general-purpose register rA and carry SR[CY] to form the result. The
	//    result is placed into general-purpose register rD.
	//
	//    The instruction will set the carry flag on unsigned overflow, and the
	//    overflow flag on signed overflow.
	CodeAddic
	// Name:
	//    And
	// Format:
	//    l.and rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA are combined with the
	//    contents of general-purpose register rB in a bit-wise logical AND
	//    operation. The result is placed into general-purpose register rD.
	CodeAnd
	// Name:
	//    And with Immediate Half Word
	// Format:
	//    l.andi rD,rA,K
	// Description:
	//    The immediate value is zero-extended and combined with the contents of
	//    general-purpose register rA in a bit-wise logical AND operation. The
	//    result is placed into general-purpose register rD.
	CodeAndi
	// Name:
	//    Branch if Flag
	// Format:
	//    l.bf N
	// Description:
	//    The immediate value is shifted left two bits, sign-extended to program
	//    counter width, and then added to the address of the branch instruction.
	//    The result is the effective address of the branch. If the flag is set,
	//    the program branches to EA. If CPUCFGR[ND] is not set, the branch
	//    occurs with a delay of one instruction.
	CodeBf
	// Name:
	//    Branch if No Flag
	// Format:
	//    l.bnf N
	// Description:
	//    The immediate value is shifted left two bits, sign-extended to program
	//    counter width, and then added to the address of the branch instruction.
	//    The result is the effective address of the branch. If the flag is
	//    cleared, the program branches to EA. If CPUCFGR[ND] is not set, the
	//    branch occurs with a delay of one instruction.
	CodeBnf
	// Name:
	//    Conditional Move
	// Format:
	//    l.cmov rD,rA,rB
	// Description:
	//    If SR[F] is set, general-purpose register rA is placed in
	//    general-purpose register rD. If SR[F] is cleared, general-purpose
	//    register rB is placed in general-purpose register rD.
	CodeCmov
	// Name:
	//    Context Syncronization
	// Format:
	//    l.csync
	// Description:
	//    Execution of context synchronization instruction results in completion
	//    of all operations inside the processor and a flush of the instruction
	//    pipelines. When all operations are complete, the RISC core resumes with
	//    an empty instruction pipeline and fresh context in all units (MMU for
	//    example).
	CodeCsync
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust1
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust1
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust2
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust2
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust3
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust3
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust4
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust4
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust5 rD,rA,rB,L,K
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust5
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust6
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust6
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust7
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust7
	// Name:
	//    Reserved for ORBIS32/64 Custom Instructions
	// Format:
	//    l.cust8
	// Description:
	//    This fake instruction only allocates instruction set space for custom
	//    instructions. Custom instructions are those that are not defined by the
	//    architecture but rather by the implementation itself.
	CodeCust8
	// Name:
	//    Divide Signed
	// Format:
	//    l.div rD,rA,rB
	// Description:
	//    The content of general-purpose register rA are divided by the content
	//    of general-purpose register rB, and the result is placed into
	//    general-purpose register rD. Both operands are treated as signed
	//    integers.
	//
	//    On divide-by zero, rD will be undefined, and the overflow flag will be
	//    set. Note that prior revisions of the manual (pre-2011) stored the
	//    divide by zero flag in SR[CY].
	CodeDiv
	// Name:
	//    Divide Unsigned
	// Format:
	//    l.divu rD,rA,rB
	// Description:
	//    The content of general-purpose register rA are divided by the content
	//    of general-purpose register rB, and the result is placed into
	//    general-purpose register rD. Both operands are treated as unsigned
	//    integers.
	//
	//    On divide-by zero, rD will be undefined, and the overflow flag will be
	//    set.
	CodeDivu
	// Name:
	//    Extend Byte with Sign
	// Format:
	//    l.extbs rD,rA
	// Description:
	//    Bit 7 of general-purpose register rA is placed in high-order bits of
	//    general-purpose register rD. The low-order eight bits of
	//    general-purpose register rA are copied into the low-order eight bits of
	//    general-purpose register rD.
	CodeExtbs
	// Name:
	//    Extend Byte with Zero
	// Format:
	//    l.extbz rD,rA
	// Description:
	//    Zero is placed in high-order bits of general-purpose register rD. The
	//    low-order eight bits of general-purpose register rA are copied into the
	//    low-order eight bits of general-purpose register rD.
	CodeExtbz
	// Name:
	//    Extend Half Word with Sign
	// Format:
	//    l.exths rD,rA
	// Description:
	//    Bit 15 of general-purpose register rA is placed in high-order bits of
	//    general-purpose register rD. The low-order 16 bits of general-purpose
	//    register rA are copied into the low-order 16 bits of general-purpose
	//    register rD.
	CodeExths
	// Name:
	//    Extend Half Word with Zero
	// Format:
	//    l.exthz rD,rA
	// Description:
	//    Zero is placed in high-order bits of general-purpose register rD. The
	//    low-order 16 bits of general-purpose register rA are copied into the
	//    low-order 16 bits of general-purpose register rD.
	CodeExthz
	// Name:
	//    Extend Word with Sign
	// Format:
	//    l.extws rD,rA
	// Description:
	//    Bit 31 of general-purpose register rA is placed in high-order bits of
	//    general-purpose register rD. The low-order 32 bits of general-purpose
	//    register rA are copied from low-order 32 bits of general-purpose
	//    register rD.
	CodeExtws
	// Name:
	//    Extend Word with Zero
	// Format:
	//    l.extwz rD,rA
	// Description:
	//    Zero is placed in high-order bits of general-purpose register rD. The
	//    low-order 32 bits of general-purpose register rA are copied into the
	//    low-order 32 bits of general-purpose register rD.
	CodeExtwz
	// Name:
	//    Find First 1
	// Format:
	//    l.ff1 rD,rA
	// Description:
	//    Position of the lowest order '1' bit is written into general-purpose
	//    register rD. Checking for bit '1' starts with bit 0 (LSB), and counting
	//    is incremented for every zero bit. If first '1' bit is discovered in
	//    LSB, one is written into rD, if first '1' bit is discovered in MSB, 32
	//    (64) is written into rD. If there is no '1' bit, zero is written in rD.
	CodeFf1
	// Name:
	//    Find Last 1
	// Format:
	//    l.fl1 rD,rA
	// Description:
	//    Position of the highest order '1' bit is written into general-purpose
	//    register rD. Checking for bit '1' starts with bit 31/63 (MSB), and
	//    counting is decremented for every zero bit until the last ‘1’ bit is
	//    found nearing the LSB. If highest order '1' bit is discovered in MSB,
	//    32 (64) is written into rD, if highest order '1' bit is discovered in
	//    LSB, one is written into rD. If there is no '1' bit, zero is written in
	//    rD.
	CodeFl1
	// Name:
	//    Jump
	// Format:
	//    l.j N
	// Description:
	//    The immediate value is shifted left two bits, sign-extended to program
	//    counter width, and then added to the address of the jump instruction.
	//    The result is the effective address of the jump. The program
	//    unconditionally jumps to EA. If CPUCFGR[ND] is not set, the jump occurs
	//    with a delay of one instruction.
	//
	//    Note that l.sys should not be placed in the delay slot after a jump.
	CodeJ
	// Name:
	//    Jump and Link
	// Format:
	//    l.jal N
	// Description:
	//    The immediate value is shifted left two bits, sign-extended to program
	//    counter width, and then added to the address of the jump instruction.
	//    The result is the effective address of the jump. The program
	//    unconditionally jumps to EA. If CPUCFGR[ND] is not set, the jump occurs
	//    with a delay of one instruction. The address of the instruction after
	//    the delay slot is placed in the link register r9.
	//
	//    The value of the link register, if read as an operand in the delay slot
	//    will be the new value, not the old value. If the link register is
	//    written in the delay slot, the value written will replace the value
	//    stored by the l.jal instruction.
	//
	//    Note that l.sys should not be placed in the delay slot after a jump.
	CodeJal
	// Name:
	//    Jump and Link Register
	// Format:
	//    l.jalr rB
	// Description:
	//    The contents of general-purpose register rB is the effective address of
	//    the jump. The program unconditionally jumps to EA. If CPUCFGR[ND] is
	//    not set, the jump occurs with a delay of one instruction. The address
	//    of the instruction after the delay slot is placed in the link register.
	//
	//    It is not allowed to specify link register r9 as rB. This is because an
	//    exception in the delay slot (including external interrupts) may cause
	//    l.jalr to be reexecuted.
	//
	//    The value of the link register, if read as an operand in the delay slot
	//    will be the new value, not the old value. If the link register is
	//    written in the delay slot, the value written will replace the value
	//    stored by the l.jalr instruction.
	//
	//    Note that l.sys should not be placed in the delay slot after a jump.
	CodeJalr
	// Name:
	//    Jump Register
	// Format:
	//    l.jr rB
	// Description:
	//    The contents of general-purpose register rB is the effective address of
	//    the jump. The program unconditionally jumps to EA. If CPUCFGR[ND] is
	//    not set, the jump occurs with a delay of one instruction.
	//
	//    Note that l.sys should not be placed in the delay slot after a jump.
	CodeJr
	// Name:
	//    Load Byte and Extend with Sign
	// Format:
	//    l.lbs rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The byte in memory addressed by EA is loaded into the low-order eight
	//    bits of general-purpose register rD. High-order bits of general-purpose
	//    register rD are replaced with bit 7 of the loaded value.
	CodeLbs
	// Name:
	//    Load Byte and Extend with Zero
	// Format:
	//    l.lbz rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The byte in memory addressed by EA is loaded into the low-order eight
	//    bits of general-purpose register rD. High-order bits of general-purpose
	//    register rD are replaced with zero.
	CodeLbz
	// Name:
	//    Load Double Word
	// Format:
	//    l.ld rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The double word in memory addressed by EA is loaded into
	//    general-purpose register rD.
	CodeLd
	// Name:
	//    Load Half Word and Extend with Sign
	// Format:
	//    l.lhs rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The half word in memory addressed by EA is loaded into the low-order
	//    16 bits of general-purpose register rD. High-order bits of
	//    general-purpose register rD are replaced with bit 15 of the loaded value.
	CodeLhs
	// Name:
	//    Load Half Word and Extend with Zero
	// Format:
	//    l.lhz rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The half word in memory addressed by EA is loaded into the low-order 16
	//    bits of general-purpose register rD. High-order bits of general-purpose
	//    register rD are replaced with zero.
	CodeLhz
	// Name:
	//    Load Single Word and Extend with Sign
	// Format:
	//    l.lws rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The single word in memory addressed by EA is loaded into the low-order
	//    32 bits of general-purpose register rD. High-order bits of
	//    general-purpose register rD are replaced with bit 31 of the loaded
	//    value.
	CodeLws
	// Name:
	//    Load Single Word and Extend with Zero
	// Format:
	//    l.lwz rD,I(rA)
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The single word in memory addressed by EA is loaded into the low-order
	//    32 bits of general-purpose register rD. High-order bits of
	//    general-purpose register rD are replaced with zero.
	CodeLwz
	// Name:
	//    Multiply and Accumulate Signed
	// Format:
	//    l.mac rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the 64 bit result is
	//    added to the special-purpose registers MACHI and MACLO. All operands
	//    are treated as signed integers.
	//
	//    The instruction will set the overflow flag if signed overflow is
	//    detecting during the addition stage.
	CodeMac
	// Name:
	//    Multiply Immediate and Accumulate Signed
	// Format:
	//    l.maci rA,I
	// Description:
	//    The immediate value and the contents of general-purpose register rA are
	//    multiplied, and the 64 bit result is added to the special-purpose
	//    registers MACHI and MACLO. All operands are treated as signed integers.
	//
	//    The instruction will set the overflow flag if signed overflow is
	//    detecting during the addition stage.
	CodeMaci
	// Name:
	//    MAC Read and Clear
	// Format:
	//    l.macrc rD
	// Description:
	//    Once all instructions in MAC pipeline are completed, the contents of
	//    MAC is placed into general-purpose register rD and MAC accumulator is
	//    cleared.
	//
	//    The MAC pipeline also synchronizes with the instruction pipeline on any
	//    access to MACLO or MACHI SPRs, so that l.mfspr can be used to read
	//    MACHI before executing l.macrc.
	CodeMacrc
	// Name:
	//    Multiply and Accumulate Unsigned
	// Format:
	//    l.macu rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the 64 bit result is
	//    added to the special-purpose registers MACHI and MACLO. All operands
	//    are treated as unsigned integers.
	//
	//    The instruction will set the overflow flag if unsigned overflow is
	//    detecting during the addition stage.
	CodeMacu
	// Name:
	//    Move From Special-Purpose Register
	// Format:
	//    l.mfspr rD,rA,K
	// Description:
	//    The contents of the special register, defined by contents of
	//    general-purpose rA logically ORed with immediate value, are moved into
	//    general-purpose register rD.
	CodeMfspr
	// Name:
	//    Move Immediate High
	// Format:
	//    l.movhi rD,K
	// Description:
	//    The 16-bit immediate value is zero-extended, shifted left by 16 bits,
	//    and placed into general-purpose register rD.
	CodeMovhi
	// Name:
	//    Multiply and Subtract Signed
	// Format:
	//    l.msb rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the 64 bit result is
	//    subtracted from the special-purpose registers MACHI and MACLO. Result
	//    of the subtraction is placed into MACHI and MACLO registers. All
	//    operands are treated as signed integers.
	//
	//    The instruction will set the overflow flag if signed overflow is
	//    detecting during the subtraction stage.
	CodeMsb
	// Name:
	//    Multiply and Subtract Unsigned
	// Format:
	//    l.msbu rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the 64 bit result is
	//    subtracted from the special-purpose registers MACHI and MACLO. Result
	//    of the subtraction is placed into MACHI and MACLO registers. All
	//    operands are treated as unsigned integers.
	//
	//    The instruction will set the overflow flag if unsigned overflow is
	//    detecting during the subtraction stage.
	CodeMsbu
	// Name:
	//    Memory Syncronization
	// Format:
	//    l.msync
	// Description:
	//    Execution of the memory synchronization instruction results in
	//    completion of all load/store operations before the RISC core continues.
	CodeMsync
	// Name:
	//    Move To Special-Purpose Register
	// Format:
	//    l.mtspr rA,rB,K
	// Description:
	//    The contents of general-purpose register rB are moved into the special
	//    register defined by contents of general-purpose register rA logically
	//    ORed with the immediate value.
	CodeMtspr
	// Name:
	//    Multiply Signed
	// Format:
	//    l.mul rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the result is truncated
	//    to destination register width and placed into general-purpose register
	//    rD. Both operands are treated as signed integers.
	//
	//    The instruction will set the overflow flag on signed overflow.
	CodeMul
	// Name:
	//    Multiply Signed to Double
	// Format:
	//    l.muld rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the result is stored in
	//    the MACHI and MACLO registers. Both operands are treated as signed
	//    integers.
	//
	//    The instruction will set the overflow flag on signed overflow.
	CodeMuld
	// Name:
	//    Multiply Unsigned to Double
	// Format:
	//    l.muldu rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the result is stored in
	//    the MACHI and MACLO registers. Both operands are treated as unsigned
	//    integers.
	//
	//    The instruction will set the overflow flag on unsigned overflow.
	CodeMuldu
	// Name:
	//    Multiply Immediate Signed
	// Format:
	//    l.muli rD,rA,I
	// Description:
	//    The immediate value and the contents of general-purpose register rA are
	//    multiplied, and the result is truncated to destination register width
	//    and placed into general-purpose register rD.
	//
	//    The instruction will set the overflow flag on signed overflow.
	CodeMuli
	// Name:
	//    Multiply Unsigned
	// Format:
	//    l.mulu rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA and the contents of
	//    general-purpose register rB are multiplied, and the result is truncated
	//    to destination register width and placed into general-purpose register
	//    rD. Both operands are treated as unsigned integers.
	//
	//    The instruction will set the carry flag on unsigned overflow.
	CodeMulu
	// Name:
	//    No Operation
	// Format:
	//    l.nop K
	// Description:
	//    This instruction does not do anything except that it takes at least one
	//    clock cycle to complete. It is often used to fill delay slot gaps.
	//    Immediate value can be used for simulation purposes.
	CodeNop
	// Name:
	//    Or
	// Format:
	//    l.or rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA are combined with the
	//    contents of general-purpose register rB in a bit-wise logical OR
	//    operation. The result is placed into general-purpose register rD.
	CodeOr
	// Name:
	//    Or with Immediate Half Word
	// Format:
	//    l.ori rD,rA,K
	// Description:
	//    The immediate value is zero-extended and combined with the contents of
	//    general-purpose register rA in a bit-wise logical OR operation. The
	//    result is placed into general-purpose register rD.
	CodeOri
	// Name:
	//    Pipeline Syncronization
	// Format:
	//    l.psync
	// Description:
	//    Execution of pipeline synchronization instruction results in completion
	//    of all instructions that were fetched before l.psync instruction. Once
	//    all instructions are completed, instructions fetched after l.psync are
	//    flushed from the pipeline and fetched again.
	CodePsync
	// Name:
	//    Return From Exception
	// Format:
	//    l.rfe
	// Description:
	//    Execution of this instruction partially restores the state of the
	//    processor prior to the exception. This instruction does not have a
	//    delay slot.
	CodeRfe
	// Name:
	//    Rotate Right
	// Format:
	//    l.ror rD,rA,rB
	// Description:
	//    General-purpose register rB specifies the number of bit positions; the
	//    contents of general-purpose register rA are rotated right. The result
	//    is written into general-purpose register rD.
	CodeRor
	// Name:
	//    Rotate Right with Immediate
	// Format:
	//    l.rori rD,rA,L
	// Description:
	//    The 6-bit immediate value specifies the number of bit positions; the
	//    contents of general-purpose register rA are rotated right. The result
	//    is written into general-purpose register rD. In 32-bit implementations
	//    bit 5 of immediate is ignored.
	CodeRori
	// Name:
	//    Store Byte
	// Format:
	//    l.sb I(rA),rB
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The low-order 8 bits of general-purpose register rB are stored to
	//    memory location addressed by EA.
	CodeSb
	// Name:
	//    Store Double Word
	// Format:
	//    l.sd I(rA),rB
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The double word in general-purpose register rB is stored to memory
	//    location addressed by EA.
	CodeSd
	// Name:
	//    Set Flag if Equal
	// Format:
	//    l.sfeq rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared. If
	//    the contents are equal, the compare flag is set; otherwise the compare
	//    flag is cleared.
	CodeSfeq
	// Name:
	//    Set Flag if Equal Immediate
	// Format:
	//    l.sfeqi rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared. If the two values are equal, the compare
	//    flag is set; otherwise the compare flag is cleared.
	CodeSfeqi
	// Name:
	//    Set Flag if Greater or Equal Than Signed
	// Format:
	//    l.sfges rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    signed integers. If the contents of the first register are greater than
	//    or equal to the contents of the second register, the compare flag is
	//    set; otherwise the compare flag is cleared.
	CodeSfges
	// Name:
	//    Set Flag if Greater or Equal Than Immediate Signed
	// Format:
	//    l.sfgesi rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as signed integers. If the contents of the
	//    first register are greater than or equal to the immediate value the
	//    compare flag is set; otherwise the compare flag is cleared.
	CodeSfgesi
	// Name:
	//    Set Flag if Greater or Equal Than Unsigned
	// Format:
	//    l.sfgeu rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    unsigned integers. If the contents of the first register are greater
	//    than or equal to the contents of the second register, the compare flag
	//    is set; otherwise the compare flag is cleared.
	CodeSfgeu
	// Name:
	//    Set Flag if Greater or Equal Than Immediate Unsigned
	// Format:
	//    l.sfgeui rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as unsigned integers. If the contents of
	//    the first register are greater than or equal to the immediate value the
	//    compare flag is set; otherwise the compare flag is cleared.
	CodeSfgeui
	// Name:
	//    Set Flag if Greater Than Signed
	// Format:
	//    l.sfgts rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    signed integers. If the contents of the first register are greater than
	//    the contents of the second register, the compare flag is set; otherwise
	//    the compare flag is cleared.
	CodeSfgts
	// Name:
	//    Set Flag if Greater Than Immediate Signed
	// Format:
	//    l.sfgtsi rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as signed integers. If the contents of the
	//    first register are greater than the immediate value the compare flag is
	//    set; otherwise the compare flag is cleared.
	CodeSfgtsi
	// Name:
	//    Set Flag if Greater Than Unsigned
	// Format:
	//    l.sfgtu rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    unsigned integers. If the contents of the first register are greater
	//    than the contents of the second register, the compare flag is set;
	//    otherwise the compare flag is cleared.
	CodeSfgtu
	// Name:
	//    Set Flag if Greater Than Immediate Unsigned
	// Format:
	//    l.sfgtui rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as unsigned integers. If the contents of
	//    the first register are greater than the immediate value the compare
	//    flag is set; otherwise the compare flag is cleared.
	CodeSfgtui
	// Name:
	//    Set Flag if Less or Equal Than Signed
	// Format:
	//    l.sfles rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    signed integers. If the contents of the first register are less than or
	//    equal to the contents of the second register, the compare flag is set;
	//    otherwise the compare flag is cleared.
	CodeSfles
	// Name:
	//    Set Flag if Less or Equal Than Immediate Signed
	// Format:
	//    l.sflesi rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as signed integers. If the contents of the
	//    first register are less than or equal to the immediate value the
	//    compare flag is set; otherwise the compare flag is cleared.
	CodeSflesi
	// Name:
	//    Set Flag if Less or Equal Than Unsigned
	// Format:
	//    l.sfleu rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    unsigned integers. If the contents of the first register are less than
	//    or equal to the contents of the second register, the compare flag is
	//    set; otherwise the compare flag is cleared.
	CodeSfleu
	// Name:
	//    Set Flag if Less or Equal Than Immediate Unsigned
	// Format:
	//    l.sfleui rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as unsigned integers. If the contents of
	//    the first register are less than or equal to the immediate value the
	//    compare flag is set; otherwise the compare flag is cleared.
	CodeSfleui
	// Name:
	//    Set Flag if Less Than Signed
	// Format:
	//    l.sflts rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    signed integers. If the contents of the first register are less than
	//    the contents of the second register, the compare flag is set; otherwise
	//    the compare flag is cleared.
	CodeSflts
	// Name:
	//    Set Flag if Less Than Immediate Signed
	// Format:
	//    l.sfltsi rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as signed integers. If the contents of the
	//    first register are less than the immediate value the compare flag is
	//    set; otherwise the compare flag is cleared.
	CodeSfltsi
	// Name:
	//    Set Flag if Less Than Unsigned
	// Format:
	//    l.sfltu rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared as
	//    unsigned integers. If the contents of the first register are less than
	//    the contents of the second register, the compare flag is set; otherwise
	//    the compare flag is cleared.
	CodeSfltu
	// Name:
	//    Set Flag if Less Than Immediate Unsigned
	// Format:
	//    l.sfltui rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared as unsigned integers. If the contents of
	//    the first register are less than the immediate value the compare flag
	//    is set; otherwise the compare flag is cleared.
	CodeSfltui
	// Name:
	//    Set Flag if Not Equal
	// Format:
	//    l.sfne rA,rB
	// Description:
	//    The contents of general-purpose registers rA and rB are compared. If
	//    the contents are not equal, the compare flag is set; otherwise the
	//    compare flag is cleared.
	CodeSfne
	// Name:
	//    Set Flag if Not Equal Immediate
	// Format:
	//    l.sfnei rA,I
	// Description:
	//    The contents of general-purpose register rA and the sign-extended
	//    immediate value are compared. If the two values are not equal, the
	//    compare flag is set; otherwise the compare flag is cleared.
	CodeSfnei
	// Name:
	//    Store Half Word
	// Format:
	//    l.sh I(rA),rB
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The low-order 16 bits of general-purpose register rB are stored to
	//    memory location addressed by EA.
	CodeSh
	// Name:
	//    Shift Left Logical
	// Format:
	//    l.sll rD,rA,rB
	// Description:
	//    General-purpose register rB specifies the number of bit positions; the
	//    contents of general-purpose register rA are shifted left, inserting
	//    zeros into the low-order bits. The result is written into
	//    general-purpose rD. In 32-bit implementations bit 5 of rB is ignored.
	CodeSll
	// Name:
	//    Shift Left Logical with Immediate
	// Format:
	//    l.slli rD,rA,L
	// Description:
	//    The immediate value specifies the number of bit positions; the contents
	//    of general-purpose register rA are shifted left, inserting zeros into
	//    the low-order bits. The result is written into general-purpose register
	//    rD. In 32-bit implementations bit 5 of immediate is ignored.
	CodeSlli
	// Name:
	//    Shift Right Arithmetic
	// Format:
	//    l.sra rD,rA,rB
	// Description:
	//    General-purpose register rB specifies the number of bit positions; the
	//    contents of general-purpose register rA are shifted right,
	//    sign-extending the high-order bits. The result is written into
	//    general-purpose register rD. In 32-bit implementations bit 5 of rB is
	//    ignored.
	CodeSra
	// Name:
	//    Shift Right Arithmetic with Immediate
	// Format:
	//    l.srai rD,rA,L
	// Description:
	//    The 6-bit immediate value specifies the number of bit positions; the
	//    contents of general-purpose register rA are shifted right,
	//    sign-extending the high-order bits. The result is written into
	//    general-purpose register rD. In 32-bit implementations bit 5 of
	//    immediate is ignored.
	CodeSrai
	// Name:
	//    Shift Right Logical
	// Format:
	//    l.srl rD,rA,rB
	// Description:
	//    General-purpose register rB specifies the number of bit positions; the
	//    contents of general-purpose register rA are shifted right, inserting
	//    zeros into the high-order bits. The result is written into
	//    general-purpose register rD. In 32-bit implementations bit 5 of rB is
	//    ignored.
	CodeSrl
	// Name:
	//    Shift Right Logical with Immediate
	// Format:
	//    l.srli rD,rA,L
	// Description:
	//    The 6-bit immediate value specifies the number of bit positions; the
	//    contents of general-purpose register rA are shifted right, inserting
	//    zeros into the high-order bits. The result is written into
	//    general-purpose register rD. In 32-bit implementations bit 5 of
	//    immediate is ignored.
	CodeSrli
	// Name:
	//    Subtract
	// Format:
	//    l.sub rD,rA,rB
	// Description:
	//    The contents of general-purpose register rB are subtracted from the
	//    contents of general-purpose register rA to form the result. The result
	//    is placed into general-purpose register rD.
	//
	//    The instruction will set the carry flag on unsigned overflow, and the
	//    overflow flag on signed overflow.
	CodeSub
	// Name:
	//    Store Single Word
	// Format:
	//    l.sw I(rA),rB
	// Description:
	//    The offset is sign-extended and added to the contents of
	//    general-purpose register rA. The sum represents an effective address.
	//    The low-order 32 bits of general-purpose register rB are stored to
	//    memory location addressed by EA.
	CodeSw
	// Name:
	//    System Call
	// Format:
	//    l.sys K
	// Description:
	//    Execution of the system call instruction results in the system call
	//    exception. The system calls exception is a request to the operating
	//    system to provide operating system services. The immediate value can be
	//    used to specify which system service is requested, alternatively a GPR
	//    defined by the ABI can be used to specify system service.
	//
	//    Because an l.sys causes an intentional exception, rather than an
	//    interruption of normal processing, the matching l.rfe returns to the
	//    next instruction. As this is considered to be the jump itself for
	//    exceptions occurring in a delay slot, l.sys should not be placed in a
	//    delay slot.
	CodeSys
	// Name:
	//    Trap
	// Format:
	//    l.trap K
	// Description:
	//    Trap exception is a request to the operating system or to the debug
	//    facility to execute certain debug services. Immediate value is used to
	//    select which SR bit is tested by trap instruction.
	CodeTrap
	// Name:
	//    Exclusive Or
	// Format:
	//    l.xor rD,rA,rB
	// Description:
	//    The contents of general-purpose register rA are combined with the
	//    contents of general-purpose register rB in a bit-wise logical XOR
	//    operation. The result is placed into general-purpose register rD.
	CodeXor
	// Name:
	//    Exclusive Or with Immediate Half Word
	// Format:
	//    l.xori rD,rA,I
	// Description:
	//    The immediate value is sign-extended and combined with the contents of
	//    general-purpose register rA in a bit-wise logical XOR operation. The
	//    result is placed into general-purpose register rD.
	CodeXori
)

// mnemonic maps op-codes to their corresponding mnemonic.
var mnemonic = map[Code]string{
	CodeAdd:    "l.add",
	CodeAddc:   "l.addc",
	CodeAddi:   "l.addi",
	CodeAddic:  "l.addic",
	CodeAnd:    "l.and",
	CodeAndi:   "l.andi",
	CodeBf:     "l.bf",
	CodeBnf:    "l.bnf",
	CodeCmov:   "l.cmov",
	CodeCsync:  "l.csync",
	CodeCust1:  "l.cust1",
	CodeCust2:  "l.cust2",
	CodeCust3:  "l.cust3",
	CodeCust4:  "l.cust4",
	CodeCust5:  "l.cust5",
	CodeCust6:  "l.cust6",
	CodeCust7:  "l.cust7",
	CodeCust8:  "l.cust8",
	CodeDiv:    "l.div",
	CodeDivu:   "l.divu",
	CodeExtbs:  "l.extbs",
	CodeExtbz:  "l.extbz",
	CodeExths:  "l.exths",
	CodeExthz:  "l.exthz",
	CodeExtws:  "l.extws",
	CodeExtwz:  "l.extwz",
	CodeFf1:    "l.ff1",
	CodeFl1:    "l.fl1",
	CodeJ:      "l.j",
	CodeJal:    "l.jal",
	CodeJalr:   "l.jalr",
	CodeJr:     "l.jr",
	CodeLbs:    "l.lbs",
	CodeLbz:    "l.lbz",
	CodeLd:     "l.ld",
	CodeLhs:    "l.lhs",
	CodeLhz:    "l.lhz",
	CodeLws:    "l.lws",
	CodeLwz:    "l.lwz",
	CodeMac:    "l.mac",
	CodeMaci:   "l.maci",
	CodeMacrc:  "l.macrc",
	CodeMacu:   "l.macu",
	CodeMfspr:  "l.mfspr",
	CodeMovhi:  "l.movhi",
	CodeMsb:    "l.msb",
	CodeMsbu:   "l.msbu",
	CodeMsync:  "l.msync",
	CodeMtspr:  "l.mtspr",
	CodeMul:    "l.mul",
	CodeMuld:   "l.muld",
	CodeMuldu:  "l.muldu",
	CodeMuli:   "l.muli",
	CodeMulu:   "l.mulu",
	CodeNop:    "l.nop",
	CodeOr:     "l.or",
	CodeOri:    "l.ori",
	CodePsync:  "l.psync",
	CodeRfe:    "l.rfe",
	CodeRor:    "l.ror",
	CodeRori:   "l.rori",
	CodeSb:     "l.sb",
	CodeSd:     "l.sd",
	CodeSfeq:   "l.sfeq",
	CodeSfeqi:  "l.sfeqi",
	CodeSfges:  "l.sfges",
	CodeSfgesi: "l.sfgesi",
	CodeSfgeu:  "l.sfgeu",
	CodeSfgeui: "l.sfgeui",
	CodeSfgts:  "l.sfgts",
	CodeSfgtsi: "l.sfgtsi",
	CodeSfgtu:  "l.sfgtu",
	CodeSfgtui: "l.sfgtui",
	CodeSfles:  "l.sfles",
	CodeSflesi: "l.sflesi",
	CodeSfleu:  "l.sfleu",
	CodeSfleui: "l.sfleui",
	CodeSflts:  "l.sflts",
	CodeSfltsi: "l.sfltsi",
	CodeSfltu:  "l.sfltu",
	CodeSfltui: "l.sfltui",
	CodeSfne:   "l.sfne",
	CodeSfnei:  "l.sfnei",
	CodeSh:     "l.sh",
	CodeSll:    "l.sll",
	CodeSlli:   "l.slli",
	CodeSra:    "l.sra",
	CodeSrai:   "l.srai",
	CodeSrl:    "l.srl",
	CodeSrli:   "l.srli",
	CodeSub:    "l.sub",
	CodeSw:     "l.sw",
	CodeSys:    "l.sys",
	CodeTrap:   "l.trap",
	CodeXor:    "l.xor",
	CodeXori:   "l.xori",
}

func (code Code) String() string {
	s, ok := mnemonic[code]
	if ok {
		return s
	}
	return fmt.Sprintf("<invalid opcode: %d>", int(code))
}
