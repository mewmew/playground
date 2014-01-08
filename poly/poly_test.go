package poly

import (
	"reflect"
	"testing"
)

type mulTest struct {
	f    Polynomial
	g    Polynomial
	want Polynomial
}

func TestPolynomialMul(t *testing.T) {
	golden := []mulTest{
		// i=0
		//    f(x) = x^2 - 1
		//    g(x) = x + 2
		//    want(x) = x^3 + 2x^2 - x - 2
		{
			f:    Polynomial{-1, 0, 1},
			g:    Polynomial{2, 1},
			want: Polynomial{-2, -1, 2, 1},
		},
		// i=1
		//    f(x) = -2x^3 - x^2 + 4x - 1
		//    g(x) = x^2 + x - 3
		//    want(x) = asdf
		{
			f:    Polynomial{-1, 4, -1, -2},
			g:    Polynomial{-3, 1, 1},
			want: Polynomial{3, -13, 6, 9, -3, -2},
		},
	}
	for i, g := range golden {
		got := g.f.Mul(g.g)
		if !reflect.DeepEqual(got, g.want) {
			t.Errorf("i=%d: expected %q, got %q.", i, g.want, got)
		}
	}
}

type divTest struct {
	f     Polynomial
	g     Polynomial
	qWant Polynomial
	rWant Polynomial
}

func TestPolynomialDiv(t *testing.T) {
	golden := []divTest{
		// i=0
		//    f(x) = x^3 + 2x^2 - x - 2
		//    g(x) = x + 2
		//    qWant(x) = x^2 - 1
		//    rWant(x) = 0
		{
			f:     Polynomial{-2, -1, 2, 1},
			g:     Polynomial{2, 1},
			qWant: Polynomial{-1, 0, 1},
			rWant: nil,
		},
		// i=1
		//    f(x) = 4x^3 + 2x^2 - x - 2
		//    g(x) = 2x + 2
		//    qWant(x) = 2x^2 - x + 1/2
		//    rWant(x) = -3
		{
			f:     Polynomial{-2, -1, 2, 4},
			g:     Polynomial{2, 2},
			qWant: Polynomial{0.5, -1, 2},
			rWant: Polynomial{-3},
		},
		// i=2
		//    f(x) = 4x^3 + 2x^2 - x - 2
		//    g(x) = 2x + 2
		//    qWant(x) = 2x^2 - x + 1/2
		//    rWant(x) = -3
		{
			f:     Polynomial{4},
			g:     Polynomial{2},
			qWant: Polynomial{2},
			rWant: nil,
		},
		// i=3
		//    f(x) = x^2 - x
		//    g(x) = x
		//    qWant(x) = x - 1
		//    rWant(x) = 0
		{
			f:     Polynomial{0, -1, 1},
			g:     Polynomial{0, 1},
			qWant: Polynomial{-1, 1},
			rWant: nil,
		},
		// i=4
		//    f(x) = x + 3
		//    g(x) = 0
		//    qWant(x) = 0
		//    rWant(x) = x + 3
		{
			f:     Polynomial{3, 1},
			g:     Polynomial{0},
			qWant: nil,
			rWant: Polynomial{3, 1},
		},
		// i=5
		//    f(x) = 5x^3 - x
		//    g(x) = 0
		//    qWant(x) = 0
		//    rWant(x) = 5x^3 - x
		{
			f:     Polynomial{0, -1, 0, 5},
			g:     nil,
			qWant: nil,
			rWant: Polynomial{0, -1, 0, 5},
		},
	}
	for i, g := range golden {
		qGot, rGot := g.f.Div(g.g)
		if !reflect.DeepEqual(qGot, g.qWant) {
			t.Errorf("i=%d: expected quotient %v, got %v.", i, g.qWant, qGot)
		}
		if !reflect.DeepEqual(rGot, g.rWant) {
			t.Errorf("i=%d: expected remainder %v, got %v.", i, g.rWant, rGot)
		}
	}
}
