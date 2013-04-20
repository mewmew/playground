package turtle

import (
	"math/rand"
	"time"
)

// NumOrder attaches the methods of sort.Interface to []*Radical, sorting in
// increasing order based on the radical's number.
type NumOrder []*Radical

// Len is the number of elements in the collection.
func (a NumOrder) Len() int {
	return len(a)
}

// Less returns whether the element with index i should sort
// before the element with index j.
func (a NumOrder) Less(i, j int) bool {
	return a[i].Num < a[j].Num
}

// Swap swaps the elements with indexes i and j.
func (a NumOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// FreqOrder attaches the methods of sort.Interface to []*Radical, sorting in
// decreasing order based on frequency.
type FreqOrder []*Radical

// Len is the number of elements in the collection.
func (a FreqOrder) Len() int {
	return len(a)
}

// Less returns whether the element with index i should sort
// before the element with index j.
func (a FreqOrder) Less(i, j int) bool {
	if a[i].Freq == a[j].Freq {
		return a[i].Num < a[j].Num
	}
	return a[i].Freq > a[j].Freq
}

// Swap swaps the elements with indexes i and j.
func (a FreqOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// rnd provides pseudo-random numbers, good enough for the random sort order.
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// RndOrder attaches the methods of sort.Interface to []*Radical, sorting in
// random order.
type RndOrder []*Radical

// Len is the number of elements in the collection.
func (a RndOrder) Len() int {
	return len(a)
}

// Less returns whether the element with index i should sort
// before the element with index j.
func (a RndOrder) Less(i, j int) bool {
	return rnd.Intn(2) == 0
}

// Swap swaps the elements with indexes i and j.
func (a RndOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
