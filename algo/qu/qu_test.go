package qu_test

import "testing"

import qu "."

func BenchmarkNew1e1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qu.New(1e1)
	}
}

func BenchmarkNew1e2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qu.New(1e2)
	}
}

func BenchmarkNew1e3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qu.New(1e3)
	}
}

func BenchmarkNew1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qu.New(1e4)
	}
}

func BenchmarkUnion128(b *testing.B) {
	n := 128
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion256(b *testing.B) {
	n := 256
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion512(b *testing.B) {
	n := 512
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion1024(b *testing.B) {
	n := 1024
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion2048(b *testing.B) {
	n := 2048
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion4096(b *testing.B) {
	n := 4096
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}

func BenchmarkUnion8192(b *testing.B) {
	n := 8192
	for i := 0; i < b.N; i++ {
		set := qu.New(n)
		for x := 1; x < n; x *= 2 {
			for j := 0; j < n; j += x * 2 {
				set.Union(j, j+x)
			}
		}
	}
}
