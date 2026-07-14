package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	current := root.Val

	// add the current element if it is within the range
	sum := 0

	if (low <= current) && (current <= high) {
		sum += current
	}

	// if current < low -> skip the left subtree, all the values there are even less
	if current >= low {
		sum += rangeSumBST(root.Left, low, high)
	}

	// if current > high -> skip the right subtree, all the values there are even bigger
	if current <= high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}

func test(arr []any, low, high int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("low: %v \n", low)
	fmt.Printf("high: %v \n", high)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := rangeSumBST(root, low, high) // iterative

	fmt.Printf("Sum of [%v; %v] elements: %v \n", low, high, result)
	fmt.Printf("Expected sum of [%v; %v] elements: %v \n", low, high, expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []any{10, 5, 15, 3, 7, 13, 18, 1, nil, 6}
	low := 6
	high := 10
	expected := 23

	test(arr, low, high, expected)
}

func test2() {
	arr := []any{10, 5, 15, 3, 7, nil, 18}
	low := 7
	high := 15
	expected := 32

	test(arr, low, high, expected)
}

func main() {
	// 938. Range Sum of BST
	test1()
	test2()
}
