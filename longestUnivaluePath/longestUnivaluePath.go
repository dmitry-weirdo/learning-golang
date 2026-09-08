package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLength := 0 // can be updated within DFS in case of node does not propagate to parent

	var dfs func(n *TreeNode) int

	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		leftPathLength := dfs(n.Left)
		rightPathLength := dfs(n.Right)

		if n.Left != nil && n.Right != nil && n.Left.Val == n.Val && n.Right.Val == n.Val {
			// left - root - right path
			// + 2 edges: root-left and root-right
			pathLength := leftPathLength + rightPathLength + 2

			maxLength = max(maxLength, pathLength)

			// the longer of "root-left-..." and "root-right-..." paths is returned to the parent
			return max(leftPathLength, rightPathLength) + 1
		}

		if n.Left != nil && n.Left.Val == n.Val {
			// root - left match
			// + 1 edge: root-left
			maxLength = max(maxLength, leftPathLength+1)
			return leftPathLength + 1
		}

		if n.Right != nil && n.Right.Val == n.Val {
			// root - right match
			// + 1 edge: root-right
			maxLength = max(maxLength, rightPathLength+1)
			return rightPathLength + 1
		}

		// no matches from the current root -> no path increase
		return 0
	}

	dfs(root)

	return maxLength
}

func test(arr []any, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	result := longestUnivaluePath(tree)

	fmt.Printf("Length of longest univalue path: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]any{5, 4, 5, 1, 1, nil, 5},
		2, // 5 - 5 - 5, we're counting the edges
	)
}

func test2() {
	test(
		[]any{1, 4, 5, 4, 4, nil, 5},
		2, // 4 - 4 - 4, we're counting the edges
	)
}

func main() {
	// 687. Longest Univalue Path
	test1()
	test2()
}
