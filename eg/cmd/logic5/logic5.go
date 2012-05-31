package main

import dbg "fmt"

func main() {
   for n := 1; n < 1000; n++ {
      for count := 2; count < 1000; count++ {
         logic5(n, count)
      }
   }
}

func logic5(n, count int) {
   if 40*n*count == 420*n-380 {
      dbg.Println("n:", n)
      dbg.Println("count:", count)
   }
}
