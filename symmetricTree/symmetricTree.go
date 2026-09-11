package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(n1 *TreeNode, n2 *TreeNode) bool {
	if (n1 == nil) && (n2 == nil) { // reached nil nodes on both sides
		return true
	}

	if (n1 == nil) || (n2 == nil) { // reached nil on only one side -> fail
		return false
	}

	// root value must be the same
	if n1.Val != n2.Val {
		return false
	}

	// values of Left/Right must match the opposites of the contrary side
	return isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left)
}

func test(arr []any, expectedResult bool) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	result := isSymmetric(tree)

	fmt.Printf("Tree is symmetric (mirrored): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]any{
			1,
			2, 2,
			3, 4, 4, 3,
		},
		true,
	)
}

func test2() {
	test(
		[]any{
			1,
			2, 2,
			3, 4, 4, 3,
			5, 6, nil, nil, nil, nil, 6, 5,
		},
		true,
	)
}

func test3() {
	test(
		[]any{
			1,
			2, 2,
			3, 4, 4, 3,
			5, 6, nil, nil, nil, nil, 5, 6,
		},
		false,
	)
}

func test4() {
	test(
		[]any{
			1,
			2, 2,
			nil, 3, nil, 3,
		},
		false,
	)
}

func main() {
	// 101. Symmetric Tree
	// Similar to "100. Same Tree", just swap Left/Right and Right/Left comparisons.
	test1()
	test2()
	test3()
	test4()
}
