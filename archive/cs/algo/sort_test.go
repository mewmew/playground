package algo

import (
	"math/rand"
	"reflect"
	"testing"
)

var goldenSort = []struct {
	in   []int
	want []int
}{
	// i = 0
	{
		in:   []int{6, 1, 4, 2, 3},
		want: []int{1, 2, 3, 4, 6},
	},
	// i = 1
	{
		in:   []int{4, 2, 5, 7, 9, 9},
		want: []int{2, 4, 5, 7, 9, 9},
	},
	// i = 2
	{
		in:   []int{6, 5, 4, 3, 2, 1},
		want: []int{1, 2, 3, 4, 5, 6},
	},
	// i = 3
	{
		in:   []int{6, 6, 5, 5, 8, 3},
		want: []int{3, 5, 5, 6, 6, 8},
	},
}

func TestInsSort(t *testing.T) {
	testSort(t, InsSort)
}

func TestSelSort(t *testing.T) {
	testSort(t, SelSort)
}

func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort)
}

func TestQuickSort(t *testing.T) {
	testSort(t, QuickSort)
}

func testSort(t *testing.T, sortFn func([]int) []int) {
	for i, g := range goldenSort {
		in := make([]int, len(g.in))
		copy(in, g.in)
		got := sortFn(in)
		if !reflect.DeepEqual(got, g.want) {
			t.Errorf("i=%d: expected %v, got %v.", i, g.want, got)
		}
	}
}

// === [ InsSort benchmark ] ===================================================

// --- [ random order ] --------------------------------------------------------

func BenchmarkInsSortRand128(b *testing.B) {
	l := genListRand(128)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortRand256(b *testing.B) {
	l := genListRand(256)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortRand512(b *testing.B) {
	l := genListRand(512)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortRand1k(b *testing.B) {
	l := genListRand(1024)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortRand2k(b *testing.B) {
	l := genListRand(2 * 1024)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortRand4k(b *testing.B) {
	l := genListRand(4 * 1024)
	benchmarkSort(b, InsSort, l)
}

// --- [ worst case ] ----------------------------------------------------------

func BenchmarkInsSortWorst128(b *testing.B) {
	l := genListDesc(128)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortWorst256(b *testing.B) {
	l := genListDesc(256)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortWorst512(b *testing.B) {
	l := genListDesc(512)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortWorst1k(b *testing.B) {
	l := genListDesc(1024)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortWorst2k(b *testing.B) {
	l := genListDesc(2 * 1024)
	benchmarkSort(b, InsSort, l)
}

func BenchmarkInsSortWorst4k(b *testing.B) {
	l := genListDesc(4 * 1024)
	benchmarkSort(b, InsSort, l)
}

// === [ SelSort benchmark ] ===================================================

// --- [ random order ] --------------------------------------------------------

func BenchmarkSelSortRand128(b *testing.B) {
	l := genListRand(128)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortRand256(b *testing.B) {
	l := genListRand(256)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortRand512(b *testing.B) {
	l := genListRand(512)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortRand1k(b *testing.B) {
	l := genListRand(1024)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortRand2k(b *testing.B) {
	l := genListRand(2 * 1024)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortRand4k(b *testing.B) {
	l := genListRand(4 * 1024)
	benchmarkSort(b, SelSort, l)
}

// --- [ worst case ] ----------------------------------------------------------

func BenchmarkSelSortWorst128(b *testing.B) {
	l := genListDesc(128)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortWorst256(b *testing.B) {
	l := genListDesc(256)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortWorst512(b *testing.B) {
	l := genListDesc(512)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortWorst1k(b *testing.B) {
	l := genListDesc(1024)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortWorst2k(b *testing.B) {
	l := genListDesc(2 * 1024)
	benchmarkSort(b, SelSort, l)
}

func BenchmarkSelSortWorst4k(b *testing.B) {
	l := genListDesc(4 * 1024)
	benchmarkSort(b, SelSort, l)
}

// === [ BubbleSort benchmark ] ===================================================

// --- [ random order ] --------------------------------------------------------

func BenchmarkBubbleSortRand128(b *testing.B) {
	l := genListRand(128)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortRand256(b *testing.B) {
	l := genListRand(256)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortRand512(b *testing.B) {
	l := genListRand(512)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortRand1k(b *testing.B) {
	l := genListRand(1024)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortRand2k(b *testing.B) {
	l := genListRand(2 * 1024)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortRand4k(b *testing.B) {
	l := genListRand(4 * 1024)
	benchmarkSort(b, BubbleSort, l)
}

// --- [ worst case ] ----------------------------------------------------------

func BenchmarkBubbleSortWorst128(b *testing.B) {
	l := genListDesc(128)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortWorst256(b *testing.B) {
	l := genListDesc(256)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortWorst512(b *testing.B) {
	l := genListDesc(512)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortWorst1k(b *testing.B) {
	l := genListDesc(1024)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortWorst2k(b *testing.B) {
	l := genListDesc(2 * 1024)
	benchmarkSort(b, BubbleSort, l)
}

func BenchmarkBubbleSortWorst4k(b *testing.B) {
	l := genListDesc(4 * 1024)
	benchmarkSort(b, BubbleSort, l)
}

// === [ QuickSort benchmark ] ===================================================

// --- [ random order ] --------------------------------------------------------

func BenchmarkQuickSortRand128(b *testing.B) {
	l := genListRand(128)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortRand256(b *testing.B) {
	l := genListRand(256)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortRand512(b *testing.B) {
	l := genListRand(512)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortRand1k(b *testing.B) {
	l := genListRand(1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortRand2k(b *testing.B) {
	l := genListRand(2 * 1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortRand4k(b *testing.B) {
	l := genListRand(4 * 1024)
	benchmarkSort(b, QuickSort, l)
}

// --- [ worst case: ascending ] -----------------------------------------------

func BenchmarkQuickSortWorstAsc128(b *testing.B) {
	l := genListAsc(128)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstAsc256(b *testing.B) {
	l := genListAsc(256)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstAsc512(b *testing.B) {
	l := genListAsc(512)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstAsc1k(b *testing.B) {
	l := genListAsc(1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstAsc2k(b *testing.B) {
	l := genListAsc(2 * 1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstAsc4k(b *testing.B) {
	l := genListAsc(4 * 1024)
	benchmarkSort(b, QuickSort, l)
}

// --- [ worst case: descending ] ----------------------------------------------

func BenchmarkQuickSortWorstDesc128(b *testing.B) {
	l := genListDesc(128)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstDesc256(b *testing.B) {
	l := genListDesc(256)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstDesc512(b *testing.B) {
	l := genListDesc(512)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstDesc1k(b *testing.B) {
	l := genListDesc(1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstDesc2k(b *testing.B) {
	l := genListDesc(2 * 1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstDesc4k(b *testing.B) {
	l := genListDesc(4 * 1024)
	benchmarkSort(b, QuickSort, l)
}

// --- [ worst case: equal ] ---------------------------------------------------

func BenchmarkQuickSortWorstEq128(b *testing.B) {
	l := genListEq(128)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstEq256(b *testing.B) {
	l := genListEq(256)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstEq512(b *testing.B) {
	l := genListEq(512)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstEq1k(b *testing.B) {
	l := genListEq(1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstEq2k(b *testing.B) {
	l := genListEq(2 * 1024)
	benchmarkSort(b, QuickSort, l)
}

func BenchmarkQuickSortWorstEq4k(b *testing.B) {
	l := genListEq(4 * 1024)
	benchmarkSort(b, QuickSort, l)
}

// genListAsc generates and returns a list in ascending order.
func genListAsc(size int) []int {
	l := make([]int, size)
	for i := range l {
		l[i] = i + 1
	}
	return l
}

// genListAsc generates and returns a list in descending order.
func genListDesc(size int) []int {
	l := make([]int, size)
	for i := range l {
		l[i] = size - i
	}
	return l
}

// genListAsc generates and returns a list in random order.
func genListRand(size int) []int {
	l := make([]int, size)
	for i := range l {
		l[i] = rand.Int()
	}
	return l
}

// genListEq generates and returns a list where all elements have the same
// value.
func genListEq(size int) []int {
	l := make([]int, size)
	x := rand.Int()
	for i := range l {
		l[i] = x
	}
	return l
}

func benchmarkSort(b *testing.B, sortFn func([]int) []int, in []int) {
	size := len(in)
	list := make([]int, size)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(list, in)
		b.StartTimer()
		sortFn(list)
	}
}
