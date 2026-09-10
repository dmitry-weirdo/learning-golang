package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
)

func minCostConnectPoints(points [][]int) int {
	// should be the same as Prim's time.
	// But space is better since we only keep edges in one direction. It is still N * (N - 1) / 2 that is O(N^2) space.
	// Is just a bit faster than Prim's (maybe just LeetCode's speed glitches)
	// Passes in 470-550 ms
	return minCostConnectPoints_kruskal(points)

	// todo: adj list in case of a very dense graph is slow and a lot of space. Optimization to O(V^2) will be using minDist[] array, which stores the minimum Manhattan distance from node i to any node already added to the MST.
	// Pures Prim's algorithm, with heap.
	// Slow, 510-560 ms
	// Same time as Dijkstra's,
	// O(E log E) = O(E log V^2) = O(2E log V) = O(E log V) = O(V^2 * log V)

	// Actually, it's also O(V^2) on constructing the adjacency list.
	//return minCostConnectPoints_prim(points)
}

func minCostConnectPoints_prim(points [][]int) int {
	// We need to get the MST tree of a fully-connected graph of points

	// Adjacency list is a full graph of points (every point [i] connected to every other point [j])
	adj := createAdjacencyListUndirectedWeighted(len(points), points)

	//fmt.Printf("Adjacency list: %v \n", adj)

	mst := getMinimumSpanningTreePrimNodesStartFrom0(adj)
	return mst.weight
}

func createAdjacencyListUndirectedWeighted(n int, points [][]int) [][][]int {
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

	for i := range n - 1 { // p1
		for j := i + 1; j < n; j++ { // p2
			from = i
			to = j
			weight = getManhattanDistance(points[i], points[j])

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
	}

	return adj
}

func getManhattanDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + // |x1 - x2|
		abs(p1[1]-p2[1]) // |y1 - y2|
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func minCostConnectPoints_kruskal(points [][]int) int {
	// We need to get the MST tree of a fully-connected graph of points
	n := len(points)

	// Edges list is a full graph of points (every point [i] connected to every other point [j])
	// We don't need both directions of edges, so it's less space than the adjacency list
	edges := getEdgesUndirected(n, points)
	//fmt.Printf("Edges: %v \n", edges)

	mst := getMinimumSpanningTreeKruskal(n, edges)

	return mst.weight
}

func getEdgesUndirected(n int, points [][]int) [][]int {
	edges := make([][]int, 0)

	from := 0
	to := 0
	weight := 0

	for i := range n - 1 { // p1
		for j := i + 1; j < n; j++ { // p2
			from = i
			to = j
			weight = getManhattanDistance(points[i], points[j])

			// we only add 1 direction, these are edges, NOT adj list
			// it will stay undirected
			edges = append(edges, []int{from, to, weight})
		}
	}

	return edges
}

// ========================= Prim's algorithm for MST begin ========================= //
type MinimumSpanningTree struct {
	edges  [][]int // every edge is from/to/weight
	weight int     // total weight for all edges
}

type Edge struct {
	from   int
	to     int
	weight int
}

func getMinimumSpanningTreePrimNodesStartFrom0(adj [][][]int) MinimumSpanningTree {
	n := len(adj)

	return getMinimumSpanningTreePrim(n, adj, 0) // start from node 0
}

func getMinimumSpanningTreePrimNodesStartFrom1(adj [][][]int) MinimumSpanningTree {
	n := len(adj)

	return getMinimumSpanningTreePrim(n+1, adj, 1) // start from node 1
}

func getMinimumSpanningTreePrim(n int, adj [][][]int, start int) MinimumSpanningTree { // start can be 0 or 1
	// Prim can works with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected.

	// If nodes are starting from 0 -> pass N, start = 0
	// If nodes are starting from 1 -> pass (N + 1), start = 1
	visited := make([]bool, n)

	nodesReached := 0 // we can stop iteration earlier if we reached all the nodes (this will happen not always)

	mst := MinimumSpanningTree{
		edges:  make([][]int, 0),
		weight: 0,
	}

	// mark start node as visited
	visited[start] = true

	pq := createMinHeapPrim()

	// add all neighbors of the start node to the heap
	for _, v := range adj[start] {
		neighbor, weight := v[0], v[1]

		if visited[neighbor] { // avoid self-cycle on the start node
			continue
		}

		heap.Push(pq, Edge{from: start, to: neighbor, weight: weight})
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

		nodesReached++

		// todo: this must be adopted if node numeration starts from 1
		// cut first -> if we reached all nodes, stop iteration
		if nodesReached >= (n - 1) {
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

// ========================= Kruskal's algorithm for MST begin ========================= //
func getMinimumSpanningTreeKruskal(n int, edges [][]int) MinimumSpanningTree { // start can be 0 or 1
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

	i := 0

	// todo: this must be adopted if node numeration starts from 1
	for (len(mst.edges) < n-1) && (i < len(edges)) {
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

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Point: %v \n", m)

	result := minCostConnectPoints(m)

	fmt.Printf("MST weight of points graph, weights are Manhattan distances : %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[][]int{
			{0, 0},
			{2, 2},
			{3, 10},
			{5, 2},
			{7, 0},
		},
		20,
	)
}

func test2() {
	test(
		[][]int{
			{3, 12},
			{-2, 5},
			{-4, 1},
		},
		18,
	)
}

func main() {
	// 1584. Min Cost to Connect All Points
	test1()
	test2()
}
