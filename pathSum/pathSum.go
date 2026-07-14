package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	//return dfs(root, targetSum, 0)
	return dfsOptimized(root, targetSum)
}

func dfsOptimized(root *TreeNode, targetSum int) bool {
	if root == nil {
		// reached an empty node -> result not reached
		return false
	}

	// we can just subtract the current node from the remaining sum
	remainingSum := targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		// leaf node reached
		return remainingSum == 0
	}

	// check whether left or right subtree have the target path
	return dfsOptimized(root.Left, remainingSum) || dfsOptimized(root.Right, remainingSum)
}

func dfs(root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}

	newSum := sum + root.Val

	if root.Left == nil && root.Right == nil {
		// leaf node reached
		return newSum == targetSum
	}

	if root.Left != nil {
		// try the left subtree
		leftWithTargetSumExists := dfs(root.Left, targetSum, newSum)

		if leftWithTargetSumExists {
			return true
		}
	}

	if root.Right != nil {
		// try the right subtree
		rightWithTargetSumExists := dfs(root.Right, targetSum, newSum)

		if rightWithTargetSumExists {
			return true
		}
	}

	return false
}

func test(arr []any, targetSum int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", targetSum)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := hasPathSum(root, targetSum)

	fmt.Printf("Root to leaf path with sum = %v exists: %v \n", targetSum, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
	targetSum := 22
	expected := true

	test(arr, targetSum, expected)
}

func test2() {
	arr := []any{1, 2, 3}
	targetSum := 5
	expected := false

	test(arr, targetSum, expected)
}

func test3() {
	arr := []any{}
	targetSum := 0
	expected := false

	test(arr, targetSum, expected)
}

func main() {
	// 112. Path Sum
	test1()
	test2()
	test3()
}
