package main

import "fmt"

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	// this fails TLE on 949 / 953
	// 50000 increasing values.
	// 50000 queries with [0, 1]
	// So we can cache the same values

	// after adding a cache, it's failing TLE on 951 / 953
	// 50000 increasing values,
	// 50000 distinct queries with [0, 2], [0, 3], [0, 4] etc
	return leftmostBuildingQueries_monotonicStack_naive(heights, queries)
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

func main() {
	// 2940. Find Building Where Alice and Bob Can Meet
	test1()
	test2()
}
