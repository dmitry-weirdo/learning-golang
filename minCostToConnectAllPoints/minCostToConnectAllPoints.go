package main

import (
	"container/heap"
	"fmt"
)

func minCostConnectPoints(points [][]int) int {

	// todo: adj list in case of a very dense graph is slow and a lot of space. Optimization to O(V^2) will be using minDist[] array, which stores the minimum Manhattan distance from node i to any node already added to the MST.
	// Pures Prim's algorithm, with heap.
	// Slow, 510-560 ms
	// Same time as Dijkstra's,
	// O(E log E) = O(E log V^2) = O(2E log V) = O(E log V) = O(V^2 * log V)
	return minCostConnectPoints_prim(points)
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
