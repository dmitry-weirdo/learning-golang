package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	minValue := min(p.Val, q.Val)
	maxValue := max(p.Val, q.Val)

	current := root

	for {
		if current.Val < minValue {
			// if (current < min) -> discard left subtree, go right
			current = current.Right
		} else if current.Val > maxValue {
			// if (current > max) -> discard right subtree, go left
			current = current.Left
		} else {
			// min <= current <= max -> nothing more to discard
			// case 1: min is left, max is right -> we're at LCA where min < LCA < max
			// case 2: min is current, max is right -> we're at LCA = min
			// case 3: max is current, min is left -> we're at LCA = max

			return current
		}
	}
}

func test(arr []any, p, q int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("p: %v \n", p)
	fmt.Printf("q: %v \n", q)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	pNode := trees.FindInTree(root, p)
	fmt.Printf("Node with value p = %v: %v \n", p, pNode)

	qNode := trees.FindInTree(root, q)
	fmt.Printf("Node with value q = %v: %v \n", q, qNode)

	result := lowestCommonAncestor(root, pNode, qNode)

	fmt.Printf("LCA of elements p = %v and q = %v: %v \n", p, q, result.Val)
	fmt.Printf("Expected LCA of elements p = %v and q = %v: %v \n", p, q, expectedResult)

	if result.Val != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result.Val)
	}
}

func test1() {
	arr := []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	p := 2
	q := 8
	expected := 6

	test(arr, p, q, expected)
}

func test2() {
	arr := []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	p := 2
	q := 4
	expected := 2

	test(arr, p, q, expected)
}

func test3() {
	arr := []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	p := 3
	q := 5
	expected := 4

	test(arr, p, q, expected)
}

func main() {
	// 235. Lowest Common Ancestor of a Binary Search Tree
	test1()
	test2()
	test3()
}
