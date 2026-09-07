package main

import (
	"fmt"
	"math"
)

const INFINITY = math.MaxInt32 / 2 // avoid overlow on addition

type Graph struct { // Floyd-Warshall implementation -> save the adj matrix
	n    int
	dist [][]int // n x n
}

func Constructor(n int, edges [][]int) Graph {
	return Graph{
		n:    n,
		dist: getShortestDistancesFloydWarshallDirected(n, edges, INFINITY),
	}
}

func (this *Graph) AddEdge(edge []int) { // O(V^2) for Floyd-Warshall
	from := edge[0]
	to := edge[1]
	weight := edge[2]

	// for every pair of nodes (i, j),
	// check whether we get an improvement with
	// dist[i][from] + weight + dist[to][j]
	for i := range this.n {
		for j := range this.n {
			// weight is the (from -> to) edge
			distanceWithNewEdge := this.dist[i][from] + weight + this.dist[to][j]

			this.dist[i][j] = min(this.dist[i][j], distanceWithNewEdge)
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int { // O(1) for Floyd-Warshall
	if this.dist[node1][node2] == INFINITY {
		return -1
	}

	return this.dist[node1][node2]
}

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

// ========================= Test methods ========================= //
func addEdge(g Graph, from, to, weight int) {
	g.AddEdge([]int{from, to, weight})

	fmt.Println()
	fmt.Printf("Added edge (%v -> %v) with weight %v. \n", from, to, weight)
}

func getShortestPath(g Graph, from, to, expectedResult int) {
	result := g.ShortestPath(from, to)

	fmt.Println()
	fmt.Printf("Shortest path (%v -> %v): %v \n", from, to, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	//Input
	//["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"]
	//[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]],
	// [3, 2], [0, 3], [[1, 3, 4]], [0, 3]]
	//
	//Output
	//[null, 6, -1, null, 6]

	edges := [][]int{
		{0, 2, 5},
		{0, 1, 2},
		{1, 2, 1},
		{3, 0, 3},
	}

	g := Constructor(4, edges)

	getShortestPath(g, 3, 2, 6)  // (3 -> 0 -> 1 -> 2) = 3 + 2 + 1 = 6
	getShortestPath(g, 0, 3, -1) // no path -> -1
	addEdge(g, 1, 3, 4)
	getShortestPath(g, 0, 3, 6) // (0 -> 1 -> 3) = 2 + 4 = 6
}

func main() {
	// 2642. Design Graph With Shortest Path Calculator

	// If we often add edges and rarely get shortest path -> use Dijkstra.
	// addEdge will be O(1) - just add an edge to the adj list
	// shortestPath will be heavy (full Dijkstra execution) - O(E * log V)

	// If we rarely add edges and often get shortest path -> use Floyd-Warshall
	// addEdge will be O(V^2) - heavy
	// getShortestPath will be O(1) - just get a value from the matrix

	// I implemented the Floyd-Warshall and yes, it CAN pass in 15 ms (beats 100%)
	test1()
}
