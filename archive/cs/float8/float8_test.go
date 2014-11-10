package float8

import (
	"math"
	"testing"
)

type testFloat8 struct {
	f8  Float8
	f32 float32
}

var golden = []testFloat8{
	// i=0
	{
		f8:  0xA9, // 10101001
		f32: -0.140625,
	},
	// i=1
	{
		f8:  0x2A, // 00101010
		f32: 0.15625,
	},
	// i=2
	{
		f8:  0x7F, // 01111111
		f32: 7.5,
	},
	// i=3
	{
		f8:  0xFF, // 11111111
		f32: -7.5,
	},
	// i=4
	{
		f8:  0x48, // 01001000
		f32: 0.5,
	},
	// i=5
	{
		f8:  0x00, // 00000000
		f32: 0,
	},
	// i=6
	{
		f8:  0xB9, // 10111001
		f32: -0.28125,
	},
	// i=7
	{
		f8:  0x3A, // 00111010
		f32: 0.3125,
	},
	// i=8
	{
		f8:  0xFE, // 11111110
		f32: -7,
	},
	// i=9
	{
		f8:  0x58, // 01011000
		f32: 1,
	},
	// i=10
	{
		f8:  0x6A, // 01101010
		f32: 2.5,
	},
	// i=11
	{
		f8:  0xCC, // 11001100
		f32: -0.75,
	},
	// i=12
	{
		f8:  0x5E, // 01011110
		f32: 1.75,
	},
	// i=13
	{
		f8:  0xAC, // 10101100
		f32: -0.1875,
	},
}

func TestNew(t *testing.T) {
	for i, g := range golden {
		f8, err := New(g.f32)
		if err != nil {
			t.Errorf("i=%d: %s", i, err)
			continue
		}
		if f8 != g.f8 {
			t.Errorf("i=%d: expected %08b, got %08b.", i, g.f8, f8)
			continue
		}
	}
}

func TestFloatFloat32(t *testing.T) {
	for i, g := range golden {
		f32 := g.f8.Float32()
		if f32 != g.f32 {
			got := math.Float32bits(f32)
			want := math.Float32bits(g.f32)
			t.Errorf("i=%d: expected %032b, got %032b.", i, want, got)
			continue
		}
	}
}

type testAdd struct {
	x, y, want Float8
}

var goldenAdd = []testAdd{
	// i=0
	{
		x:    0x00, // 0
		y:    0x00, // 0
		want: 0x00, // 0
	},
	// i=1
	{
		x:    0xA9, // -0.140625
		y:    0x00, // 0
		want: 0xA9, // -0.140625
	},
	// i=2
	{
		x:    0xA9, // -0.140625
		y:    0xA9, // -0.140625
		want: 0xB9, // -0.28125
	},
	// i=3
	{
		x:    0x2A, // 0.15625
		y:    0x2A, // 0.15625
		want: 0x3A, // 0.3125
	},
	// i=4
	{
		x:    0x7F, // 7.5
		y:    0xFF, // -7.5
		want: 0x00, // 0
	},
	// i=5
	{
		x:    0xFF, // -7.5
		y:    0x48, // 0.5
		want: 0xFE, // -7
	},
	// i=6
	{
		x:    0x48, // 0.5
		y:    0x48, // 0.5
		want: 0x58, // 1
	},
	// i=7
	{
		x:    0x6A, // 2.5
		y:    0xCC, // -0.75
		want: 0x5E, // 1.75
	},
}

func TestAdd(t *testing.T) {
	for i, g := range goldenAdd {
		got, err := Add(g.x, g.y)
		if err != nil {
			t.Errorf("i=%d: %s", i, err)
			continue
		}
		if got != g.want {
			t.Errorf("i=%d: expected %08b, got %08b.", i, g.want, got)
			continue
		}
	}
}

func TestAddNative(t *testing.T) {
	for i, g := range goldenAdd {
		got, err := AddNative(g.x, g.y)
		if err != nil {
			t.Errorf("i=%d: %s", i, err)
			continue
		}
		if got != g.want {
			t.Errorf("i=%d: expected %08b, got %08b.", i, g.want, got)
			continue
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(0xFF, 0x48)
	}
}

func BenchmarkAddNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddNative(0xFF, 0x48)
	}
}
