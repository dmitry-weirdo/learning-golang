package heapCommon

import (
	"container/heap"
	"fmt"
)

// todo: implement a template PriorityQueue, it's not always int in the head (e.g. for Dijkstra, we need (distance + node) pairs

func heapCheatSheet() {
	pq := createMinHeap()

	// push to heap
	heap.Push(pq, 100)
	heap.Push(pq, 10)

	// pop from heap - returns any, so we need to case
	for pq.Len() > 0 {
		minValue := heap.Pop(pq)
		fmt.Printf("Popped Heap min value in O(log N): %v \n", minValue)
	}
}

func createMinHeap() *PriorityQueue {
	return &PriorityQueue{
		less: func(a, b int) bool {
			// min heap
			return a < b
		},
	}
}

func createMinHeapWithValues(values []int) *PriorityQueue {
	pq := createMinHeap()

	// heapify - O(n)
	pq.items = values
	heap.Init(pq)

	return pq
}

func createMaxHeap() *PriorityQueue {
	return &PriorityQueue{
		less: func(a, b int) bool {
			// max heap
			return a > b
		},
	}
}

func createMaxHeapWithValues(values []int) *PriorityQueue {
	pq := createMaxHeap()

	// heapify - O(n)
	pq.items = values
	heap.Init(pq)

	return pq
}

type PriorityQueue struct {
	items []int
	less  func(a, b int) bool // comparator function, returns boolean, not integer!
}

// implementation of sort.Interface
func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

// implementation of sort.Interface
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

// implementation of sort.Interface
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// implementation of heap.Interface
func (pq *PriorityQueue) Push(x any) { // interface needs `x any`, else the override will not work
	pq.items = append(pq.items, x.(int))
}

// implementation of heap.Interface
func (pq *PriorityQueue) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

// helper function -> get the top of the heap without removing it
func (pq *PriorityQueue) Peek() int {
	return pq.items[0] // return the root
}
