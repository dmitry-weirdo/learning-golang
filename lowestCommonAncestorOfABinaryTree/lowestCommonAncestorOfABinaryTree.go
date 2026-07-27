package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// this is a true DFS that really searches for P and Q presence
	return lca_dfs(root, p, q)

	// this will return P or Q if the 2nd node does not exist,
	// since this version implies if we found P we don't search below from it and Q will be there if not found elsewhere
	//return lowestCommonAncestor_old(root, p, q)
}

func lca_dfs(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode = nil

	var dfs func(node *TreeNode) (foundP bool, foundQ bool)

	dfs = func(node *TreeNode) (foundP bool, foundQ bool) {
		if node == nil { // reached leaf, nothing found
			return false, false
		}

		if lca != nil { // LCA already found earlier -> do nothing
			return false, false
		}

		leftFoundP, leftFoundQ := dfs(node.Left)
		rightFoundP, rightFoundQ := dfs(node.Right)

		// this assures the true "found P" and "found Q", no implications that one will be under another
		foundP = leftFoundP || rightFoundP || node == p
		foundQ = leftFoundQ || rightFoundQ || node == q

		// we're the lowest LCA -> save this node to LCA
		if lca == nil && foundP && foundQ {
			lca = node
		}

		return foundP, foundQ
	}

	dfs(root)

	return lca
}

func lowestCommonAncestor_old(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		// reached the leaf and not found neither p nor q
		return nil
	}

	if root == p || root == q {
		// reached one of the target nodes -> immediately return it
		// - if other target node is below it -> this is LCA
		// - if other node is NOT below it -> we have to only search above the current node
		// !!! this implies that both P and Q nodes exist in the tree (by problem description)
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

	var resultValue any
	if result == nil {
		resultValue = nil
	} else {
		resultValue = result.Val
	}

	fmt.Printf("LCA of elements p = %v and q = %v: %v \n", p, q, resultValue)
	fmt.Printf("Expected LCA of elements p = %v and q = %v: %v \n", p, q, expectedResult)

	if resultValue != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, resultValue)
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

func test4() {
	// hacky test - Q does not exist. By the problem conditions, we will return the P node
	arr := []any{1, 2, 3}
	p := 2
	q := 10
	expected := 2 // P will be returned even if Q does not exist

	test(arr, p, q, expected)
}

func main() {
	// 236. Lowest Common Ancestor of a Binary Tree
	test1()
	test2()
	test3()
	test4()
}
