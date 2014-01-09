// Package or1k provides access to the 32-bit version of the Open RISC 1000
// instruction sets.
package or1k

import (
	"errors"

	"github.com/mewmew/playground/archive/openrisc/or1k-32/orbis"
)

// Decode decodes the 32 bit representation of an instruction and returns it.
func Decode(buf uint32) (inst interface{}, err error) {
	switch {
	// l.j
	case buf&0xFC000000 == 0x00000000:
		// 000000NNNNNNNNNNNNNNNNNNNNNNNNNN
		n := buf & 0x03FFFFFF
		inst = &orbis.J{
			Code: orbis.CodeJ,
			Off:  orbis.Val(n),
		}

	// l.jal
	case buf&0xFC000000 == 0x04000000:
		// 000001NNNNNNNNNNNNNNNNNNNNNNNNNN
		n := buf & 0x03FFFFFF
		inst = &orbis.Jal{
			Code: orbis.CodeJal,
			Off:  orbis.Val(n),
		}

	// l.bnf
	case buf&0xFC000000 == 0x0C000000:
		// 000011NNNNNNNNNNNNNNNNNNNNNNNNNN
		n := buf & 0x03FFFFFF
		inst = &orbis.Bnf{
			Code: orbis.CodeBnf,
			Off:  orbis.Val(n),
		}

	// l.bf
	case buf&0xFC000000 == 0x10000000:
		// 000100NNNNNNNNNNNNNNNNNNNNNNNNNN
		n := buf & 0x03FFFFFF
		inst = &orbis.Bf{
			Code: orbis.CodeBf,
			Off:  orbis.Val(n),
		}

	// l.nop
	case buf&0xFF000000 == 0x15000000:
		// 00010101--------KKKKKKKKKKKKKKKK
		if buf&0x00FF0000 != 0 {
			return nil, errors.New("invalid padding.")
		}
		k := buf & 0x0000FFFF
		inst = &orbis.Nop{
			Code: orbis.CodeNop,
			Val:  orbis.Val(k),
		}

	// l.movhi
	case buf&0xFC010000 == 0x18000000:
		// 000110DDDDD----0KKKKKKKKKKKKKKKK
		if buf&0x001E0000 != 0 {
			return nil, errors.New("invalid padding.")
		}
		d := buf & 0x03E00000 >> 21
		k := buf & 0x0000FFFF
		inst = &orbis.Movhi{
			Code: orbis.CodeMovhi,
			Dst:  orbis.Reg(d),
			Src:  orbis.Val(k),
		}

	// l.macrc
	case buf&0xFC01FFFF == 0x18010000:
		// 000110DDDDD----10000000000000000
		if buf&0x001E0000 != 0 {
			return nil, errors.New("invalid padding.")
		}
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Macrc{
			Code: orbis.CodeMacrc,
			Dst:  orbis.Reg(d),
		}

	// l.sys
	case buf&0xFFFF0000 == 0x20000000:
		// 0010000000000000KKKKKKKKKKKKKKKK
		k := buf & 0x0000FFFF
		inst = &orbis.Sys{
			Code: orbis.CodeSys,
			Val:  orbis.Val(k),
		}

	// l.trap
	case buf&0xFFFF0000 == 0x21000000:
		// 0010000100000000KKKKKKKKKKKKKKKK
		k := buf & 0x0000FFFF
		inst = &orbis.Trap{
			Code: orbis.CodeTrap,
			Val:  orbis.Val(k),
		}

	// l.msync
	case buf&0xFFFFFFFF == 0x22000000:
		// 00100010000000000000000000000000
		inst = &orbis.Msync{
			Code: orbis.CodeMsync,
		}

	// l.psync
	case buf&0xFFFFFFFF == 0x22800000:
		// 00100010100000000000000000000000
		inst = &orbis.Psync{
			Code: orbis.CodePsync,
		}

	// l.csync
	case buf&0xFFFFFFFF == 0x23000000:
		// 00100011000000000000000000000000
		inst = &orbis.Csync{
			Code: orbis.CodeCsync,
		}

	// l.rfe
	case buf&0xFC000000 == 0x24000000:
		// 001001--------------------------
		if buf&0x03FFFFFF != 0 {
			return nil, errors.New("invalid padding.")
		}
		inst = &orbis.Rfe{
			Code: orbis.CodeRfe,
		}

	/*
	   // lv.cust1
	   case buf&0xFC0000F0 == 0x280000C0:
	      // 001010------------------1100----
	      if buf&0x03FFFF0F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	*/

	/*
	   // lv.cust2
	   case buf&0xFC0000F0 == 0x280000D0:
	      // 001010------------------1101----
	      if buf&0x03FFFF0F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	*/

	/*
	   // lv.cust3
	   case buf&0xFC0000F0 == 0x280000E0:
	      // 001010------------------1110----
	      if buf&0x03FFFF0F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	*/

	/*
	   // lv.cust4
	   case buf&0xFC0000F0 == 0x280000F0:
	      // 001010------------------1111----
	      if buf&0x03FFFF0F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	*/

	/*
	   // lv.all_eq.b
	   case buf&0xFC0000FF == 0x28000010:
	      // 001010DDDDDAAAAABBBBB---00010000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_eq.h
	   case buf&0xFC0000FF == 0x28000011:
	      // 001010DDDDDAAAAABBBBB---00010001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_ge.b
	   case buf&0xFC0000FF == 0x28000012:
	      // 001010DDDDDAAAAABBBBB---00010010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_ge.h
	   case buf&0xFC0000FF == 0x28000013:
	      // 001010DDDDDAAAAABBBBB---00010011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_gt.b
	   case buf&0xFC0000FF == 0x28000014:
	      // 001010DDDDDAAAAABBBBB---00010100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_gt.h
	   case buf&0xFC0000FF == 0x28000015:
	      // 001010DDDDDAAAAABBBBB---00010101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_le.b
	   case buf&0xFC0000FF == 0x28000016:
	      // 001010DDDDDAAAAABBBBB---00010110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_le.h
	   case buf&0xFC0000FF == 0x28000017:
	      // 001010DDDDDAAAAABBBBB---00010111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_lt.b
	   case buf&0xFC0000FF == 0x28000018:
	      // 001010DDDDDAAAAABBBBB---00011000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_lt.h
	   case buf&0xFC0000FF == 0x28000019:
	      // 001010DDDDDAAAAABBBBB---00011001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_ne.b
	   case buf&0xFC0000FF == 0x2800001A:
	      // 001010DDDDDAAAAABBBBB---00011010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.all_ne.h
	   case buf&0xFC0000FF == 0x2800001B:
	      // 001010DDDDDAAAAABBBBB---00011011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_eq.b
	   case buf&0xFC0000FF == 0x28000020:
	      // 001010DDDDDAAAAABBBBB---00100000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_eq.h
	   case buf&0xFC0000FF == 0x28000021:
	      // 001010DDDDDAAAAABBBBB---00100001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_ge.b
	   case buf&0xFC0000FF == 0x28000022:
	      // 001010DDDDDAAAAABBBBB---00100010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_ge.h
	   case buf&0xFC0000FF == 0x28000023:
	      // 001010DDDDDAAAAABBBBB---00100011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_gt.b
	   case buf&0xFC0000FF == 0x28000024:
	      // 001010DDDDDAAAAABBBBB---00100100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_gt.h
	   case buf&0xFC0000FF == 0x28000025:
	      // 001010DDDDDAAAAABBBBB---00100101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_le.b
	   case buf&0xFC0000FF == 0x28000026:
	      // 001010DDDDDAAAAABBBBB---00100110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_le.h
	   case buf&0xFC0000FF == 0x28000027:
	      // 001010DDDDDAAAAABBBBB---00100111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_lt.b
	   case buf&0xFC0000FF == 0x28000028:
	      // 001010DDDDDAAAAABBBBB---00101000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_lt.h
	   case buf&0xFC0000FF == 0x28000029:
	      // 001010DDDDDAAAAABBBBB---00101001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_ne.b
	   case buf&0xFC0000FF == 0x2800002A:
	      // 001010DDDDDAAAAABBBBB---00101010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.any_ne.h
	   case buf&0xFC0000FF == 0x2800002B:
	      // 001010DDDDDAAAAABBBBB---00101011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.add.b
	   case buf&0xFC0000FF == 0x28000030:
	      // 001010DDDDDAAAAABBBBB---00110000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.add.h
	   case buf&0xFC0000FF == 0x28000031:
	      // 001010DDDDDAAAAABBBBB---00110001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.adds.b
	   case buf&0xFC0000FF == 0x28000032:
	      // 001010DDDDDAAAAABBBBB---00110010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.adds.h
	   case buf&0xFC0000FF == 0x28000033:
	      // 001010DDDDDAAAAABBBBB---00110011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.addu.b
	   case buf&0xFC0000FF == 0x28000034:
	      // 001010DDDDDAAAAABBBBB---00110100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.addu.h
	   case buf&0xFC0000FF == 0x28000035:
	      // 001010DDDDDAAAAABBBBB---00110101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.addus.b
	   case buf&0xFC0000FF == 0x28000036:
	      // 001010DDDDDAAAAABBBBB---00110110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.addus.h
	   case buf&0xFC0000FF == 0x28000037:
	      // 001010DDDDDAAAAABBBBB---00110111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.and
	   case buf&0xFC0000FF == 0x28000038:
	      // 001010DDDDDAAAAABBBBB---00111000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.avg.b
	   case buf&0xFC0000FF == 0x28000039:
	      // 001010DDDDDAAAAABBBBB---00111001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.avg.h
	   case buf&0xFC0000FF == 0x2800003A:
	      // 001010DDDDDAAAAABBBBB---00111010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_eq.b
	   case buf&0xFC0000FF == 0x28000040:
	      // 001010DDDDDAAAAABBBBB---01000000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_eq.h
	   case buf&0xFC0000FF == 0x28000041:
	      // 001010DDDDDAAAAABBBBB---01000001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_ge.b
	   case buf&0xFC0000FF == 0x28000042:
	      // 001010DDDDDAAAAABBBBB---01000010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_ge.h
	   case buf&0xFC0000FF == 0x28000043:
	      // 001010DDDDDAAAAABBBBB---01000011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_gt.b
	   case buf&0xFC0000FF == 0x28000044:
	      // 001010DDDDDAAAAABBBBB---01000100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_gt.h
	   case buf&0xFC0000FF == 0x28000045:
	      // 001010DDDDDAAAAABBBBB---01000101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_le.b
	   case buf&0xFC0000FF == 0x28000046:
	      // 001010DDDDDAAAAABBBBB---01000110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_le.h
	   case buf&0xFC0000FF == 0x28000047:
	      // 001010DDDDDAAAAABBBBB---01000111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_lt.b
	   case buf&0xFC0000FF == 0x28000048:
	      // 001010DDDDDAAAAABBBBB---01001000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_lt.h
	   case buf&0xFC0000FF == 0x28000049:
	      // 001010DDDDDAAAAABBBBB---01001001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_ne.b
	   case buf&0xFC0000FF == 0x2800004A:
	      // 001010DDDDDAAAAABBBBB---01001010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.cmp_ne.h
	   case buf&0xFC0000FF == 0x2800004B:
	      // 001010DDDDDAAAAABBBBB---01001011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.madds.h
	   case buf&0xFC0000FF == 0x28000054:
	      // 001010DDDDDAAAAABBBBB---01010100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.max.b
	   case buf&0xFC0000FF == 0x28000055:
	      // 001010DDDDDAAAAABBBBB---01010101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.max.h
	   case buf&0xFC0000FF == 0x28000056:
	      // 001010DDDDDAAAAABBBBB---01010110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.merge.b
	   case buf&0xFC0000FF == 0x28000057:
	      // 001010DDDDDAAAAABBBBB---01010111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.merge.h
	   case buf&0xFC0000FF == 0x28000058:
	      // 001010DDDDDAAAAABBBBB---01011000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.min.b
	   case buf&0xFC0000FF == 0x28000059:
	      // 001010DDDDDAAAAABBBBB---01011001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.min.h
	   case buf&0xFC0000FF == 0x2800005A:
	      // 001010DDDDDAAAAABBBBB---01011010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.msubs.h
	   case buf&0xFC0000FF == 0x2800005B:
	      // 001010DDDDDAAAAABBBBB---01011011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.muls.h
	   case buf&0xFC0000FF == 0x2800005C:
	      // 001010DDDDDAAAAABBBBB---01011100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.nand
	   case buf&0xFC0000FF == 0x2800005D:
	      // 001010DDDDDAAAAABBBBB---01011101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.nor
	   case buf&0xFC0000FF == 0x2800005E:
	      // 001010DDDDDAAAAABBBBB---01011110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.or
	   case buf&0xFC0000FF == 0x2800005F:
	      // 001010DDDDDAAAAABBBBB---01011111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.pack.b
	   case buf&0xFC0000FF == 0x28000060:
	      // 001010DDDDDAAAAABBBBB---01100000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.pack.h
	   case buf&0xFC0000FF == 0x28000061:
	      // 001010DDDDDAAAAABBBBB---01100001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.packs.b
	   case buf&0xFC0000FF == 0x28000062:
	      // 001010DDDDDAAAAABBBBB---01100010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.packs.h
	   case buf&0xFC0000FF == 0x28000063:
	      // 001010DDDDDAAAAABBBBB---01100011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.packus.b
	   case buf&0xFC0000FF == 0x28000064:
	      // 001010DDDDDAAAAABBBBB---01100100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.packus.h
	   case buf&0xFC0000FF == 0x28000065:
	      // 001010DDDDDAAAAABBBBB---01100101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.perm.n
	   case buf&0xFC0000FF == 0x28000066:
	      // 001010DDDDDAAAAABBBBB---01100110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.rl.b
	   case buf&0xFC0000FF == 0x28000067:
	      // 001010DDDDDAAAAABBBBB---01100111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.rl.h
	   case buf&0xFC0000FF == 0x28000068:
	      // 001010DDDDDAAAAABBBBB---01101000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sll.b
	   case buf&0xFC0000FF == 0x28000069:
	      // 001010DDDDDAAAAABBBBB---01101001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sll.h
	   case buf&0xFC0000FF == 0x2800006A:
	      // 001010DDDDDAAAAABBBBB---01101010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sll
	   case buf&0xFC0000FF == 0x2800006B:
	      // 001010DDDDDAAAAABBBBB---01101011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.srl.b
	   case buf&0xFC0000FF == 0x2800006C:
	      // 001010DDDDDAAAAABBBBB---01101100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.srl.h
	   case buf&0xFC0000FF == 0x2800006D:
	      // 001010DDDDDAAAAABBBBB---01101101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sra.b
	   case buf&0xFC0000FF == 0x2800006E:
	      // 001010DDDDDAAAAABBBBB---01101110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sra.h
	   case buf&0xFC0000FF == 0x2800006F:
	      // 001010DDDDDAAAAABBBBB---01101111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.srl
	   case buf&0xFC0000FF == 0x28000070:
	      // 001010DDDDDAAAAABBBBB---01110000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sub.b
	   case buf&0xFC0000FF == 0x28000071:
	      // 001010DDDDDAAAAABBBBB---01110001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.sub.h
	   case buf&0xFC0000FF == 0x28000072:
	      // 001010DDDDDAAAAABBBBB---01110010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subs.b
	   case buf&0xFC0000FF == 0x28000073:
	      // 001010DDDDDAAAAABBBBB---01110011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subs.h
	   case buf&0xFC0000FF == 0x28000074:
	      // 001010DDDDDAAAAABBBBB---01110100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subu.b
	   case buf&0xFC0000FF == 0x28000075:
	      // 001010DDDDDAAAAABBBBB---01110101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subu.h
	   case buf&0xFC0000FF == 0x28000076:
	      // 001010DDDDDAAAAABBBBB---01110110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subus.b
	   case buf&0xFC0000FF == 0x28000077:
	      // 001010DDDDDAAAAABBBBB---01110111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.subus.h
	   case buf&0xFC0000FF == 0x28000078:
	      // 001010DDDDDAAAAABBBBB---01111000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.unpack.b
	   case buf&0xFC0000FF == 0x28000079:
	      // 001010DDDDDAAAAABBBBB---01111001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.unpack.h
	   case buf&0xFC0000FF == 0x2800007A:
	      // 001010DDDDDAAAAABBBBB---01111010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lv.xor
	   case buf&0xFC0000FF == 0x2800007B:
	      // 001010DDDDDAAAAABBBBB---01111011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	// l.jr
	case buf&0xFC000000 == 0x44000000:
		// 010001----------BBBBB-----------
		if buf&0x03FF07FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Jr{
			Code: orbis.CodeJr,
			Addr: orbis.Reg(b),
		}

	// l.jalr
	case buf&0xFC000000 == 0x48000000:
		// 010010----------BBBBB-----------
		if buf&0x03FF07FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Jalr{
			Code: orbis.CodeJalr,
			Addr: orbis.Reg(b),
		}

	// l.maci
	case buf&0xFC000000 == 0x4C000000:
		// 010011-----AAAAAIIIIIIIIIIIIIIII
		if buf&0x03E00000 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Maci{
			Code: orbis.CodeMaci,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.cust1
	case buf&0xFC000000 == 0x70000000:
		// 011100--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust1{
			Code: orbis.CodeCust1,
			Buf:  v,
		}

	// l.cust2
	case buf&0xFC000000 == 0x74000000:
		// 011101--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust2{
			Code: orbis.CodeCust2,
			Buf:  v,
		}

	// l.cust3
	case buf&0xFC000000 == 0x78000000:
		// 011110--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust3{
			Code: orbis.CodeCust3,
			Buf:  v,
		}

	// l.cust4
	case buf&0xFC000000 == 0x7C000000:
		// 011111--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust4{
			Code: orbis.CodeCust4,
			Buf:  v,
		}

	// l.ld
	case buf&0xFC000000 == 0x80000000:
		// 100000DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Ld{
			Code: orbis.CodeLd,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lwz
	case buf&0xFC000000 == 0x84000000:
		// 100001DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lwz{
			Code: orbis.CodeLwz,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lws
	case buf&0xFC000000 == 0x88000000:
		// 100010DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lws{
			Code: orbis.CodeLws,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lbz
	case buf&0xFC000000 == 0x8C000000:
		// 100011DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lbz{
			Code: orbis.CodeLbz,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lbs
	case buf&0xFC000000 == 0x90000000:
		// 100100DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lbs{
			Code: orbis.CodeLbs,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lhz
	case buf&0xFC000000 == 0x94000000:
		// 100101DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lhz{
			Code: orbis.CodeLhz,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.lhs
	case buf&0xFC000000 == 0x98000000:
		// 100110DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Lhs{
			Code: orbis.CodeLhs,
			Dst:  orbis.Reg(d),
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
		}

	// l.addi
	case buf&0xFC000000 == 0x9C000000:
		// 100111DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Addi{
			Code: orbis.CodeAddi,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.addic
	case buf&0xFC000000 == 0xA0000000:
		// 101000DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Addic{
			Code: orbis.CodeAddic,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.andi
	case buf&0xFC000000 == 0xA4000000:
		// 101001DDDDDAAAAAKKKKKKKKKKKKKKKK
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		k := buf & 0x0000FFFF
		inst = &orbis.Andi{
			Code: orbis.CodeAndi,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(k),
		}

	// l.ori
	case buf&0xFC000000 == 0xA8000000:
		// 101010DDDDDAAAAAKKKKKKKKKKKKKKKK
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		k := buf & 0x0000FFFF
		inst = &orbis.Ori{
			Code: orbis.CodeOri,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(k),
		}

	// l.xori
	case buf&0xFC000000 == 0xAC000000:
		// 101011DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Xori{
			Code: orbis.CodeXori,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.muli
	case buf&0xFC000000 == 0xB0000000:
		// 101100DDDDDAAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		i := buf & 0x0000FFFF
		inst = &orbis.Muli{
			Code: orbis.CodeMuli,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.mfspr
	case buf&0xFC000000 == 0xB4000000:
		// 101101DDDDDAAAAAKKKKKKKKKKKKKKKK
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		k := buf & 0x0000FFFF
		inst = &orbis.Mfspr{
			Code: orbis.CodeMfspr,
			Dst:  orbis.Reg(d),
			Spr:  orbis.Reg(a),
			SprN: orbis.Val(k),
		}

	// l.slli
	case buf&0xFC0000C0 == 0xB8000000:
		// 101110DDDDDAAAAA--------00LLLLLL
		if buf&0x0000FF00 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		l := buf & 0x0000003F
		inst = &orbis.Slli{
			Code: orbis.CodeSlli,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(l),
		}

	// l.srli
	case buf&0xFC0000C0 == 0xB8000040:
		// 101110DDDDDAAAAA--------01LLLLLL
		if buf&0x0000FF00 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		l := buf & 0x0000003F
		inst = &orbis.Srli{
			Code: orbis.CodeSrli,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(l),
		}

	// l.srai
	case buf&0xFC0000C0 == 0xB8000080:
		// 101110DDDDDAAAAA--------10LLLLLL
		if buf&0x0000FF00 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		l := buf & 0x0000003F
		inst = &orbis.Srai{
			Code: orbis.CodeSrai,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(l),
		}

	// l.rori
	case buf&0xFC0000C0 == 0xB80000C0:
		// 101110DDDDDAAAAA--------11LLLLLL
		if buf&0x0000FF00 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		l := buf & 0x0000003F
		inst = &orbis.Rori{
			Code: orbis.CodeRori,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Val(l),
		}

	// l.sfeqi
	case buf&0xFFE00000 == 0xBC000000:
		// 10111100000AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfeqi{
			Code: orbis.CodeSfeqi,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfnei
	case buf&0xFFE00000 == 0xBC200000:
		// 10111100001AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfnei{
			Code: orbis.CodeSfnei,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfgtui
	case buf&0xFFE00000 == 0xBC400000:
		// 10111100010AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfgtui{
			Code: orbis.CodeSfgtui,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfgeui
	case buf&0xFFE00000 == 0xBC600000:
		// 10111100011AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfgeui{
			Code: orbis.CodeSfgeui,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfltui
	case buf&0xFFE00000 == 0xBC800000:
		// 10111100100AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfltui{
			Code: orbis.CodeSfltui,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfleui
	case buf&0xFFE00000 == 0xBCA00000:
		// 10111100101AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfleui{
			Code: orbis.CodeSfleui,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfgtsi
	case buf&0xFFE00000 == 0xBD400000:
		// 10111101010AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfgtsi{
			Code: orbis.CodeSfgtsi,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfgesi
	case buf&0xFFE00000 == 0xBD600000:
		// 10111101011AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfgesi{
			Code: orbis.CodeSfgesi,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sfltsi
	case buf&0xFFE00000 == 0xBD800000:
		// 10111101100AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sfltsi{
			Code: orbis.CodeSfltsi,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.sflesi
	case buf&0xFFE00000 == 0xBDA00000:
		// 10111101101AAAAAIIIIIIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		i := buf & 0x0000FFFF
		inst = &orbis.Sflesi{
			Code: orbis.CodeSflesi,
			Src1: orbis.Reg(a),
			Src2: orbis.Val(i),
		}

	// l.mtspr
	case buf&0xFC000000 == 0xC0000000:
		// 110000KKKKKAAAAABBBBBKKKKKKKKKKK
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		k := buf & 0x03E00000 >> 11
		k |= buf & 0x000007FF
		inst = &orbis.Mtspr{
			Code: orbis.CodeMtspr,
			Spr:  orbis.Reg(a),
			SprN: orbis.Val(k),
			Src:  orbis.Reg(b),
		}

	// l.mac
	case buf&0xFC00000F == 0xC4000001:
		// 110001-----AAAAABBBBB-------0001
		if buf&0x03E007F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Mac{
			Code: orbis.CodeMac,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.macu
	case buf&0xFC00000F == 0xC4000003:
		// 110001-----AAAAABBBBB-------0011
		if buf&0x03E007F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Macu{
			Code: orbis.CodeMacu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.msb
	case buf&0xFC00000F == 0xC4000002:
		// 110001-----AAAAABBBBB-------0010
		if buf&0x03E007F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Msb{
			Code: orbis.CodeMsb,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.msbu
	case buf&0xFC00000F == 0xC4000004:
		// 110001-----AAAAABBBBB-------0100
		if buf&0x03E007F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Msbu{
			Code: orbis.CodeMsbu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	/*
	   // lf.sfeq.s
	   case buf&0xFC0000FF == 0xC8000008:
	      // 110010-----AAAAABBBBB---00001000
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfne.s
	   case buf&0xFC0000FF == 0xC8000009:
	      // 110010-----AAAAABBBBB---00001001
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfgt.s
	   case buf&0xFC0000FF == 0xC800000A:
	      // 110010-----AAAAABBBBB---00001010
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfge.s
	   case buf&0xFC0000FF == 0xC800000B:
	      // 110010-----AAAAABBBBB---00001011
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sflt.s
	   case buf&0xFC0000FF == 0xC800000C:
	      // 110010-----AAAAABBBBB---00001100
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfle.s
	   case buf&0xFC0000FF == 0xC800000D:
	      // 110010-----AAAAABBBBB---00001101
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfeq.d
	   case buf&0xFC0000FF == 0xC8000018:
	      // 110010-----AAAAABBBBB---00011000
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfne.d
	   case buf&0xFC0000FF == 0xC8000019:
	      // 110010-----AAAAABBBBB---00011001
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfgt.d
	   case buf&0xFC0000FF == 0xC800001A:
	      // 110010-----AAAAABBBBB---00011010
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfge.d
	   case buf&0xFC0000FF == 0xC800001B:
	      // 110010-----AAAAABBBBB---00011011
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sflt.d
	   case buf&0xFC0000FF == 0xC800001C:
	      // 110010-----AAAAABBBBB---00011100
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.sfle.d
	   case buf&0xFC0000FF == 0xC800001D:
	      // 110010-----AAAAABBBBB---00011101
	      if buf&0x03E00700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.cust1.s
	   case buf&0xFC0000F0 == 0xC80000D0:
	      // 110010-----AAAAABBBBB---1101----
	      if buf&0x03E0070F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.cust1.d
	   case buf&0xFC0000F0 == 0xC80000E0:
	      // 110010-----AAAAABBBBB---1110----
	      if buf&0x03E0070F != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	*/

	/*
	   // lf.itof.s
	   case buf&0xFC00F8FF == 0xC8000004:
	      // 110010DDDDDAAAAA00000---00000100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.ftoi.s
	   case buf&0xFC00F8FF == 0xC8000005:
	      // 110010DDDDDAAAAA00000---00000101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.itof.d
	   case buf&0xFC00F8FF == 0xC8000014:
	      // 110010DDDDDAAAAA00000---00010100
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.ftoi.d
	   case buf&0xFC00F8FF == 0xC8000015:
	      // 110010DDDDDAAAAA00000---00010101
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.add.s
	   case buf&0xFC0000FF == 0xC8000000:
	      // 110010DDDDDAAAAABBBBB---00000000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.sub.s
	   case buf&0xFC0000FF == 0xC8000001:
	      // 110010DDDDDAAAAABBBBB---00000001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.mul.s
	   case buf&0xFC0000FF == 0xC8000002:
	      // 110010DDDDDAAAAABBBBB---00000010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.div.s
	   case buf&0xFC0000FF == 0xC8000003:
	      // 110010DDDDDAAAAABBBBB---00000011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.rem.s
	   case buf&0xFC0000FF == 0xC8000006:
	      // 110010DDDDDAAAAABBBBB---00000110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.madd.s
	   case buf&0xFC0000FF == 0xC8000007:
	      // 110010DDDDDAAAAABBBBB---00000111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.add.d
	   case buf&0xFC0000FF == 0xC8000010:
	      // 110010DDDDDAAAAABBBBB---00010000
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.sub.d
	   case buf&0xFC0000FF == 0xC8000011:
	      // 110010DDDDDAAAAABBBBB---00010001
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.mul.d
	   case buf&0xFC0000FF == 0xC8000012:
	      // 110010DDDDDAAAAABBBBB---00010010
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.div.d
	   case buf&0xFC0000FF == 0xC8000013:
	      // 110010DDDDDAAAAABBBBB---00010011
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.rem.d
	   case buf&0xFC0000FF == 0xC8000016:
	      // 110010DDDDDAAAAABBBBB---00010110
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	/*
	   // lf.madd.d
	   case buf&0xFC0000FF == 0xC8000017:
	      // 110010DDDDDAAAAABBBBB---00010111
	      if buf&0x00000700 != 0 {
	         return nil, errors.New("invalid padding.")
	      }
	      a := buf&0x001F0000 >> 16
	      b := buf&0x0000F800 >> 11
	      d := buf&0x03E00000 >> 21
	*/

	// l.sd
	case buf&0xFC000000 == 0xD0000000:
		// 110100IIIIIAAAAABBBBBIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		i := buf & 0x03E00000 >> 11
		i |= buf & 0x000007FF
		inst = &orbis.Sd{
			Code: orbis.CodeSd,
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
			Src:  orbis.Reg(b),
		}

	// l.sw
	case buf&0xFC000000 == 0xD4000000:
		// 110101IIIIIAAAAABBBBBIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		i := buf & 0x03E00000 >> 11
		i |= buf & 0x000007FF
		inst = &orbis.Sw{
			Code: orbis.CodeSw,
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
			Src:  orbis.Reg(b),
		}

	// l.sb
	case buf&0xFC000000 == 0xD8000000:
		// 110110IIIIIAAAAABBBBBIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		i := buf & 0x03E00000 >> 11
		i |= buf & 0x000007FF
		inst = &orbis.Sb{
			Code: orbis.CodeSb,
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
			Src:  orbis.Reg(b),
		}

	// l.sh
	case buf&0xFC000000 == 0xDC000000:
		// 110111IIIIIAAAAABBBBBIIIIIIIIIII
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		i := buf & 0x03E00000 >> 11
		i |= buf & 0x000007FF
		inst = &orbis.Sh{
			Code: orbis.CodeSh,
			Addr: orbis.Reg(a),
			Off:  orbis.Val(i),
			Src:  orbis.Reg(b),
		}

	// l.exths
	case buf&0xFC0003CF == 0xE000000C:
		// 111000DDDDDAAAAA------0000--1100
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Exths{
			Code: orbis.CodeExths,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.extws
	case buf&0xFC0003CF == 0xE000000D:
		// 111000DDDDDAAAAA------0000--1101
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Extws{
			Code: orbis.CodeExtws,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.extbs
	case buf&0xFC0003CF == 0xE000004C:
		// 111000DDDDDAAAAA------0001--1100
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Extbs{
			Code: orbis.CodeExtbs,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.extwz
	case buf&0xFC0003CF == 0xE000004D:
		// 111000DDDDDAAAAA------0001--1101
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Extwz{
			Code: orbis.CodeExtwz,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.exthz
	case buf&0xFC0003CF == 0xE000008C:
		// 111000DDDDDAAAAA------0010--1100
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Exthz{
			Code: orbis.CodeExthz,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.extbz
	case buf&0xFC0003CF == 0xE00000CC:
		// 111000DDDDDAAAAA------0011--1100
		if buf&0x0000FC30 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Extbz{
			Code: orbis.CodeExtbz,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.add
	case buf&0xFC00030F == 0xE0000000:
		// 111000DDDDDAAAAABBBBB-00----0000
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Add{
			Code: orbis.CodeAdd,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.addc
	case buf&0xFC00030F == 0xE0000001:
		// 111000DDDDDAAAAABBBBB-00----0001
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Addc{
			Code: orbis.CodeAddc,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sub
	case buf&0xFC00030F == 0xE0000002:
		// 111000DDDDDAAAAABBBBB-00----0010
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Sub{
			Code: orbis.CodeSub,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.and
	case buf&0xFC00030F == 0xE0000003:
		// 111000DDDDDAAAAABBBBB-00----0011
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.And{
			Code: orbis.CodeAnd,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.or
	case buf&0xFC00030F == 0xE0000004:
		// 111000DDDDDAAAAABBBBB-00----0100
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Or{
			Code: orbis.CodeOr,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.xor
	case buf&0xFC00030F == 0xE0000005:
		// 111000DDDDDAAAAABBBBB-00----0101
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Xor{
			Code: orbis.CodeXor,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.cmov
	case buf&0xFC00030F == 0xE000000E:
		// 111000DDDDDAAAAABBBBB-00----1110
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Cmov{
			Code: orbis.CodeCmov,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.ff1
	case buf&0xFC00030F == 0xE000000F:
		// 111000DDDDDAAAAA------00----1111
		if buf&0x0000FCF0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Ff1{
			Code: orbis.CodeFf1,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.sll
	case buf&0xFC0003CF == 0xE0000008:
		// 111000DDDDDAAAAABBBBB-0000--1000
		if buf&0x00000430 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Sll{
			Code: orbis.CodeSll,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.srl
	case buf&0xFC0003CF == 0xE0000048:
		// 111000DDDDDAAAAABBBBB-0001--1000
		if buf&0x00000430 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Srl{
			Code: orbis.CodeSrl,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sra
	case buf&0xFC0003CF == 0xE0000088:
		// 111000DDDDDAAAAABBBBB-0010--1000
		if buf&0x00000430 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Sra{
			Code: orbis.CodeSra,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.ror
	case buf&0xFC0003CF == 0xE00000C8:
		// 111000DDDDDAAAAABBBBB-0011--1000
		if buf&0x00000430 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Ror{
			Code: orbis.CodeRor,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.fl1
	case buf&0xFC00030F == 0xE000010F:
		// 111000DDDDDAAAAA------01----1111
		if buf&0x0000FCF0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Fl1{
			Code: orbis.CodeFl1,
			Dst:  orbis.Reg(d),
			Src:  orbis.Reg(a),
		}

	// l.mul
	case buf&0xFC00030F == 0xE0000306:
		// 111000DDDDDAAAAABBBBB-11----0110
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Mul{
			Code: orbis.CodeMul,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.muld
	case buf&0xFC00030F == 0xE0000307:
		// 111000-----AAAAABBBBB-11----0111
		if buf&0x03E004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Muld{
			Code: orbis.CodeMuld,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.div
	case buf&0xFC00030F == 0xE0000309:
		// 111000DDDDDAAAAABBBBB-11----1001
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Div{
			Code: orbis.CodeDiv,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.divu
	case buf&0xFC00030F == 0xE000030A:
		// 111000DDDDDAAAAABBBBB-11----1010
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Divu{
			Code: orbis.CodeDivu,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.mulu
	case buf&0xFC00030F == 0xE000030B:
		// 111000DDDDDAAAAABBBBB-11----1011
		if buf&0x000004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		d := buf & 0x03E00000 >> 21
		inst = &orbis.Mulu{
			Code: orbis.CodeMulu,
			Dst:  orbis.Reg(d),
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.muldu
	case buf&0xFC00030F == 0xE000030C:
		// 111000-----AAAAABBBBB-11----1100
		if buf&0x03E004F0 != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Muldu{
			Code: orbis.CodeMuldu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfeq
	case buf&0xFFE00000 == 0xE4000000:
		// 11100100000AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfeq{
			Code: orbis.CodeSfeq,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfne
	case buf&0xFFE00000 == 0xE4200000:
		// 11100100001AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfne{
			Code: orbis.CodeSfne,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfgtu
	case buf&0xFFE00000 == 0xE4400000:
		// 11100100010AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfgtu{
			Code: orbis.CodeSfgtu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfgeu
	case buf&0xFFE00000 == 0xE4600000:
		// 11100100011AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfgeu{
			Code: orbis.CodeSfgeu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfltu
	case buf&0xFFE00000 == 0xE4800000:
		// 11100100100AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfltu{
			Code: orbis.CodeSfltu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfleu
	case buf&0xFFE00000 == 0xE4A00000:
		// 11100100101AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfleu{
			Code: orbis.CodeSfleu,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfgts
	case buf&0xFFE00000 == 0xE5400000:
		// 11100101010AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfgts{
			Code: orbis.CodeSfgts,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfges
	case buf&0xFFE00000 == 0xE5600000:
		// 11100101011AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfges{
			Code: orbis.CodeSfges,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sflts
	case buf&0xFFE00000 == 0xE5800000:
		// 11100101100AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sflts{
			Code: orbis.CodeSflts,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.sfles
	case buf&0xFFE00000 == 0xE5A00000:
		// 11100101101AAAAABBBBB-----------
		if buf&0x000007FF != 0 {
			return nil, errors.New("invalid padding.")
		}
		a := buf & 0x001F0000 >> 16
		b := buf & 0x0000F800 >> 11
		inst = &orbis.Sfles{
			Code: orbis.CodeSfles,
			Src1: orbis.Reg(a),
			Src2: orbis.Reg(b),
		}

	// l.cust5
	case buf&0xFC000000 == 0xF0000000:
		// 111100--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust5{
			Code: orbis.CodeCust5,
			Buf:  v,
		}

	// l.cust6
	case buf&0xFC000000 == 0xF4000000:
		// 111101--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust6{
			Code: orbis.CodeCust6,
			Buf:  v,
		}

	// l.cust7
	case buf&0xFC000000 == 0xF8000000:
		// 111110--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust7{
			Code: orbis.CodeCust7,
			Buf:  v,
		}

	// l.cust8
	case buf&0xFC000000 == 0xFC000000:
		// 111111--------------------------
		v := buf & 0x03FFFFFF
		inst = &orbis.Cust8{
			Code: orbis.CodeCust8,
			Buf:  v,
		}
	}

	return inst, nil
}
