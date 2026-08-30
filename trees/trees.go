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

type NodeType interface {
	*Node | *TreeNode
}

// todo: adopt to use a common interface instead of methods duplication
type GenericNode interface {
	Val() int
	LeftChild() GenericNode
	RightChild() GenericNode
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

func TreeHeightWithParent(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + max(
		TreeHeightWithParent(node.Left),
		TreeHeightWithParent(node.Right),
	)
}

func TreeFromArray(arr []any) *TreeNode { // constructs from an array of BST (level-by-level) traversal
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

func TreeFromArrayWithParent(arr []any) *Node {
	if len(arr) < 1 || arr[0] == nil { // empty array -> empty tree
		return nil
	}

	// we do kind of BST to construct the tree
	// first we append 0-th level (root), then 1st level, etc.
	root := &Node{Val: arr[0].(int)}

	// todo: we can use list (LinkedList) instead of slice
	// 0th level is just the root
	queue := []*Node{root}

	i := 1 // skip the root

	for (len(queue) > 0) && (i < len(arr)) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			node.Left = &Node{Val: arr[i].(int), Parent: node}
			queue = append(queue, node.Left)
		}

		i++

		if i < len(arr) && arr[i] != nil {
			node.Right = &Node{Val: arr[i].(int), Parent: node}
			queue = append(queue, node.Right)
		}

		i++
	}

	return root
}

func TreeToArray(node *TreeNode) []any { // returns as full binary BFS tree with nulls when nodes are not present in position
	if node == nil {
		return []any{}
	}

	arr := make([]any, 0) // int or null

	// BFS the tree, including null values
	queue := make([]*TreeNode, 0) // *TreeNode can be nil
	queue = append(queue, node)

	//level := 0

	for len(queue) > 0 {
		// handle current level
		currentLevelSize := len(queue)

		//fmt.Printf("Current level: %v, elements on this level: %v \n", level, currentLevelSize)
		//level++

		nonNullElementsOnNextLevel := 0

		for range currentLevelSize {
			n := queue[0]
			queue = queue[1:] // remove 1st from the queue

			if n == nil {
				arr = append(arr, nil)

				queue = append(queue, nil) // push nil as left
				queue = append(queue, nil) // push nil as right
			} else {
				arr = append(arr, n.Val)

				queue = append(queue, n.Left)  // we also append nil values
				queue = append(queue, n.Right) // we also append nil values

				if n.Left != nil {
					nonNullElementsOnNextLevel++
				}

				if n.Right != nil {
					nonNullElementsOnNextLevel++
				}
			}
		}

		if nonNullElementsOnNextLevel <= 0 { // no non-empty elements on the next level
			break
		}
	}

	return arr
}

func PrintTreeTopDown(root *TreeNode) {
	if root == nil {
		return
	}

	height := TreeHeight(root)
	width := (1 << height) * 2

	current := []*TreeNode{root}

	for level := range height {
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

func PrintTreeTopDownWithParent(root *Node) {
	if root == nil {
		return
	}

	height := TreeHeightWithParent(root)
	width := (1 << height) * 2

	current := []*Node{root}

	for level := 0; level < height; level++ {
		gap := width / (1 << (level + 1))

		next := make([]*Node, 0)

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

func FindInTreeWithParent(root *Node, value int) *Node { // not assuming it is a BST, just searching the complete tree
	if root == nil {
		return nil
	}

	if root.Val == value {
		return root
	}

	left := FindInTreeWithParent(root.Left, value)
	if left != nil {
		return left
	}

	return FindInTreeWithParent(root.Right, value)
}
