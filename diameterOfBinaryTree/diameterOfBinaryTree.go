package main

import (
	"demo/trees"   // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := 0

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int { // returns longest path to parent
		if node == nil { // nil node has not path
			return 0
		}

		leftLongestPath := dfs(node.Left)
		rightLongestPath := dfs(node.Right)

		// the current node path is leftPath -> node -> rightPath
		// !!! we're not adding + 1 for the node since we're counting edges. Children already returned their count of edges up to the current node.
		currentDiameter := leftLongestPath + rightLongestPath
		diameter = max(diameter, currentDiameter)

		// to the parent node, we return the longest path either from left or from right + 1 for node itself (for the edge between node and parent)
		return max(leftLongestPath, rightLongestPath) + 1
	}

	dfs(root)

	return diameter
}

func test(arr []any, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	root := trees.TreeFromArray(arr)

	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := diameterOfBinaryTree(root)

	fmt.Printf("Diameter of the tree (count of edges in the longest path): %v \n", result)
	fmt.Printf("Expected diameter: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []any{1, 2, 3, 4, 5}
	expected := 3

	test(arr, expected)
}

func test2() {
	arr := []any{1, 2}
	expected := 1

	test(arr, expected)
}

func test3() {
	arr := []any{1}
	expected := 0 // no edges

	test(arr, expected)
}

func main() {
	// 543. Diameter of Binary Tree
	test1()
	test2()
	test3()
}
