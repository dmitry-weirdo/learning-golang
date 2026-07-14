package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
	"math"
)

var maxPath int

func maxPathSum(root *TreeNode) int {
	// we start with -Infinity, not with 0 - since we can have all negative nodes in the tree
	// Since min node value is -1000, we can actually start with -1001
	// todo: we can actually start with -1001
	maxPath = math.MinInt

	dfs(root)

	return maxPath
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// if the subtree returns negative maximum, we use 0 that means not including this subtree at all
	maxLeft := max(0, dfs(root.Left))
	maxRight := max(0, dfs(root.Right))

	// we can use both left and right subtrees, but then we stop appending the current node to its parent
	maxWithBothLeftAndRight := root.Val + maxLeft + maxRight
	maxPath = max(maxPath, maxWithBothLeftAndRight)

	// to include the current node to the parent, we can only select either the left of the right subtree of the current node
	return root.Val + max(maxLeft, maxRight)
}

func test(arr []any, expectedResult int) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := maxPathSum(root) // iterative

	fmt.Printf("Max path sum: %v \n", result)
	fmt.Printf("Expected max path sum: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []any{1, 2, 3}
	expected := 6

	test(arr, expected)
}

func test2() {
	arr := []any{-10, 9, 20, nil, nil, 15, 7}
	expected := 42

	test(arr, expected)
}

/*func test3() {
	// test-case taken from 543. Diameter of Binary Tree
	arr := []any{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2}
	expected := -1

	test(arr, expected)
}
*/

func main() {
	// 124. Binary Tree Maximum Path Sum
	test1()
	test2()
	//test3()
}
