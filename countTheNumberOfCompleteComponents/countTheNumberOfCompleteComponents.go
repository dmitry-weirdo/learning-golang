package main

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
	// passes in 5-15 ms -> ok
	return countCompleteComponents_unionFind(n, edges)
}

func countCompleteComponents_unionFind(n int, edges [][]int) int {
	// group nodes into connected components
	uf := newUnionFind(n)

	for _, v := range edges {
		uf.union(v[0], v[1])
	}

	//fmt.Printf("Parents of nodes: %v \n", uf.parents)
	//fmt.Printf("Sizes of groups: %v \n", uf.sizes)

	// group to edges count
	groupToEdgesCount := make(map[int]int)

	for _, v := range edges {
		group := uf.find(v[0]) // we count just for one of the (from, to) nodes
		groupToEdgesCount[group]++
	}

	//fmt.Printf("Count of edges by group: %v \n", groupToEdgesCount)

	// count of nodes in every group
	groupSizes := uf.getGroupsSizes()
	//fmt.Printf("Group sizes: %v \n", groupSizes)

	totalCompleteGroups := 0

	for group, nodesInGroup := range groupSizes {
		// Full graph will contain n * (n - 1) / 2 edges.
		// Every node of N nodes connected to (n - 1) other nodes,
		// divide by 2 since we counted every edge twice for both its nodes.
		edgesCountForCompleteGroup := nodesInGroup * (nodesInGroup - 1) / 2

		if groupToEdgesCount[group] == edgesCountForCompleteGroup {
			totalCompleteGroups++
		}
	}

	return totalCompleteGroups
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

	fmt.Printf("Total vertices (from 0 to %v): %v \n", n-1, n)
	fmt.Printf("Edges: %v \n", edges)

	result := countCompleteComponents(n, edges)

	fmt.Printf("Complete components count: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 6

	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
	}

	expected := 3

	test(n, edges, expected)
}

func test2() {
	n := 6

	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
		{3, 5},
	}

	expected := 1

	test(n, edges, expected)
}

func main() {
	// 2685. Count the Number of Complete Components
	test1()
	test2()
}
