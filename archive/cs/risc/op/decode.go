package op

import "fmt"

// Decode decodes the 16 bit representation of an instruction and returns it.
func Decode(buf uint16) (inst interface{}, err error) {
	code := Code(buf & 0xF000 >> 12)
	switch code {
	case CodeNop:
		// operand: 000

		// padding.
		pad := buf & 0x0FFF
		if pad != 0 {
			return nil, fmt.Errorf("op.Decode: invalid padding (%X) in %04X (%s).", pad, buf, code)
		}

		inst = &Nop{
			Code: code,
		}

	case CodeLoadMem:
		// operand: RXY
		//    R refers to the dst register.
		//    XY refers to the src memory address.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src memory address.
		src := Addr(buf & 0x0FF)

		inst = &LoadMem{
			Code: code,
			Src:  src,
			Dst:  dst,
		}

	case CodeLoadVal:
		// operand: RXY
		//    R refers to the dst register.
		//    XY refers to the src immediate value.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src immediate value.
		src := Val(buf & 0x0FF)

		inst = &LoadVal{
			Code: code,
			Src:  src,
			Dst:  dst,
		}

	case CodeStore:
		// operand: RXY
		//    R refers to the src register.
		//    XY refers to the dst memory address.

		// src register.
		src := Reg(buf & 0x0F00 >> 8)
		if src >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src register (%d) in %04X (%s).", src, buf, code)
		}

		// dst memory address.
		dst := Addr(buf & 0x0FF)

		inst = &Store{
			Code: code,
			Src:  src,
			Dst:  dst,
		}

	case CodeMove:
		// operand: 0RS
		//    R refers to the src register.
		//    S refers to the dst register.

		// padding.
		pad := buf & 0x0F00 >> 8
		if pad != 0 {
			return nil, fmt.Errorf("op.Decode: invalid padding (%X) in %04X (%s).", pad, buf, code)
		}

		// src register.
		src := Reg(buf & 0x00F0 >> 4)
		if src >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src register (%d) in %04X (%s).", src, buf, code)
		}

		// dst register.
		dst := Reg(buf & 0x000F)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		inst = &Move{
			Code: code,
			Src:  src,
			Dst:  dst,
		}

	case CodeAdd:
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src1 register.
		src1 := Reg(buf & 0x00F0 >> 4)
		if src1 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src1 register (%d) in %04X (%s).", src1, buf, code)
		}

		// src2 register.
		src2 := Reg(buf & 0x000F)
		if src2 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src2 register (%d) in %04X (%s).", src2, buf, code)
		}

		inst = &Add{
			Code: code,
			Src1: src1,
			Src2: src2,
			Dst:  dst,
		}

	case CodeAddFloat:
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src1 register.
		src1 := Reg(buf & 0x00F0 >> 4)
		if src1 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src1 register (%d) in %04X (%s).", src1, buf, code)
		}

		// src2 register.
		src2 := Reg(buf & 0x000F)
		if src2 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src2 register (%d) in %04X (%s).", src2, buf, code)
		}

		inst = &AddFloat{
			Code: code,
			Src1: src1,
			Src2: src2,
			Dst:  dst,
		}

	case CodeOr:
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src1 register.
		src1 := Reg(buf & 0x00F0 >> 4)
		if src1 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src1 register (%d) in %04X (%s).", src1, buf, code)
		}

		// src2 register.
		src2 := Reg(buf & 0x000F)
		if src2 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src2 register (%d) in %04X (%s).", src2, buf, code)
		}

		inst = &Or{
			Code: code,
			Src1: src1,
			Src2: src2,
			Dst:  dst,
		}

	case CodeAnd:
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src1 register.
		src1 := Reg(buf & 0x00F0 >> 4)
		if src1 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src1 register (%d) in %04X (%s).", src1, buf, code)
		}

		// src2 register.
		src2 := Reg(buf & 0x000F)
		if src2 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src2 register (%d) in %04X (%s).", src2, buf, code)
		}

		inst = &And{
			Code: code,
			Src1: src1,
			Src2: src2,
			Dst:  dst,
		}

	case CodeXor:
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.

		// dst register.
		dst := Reg(buf & 0x0F00 >> 8)
		if dst >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid dst register (%d) in %04X (%s).", dst, buf, code)
		}

		// src1 register.
		src1 := Reg(buf & 0x00F0 >> 4)
		if src1 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src1 register (%d) in %04X (%s).", src1, buf, code)
		}

		// src2 register.
		src2 := Reg(buf & 0x000F)
		if src2 >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid src2 register (%d) in %04X (%s).", src2, buf, code)
		}

		inst = &Xor{
			Code: code,
			Src1: src1,
			Src2: src2,
			Dst:  dst,
		}

	case CodeRor:
		// operand: R0X
		//    R refers to the register.
		//    X refers to the immediate value x.

		// register.
		reg := Reg(buf & 0x0F00 >> 8)
		if reg >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid register (%d) in %04X (%s).", reg, buf, code)
		}

		// padding.
		pad := buf & 0x00F0 >> 4
		if pad != 0 {
			return nil, fmt.Errorf("op.Decode: invalid padding (%X) in %04X (%s).", pad, buf, code)
		}

		// immediate value x.
		x := Val(buf & 0x000F)
		if x >= RegSize {
			return nil, fmt.Errorf("op.Decode: invalid x (%d) in %04X (%s); above %d.", x, buf, code, RegSize-1)
		}

		inst = &Ror{
			Code: code,
			Reg:  reg,
			X:    x,
		}

	case CodeCmpBranch:
		// operand: RXY
		//    R refers to the cmp register.
		//    XY refers to the memory address addr.

		// cmp register.
		cmp := Reg(buf & 0x0F00 >> 8)
		if cmp >= RegCount {
			return nil, fmt.Errorf("op.Decode: invalid cmp register (%d) in %04X (%s).", cmp, buf, code)
		}

		// memory address addr.
		addr := Addr(buf & 0x00FF)

		inst = &CmpBranch{
			Code: code,
			Addr: addr,
			Cmp:  cmp,
		}

	case CodeHalt:
		// operand: 000

		// padding.
		pad := buf & 0x0FFF
		if pad != 0 {
			return nil, fmt.Errorf("op.Decode: invalid padding (%X) in %04X (%s).", pad, buf, code)
		}

		inst = &Halt{
			Code: code,
		}

	default:
		return nil, fmt.Errorf("op.Decode: invalid code (%d) in %04X.", code, buf)
	}

	return inst, nil

}
