// Package uf implements a union-find algorithm.
//
//    speed: ###
package uf

// A UF is a set which implements a union-find algorithm.
type UF struct {
	// o2id is a map from object num to component id.
	o2id map[int]int
	// id2os is a map from component id to object num.
	id2os map[int][]int
}

// New initializes and returns a new union-find set.
//
//    speed: ###
func New(n int) (set *UF) {
	set = &UF{
		o2id:  make(map[int]int, n),
		id2os: make(map[int][]int, n),
	}
	// The object num and component id are initially the same for each individual
	// object.
	for i := 0; i < n; i++ {
		set.o2id[i] = i
		set.id2os[i] = []int{i}
	}
	return set
}

// Union connects object a and b.
//
//    speed: ###
func (set *UF) Union(a, b int) {
	aId := set.o2id[a]
	bId := set.o2id[b]
	if aId == bId {
		return
	}
	aLen := len(set.id2os[aId])
	bLen := len(set.id2os[bId])
	var smallId, bigId int
	if aLen > bLen {
		bigId = aId
		smallId = bId
	} else {
		bigId = bId
		smallId = aId
	}
	// Update the component ids of the smallest group.
	for _, o := range set.id2os[smallId] {
		set.o2id[o] = bigId
	}
	// Append objects from the smallest to the largest group.
	set.id2os[bigId] = append(set.id2os[bigId], set.id2os[smallId]...)
	delete(set.id2os, smallId)
}

// IsConnected returns true if a and b are in the same component.
//
//    speed: ###
func (set *UF) IsConnected(a, b int) bool {
	return set.o2id[a] == set.o2id[b]
}

// Find returns the component id of a.
func (set *UF) Find(a int) int {
	return set.o2id[a]
}

// Count returns the number of components in the set.
func (set *UF) Count() int {
	return len(set.id2os)
}
