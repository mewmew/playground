package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/playground/archive/fractal/mset"
)

func main() {
	colors := make([]color.Color, mset.Iterations)
	var c color.RGBA
	// Create a transition from black to blue to white.
	for i := range colors {
		x := uint8(float64(i) * 256 / float64(mset.Iterations/2))
		if i < mset.Iterations/2 {
			// Slow transition from black to blue.
			c = color.RGBA{0, 0, x, 255}
		} else {
			// Slow transition from blue to white.
			c = color.RGBA{x, x, 255, 255}
		}
		colors[i] = c
	}
	mset.Colors = colors
	set := mset.New(1280, 1024)
	filePath := "pretty.png"
	err := imgutil.WriteFile(filePath, set)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Created:", filePath)
}
