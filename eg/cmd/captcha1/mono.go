package main

import "image"
import "image/color"

import "github.com/mewmew/playground/pic"

// getMono returns a monochrome version of orig where the background gradient
// has been replaced with transparent pixels and the rest with black pixels.
func getMono(orig image.Image) (mono pic.SubImager) {
   rect := orig.Bounds()
   dst := image.NewRGBA(rect)
   for y := rect.Min.Y; y < rect.Max.Y; y++ {
      // c is the background gradient color of the current horizontal line.
      c := orig.At(0, y)
      for x := rect.Min.X; x < rect.Max.X; x++ {
         if orig.At(x, y) == c {
            // part of the gradient.
            dst.Set(x, y, color.Transparent)
         } else {
            // not part of the gradient.
            dst.Set(x, y, color.Black)
         }
      }
   }
   return dst
}
