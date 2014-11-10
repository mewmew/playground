// Package float8 implements values in 8-bit floating-point notation.
package float8

import (
	"fmt"
	"math"
)

// Float8 represents an 8-bit value stored in floating-point notation.
//
// It consists of a sign bit (1 bit), an exponent (3 bits) and a mantissa (4
// bits).
//
// The value is negative when the sign bit is true, and positive otherwise.
//
// The mantissa corresponds to a 4-bit fractional value with an implicit radix
// point on its left side.
//
// The exponent field is represented in excess four notation. It tells how many
// positions (or bits) the radix point should be moved to the right. Note that a
// negative exponent will move the radix point to the left.
//
// To eliminate the possibility of multiple representations for the same value a
// normalized form has been established. The rule is simple; the leftmost bit of
// the mantissa should always be a 1. The only exception to this rule is the
// floating-point value 0, in which the bits of the sign bit, the exponent and
// the mantissa are all 0.
//
// Example:
//
//    10101001
//
//    sign bit:    1 = true
//    exponent:  010 = 2 - 4 = -2
//    mantissa: 1001 = 0.1001
//
// The value is negative since the sign bit is set.
//
// The mantissa corresponds to the 4-bit floating-point value "0.1001".
//
// The radix point should be moved 2 positions to the left, since the exponent
// is negative.
//
// The result is displayed below:
//    -0.001001 (base 2) = -(1/8 + 1/64) = -9/64 = -0.140625 (base 10)
type Float8 uint8

// New converts the provided float32 into an 8-bit floating-point value.
//
// The 32-bit floating-point value consists of a sign bit (1 bit), an exponent
// (8 bits) and a mantissa (23 bits).
//
// The value is negative when the sign bit is true, and positive otherwise.
//
// The mantissa corresponds to a 23-bit fractional value with an implicit radix
// point on its left side. Further more, it has an implicit 1 on the integer
// side of the radix point.
//
// The exponent field is represented in excess 127 notation (2^(n-1)-1). It
// tells how many positions (or bits) the radix point should be moved to the
// right. Note that a negative exponent will move the radix point to the left.
//
// Example:
//
//    01000001110010000000000000000000
//
//    sign bit:                       0 = false
//    exponent:                10000011 = 131 - 127 = 4
//    mantissa: 10010000000000000000000 = 1.1001 (implicit 1 added on left side)
//
// The value is positive since the sign bit isn't set.
//
// The mantissa corresponds to the 24-bit floating-point value "1.1001". Note
// that and implicit 1 has been added on the integer side of the radix point.
//
// The radix point should be moved 4 positions to the right.
//
// The result is displayed below:
//    11001.0 (base 2) = 25.0 (base 10)
//
// ref: https://en.wikipedia.org/wiki/Single-precision_floating-point_format
func New(x float32) (f Float8, err error) {
	if x == 0 {
		// Return the floating-point value 0.
		return 0, nil
	}

	bits := math.Float32bits(x)

	// Sign bit.
	sign := bits&0x80000000 != 0
	if sign {
		// Encode sign bit.
		f |= 0x80
	}

	// Exponent.
	exp := int(bits & 0x7F800000 >> 23)
	exp -= 127 // excess 127 notation.
	exp += 1   // adjust exponent for the implicit 1.
	if exp < -4 {
		return 0, fmt.Errorf("float8.New: invalid exponent (%d) for %v; below -4.", exp, x)
	} else if exp > 3 {
		return 0, fmt.Errorf("float8.New: invalid exponent (%d) for %v; above 3.", exp, x)
	}

	// Encode exponent.
	fexp := Float8(exp)
	fexp += 4 // excess four notation.
	f |= fexp << 4

	// Mantissa.
	mantissa := bits & 0x007FFFFF
	// add the implicit 1 to the mantissa.
	mantissa |= 0x00800000

	// Encode the mantissa.
	// preserve 4 bits of the mantissa.
	mantissa >>= 20
	f |= Float8(mantissa)

	return f, nil
}

// Float32 converts the 8-bit floating-point value to a float32. The 8-bit and
// the 32-bit floating-point notation formats are described at Float8 and New
// respectively.
func (f Float8) Float32() float32 {
	if f == 0 {
		// Return the floating-point value 0.
		return 0
	}

	// bits is the bit representation of the float32 value.
	var bits uint32

	// Sign bit.
	if f.Sign() {
		bits |= 0x80000000
	}

	// Exponent.
	exp := f.Exp()
	exp -= 1 // adjust exponent for the implicit 1.

	// Encode exponent.
	xexp := uint8(exp)
	xexp += 127 // excess 127 notation.
	bits |= uint32(xexp) << 23

	// Mantissa.
	mantissa := f.Mantissa()
	// remove the implicit 1 from the mantissa.
	mantissa &^= 0x8

	// Encode the mantissa.
	bits |= uint32(mantissa) << 20

	return math.Float32frombits(bits)
}

func (f Float8) String() string {
	return fmt.Sprint(f.Float32())
}

// Sign returns true if f is negative, and false otherwise.
func (f Float8) Sign() bool {
	sign := f&0x80 != 0
	return sign
}

// Exp returns the exponent of f.
func (f Float8) Exp() int {
	exp := int(f & 0x70 >> 4)
	exp -= 4 // excess four notation.
	return exp
}

// Mantissa returns the mantissa of f.
func (f Float8) Mantissa() uint {
	mantissa := uint(f & 0xF)
	if mantissa != 0 && mantissa&0x8 == 0 {
		// The mantissa is not represented in normalized form.
		panic(fmt.Sprintf("mantissa (%04b) in f (%08b) is not represented in normalized form.", mantissa, f))
	}
	return mantissa
}
