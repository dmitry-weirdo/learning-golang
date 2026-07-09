package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}

func insert(root *TreeNode, val int) *TreeNode {
	// this is a trivial solution without any rebalancing
	// it also implies that the value does not exist in the tree (it's the problem condition).
	if root == nil { // this also handles the case if initial root was null
		// reached leaf position -> add a new node
		return &TreeNode{Val: val}
	}

	if val < root.Val { // go left
		// will only change if root.Left was null -> then the new node will be set
		root.Left = insert(root.Left, val)
	} else if val > root.Val { // go right
		// will only change if root.Right was null -> then the new node will be set
		root.Right = insert(root.Right, val)
	}

	// root was existing -> return it, so that it doesn't change in the parent
	return root
}

func main() {
	// 701. Insert into a Binary Search Tree
}
