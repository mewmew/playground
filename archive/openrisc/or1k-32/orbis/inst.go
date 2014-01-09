package orbis

import (
	"fmt"
	"strconv"
)

// Reg represents a register between 0 and RegCount-1, normally r0 through r31.
type Reg uint8

func (reg Reg) String() string {
	return fmt.Sprintf("r%d", uint8(reg))
}

// Val represents an immediate value between 0 and 2^32-1.
type Val uint32

func (val Val) String() string {
	return strconv.Itoa(int(val))
}

// l.add rD,rA,rB
type Add struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Add) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.addc rD,rA,rB
type Addc struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Addc) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.addi rD,rA,I
type Addi struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Addi) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.addic rD,rA,I
type Addic struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Addic) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.and rD,rA,rB
type And struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *And) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.andi rD,rA,K
type Andi struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Andi) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.bf N
type Bf struct {
	Code Code
	Off  Val
}

func (inst *Bf) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Off)
}

// l.bnf N
type Bnf struct {
	Code Code
	Off  Val
}

func (inst *Bnf) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Off)
}

// l.cmov rD,rA,rB
type Cmov struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Cmov) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.csync
type Csync struct {
	Code Code
}

func (inst *Csync) String() string {
	return inst.Code.String()
}

// l.cust1
type Cust1 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust1) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust2
type Cust2 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust2) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust3
type Cust3 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust3) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust4
type Cust4 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust4) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust5 rD,rA,rB,L,K
type Cust5 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust5) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust6
type Cust6 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust6) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust7
type Cust7 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust7) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.cust8
type Cust8 struct {
	Code Code
	Buf  uint32
}

func (inst *Cust8) String() string {
	return fmt.Sprintf("%-16s0x%06X", inst.Code, inst.Buf)
}

// l.div rD,rA,rB
type Div struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Div) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.divu rD,rA,rB
type Divu struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Divu) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.extbs rD,rA
type Extbs struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Extbs) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.extbz rD,rA
type Extbz struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Extbz) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.exths rD,rA
type Exths struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Exths) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.exthz rD,rA
type Exthz struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Exthz) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.extws rD,rA
type Extws struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Extws) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.extwz rD,rA
type Extwz struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Extwz) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.ff1 rD,rA
type Ff1 struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Ff1) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.fl1 rD,rA
type Fl1 struct {
	Code Code
	Dst  Reg
	Src  Reg
}

func (inst *Fl1) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.j N
type J struct {
	Code Code
	Off  Val
}

func (inst *J) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Off)
}

// l.jal N
type Jal struct {
	Code Code
	Off  Val
}

func (inst *Jal) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Off)
}

// l.jalr rB
type Jalr struct {
	Code Code
	Addr Reg
}

func (inst *Jalr) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Addr)
}

// l.jr rB
type Jr struct {
	Code Code
	Addr Reg
}

func (inst *Jr) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Addr)
}

// l.lbs rD,I(rA)
type Lbs struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lbs) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.lbz rD,I(rA)
type Lbz struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lbz) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.ld rD,I(rA)
type Ld struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Ld) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.lhs rD,I(rA)
type Lhs struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lhs) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.lhz rD,I(rA)
type Lhz struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lhz) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.lws rD,I(rA)
type Lws struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lws) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.lwz rD,I(rA)
type Lwz struct {
	Code Code
	Dst  Reg
	Addr Reg
	Off  Val
}

func (inst *Lwz) String() string {
	return fmt.Sprintf("%-16s%s%s(%s)", inst.Code, inst.Dst, inst.Off, inst.Addr)
}

// l.mac rA,rB
type Mac struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Mac) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.maci rA,I
type Maci struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Maci) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.macrc rD
type Macrc struct {
	Code Code
	Dst  Reg
}

func (inst *Macrc) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Dst)
}

// l.macu rA,rB
type Macu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Macu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.mfspr rD,rA,K
type Mfspr struct {
	Code Code
	Dst  Reg
	Spr  Reg
	SprN Val
}

func (inst *Mfspr) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Spr, inst.SprN)
}

// l.movhi rD,K
type Movhi struct {
	Code Code
	Dst  Reg
	Src  Val
}

func (inst *Movhi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Dst, inst.Src)
}

// l.msb rA,rB
type Msb struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Msb) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.msbu rA,rB
type Msbu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Msbu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.msync
type Msync struct {
	Code Code
}

func (inst *Msync) String() string {
	return inst.Code.String()
}

// l.mtspr rA,rB,K
type Mtspr struct {
	Code Code
	Spr  Reg
	SprN Val
	Src  Reg
}

func (inst *Mtspr) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Spr, inst.Src, inst.SprN)
}

// l.mul rD,rA,rB
type Mul struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Mul) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.muld rA,rB
type Muld struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Muld) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.muldu rA,rB
type Muldu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Muldu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.muli rD,rA,I
type Muli struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Muli) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.mulu rD,rA,rB
type Mulu struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Mulu) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.nop K
type Nop struct {
	Code Code
	Val  Val
}

func (inst *Nop) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Val)
}

// l.or rD,rA,rB
type Or struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Or) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.ori rD,rA,K
type Ori struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Ori) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.psync
type Psync struct {
	Code Code
}

func (inst *Psync) String() string {
	return inst.Code.String()
}

// l.rfe
type Rfe struct {
	Code Code
}

func (inst *Rfe) String() string {
	return inst.Code.String()
}

// l.ror rD,rA,rB
type Ror struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Ror) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.rori rD,rA,L
type Rori struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Rori) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.sb I(rA),rB
type Sb struct {
	Code Code
	Addr Reg
	Off  Val
	Src  Reg
}

func (inst *Sb) String() string {
	return fmt.Sprintf("%-16s%s(%s), %s", inst.Code, inst.Off, inst.Addr, inst.Src)
}

// l.sd I(rA),rB
type Sd struct {
	Code Code
	Addr Reg
	Off  Val
	Src  Reg
}

func (inst *Sd) String() string {
	return fmt.Sprintf("%-16s%s(%s), %s", inst.Code, inst.Off, inst.Addr, inst.Src)
}

// l.sfeq rA,rB
type Sfeq struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfeq) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfeqi rA,I
type Sfeqi struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfeqi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfges rA,rB
type Sfges struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfges) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgesi rA,I
type Sfgesi struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfgesi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgeu rA,rB
type Sfgeu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfgeu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgeui rA,I
type Sfgeui struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfgeui) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgts rA,rB
type Sfgts struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfgts) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgtsi rA,I
type Sfgtsi struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfgtsi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgtu rA,rB
type Sfgtu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfgtu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfgtui rA,I
type Sfgtui struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfgtui) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfles rA,rB
type Sfles struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfles) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sflesi rA,I
type Sflesi struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sflesi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfleu rA,rB
type Sfleu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfleu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfleui rA,I
type Sfleui struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfleui) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sflts rA,rB
type Sflts struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sflts) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfltsi rA,I
type Sfltsi struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfltsi) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfltu rA,rB
type Sfltu struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfltu) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfltui rA,I
type Sfltui struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfltui) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfne rA,rB
type Sfne struct {
	Code Code
	Src1 Reg
	Src2 Reg
}

func (inst *Sfne) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sfnei rA,I
type Sfnei struct {
	Code Code
	Src1 Reg
	Src2 Val
}

func (inst *Sfnei) String() string {
	return fmt.Sprintf("%-16s%s, %s", inst.Code, inst.Src1, inst.Src2)
}

// l.sh I(rA),rB
type Sh struct {
	Code Code
	Addr Reg
	Off  Val
	Src  Reg
}

func (inst *Sh) String() string {
	return fmt.Sprintf("%-16s%s(%s), %s", inst.Code, inst.Off, inst.Addr, inst.Src)
}

// l.sll rD,rA,rB
type Sll struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Sll) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.slli rD,rA,L
type Slli struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Slli) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.sra rD,rA,rB
type Sra struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Sra) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.srai rD,rA,L
type Srai struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Srai) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.srl rD,rA,rB
type Srl struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Srl) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.srli rD,rA,L
type Srli struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Srli) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.sub rD,rA,rB
type Sub struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Sub) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.sw I(rA),rB
type Sw struct {
	Code Code
	Addr Reg
	Off  Val
	Src  Reg
}

func (inst *Sw) String() string {
	return fmt.Sprintf("%-16s%s(%s), %s", inst.Code, inst.Off, inst.Addr, inst.Src)
}

// l.sys K
type Sys struct {
	Code Code
	Val  Val
}

func (inst *Sys) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Val)
}

// l.trap K
type Trap struct {
	Code Code
	Val  Val
}

func (inst *Trap) String() string {
	return fmt.Sprintf("%-16s%s", inst.Code, inst.Val)
}

// l.xor rD,rA,rB
type Xor struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Reg
}

func (inst *Xor) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}

// l.xori rD,rA,I
type Xori struct {
	Code Code
	Dst  Reg
	Src1 Reg
	Src2 Val
}

func (inst *Xori) String() string {
	return fmt.Sprintf("%-16s%s, %s, %s", inst.Code, inst.Dst, inst.Src1, inst.Src2)
}
