package main

import (
	"fmt"
)

func validPath(n int, edges [][]int, source int, destination int) bool {
	// union-find should be the fastest!
	// yes, it's very fast, passes in 12-25 ms
	return validPath_unionFind(n, edges, source, destination)

	// todo: we can do iterative DFS with fast return

	// BFS is much faster
	// passes in 100-110 ms
	//return validPath_bfs(n, edges, source, destination)

	// trivial DFS, no backtracking, marking visited and moving to visited again
	// passes in 215-280 ms
	//return validPath_dfs_recursive(n, edges, source, destination)
}

func validPath_unionFind(n int, edges [][]int, source int, destination int) bool {
	uf := newUnionFind(n)

	for _, v := range edges {
		uf.union(v[0], v[1])
	}

	return uf.find(source) == uf.find(destination)
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

type UnionFind struct {
	parents []int // if parent[i] = i, it is the root, else it's the index of the parent
	sizes   []int // sizes of the tree for every element
}

func newUnionFind(n int) UnionFind {
	parents := make([]int, n)
	sizes := make([]int, n)

	for i := range n {
		// every group is just a root
		parents[i] = i

		// every group has a size of 1
		sizes[i] = 1
	}

	return UnionFind{
		parents: parents,
		sizes:   sizes,
	}
}

func (uf UnionFind) find(x int) int { // recursive version
	if uf.parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	uf.parents[x] = uf.find(uf.parents[x])

	return uf.parents[x]
}

func (uf UnionFind) print() {
	fmt.Printf("Parents: %v \n", uf.parents)
	fmt.Printf("Sizes: %v \n", uf.sizes)
}

func (uf UnionFind) union(x, y int) bool { // returns false if they're already in the same set
	// these find will perform path compression
	rootX := uf.find(x)
	rootY := uf.find(y)

	//fmt.Printf("root of %d: %d, root of %d: %d\n", x, rootX, y, rootY)

	// x and y are already in the same set -> nothing to merge
	if rootX == rootY {
		//fmt.Printf("Element %v and %v already belong to the same root %v. Nothing to merge. \n", x, y, rootX)
		return false
	}

	// merge the smaller group into the bigger group
	// todo: ideally, we should merge the tree with smaller depth into the tree with bigger depth
	if uf.sizes[rootX] < uf.sizes[rootY] { // merge x into y
		//fmt.Printf("sizes[%v] = %v < sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootX, rootY)

		uf.parents[rootX] = rootY
		uf.sizes[rootY] += uf.sizes[rootX]
	} else { // merge y into x
		//fmt.Printf("sizes[%v] = %v >= sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootY, rootX)

		uf.parents[rootY] = rootX
		uf.sizes[rootX] += uf.sizes[rootY]
	}

	return true
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
