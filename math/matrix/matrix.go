// Package matrix implements common matrix operations.
package matrix

import (
	"fmt"
)

// Matrix is an m by n matrix of m rows and n columns.
type Matrix [][]float64

// New creates and returns a new m by n matrix of m rows and n columns.
func New(m, n int) (A Matrix) {
	if m < 1 {
		panic("matrix.New: at least one row is required")
	}
	if n < 1 {
		panic("matrix.New: at least one column is required")
	}
	A = make([][]float64, m)
	for i := range A {
		A[i] = make([]float64, n)
	}
	return A
}

// Rows returns the number of rows in A.
func (A Matrix) Rows() (m int) {
	return len(A)
}

// Cols returns the number of columns in A.
func (A Matrix) Cols() (n int) {
	return len(A[0])
}

// ForEach calls f for each entry (at row i and column j) in A.
func (A Matrix) ForEach(f func(i, j int)) {
	for i := range A {
		for j := range A[0] {
			f(i, j)
		}
	}
}

// ScalarMul returns the result of multiplying A by the scalar b.
//
//    C = bA
func (A Matrix) ScalarMul(b float64) (C Matrix) {
	C = New(A.Rows(), A.Cols())
	f := func(i, j int) {
		C[i][j] = b * A[i][j]
	}
	C.ForEach(f)
	return C
}

// Add returns the result of adding A to B.
//
//    C = A + B
func (A Matrix) Add(B Matrix) (C Matrix) {
	C = New(A.Rows(), A.Cols())
	f := func(i, j int) {
		C[i][j] = A[i][j] + B[i][j]
	}
	C.ForEach(f)
	return A
}

// Sub returns the result of subtracting A from B.
//
//    C = A - B
func (A Matrix) Sub(B Matrix) (C Matrix) {
	C = New(A.Rows(), A.Cols())
	f := func(i, j int) {
		C[i][j] = A[i][j] - B[i][j]
	}
	C.ForEach(f)
	return A
}

// Mul returns the result of multiplying A by B.
//
//    C = A * B
func (A Matrix) Mul(B Matrix) (C Matrix) {
	if A.Cols() != B.Rows() {
		panic(fmt.Sprintf("Matrix.Mul: undefined multiplication; the number of columns in A (%d) not equal to the number of rows in B (%d)", A.Cols(), B.Rows()))
	}
	C = New(A.Rows(), B.Cols())
	f := func(i, j int) {
		for k := 0; k < A.Cols(); k++ {
			C[i][j] += A[i][k] * B[k][j]
		}
	}
	C.ForEach(f)
	return C
}

// Trans returns the transpose of the matrix A.
//
//    B = trans(A)
func (A Matrix) Trans() (B Matrix) {
	B = New(A.Cols(), B.Rows())
	f := func(i, j int) {
		B[i][j] = A[j][i]
	}
	B.ForEach(f)
	return B
}

// Det returns the determinant of the matrix A.
func (A Matrix) Det() float64 {
	panic("Matrix.Det: not yet implemented.")
}

// Adj returns the adjugate of the matrix A.
func (A Matrix) Adj() (B Matrix) {
	panic("Matrix.Adj: not yet implemented.")
}

// Inv returns the inverse of the matrix A.
func (A Matrix) Inv() (B Matrix) {
	panic("Matrix.Inv: not yet implemented.")
}
