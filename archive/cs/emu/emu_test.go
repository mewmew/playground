package emu

import (
	"os"
	"reflect"
	"testing"
)

type testSystemRun struct {
	path string
	want *System
}

func TestSystemRun(t *testing.T) {
	golden := []testSystemRun{
		// i=0
		{
			path: "testdata/addhalt.bin",
			want: &System{
				PC: 0x08,
				Regs: [16]uint8{
					0: 0xC0,
					1: 0x5F,
					2: 0x61,
				},
				Mem: [256]uint8{
					0x00: 0x21,
					0x01: 0x5F,
					0x02: 0x22,
					0x03: 0x61,
					0x04: 0x50,
					0x05: 0x12,
					0x06: 0x30,
					0x07: 0x08,
					0x08: 0xC0, // Added during runtime: Mem[8] = 0x5F + 0x61 = 0xC0
				},
			},
		},
		// i=1
		{
			path: "testdata/2.3.1.bin",
			want: &System{
				PC: 0x04,
				Regs: [16]uint8{
					4: 0x34,
				},
				Mem: [256]uint8{
					0x00: 0x14,
					0x01: 0x02,
					0x02: 0x34,
					0x03: 0x17,
					0x04: 0xC0,
					0x17: 0x34,
				},
			},
		},
		// i=2
		{
			path: "testdata/2.3.2.bin",
			want: &System{
				PC: 0xB6,
				Regs: [16]uint8{
					3: 0xC3,
				},
				Mem: [256]uint8{
					0xB0: 0x13,
					0xB1: 0xB8,
					0xB2: 0xA3,
					0xB3: 0x02,
					0xB4: 0x33,
					0xB5: 0xB8,
					0xB6: 0xC0,
					0xB8: 0xC3,
				},
			},
		},
		// i=3
		{
			path: "testdata/2.3.3.bin",
			want: &System{
				PC: 0xB0,
				Regs: [16]uint8{
					0: 0x03,
					1: 0x03,
					2: 0x01,
				},
				Mem: [256]uint8{
					0xA4: 0x20,
					0xA6: 0x21,
					0xA7: 0x03,
					0xA8: 0x22,
					0xA9: 0x01,
					0xAA: 0xB1,
					0xAB: 0xB0,
					0xAC: 0x50,
					0xAD: 0x02,
					0xAE: 0xB0,
					0xAF: 0xAA,
					0xB0: 0xC0,
				},
			},
		},
		// i=4
		{
			path: "testdata/2.3.4.bin",
			want: &System{
				PC: 0xF8,
				Mem: [256]uint8{
					0xF0: 0x20,
					0xF1: 0xC0,
					0xF2: 0x30,
					0xF3: 0xF8,
					0xF4: 0x20,
					0xF6: 0x30,
					0xF7: 0xF9,
					0xF8: 0xC0,
				},
			},
		},
	}

	for i, g := range golden {
		f, err := os.Open(g.path)
		if err != nil {
			t.Fatalf("i=%d: %s", i, err)
		}
		defer f.Close()

		sys, err := New(f)
		if err != nil {
			t.Fatalf("i=%d: %s", i, err)
		}
		err = sys.Run()
		if err != nil {
			t.Fatalf("i=%d: %s", i, err)
		}
		if !reflect.DeepEqual(g.want, sys) {
			t.Errorf("i=%d: expected %#v, got %#v.", i, g.want, sys)
		}
	}
}
