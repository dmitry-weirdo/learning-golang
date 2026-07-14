package main

import (
	"container/list"
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
	"fmt"
)

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	// todo: we can use a slice instead of list.List to decrease the memory overhead
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		currentLevelCount := queue.Len()
		currentLevel := make([]int, currentLevelCount)

		for i := 0; i < currentLevelCount; i++ {
			// handle just the elements of the current level
			frontElement := queue.Front()
			queue.Remove(frontElement)

			node := frontElement.Value.(*TreeNode)

			currentLevel[i] = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func test(arr []any) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	root := trees.TreeFromArray(arr)
	fmt.Printf("Tree: \n")
	trees.PrintTreeTopDown(root)

	result := levelOrder(root)

	fmt.Printf("BFS level-by-level traversal: %v \n", result)
}

func test1() {
	arr := []any{3, 9, 20, nil, nil, 15, 7}

	test(arr)
}

func test2() {
	arr := []any{1}

	test(arr)
}

func test3() {
	arr := []any{}

	test(arr)
}

func main() {
	// 102. Binary Tree Level Order Traversal
	test1()
	test2()
	test3()
}
