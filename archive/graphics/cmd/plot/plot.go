package main

import (
	"fmt"
	"log"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewmew/playground/archive/graphics/bezier"
)

func main() {
	for p0 := 0.0; p0 <= 1.0; p0 += 0.1 {
		for p1 := 0.0; p1 <= 1.0; p1 += 0.1 {
			for p2 := 0.0; p2 <= 1.0; p2 += 0.1 {
				for p3 := 0.0; p3 <= 1.0; p3 += 0.1 {
					p := [4]float64{p0, p1, p2, p3}
					img := bezier.Plot(p)
					path := fmt.Sprintf("plot_%.1f_%.1f_%.1f_%.1f.png", p0, p1, p2, p3)
					err := imgutil.WriteFile(path, img)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
	}
}
