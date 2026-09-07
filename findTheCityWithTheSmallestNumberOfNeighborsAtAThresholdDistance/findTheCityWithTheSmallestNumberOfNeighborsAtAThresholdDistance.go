package main

import (
	"container/heap"
	"fmt"
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// O(V * E) = O(V^3) - for EVERY starting node
	// O(V^2) to count the number of nodes within distanceThreshold for every node.
	// Total: O(V^4)
	// Passes in 90-105 ms
	return findTheCity_bellmanFord(n, edges, distanceThreshold)

	// O(V^3) to calculate distances.
	// O(V^2) to count the number of nodes within distanceThreshold for every node.
	// Total: O(V^3 + V^2) = O(V^3)
	// Passes in 4-5 ms
	//return findTheCity_floydWarshall(n, edges, distanceThreshold)

	// O(E * log V) = O(V^2 * log V) for every distance -> O(V^3 * log V)
	// and O(V) to count the number of nodes within distanceThreshold for every node -> O(V^2)
	// Total: O(V^3 * log V + V^2) = O(V^3 * log V)
	// passes in 62-72 ms
	//return findTheCity_dijkstraNaive(n, edges, distanceThreshold)
}

func findTheCity_bellmanFord(n int, edges [][]int, distanceThreshold int) int {
	const infinity = math.MaxInt32 / 2 // avoid overlow on addition
	const DISTANCE_NOT_FOUND = infinity

	minNodes := n + 1
	minIndex := -1

	// calculate shortest for every node
	for i := range n {
		// run Bellman-Ford
		shortestDistances := getShortestDistancesBellmanFord(n, edges, i, infinity)

		nodesReachableWithinDistance := 0

		for j, v := range shortestDistances {
			if v == DISTANCE_NOT_FOUND { // do not count unreachable nodes
				continue
			}

			if j == i { // do not count 0 distance to node itself // todo: this can be skipped
				continue
			}

			if v <= distanceThreshold {
				nodesReachableWithinDistance++
			}
		}

		if nodesReachableWithinDistance <= minNodes {
			minIndex = i
			minNodes = nodesReachableWithinDistance
		}
	}

	return minIndex
}

// ========================= Bellman-Ford shortest paths from 1 starting node begin ========================= //
func getShortestDistancesBellmanFord(n int, edges [][]int, start int, infinity int) []int {
	// run Bellman-Ford
	dist := createIntArrayWithDefaultValues(n, infinity)
	dist[start] = 0 // distance to the starting node is 0

	for range n {
		// for ever node, iterate every edge in the graph,
		// !!! Not just node's edges, ALL the edges
		for _, edge := range edges {
			from := edge[0]
			to := edge[1]
			weight := edge[2]

			// undirected -> use edges in both ways (from -> to and to -> from)
			dist[to] = min(dist[to], dist[from]+weight)   // from -> to
			dist[from] = min(dist[from], dist[to]+weight) // to -> from
		}

		// todo: detect nodes affected by the  negative cycles and set their dist[i] to NEGATIVE_INFINITY - see https://www.youtube.com/watch?v=lyw4FaxrwHg
	}

	return dist
}

// ========================= Bellman-Ford shortest paths from 1 starting node end ========================= //

func findTheCity_floydWarshall(n int, edges [][]int, distanceThreshold int) int {
	infinity := math.MaxInt32 / 2 // avoid overlow on addition

	// executed Floyd-Warshall
	m := getShortestDistancesFloydWarshallUndirected(n, edges, infinity)

	// calculate the result
	minNodes := n + 1
	minIndex := -1

	for i := range n { // every node

		nodesReachableWithinDistance := 0

		for j, v := range m[i] { // distances from node[i] to all other nodes
			if v == infinity { // do not count unreachable nodes
				continue
			}

			if j == i { // do not count 0 distance to node itself // todo: this can be skipped

			}

			if v <= distanceThreshold {
				nodesReachableWithinDistance++
			}
		}

		if nodesReachableWithinDistance <= minNodes {
			minIndex = i
			minNodes = nodesReachableWithinDistance
		}
	}

	return minIndex
}

func getShortestDistancesFloydWarshallUndirected(n int, edges [][]int, infinity int) [][]int {
	m := createFloydWarshallInitialMatrixUndirected(n, edges, infinity)

	//fmt.Printf("Initial matrix for Floyd-Warshall: \n")
	//PrintIntMatrix(m)

	// todo: save paths if required, see https://www.youtube.com/watch?v=oNI0rf2P9gE

	// Floyd-Warshall execution
	// No path, no negative weight cycles
	// O(V^3)
	for k := range n {
		for i := range n {
			for j := range n {
				m[i][j] = min(m[i][j], m[i][k]+m[k][j])
			}
		}
	}

	// todo: exclude negative paths if required, see https://www.youtube.com/watch?v=oNI0rf2P9gE

	return m
}

func createFloydWarshallInitialMatrixUndirected(n int, edges [][]int, infinity int) [][]int {
	// Initial n x n matrix:
	// - distance to self = 0
	// - directly connected nodes -> set edge weight
	// - no direct edge -> set infinity
	m := createIntMatrixWithDefaultValues(n, n, infinity)

	for i := range n { // distance to self is 0
		m[i][i] = 0
	}

	from := 0
	to := 0
	weight := 0

	// undirected -> add to both sides
	for _, v := range edges {
		from = v[0]
		to = v[1]
		weight = v[2]

		m[from][to] = weight
		m[to][from] = weight
	}

	return m
}

func createIntMatrixWithDefaultValues(rows, columns int, defaultValue int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)

		for j := range columns { // !!! note that this is slow, will take O(m * n) additional operations :(
			m[i][j] = defaultValue
		}
	}

	return m
}

func PrintIntMatrix(mat [][]int) {
	rows, columns := getRowsAndColumns(mat)

	for i := range rows {
		for j := range columns {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func findTheCity_dijkstraNaive(n int, edges [][]int, distanceThreshold int) int {
	// !!! unweighted graph -> we add both directions for every edge
	adj := createAdjacencyListUndirectedWeighted(n, edges)

	//distances := make([][]int, n)

	minNodes := n + 1
	minIndex := -1

	const DISTANCE_NOT_FOUND = -1

	// calculate shortest for every node
	for i := range n {
		shortestDistances := getShortestDistancesDijkstraNodesStartFrom0(adj, i, DISTANCE_NOT_FOUND)
		//distances[i] = shortestDistances

		nodesReachableWithinDistance := 0

		for j, v := range shortestDistances {
			if v == DISTANCE_NOT_FOUND { // do not count unreachable nodes
				continue
			}

			if j == i { // do not count 0 distance to node itself // todo: this can be skipped
				continue
			}

			if v <= distanceThreshold {
				nodesReachableWithinDistance++
			}
		}

		if nodesReachableWithinDistance <= minNodes {
			minIndex = i
			minNodes = nodesReachableWithinDistance
		}
	}

	return minIndex
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

// ========================= Dijkstra shortest paths begin ========================= //
func getShortestDistancesDijkstraNodesStartFrom0(adj [][][]int, start int, distanceNotFoundValue int) []int {
	n := len(adj)

	return getShortestDistancesDijkstra(n, adj, start, distanceNotFoundValue)
}

func getShortestDistancesDijkstraNodesStartFrom1(adj [][][]int, start int, distanceNotFoundValue int) []int {
	n := len(adj)

	return getShortestDistancesDijkstra(n+1, adj, start, distanceNotFoundValue)
}

func getShortestDistancesDijkstra(n int, adj [][][]int, start int, distanceNotFoundValue int) []int {
	// Dijkstra only works for non-negative edge weights
	// Returns an array of distances from start node to all N nodes.
	// Start node will have distance 0
	// Distances[i] will be -1 if node [i] is not reachable from start.

	// -1 means "no value", we assume all weights are non-negative
	DISTANCE_NOT_FOUND := distanceNotFoundValue

	// If nodes are starting from 0 -> pass N
	// If nodes are starting from 1 -> pass (N + 1)
	distances := createIntArrayWithDefaultValues(n, DISTANCE_NOT_FOUND)

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

func test(n int, m [][]int, distanceThreshold int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of nodes: %v \n", n)
	fmt.Printf("Edges: %v \n", m)
	fmt.Printf("Distance threshold: %v \n", distanceThreshold)

	result := findTheCity(n, m, distanceThreshold)

	fmt.Printf("City with smallest edges reachable within distance %v: %v \n", distanceThreshold, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 4

	edges := [][]int{
		{0, 1, 3},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}

	distanceThreshold := 4

	expected := 3

	test(n, edges, distanceThreshold, expected)
}

func test2() {
	n := 5

	edges := [][]int{
		{0, 1, 2},
		{0, 4, 8},
		{1, 2, 3},
		{1, 4, 2},
		{2, 3, 1},
		{3, 4, 1},
	}

	distanceThreshold := 2

	expected := 0

	test(n, edges, distanceThreshold, expected)
}

func main() {
	// 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance
	test1()
	test2()
}
