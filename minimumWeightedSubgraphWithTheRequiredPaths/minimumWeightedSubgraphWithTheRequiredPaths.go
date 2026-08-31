package main

import (
	"container/heap"
	"fmt"
)

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// todo: possible to implement with just 1 Dijkstra run, see https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/solutions/1844479/simultaneous-dijkstra-beats-100-only-1-d-1mat/

	// pretty slow, passes in 160-220 ms
	return minimumWeight_3DijkstraRuns(n, edges, src1, src2, dest)
}

func minimumWeight_3DijkstraRuns(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// For more optimal fail-fast, we can run adjReversed first and check dist1 & dist2 here
	// It means calculations from src1 and src2 will not be executed at all if there is no path to dest.

	// Count distances from target on the reversed graph.
	// It basically means "distances from all nodes to dest" on the original graph
	adjReversed := createAdjacencyListDirectedWeightedReversed(n, edges)
	//fmt.Printf("Adjacency list reversed: %d \n", adjReversed)

	distDest := getShortestDistancesDijkstra(adjReversed, dest)
	if distDest[src1] < 0 || distDest[src2] < 0 {
		// dest non-reachable from src1 or src2 -> unable to build the target subgraph
		return -1
	}

	// nodes are [0; n - 1]
	adj := createAdjacencyListDirectedWeighted(n, edges)
	//fmt.Printf("Adjacency list: %d \n", adj)

	// count distances from src1 and src2 on the original graph
	distSrc1 := getShortestDistancesDijkstra(adj, src1)
	distSrc2 := getShortestDistancesDijkstra(adj, src2)

	// For all the nodes K, find the min of (src1 -> K) + (src2 -> K) + (K -> dest)
	minDistance := -1

	for k := range n {
		// if K is non-reachable from any of the nodes -> skip it
		if (distSrc1[k] < 0) ||
			(distSrc2[k] < 0) ||
			(distDest[k] < 0) {
			continue
		}

		dist := distSrc1[k] + // src1 -> k
			distSrc2[k] + // src2 -> k
			distDest[k] // k -> dest (= dest-> k on the reversed graph)

		//fmt.Println()
		//fmt.Printf("src1 -> %v: %v \n", k, distSrc1[k])
		//fmt.Printf("src2 -> %v: %v \n", k, distSrc2[k])
		//fmt.Printf("%v -> dest: %v \n", k, distDest[k])
		//fmt.Printf("Total dist: %v \n", dist)

		if minDistance == -1 {
			minDistance = dist
		} else {
			minDistance = min(minDistance, dist)
		}
	}

	return int64(minDistance)
}

func createAdjacencyListDirectedWeighted(n int, edges [][]int) [][][]int {
	// todo: we can return an array of {node, weight} structs instead of 2-elements array
	// adj[i][j][0] - "to" node
	// adj[i][j][1] - weight of "from-to" edge
	// we're assuming there are no duplicate parallel edges for the same "from + to"

	adj := make([][][]int, n)

	from := 0
	to := 0
	weight := 0
	toAndWeight := []int{}

	for _, v := range edges {
		from = v[0]
		to = v[1]
		weight = v[2]

		toAndWeight = []int{to, weight}

		// add v2 to v1
		if adj[from] == nil {
			adj[from] = [][]int{toAndWeight}
		} else {
			adj[from] = append(adj[from], toAndWeight)
		}
	}

	return adj
}

func createAdjacencyListDirectedWeightedReversed(n int, edges [][]int) [][][]int { // from -> to is reversed
	// todo: we can return an array of {node, weight} structs instead of 2-elements array
	// adj[i][j][0] - "to" node
	// adj[i][j][1] - weight of "from-to" edge
	// we're assuming there are no duplicate parallel edges for the same "from + to"

	adj := make([][][]int, n)

	from := 0
	to := 0
	weight := 0
	toAndWeight := []int{}

	for _, v := range edges {
		from = v[1]   // !!! we reverse from and to
		to = v[0]     // !!! we reverse from and to
		weight = v[2] // weight stays the same

		toAndWeight = []int{to, weight}

		// add v2 to v1
		if adj[from] == nil {
			adj[from] = [][]int{toAndWeight}
		} else {
			adj[from] = append(adj[from], toAndWeight)
		}
	}

	return adj
}

// ========================= Dijkstra shortest paths begin ========================= //
func getShortestDistancesDijkstra(adj [][][]int, start int) []int {
	// Dijkstra only works for non-negative edge weights
	// Returns an array of distances from start node to all N nodes.
	// Start node will have distance 0
	// Distances[i] will be -1 if node [i] is not reachable from start.
	n := len(adj)

	// -1 means "no value", we assume all weights are non-negative
	const DISTANCE_NOT_FOUND = -1
	distances := createIntArrayWithDefaultValues(n+1, DISTANCE_NOT_FOUND) // nodes are starting from 1 -> use (n + 1)

	maxDistance := -1

	nodesReached := 0 // we can stop iteration earlier if we reached all the nodes (this will happen not always)

	pq := createMinHeap()
	heap.Push(pq, NodeWeight{start, 0}) // start node is 0 weight

	for pq.Len() > 0 {
		nodeWeight := heap.Pop(pq).(NodeWeight)

		if distances[nodeWeight.node] != DISTANCE_NOT_FOUND {
			// node was already reached -> do not handle it again
			continue
		}

		// update if distance for this node is not yet found
		distances[nodeWeight.node] = nodeWeight.distance

		// todo: for the pureness of Dijkstra, we can calculate maxDistance after the iteration.
		// We calc maxDistance withing Dijkstra iteration to save time on distances array iteration.
		maxDistance = max(maxDistance, nodeWeight.distance)

		nodesReached++

		// cut first -> if we reached all nodes, stop iteration
		if nodesReached >= n {
			break
		}

		// Add all neighbors of this node to the heap.
		// !!! We're NOT skipping the nodes already in the heap.
		// The trick is - we can push same node multiple times, but the min-heap will select the shortest distance first
		for _, v := range adj[nodeWeight.node] {
			neighbor, weight := v[0], v[1]

			// neighbor was already reached with a shorter distance -> no reason to put it again
			if distances[neighbor] != DISTANCE_NOT_FOUND {
				// node was already reached -> do not handle it again
				continue
			}

			nw := NodeWeight{
				neighbor,                     // neighbor
				nodeWeight.distance + weight, // neighbor distance = currentNodeDistance + (current -> neighbor) edge weight
			}

			heap.Push(pq, nw)
		}
	}

	return distances
}

type NodeWeight struct { // what we push to heap
	node     int // target "to" node
	distance int // total distance up to this node
}

func createIntArrayWithDefaultValues(n int, defaultValue int) []int {
	a := make([]int, n)

	for i := range n {
		a[i] = defaultValue
	}

	return a
}

// heap with <cost, value> struct, contains NodeWeight value
func createMinHeap() *PriorityQueue {
	return &PriorityQueue{
		less: func(a, b NodeWeight) bool {
			// min heap
			return a.distance < b.distance
		},
	}
}

type PriorityQueue struct {
	items []NodeWeight
	less  func(a, b NodeWeight) bool // comparator function, returns boolean, not integer!
}

// implementation of sort.Interface
func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

// implementation of sort.Interface
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

// implementation of sort.Interface
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// implementation of heap.Interface
func (pq *PriorityQueue) Push(x any) { // interface needs `x any`, else the override will not work
	pq.items = append(pq.items, x.(NodeWeight))
}

// implementation of heap.Interface
func (pq *PriorityQueue) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

// helper function -> get the top of the heap without removing it
func (pq *PriorityQueue) Peek() NodeWeight {
	return pq.items[0] // return the root
}

// ========================= Dijkstra shortest paths end ========================= //

func test(n int, m [][]int, src1, src2, dest int, expectedResult int64) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Total nodes: %v \n", n)
	fmt.Printf("Starting node src1: %v \n", src1)
	fmt.Printf("Starting node src2: %v \n", src2)
	fmt.Printf("Target node dest: %v \n", dest)
	fmt.Printf("Edges with weights: %v \n", m)

	result := minimumWeight(n, m, src1, src2, dest)

	fmt.Printf("Min weight of subgraph where it is possible to reach (%v -> %v) and (%v -> %v): %v \n", src1, dest, src2, dest, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 6
	src1 := 0
	src2 := 1
	dest := 5

	edges := [][]int{
		{0, 2, 2},
		{0, 5, 6},
		{1, 0, 3},
		{1, 4, 5},
		{2, 1, 1},
		{2, 3, 3},
		{2, 3, 4},
		{3, 4, 2},
		{4, 5, 1},
	}

	expected := int64(9)

	test(n, edges, src1, src2, dest, expected)
}

func test2() {
	n := 3
	src1 := 0
	src2 := 1
	dest := 2

	edges := [][]int{
		{0, 1, 1},
		{2, 1, 1},
	}

	expected := int64(-1) // no possible subgraph to reach src1 -> dest and src2 -> dest

	test(n, edges, src1, src2, dest, expected)
}

func test3() {
	// failing test-case 21/88
	n := 8
	src1 := 4
	src2 := 1
	dest := 6

	edges := [][]int{
		{4, 7, 24},
		{1, 3, 30},
		{4, 0, 31},
		{1, 2, 31},
		{1, 5, 18},
		{1, 6, 19},
		{4, 6, 25},
		{5, 6, 32},
		{0, 6, 50},
	}

	expected := int64(44)

	test(n, edges, src1, src2, dest, expected)
}

func main() {
	// 2203. Minimum Weighted Subgraph With the Required Paths
	test1()
	test2()
	test3()
}
