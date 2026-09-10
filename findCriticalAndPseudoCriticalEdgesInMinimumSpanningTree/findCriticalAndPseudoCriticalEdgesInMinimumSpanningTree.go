package main

import (
	"cmp"
	"demo/matrixCommon"
	"fmt"
	"slices"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// todo: this can be space and time optimized, since we don't need to collect the MST edges here, so we can just count the edges and calculate only weight.

	startIndex := 0 // node numeration stats from 0

	// Since we need to return edge indexes, we generate new edges array with v[3] = index in the original edges
	indexedEdges := createIndexedEdges(edges)
	//fmt.Printf("Edges with edge indexes added: %v \n", indexedEdges)

	// Since we will sort edges only based on (v[0], v[1], v[2]),
	// [v3] will remain there and can be used to include in the result.
	sortEdgesForKruskal(indexedEdges)

	// next we use variations of Kruskal's algorithm:
	// - Usual Kruskal's to find the MST weight
	// - For every edge[i]
	//   - EXCLUDE this edge from Kruskal's union-find. If it leads to MST weight increase (or MST cannot proceed up to N - 1 edges), then this edge is critical.
	//   - If edge is not critical, we FORCE this edge in Kruskal's union-find and check whether MST has still the same weight -> then this edge is pseudo-critical.
	mstWithAllEdges := getMinimumSpanningTreeKruskalWithSortedEdges(n, startIndex, indexedEdges)
	//fmt.Printf("MST with all edges weight: %v \n", mstWithAllEdges.weight)

	if !mstWithAllEdges.allNodesUsed { // this must NEVER happen in this task -> graph must be connected
		panic("MST with all edges is NOT connected.")
	}

	criticalEdges := make([]int, 0)
	pseudoCriticalEdges := make([]int, 0)

	// for every edge, we check whether it's excluded
	for i, e := range indexedEdges {
		mstWithEdgeIgnored := getMinimumSpanningTreeKruskalWithSortedEdgesAndExcludedEdge(n, startIndex, indexedEdges, i)

		if !mstWithEdgeIgnored.allNodesUsed || (mstWithEdgeIgnored.weight > mstWithAllEdges.weight) {
			// edge is critical -> add its index to the critical edges result
			criticalEdges = append(criticalEdges, e[3])
			continue
		}

		// edge is non-critical -> force it and check whether there will be MST of the same weight and still all nodes
		mstWithEdgeForced := getMinimumSpanningTreeKruskalWithSortedEdgesAndForcedEdge(n, startIndex, indexedEdges, i)

		if mstWithEdgeForced.allNodesUsed && (mstWithEdgeForced.weight == mstWithAllEdges.weight) {
			pseudoCriticalEdges = append(pseudoCriticalEdges, e[3])
			continue
		}
	}

	return [][]int{
		criticalEdges,
		pseudoCriticalEdges,
	}
}

func createIndexedEdges(edges [][]int) [][]int {
	// from = v[0]
	// to = v[1]
	// weight = v[2]
	// edgeIndex = v[3]
	indexesEdges := make([][]int, len(edges))

	for i, v := range edges {
		indexesEdges[i] = []int{v[0], v[1], v[2], i}
	}

	return indexesEdges
}

// ========================= Kruskal's algorithm modifications for this problem ========================= //
func getMinimumSpanningTreeKruskalWithSortedEdges(n int, startIndex int, edges [][]int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected. It's necessary for Kruskal's since we're uniting 2 nodes with Union-Find

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

func getMinimumSpanningTreeKruskalWithSortedEdgesAndExcludedEdge(n int, startIndex int, edges [][]int, excludedEdgeIndex int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected. It's necessary for Kruskal's since we're uniting 2 nodes with Union-Find

	// union-find controls the visited by including them in the same node set
	uf := NewUnionFind(n)

	mst := MinimumSpanningTree{
		edges:  make([][]int, 0),
		weight: 0,
	}

	// index in edges[] array
	i := 0

	for (len(mst.edges) < n-startIndex-1) && (i < len(edges)) {
		if i == excludedEdgeIndex { // exclude the ignored edge from MST
			i++
			continue
		}

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

func getMinimumSpanningTreeKruskalWithSortedEdgesAndForcedEdge(n int, startIndex int, edges [][]int, forcedEdgeIndex int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected. It's necessary for Kruskal's since we're uniting 2 nodes with Union-Find

	// union-find controls the visited by including them in the same node set
	uf := NewUnionFind(n)

	mst := MinimumSpanningTree{
		edges:  make([][]int, 0),
		weight: 0,
	}

	// !!! force the edge, i.e. we are uniting its from and to nodes
	forcedEdge := edges[forcedEdgeIndex]
	forcedEdgeFrom := forcedEdge[0]
	forcedEdgeTo := forcedEdge[1]
	forcedEdgeWeight := forcedEdge[2]

	uf.Union(forcedEdgeFrom, forcedEdgeTo)
	mst.edges = append(mst.edges, []int{forcedEdgeFrom, forcedEdgeTo, forcedEdgeWeight})
	mst.weight += forcedEdgeWeight

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

func test(n int, m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of nodes: %v \n", n)

	fmt.Printf("Edges: \n")
	matrixCommon.PrintIntMatrix(m)

	result := findCriticalAndPseudoCriticalEdges(n, m)

	// order in the result is not guaranteed -> let's order result[0] and result[1], so that we can compare with the expected values that will be ordered by edge index
	slices.Sort(result[0])
	slices.Sort(result[1])

	fmt.Printf("Critical and non-critical edges: \n")
	matrixCommon.PrintIntMatrixWithDifferentColumnsCount(result)

	fmt.Printf("Expected result: \n")
	matrixCommon.PrintIntMatrixWithDifferentColumnsCount(expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	n := 5

	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 2},
		{0, 3, 2},
		{0, 4, 3},
		{3, 4, 3},
		{1, 4, 6},
	}

	expected := [][]int{
		{0, 1},       // critical edge indices
		{2, 3, 4, 5}, // pseudo-critical edge indices
	}

	test(n, edges, expected)
}

func test2() {
	n := 4

	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{0, 3, 1},
	}

	expected := [][]int{
		{},           // critical edge indices
		{0, 1, 2, 3}, // pseudo-critical edge indices
	}

	test(n, edges, expected)
}

func main() {
	// 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
	test1()
	test2()
}
