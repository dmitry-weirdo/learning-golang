package main

import (
	"demo/listsCommon"
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func nextLargerNodes(head *ListNode) []int {
	// we convert list to array, and then it's a nextGreaterElement task
	arr := listToArray(head)

	return GetNextGreater(arr, 0) // if no nextElement -> return 0
}

func listToArray(head *ListNode) []int {
	n := head

	result := make([]int, 0)

	for n != nil {
		result = append(result, n.Val)

		n = n.Next
	}

	return result
}

func GetNextGreater(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := createStack()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmpty(stack) && (getStackTop(stack) <= v) {
			removeFromStack(stack)
		}

		if stackIsNotEmpty(stack) && (getStackTop(stack) > v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

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

func test(arr []int, expectedResult []int) { // linked list to linked list
	fmt.Println()
	fmt.Println("========================")

	list := listsCommon.ArrayToList(arr)

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("List from array: \n")
	listsCommon.PrintList(list)

	result := nextLargerNodes(list) // todo: replace with your function

	fmt.Printf("Result: %v \n", result)
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
	arr := []int{2, 1, 5}
	expected := []int{5, 5, 0}

	test(arr, expected)
}

func test2() {
	arr := []int{2, 7, 4, 3, 5}
	expected := []int{7, 0, 5, 5, 0}

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := []int{0}

	test(arr, expected)
}

func test4() {
	arr := []int{}
	expected := []int{}

	test(arr, expected)
}

func main() {
	// 1019. Next Greater Node In Linked List
	// When we convert a list to array, it's the usual "nextGreaterElement" task with noElementValue = 0
	// See e.g. "496. Next Greater Element I"
	test1()
	test2()
	test3()
	test4()
}
