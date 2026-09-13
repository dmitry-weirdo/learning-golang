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

	// todo: generalize how get Val from TreeNode! It's not always TreeNode.Val

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
			//fmt.Printf("Root: %v, left: %v \n", root.Val, n.Left.Val)

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
			//fmt.Printf("Root: %v, right: %v \n", root.Val, n.Right.Val)

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

func getLca(up [][]int, levels []int, a, b int) int {
	// define what node is deeper in the tree
	lower, upper := a, b

	if levels[a] < levels[b] {
		lower, upper = b, a
	}

	levelDiff := levels[lower] - levels[upper]

	//fmt.Printf("Lower node: %v, upper node: %v, level difference: %v \n", lower, upper, levelDiff)

	// move from the lower (deeper) node to the same level as the upper (shallower) node
	lower = GetKthAncestor(up, lower, levelDiff)

	//fmt.Printf("Lower node moved to the same level %v as upper node %v. Lower node moved up to %v. \n", levels[upper], upper, lower)

	if lower == upper { // at the same level, nodes are the same -> upper node is the LCA
		return lower
	}

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	n := len(up)

	log := log2(n) + 1

	for i := log - 1; i >= 0; i-- {
		// if the ancestor of this level is the same, continue to the next level
		// I.e. this level is LCA or above
		if up[lower][i] == up[upper][i] {
			continue
		}

		// Ancestors of this level is different -> this level is below LCA
		// Move to this level (to the power of 2)
		// LCA will be still above.
		lower = up[lower][i]
		upper = up[upper][i]
	}

	// both nodes will be directly below their LCA
	return up[lower][0]
}

// todo: get K-th node up from 1483. Kth Ancestor of a Tree Node should be a common function with tests as well
func GetKthAncestor(up [][]int, node int, k int) int {
	// todo: handle -1 specially?
	n := len(up)

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	log := log2(n) + 1

	current := node

	for i := log - 1; i >= 0; i-- {
		if current == -1 { // no need to traverse further if we're above the root
			return -1
		}

		// k = 100 -> we'll go
		powerOf2 := 1 << i
		//fmt.Printf("Power of 2^%v = %v \n", i, powerOf2)

		if k >= powerOf2 {
			// go up 2^i levels, decreasing K
			current = up[current][i]

			k -= powerOf2
		}
	}

	return current
}

// =========================== tests =========================== //
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

func testBinaryLiftingDfs2() {
	n := 5

	arr := []any{ // test if some nodes are not present
		2,
		3, 4,
	}

	expectedBinaryLifting := [][]int{
		{-1, -1, -1}, // 0 // todo: do we want -1 / - 1 / - 1, same as for root?
		{-1, -1, -1}, // 1 // todo: do we want -1 / - 1 / - 1, same as for root?
		{-1, -1, -1}, // 2 // root
		{2, -1, -1},  // 3 - parent is 2
		{2, -1, -1},  // 4 - parent is 2
	}

	expectedLevels := []int{-1, -1, 0, 1, 1}

	testBinaryLiftingDfs(n, arr, expectedBinaryLifting, expectedLevels)
}

func testLca(n int, arr []any, node1, node2 int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of nodes in the tree: %v \n", n)
	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	fmt.Printf("Nodes to find LCA: %v and %v \n", node1, node2)

	// we don't test this method here
	up, levels := getBinaryLiftingDfs(n, tree)

	result := getLca(up, levels, node1, node2)
	fmt.Printf("LCA of %v and %v: %v \n", node1, node2, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func testLca1() {
	// taken from https://leetcode.com/problems/kth-ancestor-of-a-tree-node/
	n := 7

	arr := []any{
		0,
		1, 2,
		3, 4, 5, 6,
	}

	testLca(n, arr, 1, 5, 0)
	testLca(n, arr, 3, 4, 1)
	testLca(n, arr, 5, 6, 2)
	testLca(n, arr, 3, 6, 0)

	// todo: this must be fixed, also need to solve if the values do not start with 0 (or some nodes within 0..n-1 are NOT present in the tree)
	//testLca(n, arr, 1, 7, -1) // node 7 is not in the tree
}

func main() {
	testBinaryLiftingDfs1()
	//testBinaryLiftingDfs2()

	testLca1()
}
