package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalIterative(root *TreeNode) []int {
	result := make([]int, 0)

	stack := make([]*TreeNode, 0)

	current := root

	for current != nil || (len(stack) > 0) {
		// go to the max-left from the current node, pushing nodes to the stack
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop the last value from the stack
		lastIndex := len(stack) - 1
		current = stack[lastIndex]
		stack = stack[0:lastIndex] // remove the last value

		// add current root to the result
		result = append(result, current.Val)

		// go to the right of the current node, then the process will continue for this right sub-node
		// do it even if the current.Right == nil, else we'll stay at the current root
		current = current.Right
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	inOrder(root, &result)

	return result
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, result)

	*result = append(*result, root.Val)

	inOrder(root.Right, result)
}

func treeFromArray(arr []any) *TreeNode {
	if len(arr) < 1 || arr[0] == nil { // empty array -> empty tree
		return nil
	}

	// we do kind of BST to construct the tree
	// first we append 0-th level (root), then 1st level, etc.
	root := &TreeNode{Val: arr[0].(int)}

	// todo: we can use list (LinkedList) instead of slice
	// 0th level is just the root
	queue := []*TreeNode{root}

	i := 1 // skip the root

	for (len(queue) > 0) && (i < len(arr)) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func PrintTreeTopDown(root *TreeNode) {
	if root == nil {
		return
	}

	height := treeHeight(root)
	width := (1 << height) * 2

	current := []*TreeNode{root}

	for level := 0; level < height; level++ {
		gap := width / (1 << (level + 1))

		next := make([]*TreeNode, 0)

		fmt.Print(strings.Repeat(" ", gap))

		for _, node := range current {
			if node == nil {
				fmt.Print(" ")
				next = append(next, nil, nil)
			} else {
				fmt.Print(node.Val)
				next = append(next, node.Left, node.Right)
			}

			fmt.Print(strings.Repeat(" ", gap*2-1))
		}

		fmt.Println()
		current = next
	}
}

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(
		treeHeight(node.Left),
		treeHeight(node.Right),
	)
}

func test(arr []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	root := treeFromArray(arr)
	fmt.Printf("Tree: \n")
	PrintTreeTopDown(root)

	//inOrderResult := inorderTraversal(root) // recursive
	inOrderResult := inorderTraversalIterative(root) // iterative

	fmt.Printf("In order traversal (left -> root -> right): %v \n", inOrderResult)
}

func test1() {
	arr := []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}

	test(arr)
}

func test2() {
	arr := []any{}

	test(arr)
}

func test3() {
	arr := []any{1}

	test(arr)
}

func test4() {
	arr := []any{1, nil, 2, 3}

	test(arr)
}

func main() {
	// 94. Binary Tree Inorder Traversal
	test1()
	test2()
	test3()
	test4()
}
