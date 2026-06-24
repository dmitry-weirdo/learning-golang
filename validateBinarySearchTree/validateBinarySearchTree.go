package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	minValue, maxValue, valid := iterate(root)

	fmt.Printf("Min value: %d, max value: %d, valid: %v \n", minValue, maxValue, valid)

	return valid
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

func main() {
	//test1()
	test2()
}
