package main

import (
	"demo/matrixCommon"
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
	"math"
)

// todo: move this function to trees.go
func getBinaryLiftingDfs(n int, root *TreeNode) (binaryLifting [][]int, levels []int) {
	// see https://www.youtube.com/watch?v=dOAxrhAUIhA

	// todo: we should also handle node indexes, not values

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	log := log2(n) + 1

	// up[v][i]
	// v - [0; n-1] - tree node values
	// i - 2^i jumps from 0 to log2(n)
	// todo: probably we need to map not to [int] but to [*TreeNode] ???
	up := createIntMatrix(n, log)

	// set -1 for root, all levels
	for i := range log {
		up[root.Val][i] = -1
	}

	// also fill the levels array
	levels = make([]int, n) // for nodes 0 to n - 1

	var dfs func(n *TreeNode, level int)

	// todo: probably we need to map not to [int] but to [*TreeNode]
	dfs = func(n *TreeNode, level int) {
		if n == nil {
			return
		}

		levels[n.Val] = level

		// left
		if n.Left != nil {
			fmt.Printf("Root: %v, left: %v \n", root.Val, n.Left.Val)

			// todo: value should be *TreeNode?
			// 2^0 - direct parent
			v := n.Left.Val
			up[v][0] = n.Val

			// fill 2^i ancestor:
			// 2^i = 2^(i - 1) + 2^(i - 1)
			for i := 1; i < log; i++ { // powers 1, 2, ..., log2(n). For 100, we will iterate from 2^1 = 2 to 2^6 = 64
				// at this 2^i jump level, fill all the nodes
				p := up[v][i-1]

				if p == -1 { // reached the parent // todo: do we need this?
					up[v][i] = -1
				} else { // from 2^(i-1) parent, get it 2^(i-1) parent. The sum will sum up to 2^i parent of the current node.
					up[v][i] = up[p][i-1]
				}
			}

			dfs(n.Left, level+1)
		}

		// right
		if n.Right != nil {
			fmt.Printf("Root: %v, right: %v \n", root.Val, n.Right.Val)

			// 2^0 - direct parent
			v := n.Right.Val
			up[v][0] = n.Val

			// fill 2^i ancestor:
			// 2^i = 2^(i - 1) + 2^(i - 1)
			for i := 1; i < log; i++ { // powers 1, 2, ..., log2(n). For 100, we will iterate from 2^1 = 2 to 2^6 = 64
				// at this 2^i jump level, fill all the nodes
				p := up[v][i-1]

				if p == -1 { // reached the parent // todo: do we need this?
					up[v][i] = -1
				} else { // from 2^(i-1) parent, get it 2^(i-1) parent. The sum will sum up to 2^i parent of the current node.
					up[v][i] = up[p][i-1]
				}
			}

			dfs(n.Right, level+1)
		}
	}

	dfs(root, 0) // root has level 0
	return up, levels
}

func log2(n int) int {
	// todo: log2 should be handled separately, it's undefined
	return int(math.Log2(float64(n)))

	// for positive integers, counting bits can be used:
	// bits.Len(uint(n)) - 1
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func testBinaryLiftingDfs(n int, arr []any, expectedResult [][]int, expectedLevels []int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of nodes in the tree: %v \n", n)
	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	result, levels := getBinaryLiftingDfs(n, tree)

	// test the Binary Lifting matrix
	fmt.Printf("Binary lifting matrix: %v \n", result)
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected result: %v \n", expectedResult)
	matrixCommon.PrintIntMatrix(expectedResult)

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

	// test the levels array
	fmt.Printf("Node levels: %v \n", levels)
	fmt.Printf("Expected result: %v \n", expectedLevels)

	if len(levels) != len(expectedLevels) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range levels {
		if v != expectedLevels[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func testBinaryLiftingDfs1() {
	// taken from https://leetcode.com/problems/kth-ancestor-of-a-tree-node/
	n := 7

	arr := []any{
		0,
		1, 2,
		3, 4, 5, 6,
	}

	expectedBinaryLifting := [][]int{
		{-1, -1, -1},
		{0, -1, -1},
		{0, -1, -1},
		{1, 0, -1},
		{1, 0, -1},
		{2, 0, -1},
		{2, 0, -1},
	}

	expectedLevels := []int{0, 1, 1, 2, 2, 2, 2}

	testBinaryLiftingDfs(n, arr, expectedBinaryLifting, expectedLevels)
}

func main() {
	testBinaryLiftingDfs1()
}
