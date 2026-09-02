package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
	"slices"
)

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		currentLevelCount := len(queue)
		currentLevel := make([]int, currentLevelCount)

		for i := range currentLevelCount {
			// handle just the elements of the current level
			node := queue[0]
			queue = queue[1:]

			currentLevel[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	// return the levels in the opposite order (bottom-to-top)
	slices.Reverse(result)
	return result
}

func test(arr []any, expectedResult [][]int) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(tree)

	result := levelOrderBottom(tree)

	fmt.Printf("Bottom-up levels, left-to-right traversal: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	a := []any{3, 9, 20, nil, nil, 15, 7}

	expected := [][]int{
		{15, 7},
		{9, 20},
		{3},
	}

	test(a, expected)
}

func test2() {
	a := []any{1}

	expected := [][]int{
		{1},
	}

	test(a, expected)
}

func test3() {
	a := []any{}

	expected := [][]int{}

	test(a, expected)
}

func main() {
	// 107. Binary Tree Level Order Traversal II
	// Basically the same as "102. Binary Tree Level Order Traversal", we just revert the result by level
	test1()
	test2()
	test3()
}
