package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// we implement heap.Interface that extends sort.Interface with Push() and Pop() methods

type PriorityQueue struct {
	items []ListNode               // todo: maybe we need pointers
	less  func(a, b ListNode) bool // comparator function, returns boolean, not integer!
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
	pq.items = append(pq.items, x.(ListNode))
}

// implementation of heap.Interface
func (pq *PriorityQueue) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

func mergeKLists(lists []*ListNode) *ListNode {
	// todo: implement method
	return nil
}

func testPriorityQueue() {
	pq := &PriorityQueue{
		less: func(a, b ListNode) bool {
			// comparator comparing the values of the nodes, for min-heap
			return a.Val < b.Val
		},
	}

	// here the compilation will fail if we haven't implemented any methods of heap.Interface
	heap.Init(pq)

	heap.Push(pq, ListNode{666, nil})
	heap.Push(pq, ListNode{1000, nil})
	heap.Push(pq, ListNode{3, nil})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(ListNode)
		fmt.Printf("Popped node: %d \n", item.Val)
	}
}

func main() {
	testPriorityQueue()
}
