package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result = make([]int, 0)

	if root == nil {
		return nil
	}

	queue := list.New()
	appendToQueue(queue, root)

	level := 0

	for queue.Len() > 0 {
		currentLevelElements := queue.Len()
		level++

		mostRightOnCurrentLevel := getFirstFromQueue(queue)
		result = append(result, mostRightOnCurrentLevel.Val)
		fmt.Printf("Most right node on level %v: %v \n", level, mostRightOnCurrentLevel.Val)

		for i := 0; i < currentLevelElements; i++ {
			node := removeFirstFromQueue(queue)
			fmt.Printf("Level: %v, node: %v \n", level, node.Val)

			// we handle right first
			if node.Right != nil {
				appendToQueue(queue, node.Right)
			}

			if node.Left != nil {
				appendToQueue(queue, node.Left)
			}
		}
	}

	return result
}

func appendToQueue(queue *list.List, s *TreeNode) {
	queue.PushBack(s)
}

func removeFirstFromQueue(queue *list.List) *TreeNode {
	return queue.Remove(queue.Front()).(*TreeNode)
}

func removeLastFromQueue(queue *list.List) *TreeNode {
	return queue.Remove(queue.Back()).(*TreeNode)
}

func getFirstFromQueue(queue *list.List) *TreeNode {
	return queue.Front().Value.(*TreeNode)
}

func main() {
	var root1 *TreeNode

	root1 = getTest1()
	root1 = getTest2()

	rightSideView(root1)
}

func getTest1() *TreeNode {
	node5RightOf2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node2LeftOf1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: node5RightOf2,
	}

	node4RightOf3 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node3RightOf1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: node4RightOf3,
	}

	root1 := &TreeNode{
		Val:   1,
		Left:  node2LeftOf1,
		Right: node3RightOf1,
	}

	return root1
}

func getTest2() *TreeNode {
	node5LeftOf4 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node4LeftOf2 := &TreeNode{
		Val:   4,
		Left:  node5LeftOf4,
		Right: nil,
	}

	node2LeftOf1 := &TreeNode{
		Val:   2,
		Left:  node4LeftOf2,
		Right: nil,
	}

	node3RightOf1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root1 := &TreeNode{
		Val:   1,
		Left:  node2LeftOf1,
		Right: node3RightOf1,
	}

	return root1
}
