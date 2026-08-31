package main

import (
	"fmt"
)

type Color int

const (
	NONE Color = iota
	RED
	BLUE
)

type NodeEdge struct {
	node  int
	color Color // previous edge color
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	adjRed := createAdjacencyListDirected(n, redEdges)
	adjBlue := createAdjacencyListDirected(n, blueEdges)

	//fmt.Printf("Adjacency list red: %v \n", adjRed)
	//fmt.Printf("Adjacency list blue: %v \n", adjBlue)

	visitedFromRed := make(map[int]bool)  // nodes visited by red incoming edge
	visitedFromBlue := make(map[int]bool) // nodes visited by blue incoming edge

	result := createIntArrayWithDefaultValues(n, -1) // set defaults to -1

	queue := make([]*NodeEdge, 0)
	queue = append(queue, &NodeEdge{0, NONE}) // we're starting from node 0
	// !!! we do NOT add 0 as visited

	level := 0

	for len(queue) > 0 {
		currentLevelSize := len(queue)

		for range currentLevelSize {
			// remove from queue
			node := queue[0]
			queue = queue[1:]

			//fmt.Printf("Level: %v, node: %v \n", level, node)

			// mark current as found distance
			if result[node.node] == -1 { // only if it wasn't visited with another incoming edge color already
				result[node.node] = level
			}

			if node.color == RED { // next edge must be blue
				visitedFromRed[node.node] = true

				// go next with blue edge
				for _, blueNeighbor := range adjBlue[node.node] {
					if visitedFromBlue[blueNeighbor] { // node already visited from blue edge -> skip
						continue
					}

					visitedFromBlue[blueNeighbor] = true
					queue = append(queue, &NodeEdge{blueNeighbor, BLUE})
				}
			} else if node.color == BLUE { // next edge must be red
				visitedFromBlue[node.node] = true

				// go next with red edge
				for _, redNeighbor := range adjRed[node.node] {
					if visitedFromRed[redNeighbor] { // node already visited from red edge -> skip
						continue
					}

					visitedFromRed[redNeighbor] = true
					queue = append(queue, &NodeEdge{redNeighbor, RED})
				}
			} else { // visited from NONE (0-th node) -> try both red and blue edges
				visitedFromRed[node.node] = true // we already run here from 0 with both red and blue edges, so no runs on node 0 anymore

				// go next with blue edge
				for _, blueNeighbor := range adjBlue[node.node] {
					if visitedFromBlue[blueNeighbor] { // node already visited from blue edge -> skip
						continue
					}

					visitedFromBlue[blueNeighbor] = true
					queue = append(queue, &NodeEdge{blueNeighbor, BLUE})
				}

				visitedFromBlue[node.node] = true // we already run here from 0 with both red and blue edges, so no runs on node 0 anymore

				// go next with red edge
				for _, redNeighbor := range adjRed[node.node] {
					if visitedFromRed[redNeighbor] { // node already visited from red edge -> skip
						continue
					}

					visitedFromRed[redNeighbor] = true
					queue = append(queue, &NodeEdge{redNeighbor, RED})
				}
			}
		}

		level++
	}

	return result
}

func createAdjacencyListDirected(n int, edges [][]int) [][]int {
	adj := make([][]int, n)

	from := 0
	to := 0

	for _, v := range edges {
		from = v[0]
		to = v[1]

		// add v2 to v1
		if adj[from] == nil {
			adj[from] = []int{to}
		} else {
			adj[from] = append(adj[from], to)
		}
	}

	return adj
}

func createIntArrayWithDefaultValues(n int, defaultValue int) []int {
	a := make([]int, n)

	for i := range n {
		a[i] = defaultValue
	}

	return a
}

func test(n int, red, blue [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Red edges: %v \n", red)
	fmt.Printf("Blue edges: %v \n", blue)

	result := shortestAlternatingPaths(n, red, blue)

	fmt.Printf("Distances (count of nodes) from 0 to i-th node: %v \n", result)
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
	n := 3

	red := [][]int{
		{0, 1},
		{1, 2},
	}

	blue := [][]int{}

	expected := []int{0, 1, -1}

	test(n, red, blue, expected)
}

func test2() {
	n := 3

	red := [][]int{
		{0, 1},
	}

	blue := [][]int{
		{2, 1},
	}

	expected := []int{0, 1, -1}

	test(n, red, blue, expected)
}

func test3() {
	// failing test-case 23/91 - do NOT visit node 0 again!
	n := 3

	red := [][]int{
		{0, 1},
		{0, 2},
	}

	blue := [][]int{
		{1, 0},
	}

	expected := []int{0, 1, 1}

	test(n, red, blue, expected)
}

func test4() {
	// failing test-case 35/91
	// Path to node 4 is: 0-1-2-3-1-2-3-4
	n := 5

	red := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}

	blue := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}

	expected := []int{0, 1, 2, 3, 7}

	test(n, red, blue, expected)
}

func main() {
	// 1129. Shortest Path with Alternating Colors
	test1()
	test2()
	test3()
	test4()
}
