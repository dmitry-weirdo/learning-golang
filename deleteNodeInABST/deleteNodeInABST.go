package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { // key not found!
		return nil
	}

	if key < root.Val { // go left
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { // go right
		root.Right = deleteNode(root.Right, key)
	} else {
		// we need to remove the current node
		if root.Left == nil {
			// simple case 1 - no left child -> just append the right child to the parent
			fmt.Printf("Node to remove %v has no left child. Linking its right child %v to the parent. \n", root.Val, getNodeValSafe(root.Right))
			return root.Right
		}

		if root.Right == nil {
			// simple case 2 - no right child -> just append the right child to the parent
			fmt.Printf("Node to remove %v has no right child. Linking its left child %v to the parent. \n", root.Val, getNodeValSafe(root.Left))
			return root.Left
		}

		// hard case 3 - both left and right child present

		// find the minimum node of root.right
		minNode := getMinNode(root.Right)
		fmt.Printf("Found min node %v in the right subtree of node %v. \n", minNode.Val, root.Val)

		// replace the root value with the min node value
		root.Val = minNode.Val
		fmt.Printf("Found deleted node with the min node value %v. \n", minNode.Val)

		// remove the old min node from the right subtree. Its value is now in the being removed node.
		root.Right = deleteNode(root.Right, minNode.Val)
		fmt.Printf("Removed the old min node %v from the right subtree of the deleted node. \n", minNode.Val)

	}

	// replaced or non-replaced, return the current node to the parent
	return root
}

func getNodeValSafe(node *TreeNode) string {
	if node == nil {
		return "[nil]"
	}

	return strconv.Itoa(node.Val)
}

func getMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// to find the min, we constantly move to the left
	node := root

	for node.Left != nil {
		node = node.Left
	}

	return node
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

func printTree(root *TreeNode) { // level by level
	if root == nil {
		fmt.Printf("[] (root is nil)")
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				fmt.Printf("%v ", node.Val)
			} else {
				fmt.Printf("%v ", ".")
			}

			//if node != nil && node.Left != nil {
			if node != nil {
				queue = append(queue, node.Left)
			}

			//if node.Right != nil {
			if node != nil {
				queue = append(queue, node.Right)
			}
		}

		fmt.Println()
	}
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

func main() {
	// 450. Delete Node in a BST
	nodes := []any{5, 3, 6, 2, 4, nil, 7}
	tree := treeFromArray(nodes)
	//printTree(tree)
	//fmt.Printf("====================")
	PrintTreeTopDown(tree)

	key := 100
	deleteNode(tree, key)

	fmt.Printf("Tree after removing the node %v: \n", key)
	PrintTreeTopDown(tree)

}
