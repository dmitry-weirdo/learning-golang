package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//fmt.Printf("Val: %v \n", root.Val)

	maxDepthLeft := maxDepth(root.Left)
	maxDepthRight := maxDepth(root.Right)

	return max(maxDepthLeft, maxDepthRight) + 1
}

func test(arr []any, expectedResult int) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(tree)

	result := maxDepth(tree)

	fmt.Printf("Max tree depth: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]any{3, 9, 20, nil, nil, 15, 7}, 3) // 3 - 20 - 15, 3- 20 - 7
}

func test2() {
	test([]any{1, nil, 2}, 2) // 1 - 2
}

func test3() {
	test([]any{}, 0) // no nodes in the tree
}

func main() {
	// 04. Maximum Depth of Binary Tree
	test1()
	test2()
	test3()
}
