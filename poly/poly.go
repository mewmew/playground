// Package poly handles polynomial arithmatic.
package poly

import (
	"fmt"
)

// A Polynomial is represented by a slice of its coefficients.
//
// For example, a polynomial f(x) of degree n can be written as follows:
//    f(x) = a_0 + a_1*x^1 + ... + a_(n-1)*x^(n-1) + a_n*x^n.
//
// The index of the constant term a_0 is 0 and the index of the coefficient a_n
// is len(f)-1.
//
// Note that the zero value (nil) of a polynomial is the zero polynomial.
type Polynomial []float64

// NewPoly returns a new polynomial of the specified degree.
func NewPoly(deg int) (f Polynomial) {
	f = make(Polynomial, deg+1)
	return f
}

func (f Polynomial) String() (s string) {
	for k := f.Deg(); k >= 0; k-- {
		a := f.At(k)
		if a == 0 {
			continue
		}
		if len(s) > 0 && k != f.Deg() {
			if a < 0 {
				s += " - "
				a *= -1
			} else {
				s += " + "
			}
		} else if a < 0 {
			s += "-"
			a *= -1
		}
		switch k {
		case 0:
			s += fmt.Sprintf("%g", a)
		case 1:
			if a == 1 {
				s += "x"
			} else {
				s += fmt.Sprintf("%gx", a)
			}
		default:
			if a == 1 {
				s += fmt.Sprintf("x^%d", k)
			} else {
				s += fmt.Sprintf("%gx^%d", a, k)
			}
		}
	}
	return s
}

// Deg returns the degree of the polynomial. The zero polynomial will have the
// degree -1.
func (f Polynomial) Deg() int {
	return len(f) - 1
}

// At returns the coefficient at the index k or 0 if no such index exist.
func (f Polynomial) At(k int) float64 {
	if k < 0 || len(f) <= k {
		return 0
	}
	return f[k]
}

// Sub returns the result when subtracting g from f.
//
//    h(x) = f(x) - g(x).
func (f Polynomial) Sub(g Polynomial) (h Polynomial) {
	deg := max(f.Deg(), g.Deg())
	for k := deg; k >= 0; k-- {
		a := f.At(k) - g.At(k)
		if a == 0 {
			continue
		}
		if h == nil {
			h = NewPoly(k)
		}
		h[k] = a
	}
	return h
}

// Add returns the result when adding g to f.
//
//    h(x) = f(x) + g(x).
func (f Polynomial) Add(g Polynomial) (h Polynomial) {
	deg := max(f.Deg(), g.Deg())
	for k := deg; k >= 0; k-- {
		a := f.At(k) + g.At(k)
		if a == 0 {
			continue
		}
		if h == nil {
			h = NewPoly(k)
		}
		h[k] = a
	}
	return h
}

// Div returns the quotient and remainder polynomials when dividing f with g.
//
//    f(x) = q(x)*g(x) + r(x) where 0 <= deg(r) < deg(g)
func (f Polynomial) Div(g Polynomial) (q, r Polynomial) {
	gDeg := g.Deg()
	if g.At(gDeg) == 0 {
		// No need to panic at division by zero.
		//    f(x) = q(x)*0 + r(x)
		//    f(x) = r(x)
		return nil, f
	}
	qDeg := f.Deg() - gDeg
	q = NewPoly(qDeg)

	r = f
	l := g.Deg()
	for k := r.Deg(); k >= l; k-- {
		a := r.At(k)
		if a == 0 {
			continue
		}
		b := g.At(l)
		c := a / b
		m := k - l
		q[m] = c
		part := q[:m+1]
		h := part.Mul(g)
		r = r.Sub(h)
	}
	return q, r
}

// Mul returns the result when multiplying f with g.
//
//    h(x) = f(x)*g(x)
func (f Polynomial) Mul(g Polynomial) (h Polynomial) {
	hDeg := f.Deg() + g.Deg()
	h = NewPoly(hDeg)
	for k := 0; k <= f.Deg(); k++ {
		for l := 0; l <= g.Deg(); l++ {
			m := k + l
			h[m] += f.At(k) * g.At(l)
		}
	}
	return h
}

// max returns the larger of a and b.
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
