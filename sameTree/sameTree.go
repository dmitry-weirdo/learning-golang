package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil) && (q == nil) { // reached nil nodes on both sides
		return true
	}

	if (p == nil) || (q == nil) { // reached nil on only one side -> fail
		return false
	}

	// root value must be the same
	if p.Val != q.Val {
		return false
	}

	// values of Left and Right must match the same Left and Right of the 2nd tree
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func test(arr1, arr2 []any, expectedResult bool) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree 1 array: %v \n", arr1)

	tree1 := trees.TreeFromArray(arr1)
	fmt.Printf("Initial tree 1: \n")
	trees.PrintTreeTopDown(tree1)

	fmt.Printf("Tree 2 array: %v \n", arr2)

	tree2 := trees.TreeFromArray(arr2)
	fmt.Printf("Initial tree 2: \n")
	trees.PrintTreeTopDown(tree2)

	result := isSameTree(tree1, tree2)

	fmt.Printf("Tree 1 and tree 2 are the same: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]any{
			1,
			2, 3,
		},
		[]any{
			1,
			2, 3,
		},
		true,
	)
}

func test2() {
	test(
		[]any{
			1,
			2, nil,
		},
		[]any{
			1,
			nil, 2,
		},
		false,
	)
}

func test3() {
	test(
		[]any{
			1,
			2, 1,
		},
		[]any{
			1,
			1, 2,
		},
		false,
	)
}

func test4() {
	test(
		[]any{
			1,
			2, nil,
		},
		[]any{
			1,
			2, nil,
		},
		true,
	)
}

func main() {
	// 100. Same Tree
	// Similar to "101. Symmetric Tree"just swap Left/Left and Right/Right comparisons.
	test1()
	test2()
	test3()
	test4()
}
