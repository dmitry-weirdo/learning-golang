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

func isPossible(target []int) bool {
	return isPossible_optimized(target) // brute-force solution, fails on big values of K, i.e. [1, 1_000_000_000]
	//return isPossible_bruteforce(target) // brute-force solution, fails on big values of K, i.e. [1, 1_000_000_000]
}

func isPossible_optimized(target []int) bool {
	// we extract the biggest element, since its increase should happen the last
	// biggest = arraySum on previous step
	// We neecd to calculate the previous value in the place of biggest (oldBiggest)

	// prevSum = (currentSum - biggest) + oldBiggest
	// biggest = prevSum
	// prevSum = curSum - prevSum + oldBiggest
	// oldBiggest = 2 * prevSum - curSum

	// x = oldBiggest
	// rest = currentSum - biggest
	// prevSum = rest + x
	// x = prevSum - rest
	// x = biggest - rest

	// since we're targeting x > 1,
	// biggest > rest

	// biggest will decrease by rest until x < rest
	// So we can instead of x = biggest - rest do x = biggest % rest

	// poll biggest
	// push oldBiggest
	// set currentSum = prevSum

	if len(target) == 1 { // corner-case of 1 element -> it will only work if the target = [1]
		return target[0] == 1
	}

	// to find the biggest element -> use max-heap
	pq := &PriorityQueue{
		less: func(a, b int) bool {
			// max heap since we need the current max
			return a > b
		},
	}

	// heapify - O(n)
	pq.items = target
	heap.Init(pq)

	sum := 0
	for _, v := range target {
		sum += v
	}

	for pq.Peek() > 1 { // we will not decrease the elements in the heap, we're just replacing biggest value with the smaller value
		biggest := heap.Pop(pq).(int)

		rest := sum - biggest

		if rest == 1 { // reached 1, this will only happen if rest is 1 element, i.e. there are n = 2 elements in the array
			return true
		}

		//oldValueOfBiggest := 2*biggest - sum
		//oldValueOfBiggest := biggest - rest
		oldValueOfBiggest := biggest % rest // !!! optimization for the quick decrease of oldBiggest

		if (oldValueOfBiggest < 1) || // we decreased to less than 1 -> fail, impossible to deconstruct
			(oldValueOfBiggest == biggest) { // biggest is replaced with the same modulus infinitely -> impossible to decrease -> return
			return false
		}

		// for the previous step, instead of biggest we had oldBiggest
		heap.Push(pq, oldValueOfBiggest)

		// sum is decreased on how much we dropped from the biggest
		sum = sum - (biggest - oldValueOfBiggest)
	}

	// if we reached 1,
	return pq.Peek() == 1
}

func isPossible_bruteforce(target []int) bool {
	// we extract the biggest element, since its increase should happen the last
	// biggest = arraySum on previous step
	// We neecd to calculate the previous value in the place of biggest (oldBiggest)

	// prevSum = (currentSum - biggest) + oldBiggest
	// biggest = prevSum
	// prevSum = curSum - prevSum + oldBiggest
	// oldBiggest = 2 * prevSum - curSum

	// poll biggest
	// push oldBiggest
	// set currentSum = prevSum

	if len(target) == 1 { // corner-case of 1 element -> it will only work if the target = [1]
		return target[0] == 1
	}

	// to find the biggest element -> use max-heap
	pq := &PriorityQueue{
		less: func(a, b int) bool {
			// max heap since we need the current max
			return a > b
		},
	}

	// heapify - O(n)
	pq.items = target
	heap.Init(pq)

	sum := 0
	for _, v := range target {
		sum += v
	}

	for pq.Peek() > 1 { // we will not decrease the elements in the heap, we're just replacing biggest value with the smaller value
		biggest := heap.Pop(pq).(int)

		oldValueOfBiggest := 2*biggest - sum

		if oldValueOfBiggest < 1 { // we decreased to less than 1 -> fail, impossible to deconstruct
			return false
		}

		// for the previous step, instead of biggest we had oldBiggest
		heap.Push(pq, oldValueOfBiggest)

		// on the previous step, the sum of the array was equal to the current biggest
		sum = biggest
	}

	// if we reached 1,
	return pq.Peek() == 1
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := isPossible(arr)

	fmt.Printf("Possible to decrease to array of 1-s: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{9, 3, 5}
	expected := true

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1, 1, 2}
	expected := false

	test(arr, expected)
}

func test3() {
	arr := []int{8, 5}
	expected := true

	test(arr, expected)
}

func test4() {
	arr := []int{1, 1_000_000_000} // 10^9 will cause TLE in case of brute-force algorithm
	expected := true

	test(arr, expected)
}

func test5() {
	arr := []int{2} // for 1 element, it will always stay 1
	expected := false

	test(arr, expected)
}

func test6() {
	arr := []int{1} // for 1 element, it will always stay 1
	expected := true

	test(arr, expected)
}

func main() {
	// 1354. Construct Target Array With Multiple Sums
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
