package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBST_iterative(root)
	//return isValidBST_old(root)
}

func isValidBST_iterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validateBstIterative(root)
}

func isValidBST_old(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//minValue, maxValue, valid := iterate(root) // my complex
	//fmt.Printf("Min value: %d, max value: %d, valid: %v \n", minValue, maxValue, valid)

	previousNode = nil // clear a global variable before every test iteration

	valid := iterateOptimal(root) // nice solution, using the value of BST: if we iterate it left -> root -> right, it will be ascending
	fmt.Printf("Valid: %v \n", valid)

	fmt.Println()
	fmt.Println("Doing a (left -> root -> right) recursive traversal:")
	leftRootRightTraversal(root)

	fmt.Println()
	fmt.Println("Doing a (left -> root -> right) iterative traversal:")
	leftRootRightTraversalIterative(root)

	return valid
}

func iterate(root *TreeNode) (minValue int, maxValue int, valid bool) {
	fmt.Printf("Current root: %v \n", root.Val)

	var minResult int
	var maxResult int

	var minLeft, maxLeft int
	var validLeft bool
	if root.Left != nil {
		minLeft, maxLeft, validLeft = iterate(root.Left)

		minResult = min(minLeft, root.Val)
	} else {
		validLeft = true
		minResult = root.Val
	}

	var minRight, maxRight int
	var validRight bool
	if root.Right != nil {
		minRight, maxRight, validRight = iterate(root.Right)

		maxResult = max(maxRight, root.Val)
	} else {
		validRight = true
		maxResult = root.Val
	}

	leftMaxValidity := (root.Left == nil) || (maxLeft < root.Val)
	rightMinValidity := (root.Right == nil) || (root.Val < minRight)

	validResult := validLeft &&
		validRight &&
		leftMaxValidity &&
		rightMinValidity

	if !validResult {
		fmt.Printf("On node %v, tree is invalid: validLeft = %v, validRight = %v, maxLeft = %v, maxRight = %v \n", root.Val, validLeft, validResult, maxLeft, maxRight)
	}

	return minResult, maxResult, validResult
}

var previousNode *TreeNode = nil

func iterateOptimal(root *TreeNode) bool {
	if root == nil {
		return true
	}

	fmt.Println("=============================")

	if previousNode != nil {
		fmt.Printf("Current root: %v, previous node: %v \n", root.Val, previousNode.Val)
	} else {
		fmt.Printf("Current root: %v, previous node: null \n", root.Val)
	}

	// (left -> root -> right) traversal in a correct BST should be always increasing

	// iterate left
	leftValid := iterateOptimal(root.Left)

	if !leftValid { // anything invalid -> the whole tree is invalid
		fmt.Printf("Current root: %v, left is invalid. Returning false. \n", root.Val)
		return false
	}

	// left has updated the previous -> check that (prev < root)
	if previousNode != nil {
		fmt.Printf("AT THE CHECKING POINT: Current root: %v, previous node: %v \n", root.Val, previousNode.Val)
	} else {
		fmt.Printf("AT THE CHECKING POINT: Current root: %v, previous node: null \n", root.Val)
	}

	if (previousNode != nil) && (previousNode.Val >= root.Val) {
		fmt.Printf("Current node %v is NOT less than previous node %v. Returning false. \n", root.Val, previousNode.Val)
		return false
	}

	// update previous to the current node
	previousNode = root
	fmt.Printf("Set previous node to %v \n", root.Val)

	// iterate right
	rightValid := iterateOptimal(root.Right)

	return rightValid
}

func leftRootRightTraversalIterative(root *TreeNode) {
	stack := list.New()

	current := root

	for stack.Len() > 0 || current != nil {
		// first we push all left nodes
		for current != nil {
			stack.PushFront(current)
			current = current.Left
		}

		//fmt.Println("Stack after adding left nodes:")
		//printList(stack)

		// pop the last left element
		current = stack.Remove(stack.Front()).(*TreeNode)
		fmt.Printf("Iterative, current node: %d \n", current.Val)

		// add the last left.right element, even if it's null
		// if we don't push, we'll iterate to the previous left again from the current node
		current = current.Right
	}
}

func validateBstIterative(root *TreeNode) bool {
	stack := list.New()

	current := root

	var prev *TreeNode = nil

	// see https://www.geeksforgeeks.org/dsa/inorder-tree-traversal-without-recursion/
	for stack.Len() > 0 || current != nil {
		// first we push all left nodes
		// Reach the leftmost node of the current node
		for current != nil {
			stack.PushFront(current)
			current = current.Left
		}

		// pop the last left element
		current = stack.Remove(stack.Front()).(*TreeNode)

		if prev != nil && current.Val <= prev.Val {
			return false
		}

		// We have visited the node and its left subtree.
		// Now, it's right subtree's turn.
		// For the right node, the same process will follow: add all leftmost nodes, etc.

		// add the last left.right element, even if it's null
		// if we don't push, we'll iterate to the previous left again from the current node
		prev = current
		current = current.Right
	}

	return true
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == nil || e.Value.(*TreeNode) == nil {
			fmt.Printf("%v ", nil)
		} else {
			fmt.Printf("%v ", e.Value.(*TreeNode).Val)
		}
	}

	fmt.Println()
}

func leftRootRightTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	leftRootRightTraversal(root.Left)

	fmt.Printf("Current root: %v \n", root.Val)

	leftRootRightTraversal(root.Right)
}

func test1() {
	r11 := &TreeNode{3, nil, nil}
	r12 := &TreeNode{6, nil, nil}

	r01 := &TreeNode{1, nil, nil}
	r02 := &TreeNode{4, r11, r12}

	root := &TreeNode{5, r01, r02}

	isValidBST(root)
}

func test2() {
	r01 := &TreeNode{1, nil, nil}
	r02 := &TreeNode{3, nil, nil}

	root := &TreeNode{2, r01, r02}

	isValidBST(root)
}

func test3() {
	// valid tree, double-check the order validation

	// 5
	// / \
	// 3   8
	// / \   \
	// 2   4   9
	r10 := &TreeNode{2, nil, nil}
	r11 := &TreeNode{4, nil, nil}
	r12 := &TreeNode{9, nil, nil}

	r00 := &TreeNode{3, r10, r11}
	r01 := &TreeNode{8, nil, r12}

	root := &TreeNode{5, r00, r01}

	isValidBST(root)
}

func main() {
	// 98. Validate Binary Search Tree
	test1()
	test2()
	test3()
}
