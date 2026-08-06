package main

import (
	"fmt"
	"slices"
)

func findBuildings(heights []int) []int {
	// looks like creating an array of size N is still O(N), so it's better to collect the stack left-to-right and only append the nonMatchingIndexes to the end of the result array instead of allocating a result array of N

	// we can optimize to O(N) instead of O(2 * N) - add to result immediately, but it's not critical since it's still O(N)
	return findBuildings_optimized(heights) // O(N)

	//return findBuildings_trivial(heights) // O(2 * N)
}

func findBuildings_optimized(heights []int) []int {
	// We can skip returning "nextGreaterOrEqual" elements altogether.
	// We need to find indexes of buildings that have no "next greater or equal".
	nonMatchingElementIndexes := GetNextGreaterOrEqualOptimized_2(heights)

	return nonMatchingElementIndexes
}

func findBuildings_trivial(heights []int) []int {
	// we need to find indexes of buildings that have no "next greater or equal"
	blockingHeights := GetNextGreaterOrEqual(heights, -1)

	result := make([]int, 0)

	for i := range heights {
		if blockingHeights[i] == -1 {
			result = append(result, i)
		}
	}

	return result
}

func GetNextGreaterOrEqualOptimized_2(a []int /*noElementValue int,*/) (nonMatchingIndexes []int) {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: < current value
	// select top as result: if >= current value
	// push current value to stack: if top of the stack == current value, replace it with the current element to update the index
	stack := createStack()

	n := len(a)

	// we're appending by 1 and then reversing -> avoid O(N) for allocation
	nonMatchingIndexes = make([]int, 0)

	//result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmpty(stack) && (getStackTop(stack) < v) {
			removeFromStack(stack)
		}

		if stackIsNotEmpty(stack) && (getStackTop(stack) >= v) {
			//result[i] = getStackTop(stack)
		} else { // no next greater element
			//result[i] = noElementValue // should default to -1

			// optimization - collect the result immediately, from the end
			nonMatchingIndexes = append(nonMatchingIndexes, i)
		}

		// if current == stack.top, we replace stack.top with the current element since current element has the new correct index
		if stackIsNotEmpty(stack) && (getStackTop(stack) == v) {
			removeFromStack(stack)
		}

		// we always push the current element since it has the updated index
		pushToStack(stack, v)
	}

	// we collected indexes from right to left -> reverse
	slices.Reverse(nonMatchingIndexes) // in-place
	return nonMatchingIndexes
}

func GetNextGreaterOrEqualOptimized(a []int /*noElementValue int,*/) (nonMatchingIndexes []int) {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: < current value
	// select top as result: if >= current value
	// push current value to stack: if top of the stack == current value, replace it with the current element to update the index
	stack := createStack()

	n := len(a)

	// we're setting a size of N to not push to the beginning and just set by index
	nonMatchingIndexes = make([]int, n)

	//result := make([]int, n)

	nonMatchingElementIndex := len(a) - 1 // we collect from end

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmpty(stack) && (getStackTop(stack) < v) {
			removeFromStack(stack)
		}

		if stackIsNotEmpty(stack) && (getStackTop(stack) >= v) {
			//result[i] = getStackTop(stack)
		} else { // no next greater element
			//result[i] = noElementValue // should default to -1

			// optimization - collect the result immediately, from the end
			nonMatchingIndexes[nonMatchingElementIndex] = i
			nonMatchingElementIndex--
		}

		// if current == stack.top, we replace stack.top with the current element since current element has the new correct index
		if stackIsNotEmpty(stack) && (getStackTop(stack) == v) {
			removeFromStack(stack)
		}

		// we always push the current element since it has the updated index
		pushToStack(stack, v)
	}

	// return only filled subarray within result
	return nonMatchingIndexes[nonMatchingElementIndex+1:]
}

func GetNextGreaterOrEqual(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: < current value
	// select top as result: if >= current value
	// push current value to stack: if top of the stack == current value, replace it with the current element to update the index
	stack := createStack()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmpty(stack) && (getStackTop(stack) < v) {
			removeFromStack(stack)
		}

		if stackIsNotEmpty(stack) && (getStackTop(stack) >= v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

		// if current == stack.top, we replace stack.top with the current element since current element has the new correct index
		if stackIsNotEmpty(stack) && (getStackTop(stack) == v) {
			removeFromStack(stack)
		}

		// we always push the current element since it has the updated index
		pushToStack(stack, v)
	}

	return result
}

func createStack() *[]int {
	stack := make([]int, 0)
	return &stack
}

func stackIsNotEmpty(stack *[]int) bool {
	return len(*stack) > 0
}

func stackIsEmpty(stack *[]int) bool {
	return len(*stack) == 0
}

func pushToStack(stack *[]int, v int) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStack(stack *[]int) int { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTop(stack *[]int) int { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of building heights: %v \n", arr)

	result := findBuildings(arr)

	fmt.Printf("Indexes of building who have no next buildings with >= height: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

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
	arr := []int{4, 2, 3, 1}
	expected := []int{0, 2, 3} // 4, 3, 1

	test(arr, expected)
}

func test2() {
	arr := []int{4, 3, 2, 1}
	expected := []int{0, 2, 3} // 4, 2, 1

	test(arr, expected)
}

func test3() {
	arr := []int{1, 3, 2, 4}
	expected := []int{3} // only last building 4

	test(arr, expected)
}

func main() {
	// 1762. Buildings With an Ocean View
	test1()
	test2()
	test3()
}
