package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func lowestCommonAncestor(p *Node, q *Node) *Node {
	// if we want O(1) space optimization logic, we can:
	// - calculate depths of both nodes
	// - go from the deeper to align the depths
	// - move up step by step until the nodes are the same
	// It will still be O(h) time complexity

	m := make(map[int]*Node)

	// go from P to the root, collect all the values in the path into a map
	current := p

	for current != nil {
		m[current.Val] = current
		current = current.Parent
	}

	// go from Q to the root. The first node in the path that was already in the path of P is the LCA
	current = q

	for current != nil {
		if _, ok := m[current.Val]; ok {
			return current
		}

		current = current.Parent
	}

	// this must never happen, at least the root should be the LCAS
	return nil
}

func test(arr []any, p, q int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("p: %v \n", p)
	fmt.Printf("q: %v \n", q)

	root := trees.TreeFromArrayWithParent(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDownWithParent(root)

	pNode := trees.FindInTreeWithParent(root, p)
	fmt.Printf("Node with value p = %v: %v \n", p, pNode)

	qNode := trees.FindInTreeWithParent(root, q)
	fmt.Printf("Node with value q = %v: %v \n", q, qNode)

	result := lowestCommonAncestor(pNode, qNode) // we don't pass root here

	fmt.Printf("LCA of elements p = %v and q = %v: %v \n", p, q, result.Val)
	fmt.Printf("Expected LCA of elements p = %v and q = %v: %v \n", p, q, expectedResult)

	if result.Val != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result.Val)
	}
}

func test1() {
	arr := []any{5, 3, 4, 2, 1}
	p := 1
	q := 2
	expected := 3

	test(arr, p, q, expected)
}

func test2() {
	arr := []any{5, 3, 4, 2, 1, nil, 9, nil, 11, 10, 12}
	p := 3
	q := 12
	expected := 3

	test(arr, p, q, expected)
}

func main() {
	// 1650. Lowest Common Ancestor of a Binary Tree III
	test1()
	test2()
}
