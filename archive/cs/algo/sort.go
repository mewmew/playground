package algo

// InsSort sorts list using the insertion sort algorithm. It works by extending
// a sorted subset of the list by one element at the time.
func InsSort(list []int) []int {
	// u represent the start position of the unsorted portion of the list. The
	// first element is always sorted, thus start checking the second element of
	// the list.
	for u := 1; u < len(list); u++ {
		// Check the element against every previous element and insert it into the
		// already sorted subset of the list.
		for i := 0; i < u; i++ {
			if list[u] < list[i] {
				// swap
				list[u], list[i] = list[i], list[u]
			}
		}
	}
	return list
}

// SelSort sorts list using the selection sort algorithm. It works by locating
// the smallest entry from an unsorted portion of the list and moving it to the
// end of the sorted portion of the list.
func SelSort(list []int) []int {
	// u represent the start position of the unsorted portion of the list.
	// Initially the entire list is unsorted.
	for u := 0; u < len(list); u++ {
		// Locate the smallest integer in the unsorted portion of the list.
		// min represent the minimal value of the unsorted portion of the list.
		min := list[u]
		// minPos represent the position of min in list.
		minPos := u
		for i := u + 1; i < len(list); i++ {
			v := list[i]
			if v < min {
				min = v
				minPos = i
			}
		}
		// Place smallest integer from the unsorted portion of the list at the end
		// of the sorted portion of the list.
		if u != minPos {
			list[u], list[minPos] = min, list[u]
		}
	}
	return list
}

// BubbleSort sorts list using the bubble sort algorithm. It works by comparing
// adjecent entites and interchanging them if they are not in the correct order
// relative to each other. Each pass will pull the smallest entity to the start
// of the unsorted portion of the list. Watching the algorithm at work, one sees
// the small entities bubble to the top of the list.
func BubbleSort(list []int) []int {
	// u represent the position of the unsorted portion of the list. Initially
	// the entire list is unsorted.
	for u := 0; u < len(list); u++ {
		for j := len(list) - 1; j > u; j-- {
			// i represent the position directly in front of j.
			i := j - 1
			// Compare the adjecent entities and swap them if they are not in the
			// correct order.
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

// TODO(u): Write a concurrent version of QuickSort that performs each recursive
// call in a goroutine.

// TODO(u): Write a version of QuickSort that selects the pivot entry based on
// the median of a few random samples.

// QuickSort sorts list in place using the quicksort algorithm. It works by
// partitioning the list around a selected pivot entry. Every element in the
// first partition of the list is less than or equal to the pivot entry and
// every element in the second partition of the list is greater than the pivot
// entry. The quicksort algorithm is then applied recursively on each partition
// until the partition length is less than or equal to 1 in which case the
// partition is always sorted.
func QuickSort(list []int) []int {
	if len(list) <= 1 {
		// A list of one element is always sorted.
		return list
	}
	// Partition the list in two.
	q := partition(list)
	// Apply the quicksort algorithm on the smaller partition.
	QuickSort(list[:q])
	// Apply the quicksort algorithm on the larger partition.
	QuickSort(list[q+1:])
	return list
}

// partition partitions the list around a selected pivot entry. The list is
// divided into three parts; list[:q], the smaller partition containing all
// elements less than or equal to the pivot entry; list[q], the pivot entry; and
// list[q+1:], the larger partition containing all elements greater than the
// pivot entry.
func partition(list []int) (q int) {
	// The last element of the list is selected as the pivot entry.
	r := len(list) - 1
	pivot := list[r]
	for j := 0; j < r; j++ {
		// All elements in the smaller partition are less than or equal to the
		// pivot entry and all elements in the larger partition are greater the pivot entry.
		if list[j] <= pivot {
			if q != j {
				//	Swap to include list[j] in the smaller partition.
				list[q], list[j] = list[j], list[q]
			}
			// Grow the smaller partition.
			q++
		}
	}
	if q != r {
		//	Swap to place the pivot entry in between the smaller and the larger
		// partitions.
		list[q], list[r] = list[r], list[q]
	}
	return q
}
