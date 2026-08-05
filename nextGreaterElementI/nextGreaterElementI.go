package main

import (
	"container/list"
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// we calculate the map of (value -> next greater element in nums2).
	// -1 if there is no next greater element.
	// then for every element in nums1, we just look up the value from the map
	m := make(map[int]int)

	// monotonic stack - top is less than next etc.
	stack := list.New()

	for i := len(nums2) - 1; i >= 0; i-- {
		v := nums2[i]

		fmt.Println()
		fmt.Printf("nums2[%v] = %v \n", i, v)

		fmt.Println("Stack: ")
		printStack(stack)

		// remove all values <= current element
		// These values are worse than the current for all the elements to the left of the current since:
		// - current element is greater
		// - current element is more at the left

		// Values bigger than the current value remain in the stack since they can be greater element
		// for the values at the left that are bigger than current.
		for (stack.Len() > 0) && (getStackTop(stack) <= v) {
			removeFromStack(stack)
		}

		fmt.Printf("Removed all values <= %v from the stack: \n", v)
		printStack(stack)

		// top of the stack is the smallest and nearest candidate AFTER the current element
		// that is bigger than the current element (elements smaller than the current we removed above)
		// (since we're going from right to left)
		if (stack.Len() > 0) && (getStackTop(stack) > v) {
			m[v] = getStackTop(stack)
		} else { // no next greater element
			m[v] = -1
		}

		// Push the current element to the stack.
		// All the remaining values (if present) are bigger than the current element,
		// So the stack remains strictly increasing from the top.
		// The current element still can be a valid "next greater" for the values at the left that are smaller.
		pushToStack(stack, v)

		fmt.Printf("Pushed current value %v to the stack: \n", v)
		printStack(stack)
	}

	// todo: we can modify nums1 itself if we want no additional space
	// fill result for every value in nums1
	result := make([]int, len(nums1))

	for i, v := range nums1 {
		result[i] = m[v]
	}

	return result
}

func pushToStack(stack *list.List, v int) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStack(stack *list.List) int { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(int)
}

func getStackTop(stack *list.List) int { // only called when stack is not empty
	return stack.Front().Value.(int)
}

func printStack(l *list.List) {
	fmt.Printf("[ ")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Printf("]")

	fmt.Println()
}

func test(a1 []int, a2 []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array 1 (elements to search): %v \n", a1)
	fmt.Printf("Array 2 (defines next greater elements): %v \n", a2)

	result := nextGreaterElement(a1, a2)

	fmt.Printf("Next greater elements in array 2 for every element of array 1: %v \n", result)
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
	a1 := []int{4, 1, 2}
	a2 := []int{1, 3, 4, 2}
	expectedResult := []int{-1, 3, -1}

	test(a1, a2, expectedResult)
}

func test2() {
	a1 := []int{2, 4}
	a2 := []int{1, 2, 3, 4}
	expectedResult := []int{3, -1}

	test(a1, a2, expectedResult)
}

func main() {
	// 496. Next Greater Element I
	test1()
	test2()
}
