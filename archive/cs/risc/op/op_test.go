package op

import (
	"reflect"
	"testing"
)

type testInst struct {
	buf  uint16
	inst interface{}
}

var golden = []testInst{
	// i=0
	{
		buf: 0x0000,
		inst: &Nop{
			Code: CodeNop,
		},
	},
	// i=1
	{
		buf: 0x14A3,
		inst: &LoadMem{
			Code: CodeLoadMem,
			Dst:  0x4,
			Src:  0xA3,
		},
	},
	// i=2
	{
		buf: 0x20A3,
		inst: &LoadVal{
			Code: CodeLoadVal,
			Dst:  0x0,
			Src:  0xA3,
		},
	},
	// i=3
	{
		buf: 0x35B1,
		inst: &Store{
			Code: CodeStore,
			Dst:  0xB1,
			Src:  0x5,
		},
	},
	// i=4
	{
		buf: 0x40A4,
		inst: &Move{
			Code: CodeMove,
			Dst:  0x4,
			Src:  0xA,
		},
	},
	// i=5
	{
		buf: 0x5726,
		inst: &Add{
			Code: CodeAdd,
			Dst:  0x7,
			Src1: 0x2,
			Src2: 0x6,
		},
	},
	// i=6
	{
		buf: 0x634E,
		inst: &AddFloat{
			Code: CodeAddFloat,
			Dst:  0x3,
			Src1: 0x4,
			Src2: 0xE,
		},
	},
	// i=7
	{
		buf: 0x7CB4,
		inst: &Or{
			Code: CodeOr,
			Dst:  0xC,
			Src1: 0xB,
			Src2: 0x4,
		},
	},
	// i=8
	{
		buf: 0x8045,
		inst: &And{
			Code: CodeAnd,
			Dst:  0x0,
			Src1: 0x4,
			Src2: 0x5,
		},
	},
	// i=9
	{
		buf: 0x95F3,
		inst: &Xor{
			Code: CodeXor,
			Dst:  0x5,
			Src1: 0xF,
			Src2: 0x3,
		},
	},
	// i=10
	{
		buf: 0xA403,
		inst: &Ror{
			Code: CodeRor,
			Reg:  0x4,
			X:    0x3,
		},
	},
	// i=11
	{
		buf: 0xB43C,
		inst: &CmpBranch{
			Code: CodeCmpBranch,
			Cmp:  0x4,
			Addr: 0x3C,
		},
	},
	// i=12
	{
		buf: 0xC000,
		inst: &Halt{
			Code: CodeHalt,
		},
	},
}

func TestDecode(t *testing.T) {
	for i, g := range golden {
		got, err := Decode(g.buf)
		if err != nil {
			t.Errorf("i=%d: %s", i, err)
			continue
		}
		if !reflect.DeepEqual(got, g.inst) {
			t.Errorf("i=%d: expected %#v, got %#v.", i, g.inst, got)
			continue
		}
	}
}

func TestEncode(t *testing.T) {
	for i, g := range golden {
		got, err := Encode(g.inst)
		if err != nil {
			t.Errorf("i=%d: %s", i, err)
			continue
		}
		if got != g.buf {
			t.Errorf("i=%d: expected 0x%02X, got 0x%02X.", i, g.buf, got)
			continue
		}
	}
}
