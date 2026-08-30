package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func minDepth(root *TreeNode) int {
	// passed in 4-8 ms. BFS should logically be faster because of immediate return instead of recursion
	return minDepth_dfs(root)
}

func minDepth_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//fmt.Printf("Val: %v \n", root.Val)

	// we can't use 0 for nil children since it must NOT count
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	// both left and right are not null -> select min of them
	minDepthLeft := minDepth(root.Left)
	minDepthRight := minDepth(root.Right)

	return min(minDepthLeft, minDepthRight) + 1
}

func test(arr []any, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Tree : \n")
	trees.PrintTreeTopDown(tree)

	result := minDepth(tree)

	fmt.Printf("Min tree depth: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]any{3, 9, 20, nil, nil, 15, 7}, 2) // 3 - 9
}

func test2() {
	test([]any{2, nil, 3, nil, 4, nil, 5, nil, 6}, 5) // 2 - 3 - 4 - 5 - 6
}

func test3() {
	test([]any{}, 0) // no nodes
}

func main() {
	// 111. Minimum Depth of Binary Tree
	test1()
	test2()
	test3()
}
