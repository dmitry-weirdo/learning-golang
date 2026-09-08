package main

import "fmt"

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	// O(log N) for every query.
	// The trick is: while calculating the monotonic stack, for every index [i],
	// the stack contains the DECREASING array of bigger values after [i].
	// It means that we can binary-search on this array for all the queries that start at index [i].

	// This passes in 85-110 ms
	return leftmostBuildingQueries_monotonicStack_optimized(heights, queries)

	// this fails TLE on 949 / 953
	// 50000 increasing values.
	// 50000 queries with [0, 1]
	// So we can cache the same values

	// after adding a cache, it's failing TLE on 951 / 953
	// 50000 increasing values,
	// 50000 distinct queries with [0, 2], [0, 3], [0, 4] etc
	//return leftmostBuildingQueries_monotonicStack_naive(heights, queries)
}

type Query struct {
	queryIndex int // index in result. We set the found index to result[i]
	height     int // we're searching for heights[j] > height
}

func leftmostBuildingQueries_monotonicStack_optimized(heights []int, queries [][]int) []int {
	result := make([]int, len(queries))

	// key: start index [i] to search.
	// value: list of values height[left] to search the next greater indexes after [i]
	m := make(map[int][]Query)

	for i, q := range queries {
		// sort query indexes to be left and right
		left, right := getMinAndMax(q[0], q[1])

		if left == right {
			result[i] = left
		} else if heights[left] < heights[right] { // can jump directly from left to right
			result[i] = right
		} else {
			// for the queries where we need the monotonic stack logic (find next greater element),
			// we calculate queries for every index [i], to find the values > height[left]
			if _, ok := m[right]; !ok {
				m[right] = []Query{}
			}

			// for position[right], we're searching for values > heights[left]
			// and set the found index to result[i]
			query := Query{queryIndex: i, height: heights[left]}
			m[right] = append(m[right], query)
		}
	}

	//fmt.Printf("Queries to get by monotonic stack: %v \n", m)

	GetNextGreaterWithIndexesOptimized(heights, -1, m, result)

	return result
}

func GetNextGreaterWithIndexesOptimized(a []int, noElementValue int, q map[int][]Query, r []int) {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := createStackWithIndex()

	n := len(a)
	//result := make([]MatchingElement, n) // we're not calculating indexes for all elements

	for i := n - 1; i >= 0; i-- {
		// this is the usual monotonic stack code
		v := a[i]

		for stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value <= v) {
			removeFromStackWithIndex(stack)
		}

		// =======================================================
		// Query execution on monotonic stack.

		// This can be done after the remove step, this should decrease the search array.
		// Checked - it doesn't significantly improve the speed.
		// But ok, let's leave this in place.

		// !!! The main trick is that the current state of the stack contains all the values after the current index[i] that are >
		// , and the values are in DECREASING order, so we can search within the stack using the binary search.

		// Execute all queries for the index [i].
		if q[i] != nil && len(q[i]) > 0 {
			for _, query := range q[i] {
				// stack is decreasing, not increasing.
				// So we're searching for the first <= element and then trying to go left

				condition := func(e MatchingElement) bool {
					return e.value <= query.height // we're searching for values <= query.height
				}

				index := binarySearchGeneric(
					*stack,
					0,           // we search the complete stack
					len(*stack), // search insert position. If not found -> result will be after the end of the array
					condition,
				)

				// set the index from the original array
				// index is in MatchingElement.index
				if index == 0 { // no element to the left
					r[query.queryIndex] = -1
				} else {
					// we find the leftmost <= element -> go one left anc check whether it is >
					index--

					if (*stack)[index].value <= query.height {
						r[query.queryIndex] = -1
					} else {
						r[query.queryIndex] = (*stack)[index].index

					}
				}

				//fmt.Printf("Stack: %v \n", stack)
				//fmt.Printf("Result[%v] set to %v. \n", query.queryIndex, r[query.queryIndex])
			}
		}

		// next is the logic of the normal monotonic stack

		/* // we're not calculating the results for this task
		if stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value > v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next greater element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}
		*/

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	//return result
}

func binarySearchGeneric(
	arr []MatchingElement,
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func(MatchingElement) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// todo: this method can return an incorrect value for the empty array

	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func leftmostBuildingQueries_monotonicStack_naive(heights []int, queries [][]int) []int {
	nextSmaller := GetNextGreaterWithIndexes(heights, -1)

	result := make([]int, len(queries))

	cache := make(map[int]map[int]int)

	for i, q := range queries {
		// sort query indexes to be left and right
		left, right := getMinAndMax(q[0], q[1])

		if leftMap, okLeft := cache[left]; okLeft {
			if v, okRight := leftMap[right]; okRight {
				result[i] = v
				continue
			}
		}

		index := -1
		if left == right {
			index = left
		} else if heights[left] < heights[right] { // can jump directly from left to right
			index = right
		} else {
			// todo: there could be a binary-search magic to speed up this search
			// from right, go to next greater until it's greater then heights[left] or we reach the end of the array
			// !!! This is potentially a O(N) operation that is slow
			index = right

			for heights[index] <= heights[left] {
				index = nextSmaller[index].index

				if index == -1 { // no next bigger or reached the end of the array -> no solution
					break
				}
			}
		}

		result[i] = index

		if _, ok := cache[left]; !ok {
			cache[left] = make(map[int]int)
		}

		cache[left][right] = index
	}

	return result
}

func getMinAndMax(a, b int) (smaller, greater int) {
	if a <= b {
		return a, b
	}

	return b, a
}

type MatchingElement struct {
	value int
	index int // index of value in the array
}

func GetNextGreaterWithIndexes(a []int, noElementValue int) []MatchingElement {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := createStackWithIndex()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value <= v) {
			removeFromStackWithIndex(stack)
		}

		if stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value > v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next greater element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

// ======== stack of MatchingElement values ====== //
func createStackWithIndex() *[]MatchingElement {
	stack := make([]MatchingElement, 0)
	return &stack
}

func stackIsNotEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) > 0
}

func stackIsEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) == 0
}

func pushToStackWithIndex(stack *[]MatchingElement, v MatchingElement) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStackWithIndex(stack *[]MatchingElement) MatchingElement { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTopWithIndex(stack *[]MatchingElement) MatchingElement { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

func test(h []int, q [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Heights: %v \n", h)
	fmt.Printf("Queries: %v \n", q)

	result := leftmostBuildingQueries(h, q)

	fmt.Printf("Leftmost buildings to meet: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		[]int{6, 4, 8, 5, 2, 7},
		[][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}},
		[]int{2, 5, -1, 5, 2},
	)
}

func test2() {
	test(
		[]int{5, 3, 8, 2, 6, 1, 4, 6},
		[][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}},
		[]int{7, 6, -1, 4, 6},
	)
}

func test3() {
	test(
		[]int{1, 2, 3, 4},
		[][]int{{0, 1}, {0, 2}},
		[]int{1, 2},
	)
}

func main() {
	// 2940. Find Building Where Alice and Bob Can Meet
	test1()
	test2()
	test3()
}
