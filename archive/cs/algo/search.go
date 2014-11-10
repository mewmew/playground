// Package algo implements various basic algorithms.
package algo

import "sort"

// SeqList is a sequential list of integers.
type SeqList []int

// NewSeqList sorts list and returns it as a sequential list of integers.
func NewSeqList(list []int) SeqList {
	sort.Ints(list)
	return list
}

// Contains locates the presence of n in list using the sequential search
// algorithm. The receiver is assumed to be a sequential list of integers.
func (list SeqList) Contains(n int) bool {
	for _, v := range list {
		if v >= n {
			if v == n {
				return true
			}
			break
		}
	}
	return false
}
