package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)

	// graph[i] is already an adj list of node [i]

	result := make([][]int, 0)

	// source is always the same
	src := 0

	// destination is always the same -> let's not add it as DFS param
	dest := n - 1

	path := []int{src} // add the starting node 0

	var dfs func(current int, path *[]int)

	dfs = func(current int, path *[]int) {
		if current == dest {
			//fmt.Printf("Destination node %v reached. Adding path %v to result. \n", dest, path)

			result = append(result, copyArray(*path))
		}

		// !!! In this task, we're guaranteed that the graph is ACG, i.e. we don't need to track visited
		// iterate all the adjacency list
		adj := graph[current]

		for _, to := range adj {
			*path = append(*path, to)

			dfs(to, path)

			// backtrack - remove the last element
			*path = (*path)[:len(*path)-1]
		}
	}

	// start from node 0, having node [0] in the path
	dfs(0, &path)

	return result
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func test(m [][]int, expectedResult [][]int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Adjacency list of ACG: \n%v \n", m)

	result := allPathsSourceTarget(m)

	fmt.Printf("All paths from [0] to [%v]: \n%v \n", len(m)-1, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

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

func test1() {
	adj := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	expected := [][]int{
		{0, 1, 3},
		{0, 2, 3},
	}

	test(adj, expected)
}

func test2() {
	adj := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}

	expected := [][]int{
		{0, 4},
		{0, 3, 4},
		{0, 1, 3, 4},
		{0, 1, 2, 3, 4},
		{0, 1, 4},
	}

	test(adj, expected)
}

func main() {
	// 797. All Paths From Source to Target
	test1()
	test2()
}
