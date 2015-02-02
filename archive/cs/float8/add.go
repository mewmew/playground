package float8

import "fmt"

// Add returns the sum of x+y in 8-bit floating-point notation.
func Add(x, y Float8) (z Float8, err error) {
	// The mantissa contains 4 significant bits.
	//    .1111
	//
	// The exponent can move the radix point at most 4 positions to the left and
	// 3 positions to the right.
	//    .00001111 (left)
	//    111.1 (right)
	//
	// Thus it requires 11 bits (4+4+3) to store this information uncompressed
	// without any loss. The result will be at most twice as large, thus
	// requiring 12 bits of storage. The sign bit will be handled separately.
	//
	// buf will contain the uncompressed floating-point value. The first 8 bits
	// contain the integer part and the following 8 bits contain the fraction
	// part.
	//
	//    00000111.10000000 (largest 8-bit floating-point value)
	//    00000000.00001111 (smallest positive 8-bit floating-point value)
	var buf uint16
	var xbuf, ybuf, zbuf int16

	// align the radix point of x's mantissa with the radix point of buf.
	buf = uint16(x.Mantissa()) << 4
	if exp := x.Exp(); exp < 0 {
		buf >>= uint(-exp)
	} else {
		buf <<= uint(exp)
	}
	xbuf = int16(buf)
	if x.Sign() {
		xbuf = -xbuf
	}

	// align the radix point of y's mantissa with the radix point of buf.
	buf = uint16(y.Mantissa()) << 4
	if exp := y.Exp(); exp < 0 {
		buf >>= uint(-exp)
	} else {
		buf <<= uint(exp)
	}
	ybuf = int16(buf)
	if y.Sign() {
		ybuf = -ybuf
	}

	// Add xbuf and ybuf together.
	zbuf = xbuf + ybuf
	if zbuf == 0 {
		// The sum of x+y is 0.
		return 0, nil
	}

	// Encode sign bit of z.
	if zbuf < 0 {
		zbuf = -zbuf
		z |= 0x80
	}

	// Locate the leftmost bit that is set. Search from bit 11 through bit 3.
	pos := 11
	for mask := int16(0x800); mask >= 0x8; mask >>= 1 {
		if zbuf&mask != 0 {
			break
		}
		pos--
	}
	// TODO(u): Allow the Float8 value to overflow and underflow?
	if pos > 10 {
		// exponent overflow.
		return 0, fmt.Errorf("float8.Add: exponent overflow; can't represent %v + %v in 8-bit floating-point notation", x, y)
	} else if pos < 3 {
		// exponent underflow.
		return 0, fmt.Errorf("float8.Add: exponent underflow; can't represent %v + %v in 8-bit floating-point notation", x, y)
	}

	// Encode exponent of z.
	exp := Float8(pos - 7)
	exp += 4 // excess four notation.
	z |= exp << 4

	// Encode mantissa of z.
	shift := uint(pos - 3)
	z |= Float8(zbuf >> shift)

	return z, nil
}

// AddNative returns the sum of x+y in 8-bit floating-point notation. It
// converts x and y to float32 values, preforms a native addition and converts
// the result back to 8-bit floating-point notation.
func AddNative(x, y Float8) (z Float8, err error) {
	return New(x.Float32() + y.Float32())
}
