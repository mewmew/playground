package main

import dbg "fmt"
import "math"

func main() {
   for c := float64(1); c < 1e5; c++ {
      for b := float64(1); b < 1e5; b++ {
         logic8(b, c)
      }
   }
}

func logic8(b, c float64) {
   ratio := c / (c + 20141)




   // left side (small triangle)
   _, frac := math.Modf(b * ratio)
   if frac != 0 {
      return
   }
   // left side (large triangle)
   _, frac = math.Modf(b / ratio)
   if frac != 0 {
      return
   }
   ///dbg.Println("b:", b)
   dbg.Println("c:", c)
}
