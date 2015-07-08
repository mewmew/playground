package bezier

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Plot plots a Bézier curve based on the provided control points, which serve
// as weights for the Bernstein cubic polynomial basis functions.
func Plot(p [4]float64) image.Image {
	const (
		width, height = 500, 500
		scale         = 1.0 / width
	)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.White), image.ZP, draw.Src)
	for x := 0; x < width; x++ {
		u := scale * float64(x)
		y := int((p[0]*b0(u) + p[1]*b1(u) + p[2]*b2(u) + p[3]*b3(u)) / scale)
		dst.Set(x, y, color.Black)
	}
	return dst
}

// b0 represents the 0th basis function of the Bernstein cubic polynomials.
//
//    b0(u) = (1-u)^3
func b0(u float64) float64 {
	return math.Pow(1-u, 3)
}

// b1 represents the 1st basis function of the Bernstein cubic polynomials.
//
//    b1(u) = 3u(1-u)^2
func b1(u float64) float64 {
	return 3 * u * math.Pow(1-u, 2)
}

// b2 represents the 2nd basis function of the Bernstein cubic polynomials.
//
//    b2(u) = 3u^2(1-u)
func b2(u float64) float64 {
	return 3 * math.Pow(u, 2) * (1 - u)
}

// b3 represents the 3rd basis function of the Bernstein cubic polynomials.
//
//    b3(u) = u^3
func b3(u float64) float64 {
	return math.Pow(u, 3)
}
