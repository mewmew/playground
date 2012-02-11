package main

import "dump"

import "flag"
import "log"

func main() {
   flag.Parse()
   for _, rawUrl := range flag.Args() {
      err := dump.Url(rawUrl)
      if err != nil {
         log.Fatalln(err)
      }
   }
}
