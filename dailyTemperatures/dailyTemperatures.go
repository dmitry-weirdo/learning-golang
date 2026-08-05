package main

import (
	"container/list"
	"fmt"
)

type GreaterElement struct {
	value int
	index int // index of value in the array
}

func dailyTemperatures(arr []int) []int {
	n := len(arr)

	// todo: we can modify nums1 itself if we want no additional space
	result := make([]int, n)

	// monotonic stack - top is less than next etc.
	// We're storing pairs of (value, index).
	// todo: we CAN only store indexes and perform all value comparisons on arr[index], but this makes code uglier and doesn't save space drastically, since O(2n) = O(n)
	stack := list.New()

	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]

		fmt.Println()
		fmt.Printf("arr[%v] = %v \n", i, v)

		fmt.Println("Stack: ")
		printStack(stack)

		// remove all values <= current element
		// These values are worse than the current for all the elements to the left of the current since:
		// - current element is greater
		// - current element is more at the left

		// Values bigger than the current value remain in the stack since they can be greater element
		// for the values at the left that are bigger than current.
		for (stack.Len() > 0) && (getStackTop(stack).value <= v) {
			removeFromStack(stack)
		}

		fmt.Printf("Removed all values <= %v from the stack: \n", v)
		printStack(stack)

		// top of the stack is the smallest and nearest candidate AFTER the current element
		// that is bigger than the current element (elements smaller than the current we removed above)
		// (since we're going from right to left)
		if (stack.Len() > 0) && (getStackTop(stack).value > v) {
			// we get diff of indices: nextGreaterElementIndex - currentElementIndex
			result[i] = getStackTop(stack).index - i
		} else { // no next greater element
			result[i] = 0
		}

		// Push the current element to the stack.
		// All the remaining values (if present) are bigger than the current element,
		// So the stack remains strictly increasing from the top.
		// The current element still can be a valid "next greater" for the values at the left that are smaller.
		greaterElement := GreaterElement{value: v, index: i}
		pushToStack(stack, greaterElement)

		fmt.Printf("Pushed current value %v to the stack: \n", greaterElement)
		printStack(stack)
	}

	return result
}

func pushToStack(stack *list.List, v GreaterElement) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStack(stack *list.List) GreaterElement { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(GreaterElement)
}

func getStackTop(stack *list.List) GreaterElement { // only called when stack is not empty
	return stack.Front().Value.(GreaterElement)
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

	fmt.Printf("Array of temperatures: %v \n", arr)

	result := dailyTemperatures(arr)

	fmt.Printf("Indexes to next greater elements: %v \n", result)
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
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expectedResult := []int{1, 1, 4, 2, 1, 1, 0, 0}

	test(arr, expectedResult)
}

func test2() {
	arr := []int{30, 40, 50, 60}
	expectedResult := []int{1, 1, 1, 0}

	test(arr, expectedResult)
}

func main() {
	// 739. Daily Temperatures
	// This is still a modification of "496. Next Greater Element I", but we need to also keep track of the indices of the next greater elements.
	test1()
	test2()
}
