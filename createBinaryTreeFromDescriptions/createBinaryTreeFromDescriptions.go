package main

import (
	"demo/trees"
	. "demo/trees" // not recommended, but ok for LeetCode -> to use TreeNode without package prefix
	"fmt"
)

func createBinaryTree(descriptions [][]int) *TreeNode {
	// todo: we can convert to adj list and execute DFS/BFS

	// using 2 hashmaps
	// slow - passes in 77-125 ms
	return createBinaryTree_maps(descriptions)
}

func createBinaryTree_maps(descriptions [][]int) *TreeNode {
	// values are unique, so we can use them as node identifiers
	m := make(map[int]*TreeNode) // value to node
	p := make(map[int]*TreeNode) // value to parent

	var parentNode *TreeNode
	var childNode *TreeNode

	// build tree
	for _, v := range descriptions {
		parent := v[0]
		child := v[1]
		left := v[2] == 1

		// get parent node
		if _, ok := m[parent]; !ok {
			// parent does not exist -> create it
			parentNode = &TreeNode{Val: parent, Left: nil, Right: nil}
			m[parent] = parentNode
		}

		parentNode = m[parent]

		// get child node
		if _, ok := m[child]; !ok {
			// child does not exist -> create it
			childNode = &TreeNode{Val: child, Left: nil, Right: nil}
			m[child] = childNode
		}

		childNode = m[child]

		// link child to parent
		if left {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		// save child -> parent link
		p[child] = parentNode
	}

	// find the root -> only node without parent
	for value, node := range m {
		// update the current root
		if _, ok := p[value]; !ok {
			return node
		}
	}

	panic("Root not found. This must never happen.")
}

func test(m [][]int, expectedResult []any) { // nodes can be nil
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m)

	result := createBinaryTree(m)

	fmt.Printf("Tree from descriptions: \n")
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
	//Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
	//Output: [50,20,80,15,17,19]

	m := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}

	expected := []any{
		50,
		20, 80,
		15, 17, 19, nil,
	}

	test(m, expected)
}

func test2() {
	//Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
	//Output: [1,2,nil,nil,3,4]

	m := [][]int{
		{1, 2, 1},
		{2, 3, 0},
		{3, 4, 1},
	}

	expected := []any{
		1,
		2, nil,
		nil, 3, nil, nil,
		nil, nil, 4, nil, nil, nil, nil, nil,
	}

	test(m, expected)
}

func test3() {
	// Failed test-case 29/85 -> probably root set incorrectly

	m := [][]int{
		{85, 74, 0},
		{38, 82, 0},
		{39, 70, 0},
		{82, 85, 0},
		{74, 13, 0},
		{13, 39, 0},
	}

	expected := []any{
		38,
		nil, 82,
		nil, nil, nil, 85,
		nil, nil, nil, nil, nil, nil, nil, 74,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 13,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 39,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 70,
	}

	test(m, expected)
}

func main() {
	// 2196. Create Binary Tree From Descriptions
	test1()
	test2()
	test3()
}
