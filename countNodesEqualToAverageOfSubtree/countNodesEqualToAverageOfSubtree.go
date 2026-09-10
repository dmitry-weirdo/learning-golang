package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func averageOfSubtree(root *TreeNode) int {
	count := 0

	var dfs func(n *TreeNode) (nodeCount int, average int)

	dfs = func(n *TreeNode) (nodeCount int, sum int) {
		if n == nil {
			return 0, 0
		}

		nodesLeft, sumLeft := dfs(n.Left)
		nodesRight, sumRight := dfs(n.Right)

		nodeCount = nodesLeft + nodesRight + 1 // including root
		sum = sumLeft + sumRight + n.Val       // including root
		average := sum / nodeCount

		//fmt.Println()
		//fmt.Printf("Root: %v. Nodes count: %v, sum: %v, average: %v. \n", n.Val, nodeCount, sum, average)

		if average == n.Val {
			//fmt.Printf("Counting %v to the result. \n", n.Val)
			count++
		}

		return nodeCount, sum
	}

	dfs(root)

	return count
}

func test(arr []any, expectedResult int) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	result := averageOfSubtree(tree)

	fmt.Printf("Count of nodes whose value are equal to the average of its subtree (including node itself): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]any{4, 8, 5, 0, 1, nil, 6},
		5,
	)
}

func test2() {
	test(
		[]any{1},
		1,
	)
}

func test3() {
	// Failing test-case 40/139
	test(
		[]any{1, nil, 3, nil, 1, nil, 3},

		/*		[]any{
					1,
					nil, 3,
					nil, nil, nil, 1,
					nil, nil, nil, nil, nil, nil, nil, 3,
				},
		*/
		1,
	)
}

func main() {
	// 2265. Count Nodes Equal to Average of Subtree
	test1()
	test2()
	test3()
}
