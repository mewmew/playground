package wqu_test

import "fmt"

import "github.com/mewmew/algo/wqu"

func ExampleWQU() {
	set := wqu.New(10)
	set.Union(4, 3)
	set.Union(3, 8)
	set.Union(6, 5)
	set.Union(9, 4)
	set.Union(2, 1)
	set.Union(5, 0)
	set.Union(7, 2)
	set.Union(6, 1)
	// The following connections have been established.
	//      4      6
	//     /|\    /|\
	//    3 8 9  0 2 5
	//            / \
	//           1   7
	fmt.Println(set.IsConnected(8, 9))
	fmt.Println(set.IsConnected(5, 4))
	// Output:
	// true
	// false
}
