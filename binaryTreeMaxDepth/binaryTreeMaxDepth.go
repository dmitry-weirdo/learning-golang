package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	fmt.Printf("Val: %v \n", root.Val)

	maxDepthLeft := maxDepth(root.Left)
	maxDepthRight := maxDepth(root.Right)

	return max(maxDepthLeft, maxDepthRight) + 1
}

func main() {
	r11 := &TreeNode{15, nil, nil}
	r12 := &TreeNode{7, nil, nil}

	r01 := &TreeNode{9, nil, nil}
	r02 := &TreeNode{20, r11, r12}

	root := &TreeNode{3, r01, r02}

	result := maxDepth(root)

	expectedResult := 3

	fmt.Printf("Expected result: %v \n", expectedResult)
	fmt.Printf("Result: %v \n", result)
}
