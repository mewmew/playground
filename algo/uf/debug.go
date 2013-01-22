package uf

import "fmt"
import "sort"

func (set *UF) String() (s string) {
	var ids []int
	for id := range set.id2os {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	s += fmt.Sprintln("=== [ groups ] =================================================================")
	s += fmt.Sprintln()
	for _, id := range ids {
		os := set.id2os[id]
		s += fmt.Sprintf("   group %d: %v\n", id, os)
	}
	s += fmt.Sprintln()

	var os []int
	for o := range set.o2id {
		os = append(os, o)
	}
	sort.Ints(os)

	s += fmt.Sprintln("=== [ objects ] ================================================================")
	s += fmt.Sprintf("\n")
	for _, o := range os {
		id := set.o2id[o]
		s += fmt.Sprintf("   object %d: %d\n", o, id)
	}
	return s
}
