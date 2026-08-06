package main

import (
	"container/list"
	"fmt"
)

func validSubarrays(nums []int) int {
	// For every subarray:
	// a[0] <= a[i], i > 0
	// Therefore, the breaking point will be the first element that is smaller than the first in the subarray.
	// i.e., we're searching for "next smallest" (its index).
	// If there is no next smallest, we count its index as N (after the end of the array).

	// Then we count subarrays starting from every position of the array.
	// It will be (nextSmallestIndex[i] - i)

	// Example 1: [1 2]
	// i = 0, right[0] = N = 2 (after the end of the array)
	// Count of subarrays starting from a[0]: 2 - 0
	// It will be subarrays [1] and [1, 2]

	// Example 2: [2 1]
	// i = 0, right[0] = 1
	// Count of subarrays starting from a[0]: 1 - 0 = 1
	// It will be subarray [2]

	// Then we just sum up values for all the starting positions in the array

	nextSmallerElements := GetNextSmallerWithIndexes(nums, -1, len(nums))
	fmt.Printf("Next smaller elements: %v \n", nextSmallerElements)

	// converting OnlyIndexes will take O(N) again, so let's just use nextSmallerElements directly

	result := 0

	for i := range nums {
		result += nextSmallerElements[i].index - i
	}

	return result
}

type MatchingElement struct {
	value int
	index int // index of value in the array
}

func OnlyIndexes(a []MatchingElement) []int {
	result := make([]int, len(a))

	for i := range a {
		result[i] = a[i].index
	}

	return result
}

// todo: option with noElementIndex must be moved also to monotonicStack.go
func GetNextSmallerWithIndexes(a []int, noElementValue int, noElementIndex int) []MatchingElement {
	// direction: right -> left
	// stack: decreasing from top to bottom
	// removal from stack: >= current value
	// select top as result: if < current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTopWithIndex(stack).value >= v) {
			removeFromStackWithIndex(stack)
		}

		if (stack.Len() > 0) && (getStackTopWithIndex(stack).value < v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no next smaller element
			result[i] = MatchingElement{value: noElementValue, index: noElementIndex} // should default to len(nums) + 1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

func pushToStackWithIndex(stack *list.List, v MatchingElement) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStackWithIndex(stack *list.List) MatchingElement { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(MatchingElement)
}

func getStackTopWithIndex(stack *list.List) MatchingElement { // only called when stack is not empty
	return stack.Front().Value.(MatchingElement)
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Heights: %v \n", arr)

	result := validSubarrays(arr)

	fmt.Printf("Total valid subarrays: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{1, 4, 2, 5, 3}
	expected := 11 // [1], [1,4], [1,4,2], [1, 4, 2, 5], [1, 4, 2, 5, 3], [4], [2], [2, 5], [2, 5, 3], [5], [3]

	test(arr, expected)
}

func test2() {
	arr := []int{3, 2, 1}
	expected := 3 // just single-element subarrays: [3], [2], [1]

	test(arr, expected)
}

func test3() {
	arr := []int{2, 2, 2}
	expected := 6 // !!! we don't exclude the same elements arrays starting with different indexes, so the answer will be: [2], [2, 2], [2, 2, 2], [2], [2, 2], [2]

	test(arr, expected)

}

func main() {
	// 1063. Number of Valid Subarrays
	test1()
	test2()
	test3()
}
