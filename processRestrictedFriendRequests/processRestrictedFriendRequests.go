package main

import "fmt"

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := newUnionFind(n)

	success := make([]bool, len(requests)) // all false by default

	for i, request := range requests {
		// check all restriction, if any is blocking
		canFriend := true

		u := request[0]
		v := request[1]

		uParent := uf.find(u)
		vParent := uf.find(v)

		requestMinParent, requestMaxParent := getMinAndMax(uParent, vParent)

		xParent := -1
		yParent := -1

		// check all the restriction whether they block U and V
		for _, restriction := range restrictions {
			xParent = uf.find(restriction[0])
			yParent = uf.find(restriction[1])

			restrictionMinParent, restrictionMaxParent := getMinAndMax(xParent, yParent)

			if (requestMinParent == restrictionMinParent) && (requestMaxParent == restrictionMaxParent) { // blocking restriction found
				canFriend = false
				break
			}
		}

		if canFriend {
			// can friend -> join groups of people in the request
			uf.union(u, v)

			success[i] = true
		}
		//} else { // it is already false by default
		//	success[i] = false
		//}
	}

	return success
}

func getMinAndMax(a, b int) (smaller, greater int) {
	if a <= b {
		return a, b
	}

	return b, a
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

func test(n int, restrictions, requests [][]int, expectedResult []bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N (total users): %v \n", n)
	fmt.Printf("Restrictions that block friend requests: %v \n", restrictions)
	fmt.Printf("Friend requests: %v \n", requests)

	result := friendRequests(n, restrictions, requests)

	fmt.Printf("Success of %v requests: %v \n", len(requests), result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	n := 3

	restrictions := [][]int{
		{0, 1},
	}

	requests := [][]int{
		{0, 2},
		{2, 1},
	}

	expected := []bool{true, false}

	test(n, restrictions, requests, expected)
}

func test2() {
	n := 3

	restrictions := [][]int{
		{0, 1},
	}

	requests := [][]int{
		{1, 2},
		{0, 2},
	}

	expected := []bool{true, false}

	test(n, restrictions, requests, expected)
}

func test3() {
	n := 5

	restrictions := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}

	requests := [][]int{
		{0, 4},
		{1, 2},
		{3, 1},
		{3, 4},
	}

	expected := []bool{true, false, true, false}

	test(n, restrictions, requests, expected)
}

func main() {
	// 2076. Process Restricted Friend Requests
	test1()
	test2()
	test3()
}
