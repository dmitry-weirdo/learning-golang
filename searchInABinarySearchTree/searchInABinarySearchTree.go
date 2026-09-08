package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	// no need for recursion, we can return immediately and not track back
	n := root

	for n != nil {
		if n.Val == val { // found node
			return n
		}

		if n.Val > val {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	return nil
}

func test(arr []any, val int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	fmt.Printf("Value to search: %v \n", val)

	result := searchBST(tree, val)

	fmt.Printf("Tree of the found node %v: \n", val)
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
	test(
		[]any{4, 2, 7, 1, 3, nil, nil},
		2,
		[]any{2, 1, 3},
	)
}

func test2() {
	test(
		[]any{4, 2, 7, 1, 3, nil, nil},
		5, // value not present
		[]any{},
	)
}

func main() {
	// 700. Search in a Binary Search Tree
	test1()
	test2()
}
