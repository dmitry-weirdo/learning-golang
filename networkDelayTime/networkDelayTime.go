package main

import (
	"container/heap"
	"fmt"
)

func networkDelayTime(times [][]int, n int, k int) int {
	// This is Dijkstra's algorithm from node K.
	// We need to return the max value of all the distances.
	// If not all the nodes are reachable, we return -1.
	// !!! Dijkstra's algorithm only works for non-negative edge weights, it's greedy on edge weights.

	// Time: we can estimate max edges = V^2 (actually V * (V - 1) / 2)
	// Since we can put all the edges to the heap,
	// and every add/remove from the heap is O(log N),
	// the complexity will be O(E * log E) = O(V^2 * log V^2) = O(V^2 * 2 log V) = O(V^2 * log V)
	// = O(E * log V)
	return networkDelayTime_dijkstra(times, n, k)
}

func networkDelayTime_dijkstra(times [][]int, n int, k int) int {
	adj := createAdjacencyListDirectedWeighted(n+1, times) // nodes are starting from 1 -> use (n + 1)

	//fmt.Printf("Adjacency list with weights: %v \n", adj)

	// -1 means "no value", we assume all weights are non-negative
	const DISTANCE_NOT_FOUND = -1
	distances := createIntArrayWithDefaultValues(n+1, DISTANCE_NOT_FOUND) // nodes are starting from 1 -> use (n + 1)

	maxDistance := -1

	nodesReached := 0 // we can stop iteration earlier if we reached all the nodes (this will happen not always)

	pq := createMinHeap()
	heap.Push(pq, NodeWeight{k, 0}) // start node is 0 weight

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

	if nodesReached < n {
		return -1
	}

	return maxDistance
}

type NodeWeight struct { // what we push to heap
	node     int // target "to" node
	distance int // total distance up to this node
}

func createAdjacencyListDirectedWeighted(n int, edges [][]int) [][][]int {
	// todo: we can return an array of structs instead of 2-elements array
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

func test(m [][]int, n, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Total nodes: %v \n", n)
	fmt.Printf("Starting node: %v \n", k)
	fmt.Printf("Edges with weights: %v \n", m)

	result := networkDelayTime(m, n, k)

	fmt.Printf("Time to reach all the edges: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	edges := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}

	n := 4
	k := 2

	expected := 2

	test(edges, n, k, expected)
}

func test2() {
	edges := [][]int{
		{1, 2, 1},
	}

	n := 2
	k := 1

	expected := 1

	test(edges, n, k, expected)
}

func test3() {
	edges := [][]int{
		{1, 2, 1},
	}

	n := 2
	k := 2

	expected := -1

	test(edges, n, k, expected)
}

func main() {
	// 743. Network Delay Time
	test1()
	test2()
	test3()
}
