package main

import (
	"demo/matrixCommon"
	"fmt"
	"math"
)

type TreeAncestor struct {
	n   int
	log int
	up  [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	// root is 0
	// ! we have parent links, NOT tree nodes

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	log := log2(n) + 1

	up := createIntMatrix(n, log)

	// 2^0 - direct parent
	for v := range n {
		up[v][0] = parent[v]
	}

	//fmt.Println()
	//fmt.Printf("Binary-lifting matrix with just direct parents filled: \n")
	//matrixCommon.PrintIntMatrix(up)

	// fill 2^i ancestor:
	// 2^i = 2^(i - 1) + 2^(i - 1)

	for i := 1; i < log; i++ { // powers 1, 2, ..., log2(n). For 100, we will iterate from 2^1 = 2 to 2^6 = 64
		// at this 2^i jump level, fill all the nodes
		for v := range n {
			p := up[v][i-1]

			if p == -1 { // reached the parent
				up[v][i] = -1
			} else { // from 2^(i-1) parent, get it 2^(i-1) parent. The sum will sum up to 2^i parent of the current node.
				up[v][i] = up[p][i-1]
			}
		}
	}

	// todo: we can also fill the levels[v] array for quick check whether K is too big to search up. This will also be required for the LCA algorithm.

	return TreeAncestor{
		n:   n,
		log: log,
		up:  up,
	}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	// todo: handle -1 specially?

	current := node

	for i := this.log - 1; i >= 0; i-- {
		if current == -1 { // no need to traverse further if we're above the root
			return -1
		}

		// k = 100 -> we'll go
		powerOf2 := 1 << i
		//fmt.Printf("Power of 2^%v = %v \n", i, powerOf2)

		if k >= powerOf2 {
			// go up 2^i levels, decreasing K
			current = this.up[current][i]

			k -= powerOf2
		}
	}

	return current
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

func testGetKthAncestor(ta TreeAncestor, node int, k int, expectedResult int) {
	fmt.Println()
	fmt.Printf("Querying the %v-th parent of node %v... \n", k, node)

	result := ta.GetKthAncestor(node, k)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 7
	parents := []int{-1, 0, 0, 1, 1, 2, 2}

	ta := Constructor(n, parents)

	fmt.Println()
	fmt.Printf("Binary lifting matrix: \n")
	matrixCommon.PrintIntMatrix(ta.up)

	testGetKthAncestor(ta, 3, 1, 1)
	testGetKthAncestor(ta, 5, 2, 0)
	testGetKthAncestor(ta, 6, 3, -1)
	testGetKthAncestor(ta, 6, 100500, -1)
}

func main() {
	// 1483. Kth Ancestor of a Tree Node
	// Binary lifting, learn it before LCA usage in "1724. Checking Existence of Edge Length Limited Paths II"
	test1()
}
