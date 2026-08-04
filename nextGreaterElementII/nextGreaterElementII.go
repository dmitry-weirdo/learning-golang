package main

import (
	"container/list"
	"fmt"
)

func nextGreaterElements(arr []int) []int {
	n := len(arr)

	// todo: we can modify nums1 itself if we want no additional space
	result := make([]int, n)

	// monotonic stack - top is strictly less than next etc.
	stack := list.New()

	// To emulate the circularity, we just append the array to itself at the right.
	// Going circular just once is enough, since we only need the positions before the current element,
	// and the array part after the current element is already handled by the non-circular handing.

	// We get the array of size 2 * n.
	// And to collect the result, we just need the left part of the array, i.e. [0; n - 1]

	for i := 2*n - 1; i >= 0; i-- {
		// We don't copy the array to save memory, we're just reusing the same array twice.
		// E.g., for an array of size 3, we'll go down from index 5, and for i = 5 we'll use a[5 % 3] = a[2]
		v := arr[i%n]

		fmt.Println()
		fmt.Printf("nums[%v] = %v \n", i, v)

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

		// !!! we only collect the result if we're in the left part of the array
		if i < n {
			if (stack.Len() > 0) && (getStackTop(stack) > v) {
				result[i] = getStackTop(stack)
			} else { // no next greater element
				result[i] = -1 // it's weird that -1 can also be a valid "next greater" element (since the element range for this ticket are -10^9...10^9)
			}
		}

		// Push the current element to the stack.
		// All the remaining values (if present) are bigger than the current element,
		// So the stack remains strictly increasing from the top.
		// The current element still can be a valid "next greater" for the values at the left that are smaller.
		pushToStack(stack, v)

		fmt.Printf("Pushed current value %v to the stack: \n", v)
		printStack(stack)
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

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array 1 (can be handled circularly): %v \n", arr)

	result := nextGreaterElements(arr)

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
	arr := []int{1, 2, 1}
	expectedResult := []int{2, -1, 2}

	test(arr, expectedResult)
}

func test2() {
	arr := []int{1, 2, 3, 4, 3}
	expectedResult := []int{2, 3, 4, -1, 4}

	test(arr, expectedResult)
}

func main() {
	// 503. Next Greater Element II
	// It is the same algorithm as "496. Next Greater Element I", we just append the array to itself to simulate the circular array.
	test1()
	test2()
}
