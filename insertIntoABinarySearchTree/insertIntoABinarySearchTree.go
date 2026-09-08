package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//return insertIterative(root, val)
	return insertRecursive(root, val)
}

func insertRecursive(root *TreeNode, val int) *TreeNode {
	// this is a trivial solution without any rebalancing
	// it also implies that the value does not exist in the tree (it's the problem condition).
	if root == nil { // this also handles the case if initial root was null
		// reached leaf position -> add a new node
		return &TreeNode{Val: val}
	}

	if val < root.Val { // go left
		// will only change if root.Left was null -> then the new node will be set
		root.Left = insertRecursive(root.Left, val)
	} else if val > root.Val { // go right
		// will only change if root.Right was null -> then the new node will be set
		root.Right = insertRecursive(root.Right, val)
	}

	// root was existing -> return it, so that it doesn't change in the parent
	return root
}

func insertIterative(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	node := root
	//parent := nil

	newNode := &TreeNode{Val: val}

	for node != nil {
		// parent = node

		if val < node.Val { // go left
			if node.Left == nil {
				node.Left = newNode
				break
			} else {
				node = node.Left
			}
		} else if val > node.Val { // go right
			if node.Right == nil {
				node.Right = newNode
				break
			} else {
				node = node.Right
			}
		}
	}

	return root
}

func test(arr []any, val int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	fmt.Printf("Value to insert: %v \n", val)

	result := insertIntoBST(tree, val)

	fmt.Printf("Tree with inserted node %v: \n", val)
	trees.PrintTreeTopDown(result)

	resultAsArray := trees.TreeToArray(result)
	fmt.Printf("Result tree as array: %v \n", resultAsArray)
	fmt.Printf("Expected result:      %v \n", expectedResult)

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
	test(
		[]any{4, 2, 7, 1, 3, nil, nil},
		5,
		[]any{4, 2, 7, 1, 3, 5, nil},
	)
}

func test2() {
	test(
		[]any{40, 20, 60, 10, 30, 50, 70},
		25,
		[]any{40, 20, 60, 10, 30, 50, 70, nil, nil, 25, nil, nil, nil, nil, nil},
	)
}

func test3() {
	test(
		[]any{4, 2, 7, 1, 3, nil, nil},
		5,
		[]any{4, 2, 7, 1, 3, 5, nil},
	)
}

func main() {
	// 701. Insert into a Binary Search Tree
	test1()
	test2()
	test3()
}
