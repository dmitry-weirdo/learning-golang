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

	return false
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

func main() {
	// 207. Course Schedule
	test1()
	test2()
	test3()

	/*	// test if the graph is split
		result := canFinishReadySolution(2, [][]int{})
		fmt.Printf("result: %v \n", result)
		// yes, it's working!
	*/
}
