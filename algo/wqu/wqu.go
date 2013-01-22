// Package wqu implements the weighted quick-union algorithm.
//
//    speed: ###
package wqu

// A WQU is a set which implements the weighted quick-union algorithm. The ids
// int slice is used as a mapping from object num to root node.
//
// Interpretation: id[a] is parent of a.
type WQU struct {
	ids  []int
	size []int
}

// New initializes and returns a new weighted quick-union set.
//
//    speed: 2N array accesses.
func New(n int) (set *WQU) {
	set = &WQU{
		ids:  make([]int, n),
		size: make([]int, n),
	}
	// Initialize the set, using the object num as component id for every object.
	for i := range set.ids {
		set.ids[i] = i
		set.size[i] = 1
	}
	return set
}

// Union connects object a and b.
//
//    speed: lg N.
func (set *WQU) Union(a, b int) {
	aRoot := set.root(a)
	bRoot := set.root(b)
	if aRoot == bRoot {
		return
	}
	// Link root of smaller tree to root of larger tree. If the trees are equal
	// in size, link b to a.
	if set.size[aRoot] < set.size[bRoot] {
		set.ids[aRoot] = bRoot
		set.size[bRoot] += set.size[aRoot]
	} else {
		set.ids[bRoot] = aRoot
		set.size[aRoot] += set.size[bRoot]
	}
}

// IsConnected returns true if a and b are in the same component.
//
//    speed: lg N.
func (set *WQU) IsConnected(a, b int) bool {
	return set.root(a) == set.root(b)
}

// root returns the root node of a.
//
//    speed: depth of a array accesses.
func (set *WQU) root(a int) int {
	// Root nodes have the same object num and component id.
	for a != set.ids[a] {
		a = set.ids[a]
	}
	return a
}

// Find returns the component id of a.
//
//    speed: depth of a array accesses.
func (set *WQU) Find(a int) int {
	return set.root(a)
}

// Count returns the number of components in the set.
func (set *WQU) Count() int {
	m := make(map[int]bool)
	// Count unique root nodes.
	for i := range set.ids {
		m[set.root(i)] = true
	}
	return len(m)
}
