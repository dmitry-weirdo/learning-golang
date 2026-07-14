package trees

import (
	"fmt"
	"strings"
)

type Node struct { // has a link to parent!
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(
		TreeHeight(node.Left),
		TreeHeight(node.Right),
	)
}

func TreeFromArray(arr []any) *TreeNode {
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

	height := TreeHeight(root)
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

func FindInTree(root *TreeNode, value int) *TreeNode { // not assuming it is a BST, just searching the complete tree
	if root == nil {
		return nil
	}

	if root.Val == value {
		return root
	}

	left := FindInTree(root.Left, value)
	if left != nil {
		return left
	}

	return FindInTree(root.Right, value)
}
