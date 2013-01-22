package qf_test

import "fmt"

import qf "."

func ExampleQF() {
	set := qf.New(10)
	set.Union(4, 3)
	set.Union(3, 8)
	set.Union(6, 5)
	set.Union(9, 4)
	set.Union(2, 1)
	set.Union(5, 0)
	set.Union(7, 2)
	set.Union(6, 1)
	set.Union(1, 0)
	// The following connections have been established.
	//    0--1--2  3--4
	//    |  |  |  |  |
	//    5--6  7  8  9
	fmt.Println(set.IsConnected(0, 7))
	fmt.Println(set.IsConnected(2, 9))
	// Output:
	// true
	// false
}
