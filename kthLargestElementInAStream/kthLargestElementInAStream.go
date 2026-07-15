package main

import (
	"container/heap"
	"fmt"
)

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

// implementation of the ticket
type KthLargest struct {
	pq *PriorityQueue
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	pq := &PriorityQueue{
		less: func(a, b int) bool {
			// comparator comparing the values of the nodes, for min-heap
			return a < b
		},
	}

	kthLargest := KthLargest{pq, k}

	/*
		// adding n elements by one will be O(n * log k)
		for _, v := range nums {
			kthLargest.Add(v)
		}
	*/

	// this should be faster - first heapify and then remove the first elements. Heapify is O(n)
	pq.items = nums
	fmt.Printf("Non-heapified array: %v \n", pq.items)

	heap.Init(pq) // heapify the pq
	fmt.Printf("Array after heapify: %v \n", pq.items)

	// remove the minimum elements until we have just k elements
	// this should be O( (n -k) * log n )
	for pq.Len() > k {
		heap.Pop(pq)
	}

	fmt.Printf("Heap array after reducing to K = %v values: %v \n", k, pq.items)

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.pq.Len() >= this.k && val <= this.pq.Peek() {
		// there are already K elements, and we're trying to add an element that is fewer than min
		// -> do nothing, just return the current minimum
		return this.pq.Peek()
	}

	heap.Push(this.pq, val) // pq must be a pointer!

	// if we have more than K values -> drop the minimum
	if this.pq.Len() > this.k { // if we exceeded the k -> remove the min value
		heap.Pop(this.pq)
	}

	// there is no separate function in heap.Interface to return the top
	return this.pq.Peek()
}

func executeAdd(kthLargest KthLargest, val int) {
	result := kthLargest.Add(val)

	fmt.Printf("Added value %d. %v-th largest element: %v \n", val, kthLargest.k, result)
}

func test1() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	executeAdd(kthLargest, 3)
	executeAdd(kthLargest, 3)
	executeAdd(kthLargest, 5)
	executeAdd(kthLargest, 10)
	executeAdd(kthLargest, 9)
	executeAdd(kthLargest, 4)
}

func test2() {
	kthLargest := Constructor(4, []int{7, 7, 7, 7, 8, 3})
	executeAdd(kthLargest, 2)
	executeAdd(kthLargest, 10)
	executeAdd(kthLargest, 9)
	executeAdd(kthLargest, 9)
}

func main() {
	// 703. Kth Largest Element in a Stream
	test1()
	test2()
}
