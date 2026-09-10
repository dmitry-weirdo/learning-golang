package main

import (
	"cmp"
	"fmt"
	"slices"
)

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	// house nodes start with 0

	// We add node 0 that has wells[i - 1] to connect to house[i]
	// This node 0 must be connected to at least 1 house (to have water at all),
	// therefore this node 0 is also a part of the MST that we need to find.
	const WATER_SOURCE_NODE_INDEX = 0

	for i := 1; i <= n; i++ {
		from := WATER_SOURCE_NODE_INDEX
		to := i
		weight := wells[i-1]

		pipes = append(pipes, []int{from, to, weight})
	}

	//fmt.Printf("Pipes with added edges from water source %v: %v \n", WATER_SOURCE_NODE_INDEX, pipes)

	// now we're just executing the MST algorithm. Pipes is edges.
	mst := getMinimumSpanningTreeKruskalNodesStartFrom0(n+1, pipes)

	// since we can connect the water source to every node, we can always connect to all the nodes from this water source node.
	return mst.weight
}

// ========================= Kruskal's algorithm for MST begin ========================= //
type MinimumSpanningTree struct {
	edges        [][]int // every edge is from/to/weight
	weight       int     // total weight for all edges
	allNodesUsed bool    // false if there are multiple non-connected components
}

func getMinimumSpanningTreeKruskalNodesStartFrom0(n int, edges [][]int) MinimumSpanningTree {
	return getMinimumSpanningTreeKruskal(n, 0, edges) // start from node 0
}

func getMinimumSpanningTreeKruskalNodesStartFrom1(n int, edges [][]int) MinimumSpanningTree {
	return getMinimumSpanningTreeKruskal(n+1, 1, edges) // start from node 1
}

func getMinimumSpanningTreeKruskal(n int, startIndex int, edges [][]int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected. It's necessary for Kruskal's since we're uniting 2 nodes with Union-Find

	// Instead of using a complex Heap structure for getting the smallest edge, we can just sort an array of edges by weight
	// Complexity should be the save O(E * log E)
	sortEdgesForKruskal(edges)

	// union-find controls the visited by including them in the same node set
	uf := NewUnionFind(n)

	mst := MinimumSpanningTree{
		edges:  make([][]int, 0),
		weight: 0,
	}

	// index in edges[] array
	i := 0

	for (len(mst.edges) < n-startIndex-1) && (i < len(edges)) {
		from := edges[i][0]
		to := edges[i][1]
		weight := edges[i][2]

		if !uf.Union(from, to) { // from and to already in the same MST set -> skip this edge
			i++
			continue
		}

		// Add the current edge to MST
		mst.edges = append(mst.edges, []int{from, to, weight})

		// Add the edge weight to total weight
		mst.weight += weight

		// go to next edge
		i++
	}

	// check whether all nodes are in the single MST component
	// We will have (N - 1) nodes in the MST in this case.
	mst.allNodesUsed = len(mst.edges) == (n - startIndex - 1)

	return mst
}

func sortEdgesForKruskal(edges [][]int) {
	slices.SortFunc(edges, func(a, b []int) int {
		// from = v[0]
		// to = v[1]
		// weight = v[2]

		if a[2] != b[2] { // different weights -> compare by weights
			return cmp.Compare(a[2], b[2])
		}

		if a[0] != b[0] { // different from -> compare by from
			return cmp.Compare(a[0], b[0])
		}

		// compare by to
		return cmp.Compare(a[1], b[1])
	})
}

type UnionFind struct {
	parents []int // if parent[i] = i, it is the root, else it's the index of the parent
	sizes   []int // sizes of the tree for every element
}

func NewUnionFind(n int) UnionFind {
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

func (uf UnionFind) Find(x int) int { // recursive version
	if uf.parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	uf.parents[x] = uf.Find(uf.parents[x])

	return uf.parents[x]
}

func (uf UnionFind) Print() {
	fmt.Printf("Parents: %v \n", uf.parents)
	fmt.Printf("Sizes: %v \n", uf.sizes)
}

func (uf UnionFind) Union(x, y int) bool { // returns false if they're already in the same set
	// these find will perform path compression
	rootX := uf.Find(x)
	rootY := uf.Find(y)

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

func (uf UnionFind) GroupSize(x int) int {
	return uf.sizes[uf.Find(x)]
}

func (uf UnionFind) GetGroupsSizes() map[int]int { // returns sizes for every group
	m := make(map[int]int)

	for i := range uf.parents {
		if uf.Find(i) == i { // root node
			// every root group will be iterated just once, no need to check whether it's already in the map
			m[i] = uf.GroupSize(i)

			/*
				if _, ok := m[i]; !ok {
					// group not yet in map -> add it
					m[i] = uf.groupSize(i)
				}
			*/
		}
	}

	return m
}

// ========================= Kruskal's algorithm for MST end ========================= //

func test(n int, wells []int, pipes [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of houses: %v \n", n)
	fmt.Printf("Wells cost: %v \n", wells)
	fmt.Printf("Edges: %v \n", pipes)

	result := minCostToSupplyWater(n, wells, pipes)

	fmt.Printf("Min cost to supply all houses with water: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 3

	wells := []int{1, 2, 3}

	pipes := [][]int{
		{1, 2, 1},
		{2, 3, 1},
	}

	expected := 3

	test(n, wells, pipes, expected)
}

func test2() {
	n := 2

	wells := []int{1, 1}

	pipes := [][]int{
		{1, 2, 1},
		{1, 2, 2},
	}

	expected := 2

	test(n, wells, pipes, expected)
}

func main() {
	// 1168. Optimize Water Distribution in a Village
	test1()
	test2()
}
