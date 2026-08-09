package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(left, right int) *TreeNode

	dfs = func(left, right int) *TreeNode {
		// base case, no elements -> return nil
		if left > right {
			//fmt.Printf("Reached nil node, returning it. \n")
			return nil
		}

		// base case, if just 1 element -> create leaf node and return it
		if left == right {
			//fmt.Printf("Reached leaf node %v, returning it. \n", nums[left])
			return &TreeNode{nums[left], nil, nil}
		}

		size := right - left + 1

		// since we're building a _balanced_ tree, we're selecting the root in the middle
		// to have close number
		rootIndex := 0

		if size%2 == 1 { // odd number of elements -> middle is obvious
			rootIndex = (left + right) / 2
		} else { // even number of elements -> we can select either more nodes on left or more nodes on right
			// more on the left
			rootIndex = (left+right)/2 + 1

			// if we want more on the right
			//rootIndex = (left + right) / 2
		}

		leftStart := left
		leftEnd := rootIndex - 1

		rightStart := rootIndex + 1
		rightEnd := right

		//fmt.Println()
		//fmt.Printf("Root: %v \n", nums[rootIndex])
		//fmt.Printf("Left tree size: %d \n", leftEnd-leftStart+1)
		//fmt.Printf("Left tree: %d \n", nums[leftStart:leftEnd+1])
		//
		//fmt.Printf("Right tree size: %d \n", rightEnd-rightStart+1)
		//fmt.Printf("Right tree: %d \n", nums[rightStart:rightEnd+1])

		leftTree := dfs(leftStart, leftEnd)
		rightTree := dfs(rightStart, rightEnd)

		return &TreeNode{nums[rootIndex], leftTree, rightTree}
	}

	return dfs(0, len(nums)-1)
}

func test(arr []int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Sorted array: %v \n", arr)

	result := sortedArrayToBST(arr)

	fmt.Printf("Built balanced BST: \n")
	trees.PrintTreeTopDown(result)

	resultAsArray := trees.TreeToArray(result)
	fmt.Printf("Tree as array:   %v \n", resultAsArray)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(resultAsArray) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(resultAsArray))
		return
	}

	for i, v := range resultAsArray {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	arr := []int{-10, -3, 0, 5, 9}
	expected := []any{0, -3, 9, -10, nil, 5, nil}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 3}
	expected := []any{3, 1, nil}

	test(arr, expected)
}

func main() {
	// 108. Convert Sorted Array to Binary Search Tree
	test1()
	test2()
}
