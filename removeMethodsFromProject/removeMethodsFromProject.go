package main

import "fmt"

func remainingMethods(n int, k int, invocations [][]int) []int {
	// todo: searching with BFS can be faster because of no function call stack used

	// todo: we can even more optimize like this:
	// - if there are few bad nodes, we can run the inDegree check.
	// - if there are many bad nodes, we can run the adj check for bad nodes.
	// inDegree check complexity is size(badNodes)
	// adj check complexity is sum(size(adj)) for good nodes
	// we can calc the adjGoodNodes by calculating total sum of adj (that is probably just invocations.size),
	// and then on every bad node added to visited, we can decrease adjGoodNodes by size(adj[badNode])

	return remainingMethods_optimized(n, k, invocations)
	//return remainingMethods_trivial(n, k, invocations)
}

func remainingMethods_optimized(n int, k int, invocations [][]int) []int {
	// nodes are identified as [0; n - 1]

	// collect adjacency list
	adj := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, v := range invocations {
		from := v[0]
		to := v[1]

		adj[from] = append(adj[from], to)
		inDegree[to]++
	}

	fmt.Printf("Adjacency list: %v \n", adj)
	fmt.Printf("InDegree counts: %v \n", inDegree)

	// collecting bad nodes, DFS from the bad node K
	visited := make(map[int]bool)

	var dfs func(node int)

	dfs = func(node int) {
		// even if we came to the visited node once again, we need to decrease inDegree for the current incoming edge
		visited[node] = true

		for _, v := range adj[node] {
			// we decrease the inDegree even if the node is already visited, since we're subtracting all the edges in the DFS bad clique
			inDegree[v]--

			if !visited[v] {
				// !!! we're not checking visited since we need to decrease inDegree for this node still
				dfs(v)
			}
		}
	}

	dfs(k)

	fmt.Printf("Bad nodes: %v \n", visited)
	fmt.Printf("InDegree counts after DFS of bad nodes: %v \n", inDegree)

	goodNodesCount := n - len(visited)
	fmt.Printf("%v of %v nodes are bad. Remaining %v nodes are good. \n", len(visited), n, goodNodesCount)

	if goodNodesCount <= 0 {
		// there are no good nodes -> we're removing all the nodes from the graph. Remaining nodes are empty.
		fmt.Printf("All the %v nodes are bad. We're removing all the nodes. Returning an empty array. \n", n)
		return []int{}
	}

	// check if any bad nodes have inDegree > 0.
	// These remaining edges weren't subtracted by DFS -> they're coming from the good nodes.
	for badNode, _ := range visited {
		if inDegree[badNode] > 0 {
			fmt.Printf("For bad node %v, inDegree[%v] = %v > 0. Found an edge from good node to bad node. Returning all nodes from 0 to n. \n", badNode, badNode, inDegree[badNode])

			// return all nodes from 0 to n
			return generateArrayOfAllNodes(n)
		}
	}

	// bad nodes must be removed -> collect good nodes, return them
	goodNodes := make([]int, goodNodesCount)
	j := 0 // index of good node

	for i := range n {
		if visited[i] {
			// bad node, skip
			continue
		}

		// add to the list of good nodes
		goodNodes[j] = i
		j++
	}

	// no incoming edges from good into bad nodes -> return good nodes only
	return goodNodes
}

func remainingMethods_trivial(n int, k int, invocations [][]int) []int {
	// nodes are identified as [0; n - 1]

	// collect adjacency list
	adj := make(map[int][]int)

	for _, v := range invocations {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	fmt.Printf("Adjacency list: %v \n", adj)

	// collecting bad nodes, DFS from the bad node K
	visited := make(map[int]bool)

	var dfs func(node int)

	dfs = func(node int) {
		visited[node] = true

		for _, v := range adj[node] {
			if !visited[v] {
				dfs(v)
			}
		}
	}

	dfs(k)

	fmt.Printf("Bad nodes: %v \n", visited)

	goodNodesCount := n - len(visited)
	fmt.Printf("%v of %v nodes are bad. Remaining %v nodes are good. \n", len(visited), n, goodNodesCount)

	if goodNodesCount <= 0 {
		// there are no good nodes -> we're removing all the nodes from the graph. Remaining nodes are empty.
		fmt.Printf("All the %v nodes are bad. We're removing all the nodes. Returning an empty array. \n", n)
		return []int{}
	}

	// check if any bad nodes have an incoming edge from good nodes
	goodNodes := make([]int, goodNodesCount)
	j := 0 // index of good node

	for i := range n {
		if visited[i] {
			// bad node -> skip
			continue
		}

		// node is good

		// add to the list of good nodes
		goodNodes[j] = i
		j++

		// check whether any edge from a good node has a bad node destination
		for _, toNode := range adj[i] {
			if visited[toNode] {
				fmt.Printf("Found an edge from good node %v to bad node %v. Returning all nodes from 0 to n. \n", i, toNode)

				// return all nodes from 0 to n
				return generateArrayOfAllNodes(n)
			}
		}
	}

	// no incoming edges from good into bad nodes -> return good nodes only
	return goodNodes
}

func generateArrayOfAllNodes(n int) []int {
	// return all nodes from 0 to n
	result := make([]int, n)
	for i := range n { // k is used for function argument
		result[i] = i
	}

	return result
}

func test(n, k int, edges [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("N (total vertices): %v \n", n)
	fmt.Printf("K (starting bad vertex): %v \n", k)
	fmt.Printf("Edges: %v \n", edges)

	result := remainingMethods(n, k, edges)

	fmt.Printf("Remaining good vertices after remaining bad vertices: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

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
	n := 4
	k := 1
	edges := [][]int{{1, 2}, {0, 1}, {3, 2}}

	expected := []int{0, 1, 2, 3}

	test(n, k, edges, expected)
}

func test2() {
	n := 5
	k := 0
	edges := [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}}

	expected := []int{3, 4} // good clique {3, 4} has no edges into bad clique {0, 1, 2}

	test(n, k, edges, expected)
}

func test3() {
	n := 3
	k := 2
	edges := [][]int{{1, 2}, {0, 1}, {2, 0}}

	expected := []int{} // all {0, 1, 2} are in a bad clique

	test(n, k, edges, expected)
}

func main() {
	// 3310. Remove Methods From Project
	test1()
	test2()
	test3()
}
