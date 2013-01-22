// Package qf implements the quick-find algorithm.
//
//    speed: N^2 (N union commands on N objects).
package qf

// A QF is a set which implements the quick-find algorithm. The int slice is
// used as a mapping from object num to component id.
//
// Interpretation: p and q are connected iff they have the same id.
type QF []int

// New initializes and returns a new quick-find set.
//
//    speed: N array accesses.
func New(n int) (ids QF) {
	ids = make([]int, n)
	// Initialize the set, using the object num as component id for every object.
	for i := range ids {
		ids[i] = i
	}
	return ids
}

// Union connects object a and b.
//
//    speed: 2N + 2 array accesses.
func (ids QF) Union(a, b int) {
	aId := ids[a]
	bId := ids[b]
	if aId == bId {
		return
	}
	// Change all objects whose id equals id[a] to id[b].
	for i := range ids {
		if ids[i] == aId {
			ids[i] = bId
		}
	}
}

// IsConnected returns true if a and b are in the same component.
//
//    speed: 2 array accesses.
func (ids QF) IsConnected(a, b int) bool {
	return ids[a] == ids[b]
}

// Find returns the component id of a.
func (ids QF) Find(a int) int {
	return ids[a]
}

// Count returns the number of components in the set.
func (ids QF) Count() int {
	m := make(map[int]bool)
	for _, id := range ids {
		m[id] = true
	}
	return len(m)
}
