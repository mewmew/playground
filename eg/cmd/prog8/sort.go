package main

import "image"
import "image/color"
import "image/draw"
import "sort"

import "github.com/mewmew/playground/pic"

type Line struct {
   y int
   c color.Color
}

type Lines []*Line

func (lines Lines) Len() int {
   return len(lines)
}

func (lines Lines) Swap(i, j int) {
   lines[i], lines[j] = lines[j], lines[i]
}

func (lines Lines) Less(i, j int) bool {
   r1, g1, b1, a1 := lines[i].c.RGBA()
   r2, g2, b2, a2 := lines[j].c.RGBA()
   if r1 < r2 {
      return true
   }
   if g1 < g2 {
      return true
   }
   if b1 < b2 {
      return true
   }
   if a1 < a2 {
      return true
   }
   return false
}

// getLineOrder returns the line order, which has been calculated by sorting the
// lines based on the gradient color.
func getLineOrder(orig image.Image) (ys []int) {
   rect := orig.Bounds()
   lines := make(Lines, 0)
   for y := rect.Min.Y; y < rect.Max.Y; y++ {
      c := orig.At(0, y)
      line := &Line{y: y, c: c}
      lines = append(lines, line)
   }
   sort.Sort(lines)
   for _, line := range lines {
      ys = append(ys, line.y)
   }
   return ys
}

// getSorted returns an image, who's horizontal lines have been sorted based on
// the background gradient.
func getSorted(orig image.Image) (sorted pic.SubImager) {
   rect := orig.Bounds()
   dst := image.NewRGBA(rect)
   rectLine := rect
   ys := getLineOrder(orig)
   for y, origY := range ys {
      rectLine.Min.Y = y
      rectLine.Max.Y = y + 1
      draw.Draw(dst, rectLine, orig, image.Pt(0, origY), draw.Src)
   }
   return dst
}
