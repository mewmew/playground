package algo

import "testing"

var golden = []struct {
	list []int
	n    int
	want bool
}{
	// i=0
	{
		list: []int{3, 2, 0, 5},
		n:    1,
		want: false,
	},
	// i=1
	{
		list: []int{3, 2, 0, 5},
		n:    0,
		want: true,
	},
	// i=2
	{
		list: []int{981829, 1281, 3771, 1, 3020, 29, 38, 95},
		n:    0,
		want: false,
	},
	// i=3
	{
		list: []int{1, 2, 3},
		n:    3,
		want: true,
	},
	// i=4
	{
		list: []int{3, 2, 1},
		n:    3,
		want: true,
	},
	// i=5
	{
		list: []int{},
		n:    9,
		want: false,
	},
}

func TestSeqListContains(t *testing.T) {
	for i, g := range golden {
		list := NewSeqList(g.list)
		got := list.Contains(g.n)
		if got != g.want {
			t.Errorf("i=%d: expected %v, got %v.", i, g.want, got)
		}
	}
}

// === [ SeqListContains benchmark ] ===========================================

func BenchmarkSeqListContains128(b *testing.B) {
	benchmarkSeqListContains(b, 128)
}

func BenchmarkSeqListContains256(b *testing.B) {
	benchmarkSeqListContains(b, 256)
}

func BenchmarkSeqListContains512(b *testing.B) {
	benchmarkSeqListContains(b, 512)
}

func BenchmarkSeqListContains1k(b *testing.B) {
	benchmarkSeqListContains(b, 1024)
}

func BenchmarkSeqListContains2k(b *testing.B) {
	benchmarkSeqListContains(b, 2*1024)
}

func BenchmarkSeqListContains4k(b *testing.B) {
	benchmarkSeqListContains(b, 4*1024)
}

func benchmarkSeqListContains(b *testing.B, size int) {
	list := make([]int, size)
	for i := 0; i < cap(list); i++ {
		list[i] = i
	}
	l := NewSeqList(list)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Contains(size - 1)
	}
}
