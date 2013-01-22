package qu_test

import "fmt"

import qu "."

func ExampleQU() {
	set := qu.New(10)
	set.Union(4, 3)
	set.Union(3, 8)
	set.Union(6, 5)
	set.Union(9, 4)
	set.Union(2, 1)
	// The following connections have been established.
	//    0  1  5  7  8
	//       |  |     |\
	//       2  6     3 9
	//                |
	//                4
	fmt.Println(set.IsConnected(8, 9))
	fmt.Println(set.IsConnected(5, 4))
	// Output:
	// true
	// false
}
