package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func kthSmallest(root *TreeNode, k int) int {
	// basically, we do iterative in-order processing
	// iterative to make fast-return instead of going up the recursion

	stack := make([]*TreeNode, 0)

	current := root

	counter := 0

	for current != nil || (len(stack) > 0) {
		// go to the max-left from the current node, pushing nodes to the stack
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop the last value from the stack
		lastIndex := len(stack) - 1
		current = stack[lastIndex]
		stack = stack[0:lastIndex] // remove the last value

		// increase counter
		counter++
		if counter == k { // k-th node reached -> return its value
			return current.Val
		}

		// go to the right of the current node, then the process will continue for this right sub-node
		// do it even if the current.Right == nil, else we'll stay at the current root
		current = current.Right
	}

	// not found -> should not happen
	return -1
}

func test(arr []any, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K: %v \n", k)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := kthSmallest(root, k) // iterative

	fmt.Printf("%v-th smallest element: %v \n", k, result)
	fmt.Printf("Expected %v-th smallest element: %v \n", k, expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []any{3, 1, 4, nil, 2}
	k := 1
	expected := 1

	test(arr, k, expected)
}

func test2() {
	arr := []any{5, 3, 6, 2, 4, nil, nil, 1}
	k := 3
	expected := 3

	test(arr, k, expected)
}

func main() {
	// 230. Kth Smallest Element in a BST
	test1()
	test2()
}
