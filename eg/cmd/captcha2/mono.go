package main

import "image"
import "image/color"

import "github.com/mewmew/playground/pic"

// getMono returns a monochrome version of orig where the background has been
// replaced with transparent pixels and the rest with black pixels.
func getMono(orig image.Image) (mono pic.SubImager) {
	rect := orig.Bounds()
	dst := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			textColor := color.RGBA{R: 41, G: 36, B: 33, A: 255}
			if orig.At(x, y) == textColor {
				// text
				dst.Set(x, y, color.Black)
			} else {
				// background
				dst.Set(x, y, color.Transparent)
			}
		}
	}
	return dst
}
