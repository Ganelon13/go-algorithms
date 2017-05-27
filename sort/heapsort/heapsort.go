//
// Package heapsort contains sorting methods based on a binary heap.
//
// Heapsort is a comparison-based sorting algorithm. Heapsort can be
// thought of as an improved selection sort: like that algorithm, it
// divides its input into a sorted and an unsorted region, and it
// iteratively shrinks the unsorted region by extracting the
// largest/smallest element and moving that to the sorted region.
// The improvement consists of the use of a heap data structure rather
// than a linear-time search to find the maximum
//
package heapsort

import (
	"github.com/Ganelon13/go-data-structures/binaryheap"
	"sort"
)

// SortInts sorting []int slice. Method independently implements a binary heap.
func SortInts(a []int) {
	for start := (len(a) - 2) / 2; start >= 0; start-- {
		siftDownInts(a, start, len(a)-1)
	}
	for end := len(a) - 1; end > 0; end-- {
		a[0], a[end] = a[end], a[0]
		siftDownInts(a, 0, end-1)
	}
}

// SortInts sorting []int slice. Method Use the implemented binary heap
func SortIntsWithHeap(a []int) {
	h := binaryheap.Init(a, binaryheap.MIN)
	for i := 0; h.Len() > 0; i++ {
		a[i], _ = h.Pop()
	}
}

// SortInterface sorting any type that implements sort.Interface.
// Method independently implements a binary heap.
func SortInterface(a sort.Interface) {
	for start := (a.Len() - 2) / 2; start >= 0; start-- {
		siftDownInterface(a, start, a.Len()-1)
	}
	for end := a.Len() - 1; end > 0; end-- {
		a.Swap(0, end)
		siftDownInterface(a, 0, end-1)
	}
}

func siftDownInterface(a sort.Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && a.Less(child, child+1) {
			child++
		}
		if !a.Less(root, child) {
			return
		}
		a.Swap(root, child)
		root = child
	}
}

func siftDownInts(a []int, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && (a[child] < a[child+1]) {
			child++
		}
		if !(a[root] < a[child]) {
			return
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}