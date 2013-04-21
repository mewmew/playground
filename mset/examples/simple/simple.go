package main

import (
	"fmt"
	"log"

	"github.com/mewkiz/pkg/geometry"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/playground/mset"
)

func main() {
	mset.Grid = geometry.Rect(-2.5, -2.5, 2.5, 2.5)
	set := mset.New(1024, 1024)
	set.White = true
	filePath := "simple.png"
	err := imgutil.WriteFile(filePath, set)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Created:", filePath)
}
