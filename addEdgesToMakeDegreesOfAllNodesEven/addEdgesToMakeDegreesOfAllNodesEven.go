package main

import "slices"

import "fmt"

func isPossible(n int, edges [][]int) bool {
	adj := make([][]int, n+1) // node numbering starts from 1

	v1 := 0
	v2 := 0

	for _, v := range edges {
		v1 = v[0]
		v2 = v[1]

		// add v2 to v1
		if adj[v1] == nil {
			adj[v1] = []int{v2}
		} else {
			adj[v1] = append(adj[v1], v2)
		}

		// add v1 to v2
		if adj[v2] == nil {
			adj[v2] = []int{v1}
		} else {
			adj[v2] = append(adj[v2], v1)
		}
	}

	fmt.Printf("Adjacency list: %v \n", adj)

	// collect odd nodes
	oddNodes := make([]int, 0)

	for node, neighbours := range adj {
		if len(neighbours)%2 == 1 {
			oddNodes = append(oddNodes, node)

			if len(oddNodes) > 4 { // > 4 nodes cannot be solved with just 2 edges
				return false
			}
		}
	}

	fmt.Printf("Total %d odd nodes: %v \n", len(oddNodes), oddNodes)

	totalOddNodes := len(oddNodes)
	if totalOddNodes == 0 { // the graph already has only the even nodes -> nothing to do
		return true
	}

	if totalOddNodes == 2 {
		odd1 := oddNodes[0]
		odd2 := oddNodes[1]

		// if we have no edge between them -> we will just connect them
		if !hasEdge(adj, odd1, odd2) {
			return true
		}

		// If we have an edge between them -> we need to find other (even) node not connected to any of the odd nodes
		// If this node exists, we will connect both odd nodes to this node.

		// We cannot use different "not connected to odd" nodes, since if we connect from odd to separate even nodes (using both 2 allowed edges),
		// these even nodes will become odd.
		for i := 1; i <= n; i++ {
			if (i != odd1) &&
				(i != odd2) &&
				!hasEdge(adj, i, odd1) &&
				!hasEdge(adj, i, odd2) {
				return true
			}
		}

		return false
	}

	if totalOddNodes == 4 {
		odd1 := oddNodes[0]
		odd2 := oddNodes[1]
		odd3 := oddNodes[2]
		odd4 := oddNodes[3]

		// We only have 2 edges to fix and 4 nodes to fix
		// -> we need to connect 2 different pairs from these 4 nodes
		// If we cannot select 2 pairs that don't have edges between them, it's impossible to fix.

		// Pairs can be:
		// 1, 2 - 3, 4
		// 1, 3 - 2, 4
		// 1, 4 - 2, 3
		if (!hasEdge(adj, odd1, odd2) && !hasEdge(adj, odd3, odd4)) || // 1, 2 - 3, 4
			(!hasEdge(adj, odd1, odd3) && !hasEdge(adj, odd2, odd4)) || // 1, 3 - 2, 4
			(!hasEdge(adj, odd1, odd4) && !hasEdge(adj, odd2, odd3)) { // 1, 4 - 2, 3
			return true
		}

		return false
	}

	// having 1 or 3 odd nodes in an undirected graph is impossible
	return false
}

func hasEdge(adj [][]int, from, to int) bool {
	return slices.Contains(adj[from], to)
}

func test(n int, edges [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - number of edges from 0 to %v: %v \n", n-1, n)
	fmt.Printf("Edges: %v \n", edges)

	result := isPossible(n, edges)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// 2 odd nodes 4 and 5, no edge between them
	n := 5

	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 2},
		{1, 4},
		{2, 5},
	}

	expected := true

	test(n, edges, expected)
}

func test2() {
	// 4 odd nodes, we can connect pairs of 1-3 and 2-4 (or 1-4 and 2-3)
	n := 4

	edges := [][]int{
		{1, 2},
		{3, 4},
	}

	expected := true

	test(n, edges, expected)
}

func test3() {
	// 4 odd nodes, we can NOT connect different pairs of them
	n := 4

	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}

	expected := false

	test(n, edges, expected)
}

func main() {
	// 2508. Add Edges to Make Degrees of All Nodes Even
	test1()
	test2()
	test3()
}
