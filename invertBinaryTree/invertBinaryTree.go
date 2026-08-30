package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	rightInverted := invertTree(root.Right)
	leftInverted := invertTree(root.Left)

	root.Left = rightInverted
	root.Right = leftInverted

	return root
}

func test(arr []any, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	result := invertTree(tree)

	fmt.Printf("Inverted tree: \n")
	trees.PrintTreeTopDown(result)

	resultAsArray := trees.TreeToArray(result)
	fmt.Printf("Result tree as array: %v \n", resultAsArray)
	fmt.Printf("Expected result:      %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	arr := []any{4, 2, 7, 1, 3, 6, 9}
	expected := []any{4, 7, 2, 9, 6, 3, 1}

	test(arr, expected)
}

func test2() {
	arr := []any{2, 1, 3}
	expected := []any{2, 3, 1}

	test(arr, expected)
}

func test3() {
	arr := []any{}
	expected := []any{}

	test(arr, expected)
}

func main() {
	// 226. Invert Binary Tree
	test1()
	test2()
	test3()
}
