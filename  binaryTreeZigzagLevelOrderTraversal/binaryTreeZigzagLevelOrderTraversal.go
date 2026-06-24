package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { // need to use pointer to modify the Stack, else it will be a copy
	s.data = append(s.data, v) // append to end
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) <= 0 {
		var zeroValue T
		return zeroValue, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[0:lastIndex] // remove the last element

	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var s1 Stack[*TreeNode]
	var s2 Stack[*TreeNode]

	var result [][]int
	// rowIndex := 0

	if root == nil {
		return nil
	}

	s1.Push(root)

	for !s1.IsEmpty() || !s2.IsEmpty() {
		var row1 []int

		for !s1.IsEmpty() {
			var node *TreeNode
			node, _ = s1.Pop()

			if node.Left != nil {
				// fmt.Printf("node.Left: %v %v \n", node.Left, node.Left.Val)

				s2.Push(node.Left)
			}

			if node.Right != nil {
				// fmt.Printf("node.Right: %v %v \n", node.Right, node.Right.Val)
				s2.Push(node.Right)
			}

			// fmt.Println(node.Val)
			row1 = append(row1, node.Val)
		}

		if len(row1) > 0 {
			result = append(result, row1)
			// result[rowIndex] = row1
			// rowIndex++
		}

		var row2 []int

		for !s2.IsEmpty() {
			node, _ := s2.Pop()

			if node.Right != nil {
				s1.Push(node.Right)
			}

			if node.Left != nil {
				s1.Push(node.Left)
			}

			// fmt.Println(node.Val)
			row2 = append(row2, node.Val)
		}

		if len(row2) > 0 {
			result = append(result, row2)
			// result[rowIndex] = row2
			// rowIndex++
		}
	}

	return result
}

func main() {
	r11 := &TreeNode{15, nil, nil}
	r12 := &TreeNode{7, nil, nil}

	r01 := &TreeNode{9, nil, nil}
	r02 := &TreeNode{20, r11, r12}

	root := &TreeNode{3, r01, r02}

	result := zigzagLevelOrder(root)

	fmt.Printf("Result: %v \n", result)
}
