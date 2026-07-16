package main

import (
	"fmt"
)

func canFinishReadySolution(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = []int{}
	}

	for _, prereq := range prerequisites {
		src, dst := prereq[0], prereq[1]
		indegree[dst]++
		adj[src] = append(adj[src], dst)
	}

	q := []int{}
	for n := 0; n < numCourses; n++ {
		if indegree[n] == 0 {
			q = append(q, n)
		}
	}

	finish := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		finish++
		for _, nei := range adj[node] {
			indegree[nei]--
			if indegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	return finish == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// this version is using the Kahn's algorithm
	// if we added all nodes to the top-sort array -> no cycles

	// every course is identified by i = [0; numCourses - 1]
	// so we can identify a course with just a position in the array

	// adjacency list -> for every vertex, we keep a list of destination from this vertex
	adj := make([][]int, numCourses)

	// initialize adj with empty lists for every vertex
	for i := 0; i < len(adj); i++ {
		adj[i] = []int{}
	}

	// how many inbound edges are there for every vertex
	inDegree := make([]int, numCourses)

	for _, v := range prerequisites {
		// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
		from := v[1]
		to := v[0]

		inDegree[to]++

		adj[from] = append(adj[from], to)
	}

	fmt.Printf("adjacency list: %v \n", adj)
	fmt.Printf("inDegree list: %v \n", inDegree)

	// queue of nodes to process
	queue := []int{}

	// add all nodes that have 0 inDegree to the queue
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	fmt.Printf("Initial nodes without incoming edges (with inDegree = 0): %v \n", queue)

	// topological sort array
	index := 0 // index in top-sort array
	order := make([]int, numCourses)

	for len(queue) > 0 {
		// pop from the queue
		node := queue[0]
		queue = queue[1:]

		// put the popped vertex to the top-sort
		order[index] = node
		index++ // will also be increased after the last node -> although this is 0-based, it will reach N (numCourses) if all nodes have been processed

		// for all nodes going out of the node, decrease their inDegree
		for _, dest := range adj[node] {
			inDegree[dest]--

			// if the neighbour going out of the node now has inDegree == 0 -> put it to the queue
			if inDegree[dest] == 0 {
				queue = append(queue, dest)
			}
		}
	}

	// If we handled all the nodes via inDegree == 0 logic, then there are no cycles.
	// If there are some cycles, nodes with inDegree > 0 will remain.
	fmt.Printf("Topological sort: %v \n", order)
	fmt.Printf("Topological sort covered %v of %v vertexes. \n", index, numCourses)
	fmt.Printf("Topological sort finished (i.e. no cycles in the graph): %v \n", index == numCourses)

	return index == numCourses
}

func test(numCourses int, prerequisites [][]int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number of courses: %v \n", numCourses)
	fmt.Printf("Prerequisites: %v \n", prerequisites)

	result := canFinish(numCourses, prerequisites)

	fmt.Printf("Can finish all courses (no cycles in the graph): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 2
	prerequisites := [][]int{
		{1, 0},
	}

	expected := true

	test(n, prerequisites, expected)
}

func test2() {
	n := 2
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}

	expected := false

	test(n, prerequisites, expected)
}

func test3() {
	n := 4

	prerequisites := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}

	expected := true

	test(n, prerequisites, expected)
}

func test4() {
	// from https://youtu.be/cIBFEhD77b4?t=486
	n := 14

	prerequisites := [][]int{
		{2, 0},
		{3, 0},
		{6, 0},
		{4, 1},
		{6, 2},
		{4, 3},
		{1, 3},
		{5, 4},
		{8, 4},
		// nothing from 5
		{7, 6},
		{11, 6},
		{4, 7},
		{12, 7},
		// nothing from 8
		{2, 9},
		{10, 9},
		{6, 10},
		{12, 11},
		{8, 12},
		// nothing from 13
	}

	expected := true

	test(n, prerequisites, expected)
}

func main() {
	// 207. Course Schedule
	test1()
	test2()
	test3()
	test4()

	/*	// test if the graph is split
		result := canFinishReadySolution(2, [][]int{})
		fmt.Printf("result: %v \n", result)
		// yes, it's working!
	*/
}
