package main

import (
	"cmp"
	"container/heap"
	"demo/matrixCommon"
	"fmt"
	"slices"
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
	edges        [][]int // every edge is from/to/weight
	weight       int     // total weight for all edges
	allNodesUsed bool    // false if there are multiple non-connected components
}

type Edge struct { // element to be stored in the heap for the Kruskal's algorithm
	from   int
	to     int
	weight int
}

func getMinimumSpanningTreePrimNodesStartFrom0(n int, adj [][][]int) MinimumSpanningTree {
	// adj will contain n elements
	return getMinimumSpanningTreePrim(n, 0, adj) // start from node 0
}

func getMinimumSpanningTreePrimNodesStartFrom1(n int, adj [][][]int) MinimumSpanningTree {
	// adj will contain (n + 1) elements
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

// ========================= Kruskal's algorithm for MST begin ========================= //
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

// ========================= Test functions ========================= //
func testPrim(n int, startIndex int, edges [][]int, expectedMstAllNodesUsed bool, expectedMstWeight int, expectedMstEdges [][]int) {
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("Testing method \"%v\"... \n", "getMinimumSpanningTreePrim")

	adj := createAdjacencyListUndirectedWeighted(n, edges)
	mst := getMinimumSpanningTreePrim(n, startIndex, adj)

	// validate MST all nodes in single component
	fmt.Printf("MST all nodes used: %v \n", mst.allNodesUsed)
	fmt.Printf("Expected MST all nodes used: %v \n", expectedMstAllNodesUsed)

	if mst.allNodesUsed != expectedMstAllNodesUsed {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedMstAllNodesUsed, mst.allNodesUsed)
	}

	// validate MST weight
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

func testPrim1() {
	n := 5
	startIndex := 0

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

	expectedMstAllNodesUsed := true

	expectedMstWeight := 10 // 3 + 4 + 1 + 2

	// from, to, weight
	expectedMstEdges := [][]int{
		{0, 2, 3},
		{2, 1, 4},
		{1, 3, 1},
		{3, 4, 2},
	}

	testPrim(n, startIndex, edges, expectedMstAllNodesUsed, expectedMstWeight, expectedMstEdges)
}

func testPrim2() {
	n := 5 // 1-4, index 0 not used
	startIndex := 1

	edges := [][]int{
		// from - to - weight
		{1, 2, 3},
		{3, 4, 4},
	}

	expectedMstAllNodesUsed := false

	expectedMstWeight := 3 // 3, not connected. Edge 4 will not be reached

	// from, to, weight
	// For non-connected graph, Prim's will only return the edges of the component where the starting node (0 or 1) belongs.
	expectedMstEdges := [][]int{
		{1, 2, 3},
		//{3, 4, 4},
	}

	testPrim(n, startIndex, edges, expectedMstAllNodesUsed, expectedMstWeight, expectedMstEdges)
}

func testPrimSuite() {
	testPrim1()
	testPrim2()
}

func testKruskal(n int, startIndex int, edges [][]int, expectedMstAllNodesUsed bool, expectedMstWeight int, expectedMstEdges [][]int) {
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("Testing method \"%v\"... \n", "getMinimumSpanningTreeKruskal")

	mst := getMinimumSpanningTreeKruskal(n, startIndex, edges)

	// validate MST all nodes in single component
	fmt.Printf("MST all nodes used: %v \n", mst.allNodesUsed)
	fmt.Printf("Expected MST all nodes used: %v \n", expectedMstAllNodesUsed)

	if mst.allNodesUsed != expectedMstAllNodesUsed {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedMstAllNodesUsed, mst.allNodesUsed)
	}

	// validate MST weight
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

func testKruskal1() {
	n := 5
	startIndex := 0

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

	expectedMstAllNodesUsed := true

	expectedMstWeight := 10 // 3 + 4 + 1 + 2

	// from, to, weight
	// For Kruskal's, we're taking the edges in weight order (if not yet in tree)
	// In our implementation, we're ordering on from-to in case of the same weight.
	expectedMstEdges := [][]int{
		{1, 3, 1},
		{3, 4, 2},
		{0, 2, 3},
		{1, 2, 4},
	}

	testKruskal(n, startIndex, edges, expectedMstAllNodesUsed, expectedMstWeight, expectedMstEdges)
}

func testKruskal2() {
	n := 5 // 1-4, index 0 not used
	startIndex := 1

	edges := [][]int{
		// from - to - weight
		{1, 2, 3},
		{3, 4, 4},
	}

	expectedMstAllNodesUsed := false

	expectedMstWeight := 7 // 3 + 4, not connected

	// from, to, weight
	// For Kruskal's, we're taking the edges in weight order (if not yet in tree)
	// In our implementation, we're ordering on from-to in case of the same weight.
	expectedMstEdges := [][]int{
		{1, 2, 3},
		{3, 4, 4},
	}

	testKruskal(n, startIndex, edges, expectedMstAllNodesUsed, expectedMstWeight, expectedMstEdges)
}

func testKruskalSuite() {
	testKruskal1()
	testKruskal2()
}

func main() {
	testPrimSuite()
	testKruskalSuite()
}
