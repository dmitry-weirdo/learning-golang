package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

type DistanceLimitedPathsExist struct {
	m map[int][]int16 // weight -> UF array for this weight
	w []int           // weights , to do binary-search for (value < weight), since we do NOT have sortedMap/treeMap in Go

	graphNodeToKrtNode map[int]*KrtNode
	ufRootToKrtRoot    map[int]*KrtNode
}

func Constructor(n int, edges [][]int) DistanceLimitedPathsExist {
	return constructor_krt(n, edges)

	// fails MLE on big counts
	//return constructor_naive(n, edges)
}

func constructor_krt(n int, edges [][]int) DistanceLimitedPathsExist {
	graphNodeToKrtNode := make(map[int]*KrtNode) // node -> tree node in KRT

	ufRootToKrtRoot := make(map[int]*KrtNode) // UF root to KRT root

	for nodeId := range n {
		nodeKrtNode := &KrtNode{
			nodeType: Node,
			id:       nodeId,
			label:    "node_" + strconv.Itoa(nodeId),
			left:     nil,
			right:    nil,
			parent:   nil,
		}

		// At the start, every node is its own KRT tree
		graphNodeToKrtNode[nodeId] = nodeKrtNode
		ufRootToKrtRoot[nodeId] = nodeKrtNode
	}

	// Simplified Kruskal
	// We don't need the MST, we only need the UF states
	startIndex := 0 // nodes start with 0

	sortEdgesForKruskal(edges)

	// union-find controls the visited by including them in the same node set
	uf := NewUnionFind(n)

	totalEdges := 0

	// index in edges[] array
	i := 0

	for (totalEdges < n-startIndex-1) && (i < len(edges)) {
		from := edges[i][0]
		to := edges[i][1]
		weight := edges[i][2]

		ufRootFrom := uf.Find(from)
		ufRootTo := uf.Find(to)

		if !uf.Union(from, to) { // from and to already in the same MST set -> skip this edge
			i++
			continue
		}

		/*		// should we still put to the map if uf.Union returned false? No, since the edge is not added
				if weight%10 == 0 {
					// with int, it fails on 1880
					// with int16, it fails on 7690 - much better, but still MLE, we need 10000 to work :(
					fmt.Printf("%v \n", weight)
				}
		*/

		newUfRoot := uf.Find(from) // after union

		krtRootFrom := ufRootToKrtRoot[ufRootFrom]
		krtRootTo := ufRootToKrtRoot[ufRootTo]

		edgeKrtNode := &KrtNode{
			nodeType: Edge,
			weight:   weight,
			label:    "edge_" + strconv.Itoa(from) + "-" + strconv.Itoa(to),
			left:     krtRootFrom,
			right:    krtRootTo,
		}

		krtRootFrom.parent = edgeKrtNode
		krtRootTo.parent = edgeKrtNode

		// !!! We only care about the roots of the UF in the ufRootToKrtRoot map.
		// In the test example, after merging nodes 2 and 3, we're setting the parent just to one node 2 that is now the parent of both 2 and 3 (in the UF).
		// ufRootToKrtRoot[3] remains unchanged in ufRootToKrtRoot, but we're not using it anymore.
		// 1 -> node_1, 2 -> edge_2-3, 3 -> node_3, 4 -> node_4, 5 -> node_5, 0 -> node_0,

		// from and to now belong to the new root of the edge
		ufRootToKrtRoot[newUfRoot] = edgeKrtNode
		//ufRootToKrtRoot[from] = edgeKrtNode
		//ufRootToKrtRoot[to] = edgeKrtNode

		fmt.Println()
		fmt.Printf("KRT root for edge %v is now the parent of %v and %v. \n", edgeKrtNode.label, krtRootFrom.label, krtRootTo.label)
		fmt.Printf("KRT root node of new UF root %v set to %v \n", newUfRoot, edgeKrtNode.label)
		//fmt.Printf("KRT root node of %v set to %v \n", from, edgeKrtNode.label)
		//fmt.Printf("KRT root node of %v set to %v \n", to, edgeKrtNode.label)

		fmt.Printf("ufRootToKrtRoot: \n")
		for k, v := range ufRootToKrtRoot {
			fmt.Printf("%v -> %v, ", k, v.label)
		}
		fmt.Println()

		// increase the MST edges counter
		totalEdges++

		// go to next edge
		i++
	}

	// todo: we need to store the UF.parents array to quickly define the component

	// todo: for every tree in the KRT forest, pre-calculate the binary lifting and the levels

	return DistanceLimitedPathsExist{
		m:                  nil, // todo: remove legacy field of other implementation
		w:                  nil, // todo: remove legacy field of other implementation
		graphNodeToKrtNode: graphNodeToKrtNode,
		ufRootToKrtRoot:    ufRootToKrtRoot,
	}
}

func constructor_naive(n int, edges [][]int) DistanceLimitedPathsExist {
	// clone of UF array for every weight
	// If multiple node have the same weights, we will override after adding more nodes.
	m := make(map[int][]int16) // weight -> UF array for this weight
	w := make([]int, 0)        // weights , to do binary-search for (value < weight), since we do NOT have sortedMap/treeMap in Go

	// Simplified Kruskal
	// We don't need the MST, we only need the UF states
	startIndex := 0 // nodes start with 0

	sortEdgesForKruskal(edges)

	// union-find controls the visited by including them in the same node set
	uf := NewUnionFind(n)

	totalEdges := 0

	// index in edges[] array
	i := 0

	for (totalEdges < n-startIndex-1) && (i < len(edges)) {
		from := edges[i][0]
		to := edges[i][1]
		weight := edges[i][2]

		if !uf.Union(from, to) { // from and to already in the same MST set -> skip this edge
			i++
			continue
		}

		// should we still put to the map if uf.Union returned false? No, since the edge is not added
		if weight%10 == 0 {
			// with int, it fails on 1880
			// with int16, it fails on 7690 - much better, but still MLE, we need 10000 to work :(
			fmt.Printf("%v \n", weight)
		}

		w = append(w, weight)
		m[weight] = copyArray(uf.parents)

		// increase the MST edges counter
		totalEdges++

		// go to next edge
		i++
	}

	// store the state for the Query execution
	return DistanceLimitedPathsExist{m: m, w: w}
}

func copyArray(arr []int) []int16 {
	// todo: we can make union-find array also int16
	arrayCopy := make([]int16, len(arr))

	for i, v := range arr {
		arrayCopy[i] = int16(v)
	}

	return arrayCopy

	/*	arrayCopy := make([]int16, len(arr))
		copy(arrayCopy, int16(arr))
		return arrayCopy*/
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
	return this.Query_Krt(p, q, limit)

	// Nice and O(1), but we will get MLE on saving all the UF arrays
	//return this.Query_Naive(p, q, limit)
}

func (this *DistanceLimitedPathsExist) Query_Krt(p int, q int, limit int) bool {
	// we need to find the LCA of p and q. This will be the minimum weight in their paths
	/*	if this.graphNodeToKrtRoot[p] != this.graphNodeToKrtRoot[q] { // nodes belong to different components
			return false
		}
	*/
	// todo: this is a stupid O(N) solution with iterating from node to its parent. This is basically the same as DFS through the MST graph
	// find the LCA
	np := this.graphNodeToKrtNode[p]
	nq := this.graphNodeToKrtNode[q]

	lca := lowestCommonAncestor_byParent(np, nq)

	if lca == nil { // this must never happen
		//panic(fmt.Sprintf("Nodes %v and %v belong to the same KRT but have no LCA.", p, q))
		fmt.Printf("LCA of nodes %v and %v not found. \n", p, q)
		return false
	}

	fmt.Printf("LCA of nodes %v and %v: %v, weight = %v (%v) \n", p, q, lca.label, lca.weight, lca)
	return lca.weight < limit
}

func lowestCommonAncestor_byParent(p *KrtNode, q *KrtNode) *KrtNode { // copied from "1650. Lowest Common Ancestor of a Binary Tree III"
	// if we want O(1) space optimization logic, we can:
	// - calculate depths of both nodes
	// - go from the deeper node up to align the depths
	// - move up step by step until the nodes are the same
	// It will still be O(h) time complexity

	m := make(map[string]*KrtNode)

	// go from P to the root, collect all the values in the path into a map
	current := p

	for current != nil {
		m[current.label] = current
		current = current.parent
	}

	// go from Q to the root. The first node in the path that was already in the path of P is the LCA
	current = q

	for current != nil {
		if _, ok := m[current.label]; ok {
			return current
		}

		current = current.parent
	}

	// this must never happen, at least the root should be the LCA
	return nil
}

func (this *DistanceLimitedPathsExist) Query_Naive(p int, q int, limit int) bool {
	// find biggest existing weight < weight
	index := searchRightmostLessThanTarget(this.w, limit)
	if index < 0 { // there is no weight less than limit
		return false
	}

	// get the union find state for the max satisfying weight
	weight := this.w[index]
	ufState := this.m[weight]

	// check whether in this state, nodes P and Q were in the same MST subgraph
	return Find(ufState, int16(p)) == Find(ufState, int16(q))
}

func Find(parents []int16, x int16) int16 { // recursive version
	if parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	parents[x] = Find(parents, parents[x])

	return parents[x]
}

func searchRightmostLessThanTarget(arr []int, target int) int {
	condition := func(x int) bool {
		return x >= target
	}

	index := binarySearchGeneric(
		arr,
		0,
		len(arr), // insert position can be after the end of the array
		condition,
	)

	// there is no previous element -> no result
	if index <= 0 {
		return -1
	}

	// go one element left -> this will be the last element < target
	index--

	if arr[index] >= target {
		return -1
	}

	return index
}

func binarySearchGeneric(
	arr []int, // todo: we can also generalize the type in the array
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func(int) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// todo: this method can return an incorrect value for the empty array

	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

// ========================= Kruskal's algorithm for MST begin ========================= //
type MinimumSpanningTree struct {
	edges        [][]int // every edge is from/to/weight
	weight       int     // total weight for all edges
	allNodesUsed bool    // false if there are multiple non-connected components
}

func getMinimumSpanningTreeKruskalNodesStartFrom0(n int, edges [][]int) MinimumSpanningTree {
	return getMinimumSpanningTreeKruskal(n, 0, edges) // start from node 0
}

func getMinimumSpanningTreeKruskalNodesStartFrom1(n int, edges [][]int) MinimumSpanningTree {
	return getMinimumSpanningTreeKruskal(n+1, 1, edges) // start from node 1
}

// todo: for many problems, we only need to connect the MST.weight. So we can save time and space on collecting the edges.
func getMinimumSpanningTreeKruskal(n int, startIndex int, edges [][]int) MinimumSpanningTree { // start can be 0 or 1
	// If nodes start with 0, set n = N, startIndex = 0
	// If nodes start with 1, set n = N + 1, startIndex = 1

	// Prim can work with negative edge weights.
	// Returns an array of edges of MST.
	// We assume that the graph is undirected. It's necessary for Kruskal's since we're uniting 2 nodes with Union-Find

	// Instead of using a complex Heap structure for getting the smallest edge, we can just sort an array of edges by weight
	// Complexity should be the save O(E * log E)
	sortEdgesForKruskal(edges)

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

func sortEdgesForKruskal(edges [][]int) {
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

// ========================= My KRT implementation begin ========================= //
type KrtNodeType int

const (
	Node KrtNodeType = iota // node of the original graph
	Edge                    // edge of the original graph
)

type KrtNode struct {
	nodeType KrtNodeType
	weight   int    // for edge only
	id       int    // for node only
	label    string // to uniquely identify the node

	left   *KrtNode // not necessary, but for consistency
	right  *KrtNode // not necessary, but for consistency
	parent *KrtNode // to iterate upwards to find the LCA
}

// ========================= My KRT implementation end ========================= //

func testQuery(d DistanceLimitedPathsExist, p, q, limit int, expectedResult bool) {
	fmt.Println()

	result := d.Query(p, q, limit)

	fmt.Printf("Query the path from %v to %v within %v edge weight limit. \n", p, q, limit)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 6

	edges := [][]int{
		{0, 2, 4},
		{0, 3, 2},
		{1, 2, 3},
		{2, 3, 1},
		{4, 5, 5},
	}

	d := Constructor(n, edges)

	testQuery(d, 2, 3, 2, true)
	testQuery(d, 1, 3, 3, false)
	testQuery(d, 2, 0, 3, true)
	testQuery(d, 0, 5, 6, false)
}

func main() {
	// 1724. Checking Existence of Edge Length Limited Paths II

	// MST is always a MBST (Minimum Bottleneck Spanning Tree).
	// See https://en.wikipedia.org/wiki/Minimum_bottleneck_spanning_tree#Properties
	// See https://stanford.edu/~rezab/discrete/Midterm/midterm2016_soln.pdf

	// Therefore, we should get the max weight of path nodeP -> nodeQ within the MST.

	// Notably, when Kruskal calculates the MST, we go from the minimal weight edges up.
	// So, we need to have the union-find state on every edge weight.
	// All the nodes connected on this step will be reachable within edge[i].weight.

	// The nice trick is to clone the UF array for every weight.
	// Then by executing Find(P) and Find(Q) on this clone, we can define whether they belong to the same MST within the weight.

	// This is a super-nice solution, but it fails MLE on Test-case 51/53, with 10000 nodes.
	// Copying and array of 10000 elements up to 10000 times (there are 10^4 edges that can have distinct lengths) is too much.

	test1()
}
