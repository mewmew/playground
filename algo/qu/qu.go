// Package qu implements the quick-union algorithm.
//
//    speed: ###
package qu

// A QU is a set which implements the quick-union algorithm. The int slice is
// used as a mapping from object num to root node.
//
// Interpretation: id[a] is parent of a.
type QU []int

// New initializes and returns a new quick-union set.
//
//    speed: N array accesses.
func New(n int) (ids QU) {
	ids = make([]int, n)
	// Initialize the set, using the object num as component id for every object.
	for i := range ids {
		ids[i] = i
	}
	return ids
}

// Union connects object a and b.
//
//    speed: depth of p and q + 1 array accesses.
func (ids QU) Union(a, b int) {
	aRoot := ids.root(a)
	bRoot := ids.root(b)
	if aRoot == bRoot {
		return
	}
	ids[aRoot] = bRoot
}

// IsConnected returns true if a and b are in the same component.
//
//    speed: depth of a and b array accesses.
func (ids QU) IsConnected(a, b int) bool {
	return ids.root(a) == ids.root(b)
}

// root returns the root node of a.
//
//    speed: depth of a array accesses.
func (ids QU) root(a int) int {
	// Root nodes have the same object num and component id.
	for a != ids[a] {
		a = ids[a]
	}
	return a
}

// Find returns the component id of a.
//
//    speed: depth of a array accesses.
func (ids QU) Find(a int) int {
	return ids.root(a)
}

// Count returns the number of components in the set.
func (ids QU) Count() int {
	m := make(map[int]bool)
	// Count unique root nodes.
	for i := range ids {
		m[ids.root(i)] = true
	}
	return len(m)
}
