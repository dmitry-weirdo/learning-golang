package main

import (
	"container/heap"
	"demo/matrixCommon"
	"fmt"
)

func createAdjacencyListUndirectedUnweighted(n int, edges [][]int) [][]int {
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

func createAdjacencyListDirectedUnweighted(n int, edges [][]int) [][]int {
	adj := make([][]int, n)

	from := 0
	to := 0

	for _, v := range edges {
		from = v[0]
		to = v[1]

		// add from -> to
		if adj[from] == nil {
			adj[from] = []int{to}
		} else {
			adj[from] = append(adj[from], to)
		}
	}

	return adj
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

// ========================= Floyd-Warshall shortest paths between all nodes begin ========================= //
func getShortestDistancesFloydWarshallDirected(n int, edges [][]int, infinity int) [][]int {
	m := createFloydWarshallInitialMatrixDirected(n, edges, infinity)

	//fmt.Printf("Initial matrix for Floyd-Warshall: \n")
	//PrintIntMatrix(m)

	// todo: save paths if required, see https://www.youtube.com/watch?v=oNI0rf2P9gE

	// Floyd-Warshall execution
	// No paths saved, no negative weight cycles handled
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

func getShortestDistancesFloydWarshallUndirected(n int, edges [][]int, infinity int) [][]int {
	m := createFloydWarshallInitialMatrixUndirected(n, edges, infinity)

	//fmt.Printf("Initial matrix for Floyd-Warshall: \n")
	//PrintIntMatrix(m)

	// todo: save paths if required, see https://www.youtube.com/watch?v=oNI0rf2P9gE

	// Floyd-Warshall execution
	// No paths saved, no negative weight cycles handled
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

func createFloydWarshallInitialMatrixDirected(n int, edges [][]int, infinity int) [][]int {
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
	}

	return m
}

func createIntMatrixWithDefaultValues(rows, columns int, defaultValue int) [][]int { // this is from matrixCommon
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)

		for j := range columns { // !!! note that this is slow, will take O(m * n) additional operations :(
			m[i][j] = defaultValue
		}
	}

	return m
}

// ========================= Floyd-Warshall shortest paths between all nodes end ========================= //

// ========================= Bellman-Ford shortest paths from 1 starting node begin ========================= //
func getShortestDistancesBellmanFord(n int, edges [][]int, start int, infinity int) []int {
	// O(V * E) = O(V^3) - for EVERY starting node
	// O(V^2) to count the number of nodes within distanceThreshold for every node.
	// Total: O(V^4)

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

// ========================= Dijkstra shortest paths from 1 starting node begin ========================= //
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

	pq := createMinHeapDijkstra()
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
func createMinHeapDijkstra() *PriorityQueueDijkstra {
	return &PriorityQueueDijkstra{
		less: func(a, b NodeWeight) bool {
			// min heap
			return a.distance < b.distance
		},
	}
}

type PriorityQueueDijkstra struct {
	items []NodeWeight
	less  func(a, b NodeWeight) bool // comparator function, returns boolean, not integer!
}

// implementation of sort.Interface
func (pq *PriorityQueueDijkstra) Len() int {
	return len(pq.items)
}

// implementation of sort.Interface
func (pq *PriorityQueueDijkstra) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

// implementation of sort.Interface
func (pq *PriorityQueueDijkstra) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// implementation of heap.Interface
func (pq *PriorityQueueDijkstra) Push(x any) { // interface needs `x any`, else the override will not work
	pq.items = append(pq.items, x.(NodeWeight))
}

// implementation of heap.Interface
func (pq *PriorityQueueDijkstra) Pop() any { // interface needs `x any`, else the override will not work
	n := len(pq.items)
	lastItem := pq.items[n-1]

	pq.items = pq.items[0 : n-1] // remove the last element

	return lastItem
}

// helper function -> get the top of the heap without removing it
func (pq *PriorityQueueDijkstra) Peek() NodeWeight {
	return pq.items[0] // return the root
}

// ========================= Dijkstra shortest paths end ========================= //

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

// ========================= Test functions ========================= //

func testPrim1() {
	edges := [][]int{
		// from - to - weight
		{0, 1, 10},
		{0, 2, 3},
		{1, 2, 4},
		{1, 3, 1},
		{2, 3, 4},
		{2, 4, 4},
		{3, 4, 2},
	}

	adj := createAdjacencyListUndirectedWeighted(5, edges)

	expectedMstWeight := 10 // 3 + 4 + 1 + 2

	// from, to, weight
	expectedMstEdges := [][]int{
		{0, 2, 3},
		{2, 1, 4},
		{1, 3, 1},
		{3, 4, 2},
	}

	mst := getMinimumSpanningTreePrimNodesStartFrom0(adj)

	fmt.Printf("MST weight: %v \n", mst.weight)
	fmt.Printf("Expected MST weight: %v \n", expectedMstWeight)

	if mst.weight != expectedMstWeight {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedMstWeight, mst.weight)
	}

	// validate edges
	result := mst.edges
	expectedResult := expectedMstEdges

	fmt.Printf("MST edges: \n")
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected MST edges: \n")
	matrixCommon.PrintIntMatrix(expectedResult)

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

func testPrimSuite() {
	testPrim1()
}

func main() {
	testPrimSuite()
}
