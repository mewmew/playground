package main

import dbg "fmt"

func main() {
	for down := float64(0); down <= 90*3; down++ {
		for flat := float64(0); flat <= 80*3; flat++ {
			for up := float64(0); up <= 72*3; up++ {
				logib6a(down, flat, up)
			}
		}
	}
}

func logib6a(down, flat, up float64) {
	if down/90+flat/80+up/72 == 3 {
		logib6b(up, flat, down)
	}
}

func logib6b(down, flat, up float64) {
	if down/90+flat/80+up/72 == 3.5 {
		dbg.Println("down:", down)
		dbg.Println("flat:", flat)
		dbg.Println("up:", up)
		dbg.Println()
	}
}
