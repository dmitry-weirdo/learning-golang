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

func lastStoneWeight(stones []int) int {
	pq := &PriorityQueue{
		less: func(a, b int) bool {
			// max heap since we need the stones with max weights
			return a > b
		},
	}

	// heapify - O(n)
	pq.items = stones
	heap.Init(pq)

	for pq.Len() > 1 {
		// x >= y
		x := heap.Pop(pq).(int)
		y := heap.Pop(pq).(int)

		if x == y {
			fmt.Printf("2 stones of same weight = %v destroyed. \n", x)
			continue
		}

		newStoneWeight := x - y // x > y
		heap.Push(pq, newStoneWeight)

		fmt.Printf("Pushed the new stone with weight = %v - %v = %v to the heap. \n", x, y, newStoneWeight)
	}

	if pq.Len() <= 0 {
		return 0
	}

	return pq.Peek()
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := lastStoneWeight(arr)

	fmt.Printf("Last stone weight: %v \n", result)
	fmt.Printf("Expected last stone weight: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{2, 7, 4, 1, 8, 1}
	expected := 1

	test(arr, expected)
}

func test2() {
	arr := []int{1}
	expected := 1

	test(arr, expected)
}

func main() {
	// 1046. Last Stone Weight
	test1()
	test2()
}
