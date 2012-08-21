// Package pic implements various image manipulation functions.
package pic

import "image"

type SubImager interface {
	image.Image
	SubImage(r image.Rectangle) image.Image
}

// Crop returns the smallest version of orig, after cropping the transparent
// pixels.
func Crop(orig SubImager) (crop SubImager) {
	rect := orig.Bounds()
	cropRect := rect
	// x start
	var x int
	for x = rect.Min.X; x < rect.Max.X; x++ {
		if !IsVLineAlpha(orig, x) {
			break
		}
	}
	cropRect.Min.X = x
	// x end
	for x = rect.Max.X - 1; x >= rect.Min.X; x-- {
		if !IsVLineAlpha(orig, x) {
			break
		}
	}
	cropRect.Max.X = x + 1
	// y start
	var y int
	for y = rect.Min.Y; y < rect.Max.Y; y++ {
		if !IsHLineAlpha(orig, y) {
			break
		}
	}
	cropRect.Min.Y = y
	// y end
	for y = rect.Max.Y - 1; y >= rect.Min.Y; y-- {
		if !IsHLineAlpha(orig, y) {
			break
		}
	}
	cropRect.Max.Y = y + 1
	crop, _ = orig.SubImage(cropRect).(SubImager)
	return crop
}

// IsHLineAlpha returns true if the horizontal line at y is transparent, and
// false otherwise.
func IsHLineAlpha(img image.Image, y int) bool {
	rect := img.Bounds()
	for x := rect.Min.X; x < rect.Max.X; x++ {
		_, _, _, a := img.At(x, y).RGBA()
		if a != 0 {
			return false
		}
	}
	return true
}

// IsVLineAlpha returns true if the vertical line at x is transparent, and false
// otherwise.
func IsVLineAlpha(img image.Image, x int) bool {
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		_, _, _, a := img.At(x, y).RGBA()
		if a != 0 {
			return false
		}
	}
	return true
}
