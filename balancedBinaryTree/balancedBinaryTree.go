package main

import (
	"demo/trees"
	"fmt"
)

import . "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix

func isBalanced(root *TreeNode) bool {
	//balanced, _ := treeIsBalanced(root)
	balanced, _ := treeIsBalancedOptimized(root)

	return balanced
}

func treeIsBalancedOptimized(node *TreeNode) (balanced bool, height int) { // optimized - fail-fast
	if node == nil {
		return true, 0
	}

	balancedLeft, heightLeft := treeIsBalancedOptimized(node.Left)
	if !balancedLeft { // return false, height does not matter
		return false, -1
	}

	balancedRight, heightRight := treeIsBalancedOptimized(node.Right)
	if !balancedRight { // return false, height does not matter
		return false, -1
	}

	heightDiff := abs(heightLeft - heightRight)
	balanced = heightDiff < 2

	rootHeight := 1 + max(heightLeft, heightRight)

	return balanced, rootHeight
}

func treeIsBalanced(node *TreeNode) (balanced bool, height int) { // will go over all nodes in the tree
	if node == nil {
		return true, 0
	}

	balancedLeft, heightLeft := treeIsBalanced(node.Left)
	balancedRight, heightRight := treeIsBalanced(node.Right)

	heightDiff := abs(heightLeft - heightRight)
	balanced = balancedLeft && balancedRight && (heightDiff < 2)

	rootHeight := 1 + max(heightLeft, heightRight)

	return balanced, rootHeight
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func test(arr []any, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	isBalancedResult := isBalanced(root) // iterative

	fmt.Printf("Is balanced: %v \n", isBalancedResult)
	fmt.Printf("Expected balanced: %v \n", expectedResult)

	if isBalancedResult != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, isBalancedResult)
	}
}

func test1() {
	arr := []any{3, 9, 20, nil, nil, 15, 7}
	expected := true

	test(arr, expected)
}

func test2() {
	arr := []any{1, 2, 2, 3, 3, nil, nil, 4, 4}
	expected := false

	test(arr, expected)
}

func test3() {
	arr := []any{1}
	expected := true

	test(arr, expected)
}

func test4() {
	arr := []any{}
	expected := true

	test(arr, expected)
}

func main() {
	// 110. Balanced Binary Tree
	test1()
	test2()
	test3()
	test4()
}
