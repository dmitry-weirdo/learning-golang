package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
	"strconv"
)

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
		minNode, minNodeParent := getMinNodeAndItsParent(root.Right, root)
		fmt.Printf("Found min node %v in the right subtree of node %v. Min node parent: %v \n", minNode.Val, root.Val, getNodeValSafe(minNodeParent))

		// replace the root value with the min node value
		root.Val = minNode.Val
		fmt.Printf("Found deleted node with the min node value %v. \n", minNode.Val)

		// remove the old min node from the right subtree. Its value is now in the being removed node.
		//root.Right = deleteNode(root.Right, minNode.Val)

		if minNodeParent != root {
			minNodeParent.Left = deleteNode(minNodeParent.Left, minNode.Val)
		} else {
			minNodeParent.Right = deleteNode(minNodeParent.Right, minNode.Val)
		}

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

func getMinNodeAndItsParent(root *TreeNode, rootParent *TreeNode) (minNode *TreeNode, minNodeParent *TreeNode) {
	if root == nil {
		return nil, nil
	}

	// to find the min, we constantly move to the left
	var parent = rootParent
	node := root

	for node.Left != nil {
		parent = node
		node = node.Left
	}

	return node, parent
}

func test(arr []any, key int, expectedResult []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Tree array: %v \n", arr)

	tree := trees.TreeFromArray(arr)
	fmt.Printf("Initial tree: \n")
	trees.PrintTreeTopDown(tree)

	fmt.Printf("Key to delete: %v \n", key)

	result := deleteNode(tree, key)

	fmt.Printf("Tree with node %v deleted: \n", key)
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
		[]any{50, 30, 70, nil, 40, 60, 80},
		50,
		[]any{60, 30, 70, nil, 40, nil, 80},
	)
}

func test2() {
	test(
		[]any{5, 3, 6, 2, 4, nil, 7},
		3,
		[]any{5, 4, 6, 2, nil, nil, 7},
	)
}

func test3() {
	test(
		[]any{5, 3, 6, 2, 4, nil, 7},
		0, // value not in the tree -> no changes
		[]any{5, 3, 6, 2, 4, nil, 7},
	)
}

func main() {
	// 450. Delete Node in a BST
	test1()
	test2()
	test3()
}
