package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	// todo: union-find should be the fastest!

	// todo: we can do iterative DFS with fast return

	// BFS is much faster
	// passes in 100-110 ms
	return validPath_bfs(n, edges, source, destination)

	// trivial DFS, no backtracking, marking visited and moving to visited again
	// passes in 215-280 ms
	//return validPath_dfs_recursive(n, edges, source, destination)
}

func validPath_bfs(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	adj := createAdjacencyListUndirected(n, edges)
	//fmt.Printf("Adjacency list: %v \n", adj)

	visited := make(map[int]bool)
	visited[source] = true

	queue := []int{source}
	level := 0

	for len(queue) > 0 {
		level++
		currentLevelNodesCount := len(queue)

		for range currentLevelNodesCount {
			// pop from queue
			current := queue[0]
			queue = queue[1:] // O(1), does NOT copy the array, just moves the pointer

			if current == destination { // dest found -> immediately return
				return true
			}

			neighbours := adj[current]

			for _, v := range neighbours {
				if v == destination { // cut earlier, no need to further process even the current level, if reached the dest
					return true
				}

				if visited[v] {
					continue
				}

				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return false
}

func validPath_dfs_recursive(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	adj := createAdjacencyListUndirected(n, edges)
	//fmt.Printf("Adjacency list: %v \n", adj)

	visited := make(map[int]bool)
	visited[source] = true

	destinationFound := false

	var dfs func(current int) // sets destinationFound, returns nothing

	dfs = func(current int) {
		if destinationFound { // already found -> nothing to do, stop iteration
			return
		}

		if current == destination { // reached the destination -> mark found, stop iteration
			destinationFound = true
			return
		}

		neighbours := adj[current]

		for _, v := range neighbours {
			if visited[v] {
				continue
			}

			visited[v] = true
			dfs(v)

			// no backtracking
		}
	}

	dfs(source)

	return destinationFound
}

func createAdjacencyListUndirected(n int, edges [][]int) [][]int {
	adj := make([][]int, n)

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

	return adj
}

func test(n int, edges [][]int, source int, destination int, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N (total vertices in the graph): %v \n", n)
	fmt.Printf("Edges: %v \n", edges)
	fmt.Printf("Source node: %v \n", source)
	fmt.Printf("Destination node: %v \n", destination)

	result := validPath(n, edges, source, destination)

	fmt.Printf("Possible to go from %v to %v: %v \n", source, destination, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 3

	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}

	source := 0
	destination := 2
	expected := true

	test(n, edges, source, destination, expected)
}

func test2() {
	n := 6

	edges := [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}

	source := 0
	destination := 5
	expected := false

	test(n, edges, source, destination, expected)
}

func main() {
	// 1971. Find if Path Exists in Graph
	test1()
	test2()
}
