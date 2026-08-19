package main

import "fmt"

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	// brute-force is working because of small constrains
	// passes in 90-95 ms.
	return maximalPathQuality_bruteForce(values, edges, maxTime)
}

func maximalPathQuality_bruteForce(values []int, edges [][]int, maxTime int) int {
	n := len(values)

	adj := createAdjacencyListUndirected(n, edges)
	//fmt.Printf("Adjacency list: %v \n", adj)

	maxValue := values[0]

	visited := make(map[int]bool)
	visited[0] = true

	var dfs func(i int, currentTime int, currentSum int)

	dfs = func(i int, currentTime int, currentSum int) {
		// currentTime already contains the path to the current node
		// currentSum already contains the current node

		if currentTime > maxTime { // already over the time -> do nothing
			return
		}

		if i == 0 { // reached the node 0
			maxValue = max(maxValue, currentSum)
		}

		for _, neighbor := range adj[i] {
			neighborId := neighbor[0]
			time := neighbor[1]

			if currentTime+time > maxTime {
				// going to this edge will go over the time -> skip
				continue
			}

			newUnvisitedNeighbor := !visited[neighborId]
			if newUnvisitedNeighbor {
				visited[neighborId] = true
			}

			newSum := currentSum
			if newUnvisitedNeighbor {
				newSum += values[neighborId]
			}

			dfs(neighborId, currentTime+time, newSum)

			// backtrack - remove the visited node
			if newUnvisitedNeighbor { // only remove from visited if it wasn't there before going to neighbor
				delete(visited, neighborId)
			}
		}
	}

	dfs(0, 0, values[0])

	return maxValue
}

func createAdjacencyListUndirected(n int, edges [][]int) [][][]int {
	adj := make([][][]int, n)

	v1 := 0
	v2 := 0
	time := 0

	for _, v := range edges {
		v1 = v[0]
		v2 = v[1]
		time = v[2]

		// add v2 to v1
		v1ToV2Edge := []int{v2, time}

		if adj[v1] == nil {
			adj[v1] = [][]int{v1ToV2Edge}
		} else {
			adj[v1] = append(adj[v1], v1ToV2Edge)
		}

		// add v1 to v2
		v2toV1Edge := []int{v1, time}

		if adj[v2] == nil {
			adj[v2] = [][]int{v2toV1Edge}
		} else {
			adj[v2] = append(adj[v2], v2toV1Edge)
		}
	}

	return adj
}

func test(values []int, edges [][]int, maxTime int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Node values: %v \n", values)
	fmt.Printf("Edges: %v \n", edges)
	fmt.Printf("Max time to travel from node 0 to node 0: %v \n", maxTime)

	result := maximalPathQuality(values, edges, maxTime)

	fmt.Printf("Max nodes sum collectible within %v time: %v \n", maxTime, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	values := []int{0, 32, 10, 43}

	edges := [][]int{
		{0, 1, 10},
		{1, 2, 15},
		{0, 3, 10},
	}

	maxTime := 49

	expected := 75

	test(values, edges, maxTime, expected)
}

func test2() {
	values := []int{5, 10, 15, 20}

	edges := [][]int{
		{0, 1, 10},
		{1, 2, 10},
		{0, 3, 10},
	}

	maxTime := 30

	expected := 25

	test(values, edges, maxTime, expected)
}

func test3() {
	values := []int{1, 2, 3, 4}

	edges := [][]int{
		{0, 1, 10},
		{1, 2, 11},
		{2, 3, 12},
		{1, 3, 13},
	}

	maxTime := 50

	expected := 7

	test(values, edges, maxTime, expected)
}

func main() {
	// 2065. Maximum Path Quality of a Graph
	test1()
	test2()
	test3()
}
