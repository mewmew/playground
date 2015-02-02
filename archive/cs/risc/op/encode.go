package op

import "fmt"

// Encode encodes the instruction and returns a 16-bit representation of it.
func Encode(inst interface{}) (buf uint16, err error) {
	switch v := inst.(type) {
	case *Nop:
		// op-code: 0
		// operand: 000

	case *LoadMem:
		// op-code: 1
		// operand: RXY
		//    R refers to the dst register.
		//    XY refers to the src memory address.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}

		// op-code
		buf |= 0x1000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src memory address.
		buf |= uint16(v.Src)

	case *LoadVal:
		// op-code: 2
		// operand: RXY
		//    R refers to the dst register.
		//    XY refers to the src immediate value.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}

		// op-code
		buf |= 0x2000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src immediate value.
		buf |= uint16(v.Src)

	case *Store:
		// op-code: 3
		// operand: RXY
		//    R refers to the src register.
		//    XY refers to the dst memory address.
		if v.Src >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src register %d in %s", v.Src, v.Code)
		}

		// op-code
		buf |= 0x3000

		// src register.
		buf |= uint16(v.Src) << 8

		// dst memory address.
		buf |= uint16(v.Dst)

	case *Move:
		// op-code: 4
		// operand: 0RS
		//    R refers to the src register.
		//    S refers to the dst register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src register %d in %s", v.Src, v.Code)
		}

		// op-code
		buf |= 0x4000

		// src register.
		buf |= uint16(v.Src) << 4

		// dst register.
		buf |= uint16(v.Dst)

	case *Add:
		// op-code: 5
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src1 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src1 register %d in %s", v.Src1, v.Code)
		}
		if v.Src2 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src2 register %d in %s", v.Src2, v.Code)
		}

		// op-code
		buf |= 0x5000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src1 register.
		buf |= uint16(v.Src1) << 4

		// src2 register.
		buf |= uint16(v.Src2)

	case *AddFloat:
		// op-code: 6
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src1 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src1 register %d in %s", v.Src1, v.Code)
		}
		if v.Src2 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src2 register %d in %s", v.Src2, v.Code)
		}

		// op-code
		buf |= 0x6000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src1 register.
		buf |= uint16(v.Src1) << 4

		// src2 register.
		buf |= uint16(v.Src2)

	case *Or:
		// op-code: 7
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src1 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src1 register %d in %s", v.Src1, v.Code)
		}
		if v.Src2 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src2 register %d in %s", v.Src2, v.Code)
		}

		// op-code
		buf |= 0x7000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src1 register.
		buf |= uint16(v.Src1) << 4

		// src2 register.
		buf |= uint16(v.Src2)

	case *And:
		// op-code: 8
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src1 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src1 register %d in %s", v.Src1, v.Code)
		}
		if v.Src2 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src2 register %d in %s", v.Src2, v.Code)
		}

		// op-code
		buf |= 0x8000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src1 register.
		buf |= uint16(v.Src1) << 4

		// src2 register.
		buf |= uint16(v.Src2)

	case *Xor:
		// op-code: 9
		// operand: RST
		//    R refers to the dst register.
		//    S refers to the src1 register.
		//    T refers to the src2 register.
		if v.Dst >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid dst register %d in %s", v.Dst, v.Code)
		}
		if v.Src1 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src1 register %d in %s", v.Src1, v.Code)
		}
		if v.Src2 >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid src2 register %d in %s", v.Src2, v.Code)
		}

		// op-code
		buf |= 0x9000

		// dst register.
		buf |= uint16(v.Dst) << 8

		// src1 register.
		buf |= uint16(v.Src1) << 4

		// src2 register.
		buf |= uint16(v.Src2)

	case *Ror:
		// op-code: A
		// operand: R0X
		//    R refers to the register.
		//    X refers to the immediate value x.
		if v.Reg >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid register %d in %s", v.Reg, v.Code)
		}
		if v.X >= RegSize {
			return 0, fmt.Errorf("op.Encode: invalid x (%d) in %s; above %d", v.X, v.Code, RegSize-1)
		}

		// op-code
		buf |= 0xA000

		// register.
		buf |= uint16(v.Reg) << 8

		// immediate value x.
		buf |= uint16(v.X)

	case *CmpBranch:
		// op-code: B
		// operand: RXY
		//    R refers to the cmp register.
		//    XY refers to the memory address addr.
		if v.Cmp >= RegCount {
			return 0, fmt.Errorf("op.Encode: invalid cmp register %d in %s", v.Cmp, v.Code)
		}

		// op-code
		buf |= 0xB000

		// cmp register.
		buf |= uint16(v.Cmp) << 8

		// memory address addr.
		buf |= uint16(v.Addr)

	case *Halt:
		// op-code: C
		// operand: 000

		// op-code
		buf |= 0xC000

	default:
		return 0, fmt.Errorf("op.Encode: unable to encode instruction (%T)", inst)
	}

	return buf, nil
}
