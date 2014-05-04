package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get input from stdin.
	var n, k int
	_, err := fmt.Fscanf(os.Stdin, "%d %d", &n, &k)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Fib(n, k))
}

// Fib returns the total number of rabbit pairs that will be present after n
// months if we begin with 1 pair and in each generation, every pair of
// production-age rabbits produce a litter of k rabbit paris.
//
// Below follows an example where n=5 and k=3. Rabbit kits are represented using
// r and rabbit adults using R.
//
// 1st month (1 rabbit)
//    r
//
// 2nd month (1 rabbit)
//    R
//
// 3rd month (1 + 1*3 = 4 rabbits)
//    R r r r
//
// 4th month (4 + 1*3 = 7 rabbits)
//    R R R R r r r
//
// 5th month (7 + 4*3 = 19 rabbits)
//    R R R R R R R r r r r r r r r r r r r
func Fib(n, k int) int {
	// Similar to the Fibbonaci sequence but calculated using:
	//    F_n = F_{n-1} = k*F_{n-2}
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, b+a*k
	}

	return b
}
