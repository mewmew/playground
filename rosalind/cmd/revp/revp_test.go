package main

import (
	"fmt"
)

func ExampleRevPal() {
	locs, ns := RevPal("TCAATGCATGCGGGTCTATATGCAT")
	for i := range locs {
		loc := locs[i]
		n := ns[i]
		fmt.Println(loc+1, n)
	}
	// Output:
	// 4 6
	// 5 4
	// 6 6
	// 7 4
	// 17 4
	// 18 4
	// 20 6
	// 21 4
}
