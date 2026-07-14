package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		// reached the leaf and not found neither p nor q
		return nil
	}

	if root == p || root == q {
		// reached one of the target nodes -> immediately return it
		// - if other target node is below it -> this is LCA
		// - if other node is NOT below it -> we have to only search above the current node
		return root
	}

	leftLca := lowestCommonAncestor(root.Left, p, q)
	rightLca := lowestCommonAncestor(root.Right, p, q)

	if leftLca != nil && rightLca != nil {
		// p and q found in different subtrees -> current node is LCA
		return root
	}

	if leftLca == nil && rightLca == nil {
		// none of the subtrees found neither p nor q
		return nil
	}

	if leftLca != nil {
		// one or two of the nodes found in the left subtree
		// it will be either LCA or one or the found nodes
		return leftLca
	} else {
		// one or two of the nodes found in the right subtree
		// it will be either LCA or one or the found nodes
		return rightLca
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
	arr := []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	p := 5
	q := 1
	expected := 3

	test(arr, p, q, expected)
}

func test2() {
	arr := []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	p := 5
	q := 4
	expected := 5

	test(arr, p, q, expected)
}

func test3() {
	// failing test-case from BST variant
	// https://neetcode.io/problems/lowest-common-ancestor-in-binary-search-tree/question
	// the problem was that nodes were not the same memory, so we have to compare values
	arr := []any{5, 3, 8, 1, 4, 7, 9, nil, 2}
	p := 3
	q := 8
	expected := 5

	test(arr, p, q, expected)
}

func main() {
	// 236. Lowest Common Ancestor of a Binary Tree
	test1()
	test2()
	test3()
}
