package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
)

func minimumCost(n int, connections [][]int) int {
	// Prim's should be faster for non-connected graphs,
	// since it will only iterate the component containing the starting node.
	// But it is actually slower since we're building an adj list and using a heap
	// passes in 30-65 ms
	return minimumCost_prim(n, connections)

	// passes in 13-32 ms
	//return minimumCost_kruskal(n, connections)
}

func minimumCost_prim(n int, connections [][]int) int {
	// edge numbers are from 1 to N

	adj := createAdjacencyListUndirectedWeighted(n+1, connections)

	// We need MST and check that all nodes are connected.
	// In Kruskal's algorithm, we can check whether union-find will connect just 1 group or more.
	mst := getMinimumSpanningTreePrimNodesStartFrom1(n, adj)

	if !mst.allNodesUsed {
		return -1
	}

	return mst.weight
}

func createAdjacencyListUndirectedWeighted(n int, edges [][]int) [][][]int {
	// todo: we can return an array of {node, weight} structs instead of 2-elements array
	// adj[i][j][0] - "to" node
	// adj[i][j][1] - weight of "from-to" edge
	// we're assuming there are no duplicate parallel edges for the same "from + to"

	adj := make([][][]int, n)

	from := 0
	to := 0
	weight := 0
	toAndWeight := []int{}
	fromAndWeight := []int{}

	for _, v := range edges {
		from = v[0]
		to = v[1]
		weight = v[2]

		toAndWeight = []int{to, weight}
		fromAndWeight = []int{from, weight}

		// add v1 to v2
		if adj[from] == nil {
			adj[from] = [][]int{toAndWeight}
		} else {
			adj[from] = append(adj[from], toAndWeight)
		}

		// add v2 to v1
		if adj[to] == nil {
			adj[to] = [][]int{fromAndWeight}
		} else {
			adj[to] = append(adj[to], fromAndWeight)
		}
	}

	return adj
}

func minimumCost_kruskal(n int, connections [][]int) int {
	// edge numbers are from 1 to N

	// We need MST and check that all nodes are connected.
	// In Kruskal's algorithm, we can check whether union-find will connect just 1 group or more.
	mst := getMinimumSpanningTreeKruskalNodesStartFrom1(n, connections)

	if !mst.allNodesUsed {
		return -1
	}

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

// ========================= Prim's algorithm for MST begin ========================= //
type Edge struct {
	from   int
	to     int
	weight int
}

func getMinimumSpanningTreePrimNodesStartFrom0(n int, adj [][][]int) MinimumSpanningTree {
	// adj will contain n elements
	return getMinimumSpanningTreePrim(n, 0, adj) // start from node 0
}

func getMinimumSpanningTreePrimNodesStartFrom1(n int, adj [][][]int) MinimumSpanningTree {
	// adj will contain n + 1 elements
	return getMinimumSpanningTreePrim(n+1, 1, adj) // start from node 1
}

func getMinimumSpanningTreePrim(n int, startIndex int, adj [][][]int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected.

	// If nodes are starting from 0 -> pass N, start = 0
	// If nodes are starting from 1 -> pass (N + 1), start = 1
	visited := make([]bool, n)

	mst := MinimumSpanningTree{
		edges:  make([][]int, 0),
		weight: 0,
	}

	// mark start node as visited
	visited[startIndex] = true

	pq := createMinHeapPrim()

	// add all neighbors of the start node to the heap
	for _, v := range adj[startIndex] {
		neighbor, weight := v[0], v[1]

		if visited[neighbor] { // avoid self-cycle on the start node
			continue
		}

		heap.Push(pq, Edge{from: startIndex, to: neighbor, weight: weight})
	}

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)

		if visited[edge.to] {
			// node was already visited -> do not handle it again
			continue
		}

		// update if distance for this node is not yet found
		visited[edge.to] = true

		// Add the current edge to MST
		mst.edges = append(mst.edges, []int{edge.from, edge.to, edge.weight})

		// Add the edge weight to total weight
		mst.weight += edge.weight

		// cut first -> if we reached all nodes, stop iteration
		if len(mst.edges) >= (n - startIndex - 1) {
			break
		}

		// Add all neighbors of this node to the heap.
		// !!! We're NOT skipping the nodes already in the heap.
		// The trick is - we can push same node multiple times, but the min-heap will select the shortest distance first
		for _, v := range adj[edge.to] {
			neighbor, weight := v[0], v[1]

			// neighbor was already reached with a shorter distance -> no reason to put it again
			if visited[neighbor] {
				// node was already reached -> do not handle it again
				continue
			}

			e := Edge{
				from:   edge.to,
				to:     neighbor, // neighbor
				weight: weight,   // unlike Dijkstra, we just add the edge weight, NOT the summary path weight
			}

			heap.Push(pq, e)
		}
	}

	// check whether all nodes are in the single MST component
	// We will have (N - 1) nodes in the MST in this case.
	mst.allNodesUsed = len(mst.edges) == (n - startIndex - 1)

	return mst
}

func createMinHeapPrim() *PriorityQueuePrim {
	return &PriorityQueuePrim{
		less: func(a, b Edge) bool {
			if a.weight == b.weight { // for stable MST return -> in case of same weight -> return earlier edge
				return a.to < b.to
			}

			// min heap
			return a.weight < b.weight
		},
	}
}

type PriorityQueuePrim struct {
	items []Edge
	less  func(a, b Edge) bool // comparator function, returns boolean, not integer!
}

// implementation of sort.Interface
func (pq *PriorityQueuePrim) Len() int {
	return len(pq.items)
}

// implementation of sort.Interface
func (pq *PriorityQueuePrim) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

// implementation of sort.Interface
func (pq *PriorityQueuePrim) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// implementation of heap.Interface
func (pq *PriorityQueuePrim) Push(x any) { // interface needs `x any`, else the override will not work
	pq.items = append(pq.items, x.(Edge))
}

// implementation of heap.Interface
func (pq *PriorityQueuePrim) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

// helper function -> get the top of the heap without removing it
func (pq *PriorityQueuePrim) Peek() Edge {
	return pq.items[0] // return the root
}

// ========================= Prim's algorithm for MST end ========================= //

func test(n int, m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Total nodes: %v \n", n)
	fmt.Printf("Edges: %v \n", m)

	result := minimumCost(n, m)

	fmt.Printf("Minimum cost to connect all %v edges: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		3,
		[][]int{
			{1, 2, 5},
			{1, 3, 6},
			{2, 3, 1},
		},
		6,
	)
}

func test2() {
	test(
		4,
		[][]int{
			{1, 2, 3},
			{3, 4, 4},
		},
		-1, // 1-2 and 3-4 are 2 separate components
	)
}

func main() {
	// 1135. Connecting Cities With Minimum Cost
	test1()
	test2()
}
