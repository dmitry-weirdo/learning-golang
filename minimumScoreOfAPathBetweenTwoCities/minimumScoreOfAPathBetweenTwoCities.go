package main

import "fmt"

func minScore(n int, roads [][]int) int {
	// in the component 1 - n, find the smallest edge
	// group nodes into connected components
	uf := newUnionFind(n + 1)

	for _, v := range roads {
		uf.union(v[0], v[1])
	}

	//fmt.Printf("Parents of nodes: %v \n", uf.parents)
	//fmt.Printf("Sizes of groups: %v \n", uf.sizes)

	group1 := uf.find(1)
	groupN := uf.find(n)

	//fmt.Printf("Group of node 1: %v \n", group1)
	//fmt.Printf("Group of node N = %v: %v \n", n, groupN)

	if group1 != groupN {
		panic(fmt.Sprintf("Nodes 1 and N = %v are in different groups %v and %v.", n, group1, groupN))
	}

	const BIGGER_THAN_MAX_EDGE = 999_999 // edge is up to 10_000

	minEdgeInGroup1 := BIGGER_THAN_MAX_EDGE

	for _, v := range roads {
		if uf.find(v[0]) != group1 { // edge not in the 1-N group -> skip
			continue
		}

		minEdgeInGroup1 = min(minEdgeInGroup1, v[2])
	}

	return minEdgeInGroup1
}

type UnionFind struct {
	parents []int // if parent[i] = i, it is the root, else it's the index of the parent
	sizes   []int // sizes of the tree for every element
}

func newUnionFind(n int) UnionFind {
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

func (uf UnionFind) find(x int) int { // recursive version
	if uf.parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	uf.parents[x] = uf.find(uf.parents[x])

	return uf.parents[x]
}

func (uf UnionFind) print() {
	fmt.Printf("Parents: %v \n", uf.parents)
	fmt.Printf("Sizes: %v \n", uf.sizes)
}

func (uf UnionFind) union(x, y int) bool { // returns false if they're already in the same set
	// these find will perform path compression
	rootX := uf.find(x)
	rootY := uf.find(y)

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

func (uf UnionFind) groupSize(x int) int {
	return uf.sizes[uf.find(x)]
}

func (uf UnionFind) getGroupsSizes() map[int]int { // returns sizes for every group
	m := make(map[int]int)

	for i := range uf.parents {
		if uf.find(i) == i { // root node
			// every root group will be iterated just once, no need to check whether it's already in the map
			m[i] = uf.groupSize(i)

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

func test(n int, edges [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - total nodes from 1 to %v: %v \n", n, n)
	fmt.Printf("Edges: %v \n", edges)

	result := minScore(n, edges)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 4

	edges := [][]int{
		{1, 2, 9},
		{2, 3, 6},
		{2, 4, 5},
		{1, 4, 7},
	}

	expected := 5

	test(n, edges, expected)
}

func test2() {
	n := 4

	edges := [][]int{
		{1, 2, 2},
		{1, 3, 4},
		{3, 4, 7},
	}

	expected := 2

	test(n, edges, expected)
}

func main() {
	// 2492. Minimum Score of a Path Between Two Cities
	test1()
	test2()
}
