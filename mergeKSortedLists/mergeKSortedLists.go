package main

import (
	"container/heap"
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

// we implement heap.Interface that extends sort.Interface with Push() and Pop() methods

type PriorityQueue struct {
	items []*ListNode
	less  func(a, b *ListNode) bool // comparator function, returns boolean, not integer!
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
	pq.items = append(pq.items, x.(*ListNode))
}

// implementation of heap.Interface
func (pq *PriorityQueue) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

func mergeKLists(lists []*ListNode) *ListNode {
	// k: = len(lists)
	dummyHead := &ListNode{Val: -666, Next: nil}
	currentNode := dummyHead

	pq := &PriorityQueue{
		less: func(a, b *ListNode) bool {
			// comparator comparing the values of the nodes, for min-heap
			return a.Val < b.Val
		},
	}

	// todo: we can use array and heapify to make this faster than O(k * log k)
	// append heads of all the lists
	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}
	}

	for pq.Len() > 0 {
		// append the top to the result list
		// heap.Pop(pq) does return the smallest element
		// However, pq.Pop does NOT return the smallest element
		smallestNode := heap.Pop(pq).(*ListNode)

		currentNode.Next = smallestNode
		currentNode = currentNode.Next

		// add smallestNode.Next to the heap, if it exists
		if smallestNode.Next != nil {
			heap.Push(pq, smallestNode.Next)
		}
	}

	// skip dummy head
	return dummyHead.Next
}

func testPriorityQueue() {
	pq := &PriorityQueue{
		less: func(a, b *ListNode) bool {
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

func test(arrays ...[]int) {
	fmt.Println()
	fmt.Println("===================")

	lists := listsCommon.ArraysToLists(arrays...) // note the ... trick, without it arrays will have type [][]int

	fmt.Println("Initial lists:")
	for _, list := range lists {
		listsCommon.PrintList(list)
	}

	mergedList := mergeKLists(lists)

	fmt.Println("Merged list: ")
	listsCommon.PrintList(mergedList)
}

func test1() {
	a1 := []int{1, 4, 5}
	a2 := []int{1, 3, 4}
	a3 := []int{2, 6}

	test(a1, a2, a3)
}

func main() {
	// 23. Merge k Sorted Lists
	// testPriorityQueue()
	test1()
}
