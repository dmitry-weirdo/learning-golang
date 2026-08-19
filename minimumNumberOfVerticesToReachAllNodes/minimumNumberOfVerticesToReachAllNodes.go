package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	// since this is a DGA, we're just collecting the vertices without inbound edges
	hasInboundEdges := make([]bool, n)

	to := 0

	for _, v := range edges {
		to = v[1] // we only care about to, not from
		hasInboundEdges[to] = true
	}

	// collect nodes without incoming edges
	result := make([]int, 0)

	for i, v := range hasInboundEdges {
		if !v {
			result = append(result, i)
		}
	}

	return result
}

func test(n int, m [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of vertices from 0 to %v: %v \n", n-1, n)
	fmt.Printf("Edges: %v \n", m)

	result := findSmallestSetOfVertices(n, m)

	fmt.Printf("Vertices without incoming edges: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	n := 6

	edges := [][]int{
		{0, 1},
		{0, 2},
		{2, 5},
		{3, 4},
		{4, 2},
	}

	expected := []int{0, 3}

	test(n, edges, expected)
}

func test2() {
	n := 5

	edges := [][]int{
		{0, 1},
		{2, 1},
		{3, 1},
		{1, 4},
		{2, 4},
	}

	expected := []int{0, 2, 3}

	test(n, edges, expected)
}

func main() {
	// 1557. Minimum Number of Vertices to Reach All Nodes
	test1()
	test2()
}
