package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iterate(root *TreeNode) (minValue int, maxValue int, valid bool) {
	fmt.Printf("Current root: %v \n", root.Val)

	var minsToCompare []int = []int{root.Val}
	var maxesToCompare []int = []int{root.Val}

	var minLeft, maxLeft int
	var validLeft bool
	if root.Left != nil {
		minLeft, maxLeft, validLeft = iterate(root.Left)

		minsToCompare = append(minsToCompare, minLeft)
	} else {
		minLeft, maxLeft, validLeft = math.MinInt64, math.MinInt64, true
	}

	var minRight, maxRight int
	var validRight bool
	if root.Right != nil {
		minRight, maxRight, validRight = iterate(root.Right)
		maxesToCompare = append(maxesToCompare, maxRight)
	} else {
		minRight, maxRight, validRight = math.MaxInt64, math.MaxInt64, true
	}

	minResult := minOfArray(minsToCompare)
	maxResult := maxOfArray(maxesToCompare)

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

func minOfArray(arr []int) int {
	// we assume a non-empty array
	result := arr[0]

	for _, v := range arr {
		if v < result {
			result = v
		}
	}

	return result
}

func maxOfArray(arr []int) int {
	// we assume a non-empty array
	result := arr[0]

	for _, v := range arr {
		if v > result {
			result = v
		}
	}

	return result
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	minValue, maxValue, valid := iterate(root)

	fmt.Printf("Min value: %d, max value: %d, valid: %v \n", minValue, maxValue, valid)

	return valid
}

func main() {
	r11 := &TreeNode{3, nil, nil}
	r12 := &TreeNode{6, nil, nil}

	r01 := &TreeNode{1, nil, nil}
	r02 := &TreeNode{4, r11, r12}

	root := &TreeNode{5, r01, r02}

	isValidBST(root)
}
