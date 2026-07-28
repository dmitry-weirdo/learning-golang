package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

var globalHead *Node
var prev *Node // previous node

func treeToDoublyList(root *Node) *Node {
	if root == nil { // tree can be empty (0 nodes)
		return nil
	}

	// reset global variables
	globalHead = nil
	prev = nil

	// we don't user Node.Parent for this ticket
	dfs(root)

	fmt.Printf("After DFS, previous node = %d \n", prev.Val)

	globalHead.Left = prev  // head.prev = tail
	prev.Right = globalHead // tail.next = head

	return globalHead
}

func dfs(node *Node) {
	if node == nil { // base case
		return
	}

	dfs(node.Left)

	// after the left sub-tree, prev will be the current.Left
	fmt.Println()
	fmt.Printf("Current node (after left visited): %v \n", node.Val)

	if prev == nil { // it will be for the left-most, i.e. the smallest node -> it is the new head
		fmt.Printf("Previous node (after left visited): %v \n", nil)

		fmt.Printf("Node %v has previous node = nil. This is the smallest element. It will be head of the new sorted list. \n", node.Val)
		globalHead = node
	} else {
		fmt.Printf("Previous node (after left visited): %v \n", prev.Val)

		// connect prev (already handled) and current node (is currently handled).
		// Since we only change current.left
		node.Left = prev  // current.prev = previous
		prev.Right = node // previous.next = current
	}

	// for the right sub-tree, prev will be current node
	prev = node

	dfs(node.Right)
}

func test(arr []any) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	root := trees.TreeFromArrayWithParent(arr)

	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDownWithParent(root)

	result := treeToDoublyList(root)

	if result == nil {
		fmt.Printf("result.head: %v \n", nil)
	} else {
		fmt.Printf("result.head: %v \n", result.Val)
	}

	if result == nil || result.Left == nil {
		fmt.Printf("result.next: %v \n", nil)
	} else {
		fmt.Printf("result.next: %v \n", result.Left.Val)
	}

	if result == nil || result.Right == nil {
		fmt.Printf("result.prev: %v \n", nil)
	} else {
		fmt.Printf("result.prev: %v \n", result.Right.Val)
	}
}

func test1() {
	arr := []any{4, 2, 5, 1, 3}

	test(arr)
}

func test2() {
	arr := []any{1}

	test(arr)
}

func test3() {
	arr := []any{}

	test(arr)
}

func test4() {
	arr := []any{2, 1, nil}

	test(arr)
}

func main() {
	// 426. Convert Binary Search Tree to Sorted Doubly Linked List
	test1()
	test2()
	test3()
	test4()
}
