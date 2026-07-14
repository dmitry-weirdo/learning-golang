package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

// map of value to index in the inorder array
var inorderMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	// build the inorder map
	inorderMap = make(map[int]int)

	for i, v := range inorder {
		inorderMap[v] = i
	}

	return dfs(preorder, inorder, 0, 0, len(preorder))
}

func dfs(preorder []int, inorder []int, preorderIndex, inorderIndex, subtreeSize int) *TreeNode {
	fmt.Println()
	fmt.Println("=================")
	fmt.Printf("preorder index: %v, inorder index: %v, subtree size: %v \n", preorderIndex, inorderIndex, subtreeSize)

	if subtreeSize <= 0 { // reached the empty subtree
		fmt.Printf("Reached an empty subtree, returning nil. \n")
		return nil
	}

	// in the given subtree range, first element in pre-order array is the root
	rootValue := preorder[preorderIndex]
	fmt.Printf("Root value is preorder[%v] = %v \n", preorderIndex, rootValue)

	// Find the current subtree root in the inorder array
	rootInorderIndex := inorderMap[rootValue]
	fmt.Printf("Root value %v is inorder[%v]. \n", rootValue, rootInorderIndex)

	// get left subtree size from the inorder array -> all elements from the start of the current subtree up to the root (non-inclusive)
	leftSubtreeSize := rootInorderIndex - inorderIndex
	fmt.Printf("Left subtree size is rootInorderIndex - inorderIndex = %v - %v = %v \n", rootInorderIndex, inorderIndex, leftSubtreeSize)

	// subtree in the inorder array looks like:
	// [ leftSubtree, root, rightSubtree]
	// subtreeSize = leftSubtreeSize + rootSize + rightSubtreeSize, root size = 1
	rightSubtreeSize := subtreeSize - leftSubtreeSize - 1
	fmt.Printf("Right subtree size is subtreeSize - leftSubtreeSize - rootSize = %v - %v - 1 = %v \n", subtreeSize, leftSubtreeSize, rightSubtreeSize)

	// in preorder, element next to the root is the root of the left subtree (we just have to skip the root)
	leftSubtreeRootIndexInPreorder := preorderIndex + 1 // this can go out of bounds
	if leftSubtreeRootIndexInPreorder < len(preorder) {
		fmt.Printf("Left subtree root is preorder[%v] = %v \n", leftSubtreeRootIndexInPreorder, preorder[leftSubtreeRootIndexInPreorder])
	}

	// in preorder, we have to skip the left subtree size and the root size (1)
	rightSubtreeRootIndexInPreorder := preorderIndex + 1 + leftSubtreeSize // this can go out of bounds
	if rightSubtreeRootIndexInPreorder < len(preorder) {
		fmt.Printf("Right subtree root is preorder[%v] = %v \n", rightSubtreeRootIndexInPreorder, preorder[rightSubtreeRootIndexInPreorder]) // todo: this can go out of bounds?
	}

	// for left subtree, inorder start index stays the same
	leftSubtreeLeftIndexInInorder := inorderIndex // this can go out of bounds
	if leftSubtreeLeftIndexInInorder < len(inorder) {
		fmt.Printf("Left subtree start position is inorder[%v] = %v \n", leftSubtreeLeftIndexInInorder, inorder[leftSubtreeLeftIndexInInorder])
	}

	// for right subtree, inorder start index is right after the root position
	rightSubtreeLeftIndexInInorder := rootInorderIndex + 1 // this can go out of bounds
	if rightSubtreeLeftIndexInInorder < len(inorder) {
		fmt.Printf("Right subtree start position is inorder[%v] = %v \n", rightSubtreeLeftIndexInInorder, inorder[rightSubtreeLeftIndexInInorder])
	}

	// just debugging purposes:
	fmt.Printf("Left subtree from inorder: %v \n", inorder[leftSubtreeLeftIndexInInorder:leftSubtreeLeftIndexInInorder+leftSubtreeSize])
	fmt.Printf("Right subtree from inorder: %v \n", inorder[rightSubtreeLeftIndexInInorder:rightSubtreeLeftIndexInInorder+rightSubtreeSize])

	// call for the left subtree
	leftSubtree := dfs(preorder, inorder, leftSubtreeRootIndexInPreorder, leftSubtreeLeftIndexInInorder, leftSubtreeSize)
	//var leftSubtree *TreeNode = nil

	// call for the right subtree
	rightSubtree := dfs(preorder, inorder, rightSubtreeRootIndexInPreorder, rightSubtreeLeftIndexInInorder, rightSubtreeSize)
	//var rightSubtree *TreeNode = nil

	return &TreeNode{
		Val:   rootValue,
		Left:  leftSubtree,
		Right: rightSubtree,
	}
}

func test(preorder []int, inorder []int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("preorder: %v \n", preorder)
	fmt.Printf("inorder: %v \n", inorder)

	result := buildTree(preorder, inorder) // iterative

	fmt.Printf("Tree built: \n")
	trees.PrintTreeTopDown(result)
}

func test1() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	test(preorder, inorder)
}

func test2() {
	preorder := []int{-1}
	inorder := []int{-1}

	test(preorder, inorder)
}

func main() {
	// 105. Construct Binary Tree from Preorder and Inorder Traversal
	test1()
	test2()
}
