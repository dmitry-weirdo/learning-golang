package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	// we need to declare if we want to call dfs within dfs
	var dfs func(inorderLeft, inorderRight, postorderLeft, postorderRight int) *TreeNode

	dfs = func(inorderLeft, inorderRight, postorderLeft, postorderRight int) *TreeNode {
		// base case, no elements -> return nil
		if inorderLeft > inorderRight {
			//fmt.Printf("Reached nil node, returning it. \n")
			return nil
		}

		// base case, if just 1 element -> create leaf node and return it
		if inorderLeft == inorderRight {
			//fmt.Printf("Reached leaf node %v, returning it. \n", inorder[inorderLeft])
			return &TreeNode{inorder[inorderLeft], nil, nil}
		}

		// in postorder, root is the last
		rootValue := postorder[postorderRight]

		// todo: any way to find it faster than full traversal?
		// find root in [inorderLeft:inorderRight]
		inorderRootIndex := -1

		for inorderIndex := inorderLeft; inorderIndex <= inorderRight; inorderIndex++ {
			if inorder[inorderIndex] == rootValue {
				inorderRootIndex = inorderIndex
				break
			}
		}

		// inorder goes [<leftTree><root><rightTree>]
		// postorder goes [<leftTree><rightTree><root>]

		inorderLeftStart := inorderLeft
		inorderLeftEnd := inorderRootIndex - 1
		leftTreeSize := inorderLeftEnd - inorderLeftStart + 1

		inorderRightStart := inorderRootIndex + 1
		inorderRightEnd := inorderRight
		rightTreeSize := inorderRightEnd - inorderRightStart + 1

		postorderRightEnd := postorderRight - 1
		postorderRightStart := postorderRightEnd - rightTreeSize + 1

		postorderLeftEnd := postorderRightStart - 1
		postorderLeftStart := postorderLeftEnd - leftTreeSize + 1

		//fmt.Printf("Left tree size: %d \n", leftTreeSize)
		//fmt.Printf("Left tree inorder: %d \n", inorder[inorderLeftStart:inorderLeftEnd+1])
		//fmt.Printf("Left tree postorder: %d \n", postorder[postorderLeftStart:postorderLeftEnd+1])
		//
		//fmt.Printf("Right tree size: %d \n", rightTreeSize)
		//fmt.Printf("Right tree inorder: %d \n", inorder[inorderRightStart:inorderRightEnd+1])
		//fmt.Printf("Right tree postorder: %d \n", postorder[postorderRightStart:postorderRightEnd+1])

		leftTree := dfs(inorderLeftStart, inorderLeftEnd, postorderLeftStart, postorderLeftEnd)
		rightTree := dfs(inorderRightStart, inorderRightEnd, postorderRightStart, postorderRightEnd)

		root := &TreeNode{rootValue, leftTree, rightTree}

		return root
	}

	return dfs(0, len(inorder)-1, 0, len(postorder)-1)
}

func test(inorder []int, postorder []int, expectedResult []any) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Inorder traversal: %v \n", inorder)
	fmt.Printf("Postorder traversal: %v \n", postorder)

	result := buildTree(inorder, postorder)

	fmt.Printf("Built tree: \n")
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
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	expected := []any{3, 9, 20, nil, nil, 15, 7}

	test(inorder, postorder, expected)
}

func test2() {
	inorder := []int{2, 1}
	postorder := []int{2, 1}
	expected := []any{1, 2, nil}

	test(inorder, postorder, expected)
}

func test3() {
	inorder := []int{-1}
	postorder := []int{-1}
	expected := []any{-1}

	test(inorder, postorder, expected)
}

func main() {
	// 106. Construct Binary Tree from Inorder and Postorder Traversal
	test1()
	test2()
	test3()
}
