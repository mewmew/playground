package main

import (
	"fmt"
	"log"

	"github.com/mewkiz/pkg/geometry"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/playground/archive/fractal/buddha"
)

func main() {
	buddha.Grid = geometry.Rect(-2.5, -2.5, 2.5, 2.5)
	set := buddha.New(1024, 1024)
	filePath := "buddha.png"
	err := imgutil.WriteFile(filePath, set)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Created:", filePath)
}
