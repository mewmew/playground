// Package mset constructs visual representations of the Mandelbrot set.
//
// The Mandelbrot set is the visual representation of an iterated function on
// the complex plane.
//    ref: http://www.mandelbrotset.net/tutorial.html
package mset

import (
	"image"
	"image/color"

	"github.com/mewkiz/pkg/geometry"
)

// New returns a new image of the Mandelbrot set, with the specified dimensions.
func New(w, h int) (img *Image) {
	img = &Image{
		w: w,
		h: h,
	}
	return img
}

// Image implements a visual representation of the Mandelbrot set.
type Image struct {
	// When White is true, color all pixels white that are outside of the circle
	// before the first iteration.
	White bool
	w, h  int
}

// ColorModel returns the Image's color model.
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

// At returns the color of the pixel at (x, y).
//
// The color of each pixel is based on the number of iterations it took to
// escape the circle. Pixels that are already outside of the circle will be
// colored white, if img.White is set to true. Pixels that doesn't escape the
// circle after the max number of iterations will be colored black.
func (img *Image) At(x, y int) color.Color {
	// Convert the x and y coordinates to a complex number.
	rFactor := Grid.Dx() / float64(img.w)
	iFactor := Grid.Dy() / float64(img.h)
	r := Grid.Min.X + float64(x)*rFactor
	i := Grid.Min.Y + float64(y)*iFactor
	c := complex(r, i)

	// Print the pixel white if it is already outside of the circle.
	if img.White && hasEscaped(c) {
		return color.White
	}

	next := gen(c)
	for i := 0; i < Iterations; i++ {
		x := next()
		if hasEscaped(x) {
			// Give the pixel a color based on how many iterations it took to
			// escape the circle.
			return Colors[i%len(Colors)]
		}
	}

	// Pain the pixel black since it didn't escape the circle during the
	// specified number of iterations.
	return color.Black
}

// hasEscaped returns true if c falls outside of a circle with radius 2 centered
// around origin.
func hasEscaped(c complex128) bool {
	return real(c)*real(c)+imag(c)*imag(c) > 4
}

// gen returns a function which on successive calls calculates and returns the
// next output, based on the formula below:
// the formula:
//    f(x) = x**2 + c
//
// gen stores c and uses it as the first input to f.
func gen(c complex128) func() complex128 {
	x := c
	next := func() complex128 {
		// f(x) = x**2 + c
		x = x*x + c
		return x
	}
	return next
}

// Grid is the bounds of the coordinate system on which the Mandelbrot set
// will be displayed. The x and y values of grid doesn't represent pixels but
// rather positions on the coordinate system.
var Grid = geometry.Rect(-2, -1.2, 1, 1.2)

// definition of a few common colors.
var (
	red    = color.RGBA{255, 0, 0, 255}
	green  = color.RGBA{0, 128, 0, 255}
	blue   = color.RGBA{0, 0, 255, 255}
	yellow = color.RGBA{255, 255, 0, 255}
)

// Colors represent the slice of colors that are used to represent how many
// iterations it took to escape the circle.
//
// Thus pixels that escape on the first iteration are colored red, while pixels
// that escape on the second iteration are colored green.
var Colors = []color.Color{
	red,
	green,
	blue,
	yellow,
}

// Iterations represent the number of iterations that are required before a
// pixel is painted black.
var Iterations = 100
