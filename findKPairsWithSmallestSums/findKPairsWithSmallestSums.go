package main

import (
	"container/heap"
	"fmt"
	"math"
)

type ElementsPair struct {
	sum int
	i   int // index in array 1
	j   int // index in array 2
	// we're not storing a1[i] and a2[j] to save memory
}

type PriorityQueue struct {
	items []ElementsPair
	less  func(a, b ElementsPair) bool // comparator function, returns boolean, not integer!
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
	pq.items = append(pq.items, x.(ElementsPair))
}

// implementation of heap.Interface
func (pq *PriorityQueue) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	return kSmallestPairs_heap_optimized(nums1, nums2, k) // minimized heap usage, no visited tracking required, auto-exclude duplicates
	//return kSmallestPairs_heap_obvious(nums1, nums2, k) // obvious heap usage
	//return kSmallestPairs_bad(nums1, nums2, k) // incorrect 2 pointers solution
}

func kSmallestPairs_heap_optimized(nums1 []int, nums2 []int, k int) [][]int {
	// matrix where row = a1[i], columns = a2[j]
	// because of both arrays are sorted, all rows and columns are sorted non-decreasing

	pq := &PriorityQueue{
		less: func(a, b ElementsPair) bool {
			// min heap by sum
			return a.sum < b.sum
		},
	}

	rowsCount := min(k, len(nums1))

	// add all pairs a1[i], a2[0] pairs
	// add not more than K rows since we only need K pairs and next rows will be bigger
	for i := range rowsCount {
		heap.Push(pq, ElementsPair{nums1[i] + nums2[0], i, 0})
	}

	result := make([][]int, 0)

	for len(result) < k {
		smallest := heap.Pop(pq).(ElementsPair)
		i := smallest.i
		j := smallest.j

		result = append(result, []int{nums1[smallest.i], nums2[smallest.j]})

		// For the popped element, take the next from the same row, i.e. [i][j + 1] (if it exists).
		// It's smaller than [i + 1][j + 1]
		// For the next rows, we already have smaller candidates in the heap
		if (j + 1) < len(nums2) {
			heap.Push(pq, ElementsPair{nums1[i] + nums2[j+1], i, j + 1})
		}
	}

	return result
}

func kSmallestPairs_heap_obvious(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{
		less: func(a, b ElementsPair) bool {
			// min heap by sum
			return a.sum < b.sum
		},
	}

	// used elements to avoid duplicates
	m := make(map[ElementsPair]bool)

	// put [0][0] as the smallest pair
	firstSmallest := ElementsPair{nums1[0] + nums2[0], 0, 0}

	heap.Push(pq, firstSmallest)
	m[firstSmallest] = true

	result := make([][]int, 0)

	for len(result) < k {
		smallest := heap.Pop(pq).(ElementsPair)
		i := smallest.i
		j := smallest.j

		result = append(result, []int{nums1[smallest.i], nums2[smallest.j]})

		// add next 2 possible smallest elements -> [i + 1, j] and [i, j + 1] (to bottom and to right in the matrix)
		// if they exist in the array and these indexes are not yet visited
		if (i + 1) < len(nums1) {
			toBottom := ElementsPair{nums1[i+1] + nums2[j], i + 1, j}

			if !m[toBottom] {
				heap.Push(pq, ElementsPair{nums1[i+1] + nums2[j], i + 1, j})
				m[toBottom] = true
			}
		}

		if (j + 1) < len(nums2) {
			toRight := ElementsPair{nums1[i] + nums2[j+1], i, j + 1}

			if !m[toRight] {
				heap.Push(pq, ElementsPair{nums1[i] + nums2[j+1], i, j + 1})
				m[toRight] = true
			}
		}
	}

	return result
}

func kSmallestPairs_bad(nums1 []int, nums2 []int, k int) [][]int {
	i := 0
	j := 0

	result := make([][]int, 0)
	result = append(result, []int{nums1[0], nums2[0]})

	for len(result) < k {
		nextLeft := math.MinInt32
		if i < len(nums1)-1 {
			nextLeft = nums1[i+1] + nums2[j]
		}

		nextRight := math.MinInt32
		if j < len(nums2)-1 {
			nextRight = nums1[i] + nums2[j+1]
		}

		if nextLeft < nextRight {
			result = append(result, []int{nums1[i+1], nums2[j]})
			i++
		} else {
			result = append(result, []int{nums1[i], nums2[j+1]})
			j++
		}
	}

	return result
}

func test(a1, a2 []int, k int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array 1: %v \n", a1)
	fmt.Printf("Array 2: %v \n", a1)
	fmt.Printf("K (count of smallest pairs): %v \n", k) // will always be 0

	result := kSmallestPairs(a1, a2, k)

	fmt.Printf("%v smallest pairs: \n%v \n", k, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := range expectedResult {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] {
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

func test1() {
	a1 := []int{1, 7, 11}
	a2 := []int{2, 4, 6}
	k := 3

	expected := [][]int{
		{1, 2},
		{1, 4},
		{1, 6},
	}

	test(a1, a2, k, expected)
}

func test2() {
	a1 := []int{1, 1, 2}
	a2 := []int{1, 2, 3}
	k := 2

	expected := [][]int{
		{1, 1},
		{1, 1},
	}

	test(a1, a2, k, expected)
}

func test3() {
	a1 := []int{1, 2, 4, 5, 6}
	a2 := []int{3, 5, 7, 9}
	k := 3

	expected := [][]int{
		{1, 3},
		{2, 3},
		{1, 5},
	}

	test(a1, a2, k, expected)
}

func test4() {
	a1 := []int{1, 2, 4, 5, 6}
	a2 := []int{3, 5, 7, 9}
	k := 20

	// todo: the ordering or the algorithms will be different for the same elements
	expected := [][]int{
		{1, 3},
		{2, 3},
		{1, 5},
		{2, 5},
		{4, 3},
		{1, 7},
		{5, 3},
		{2, 7},
		{4, 5},
		{6, 3},
		{1, 9},
		{5, 5},
		{2, 9},
		{4, 7},
		{6, 5},
		{5, 7},
		{4, 9},
		{6, 7},
		{5, 9},
		{6, 9},
	}

	test(a1, a2, k, expected)
}

func main() {
	// 373. Find K Pairs with Smallest Sums
	test1()
	test2()
	test3()
	test4()
}
