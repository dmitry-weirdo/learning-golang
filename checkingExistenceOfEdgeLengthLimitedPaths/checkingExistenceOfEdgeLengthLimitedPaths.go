package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// todo: we can solve without KRT, it by sorting queries by weight and checking uf.find(p) == uf.find(q) after all edges with edge.weight < query.weight have been processed

	// Using the more complex generic solution from. Pre-calculating KRT and then querying by it.
	// Passes, but slowly, 200-230 ms
	return distanceLimitedPathsExist_krt(n, edgeList, queries)
}

func distanceLimitedPathsExist_krt(n int, edgeList [][]int, queries [][]int) []bool {
	// this uses a more generic and complex solution of "1724. Checking Existence of Edge Length Limited Paths II"
	// Should also pass for 10^5 constraints of this task.

	dl := Constructor(n, edgeList)

	result := make([]bool, len(queries))

	for i, q := range queries {
		result[i] = dl.Query(q[0], q[1], q[2])
	}

	return result
}

type DistanceLimitedPathsExist struct {
	m map[int][]int16 // weight -> UF array for this weight
	w []int           // weights , to do binary-search for (value < weight), since we do NOT have sortedMap/treeMap in Go

	graphNodeToKrtNode map[int]*KrtNode
	ufRootToKrtRoot    map[int]*KrtNode
	indexToKrtIndex    []int
	krtData            map[string]*KrtTreeData
}

type KrtTreeData struct {
	binaryLifting [][]int
	levels        []int
	valueToNode   []*KrtNode // index [0; n-1] to node of this tree.
}

func Constructor(n int, edges [][]int) DistanceLimitedPathsExist {
	return constructor_krt(n, edges)

	// fails MLE on big counts
	//return constructor_naive(n, edges)
}

func constructor_krt(n int, edges [][]int) DistanceLimitedPathsExist {
	const VALUE_NOT_SET = -1

	graphNodeToKrtNode := make(map[int]*KrtNode) // node -> tree node in KRT

	ufRootToKrtRoot := make(map[int]*KrtNode) // UF root to KRT root

	for nodeId := range n {
		nodeKrtNode := &KrtNode{
			nodeType: Node,
			id:       nodeId,
			label:    "node_" + strconv.Itoa(nodeId),
			value:    VALUE_NOT_SET, // not set yet
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
			value:    VALUE_NOT_SET, // not set yet
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

		//fmt.Println()
		//fmt.Printf("KRT root for edge %v is now the parent of %v and %v. \n", edgeKrtNode.label, krtRootFrom.label, krtRootTo.label)
		//fmt.Printf("KRT root node of new UF root %v set to %v \n", newUfRoot, edgeKrtNode.label)

		//fmt.Printf("KRT root node of %v set to %v \n", from, edgeKrtNode.label)
		//fmt.Printf("KRT root node of %v set to %v \n", to, edgeKrtNode.label)

		/*
			fmt.Printf("ufRootToKrtRoot: \n")
			for k, v := range ufRootToKrtRoot {
				fmt.Printf("%v -> %v, ", k, v.label)
			}
			fmt.Println()
		*/

		// increase the MST edges counter
		totalEdges++

		// go to next edge
		i++
	}

	// todo: split this hell into separate methods
	// map ufRootToKrtRoot for every node to its actual tree root

	// collect KRt-tree identifiers
	krtTreeLabels := make(map[string]*KrtNode) // KRT-tree label →

	for i = range n {
		parent := uf.Find(i)
		ufRootToKrtRoot[i] = ufRootToKrtRoot[parent]

		krtTreeLabels[ufRootToKrtRoot[parent].label] = ufRootToKrtRoot[parent]
	}

	// for every KRT tree, pre-calculate binaryLifting and levels

	// !!! KRT trees do NOT have nodes 0...n -1:
	// - There can be multiple components
	// - There will be nodes for the edges

	// Therefore, to have every tree to have [0; n - 1] numeration,
	// let's assign a special Val field with the increasing counter while iterating the tree.
	// We will also need to map the original node indexes/values to the KRT nodes.

	// KrtRoot.label identifies the KRT tree
	// indexes can be repeated for different trees, so we have to check whether the nodes belong to the same component (i.e. KRT tree).
	indexToKrtIndex := make([]int, n)

	// for every tree
	componentsCount := len(krtTreeLabels)

	//fmt.Println()
	//fmt.Printf("Total graph components (KRT trees): %v \n", componentsCount)

	// we need to know the count of nodes for every KRT
	krtNodesCount := make(map[string]int, componentsCount)

	// krtLabel -> mapping of krtNodeIndex to the node
	krtNodeMappings := make(map[string][]*KrtNode)

	for krtLabel, krtTreeRoot := range krtTreeLabels {
		//fmt.Printf("KRT tree root: %v \n", krtTreeRoot.label)

		// dfs KRT tree from the root
		// set KrtNode.value and save the mapping from the original node

		value := 0

		indexToNode := make([]*KrtNode, 0)

		var dfs func(kn *KrtNode)

		dfs = func(kn *KrtNode) {
			if kn == nil {
				return
			}

			if kn.value == VALUE_NOT_SET {
				// value is set for both nodes and edges
				kn.value = value

				if kn.nodeType == Node { // save the mapping from the original node index only for edges
					indexToKrtIndex[kn.id] = value
				}

				indexToNode = append(indexToNode, kn)

				value++
			}

			dfs(kn.left)
			dfs(kn.right)
		}

		dfs(krtTreeRoot)

		// save N for every KRT
		krtNodesCount[krtLabel] = value
		//fmt.Printf("Total nodes in KRT[%v]: %v \n", krtTreeRoot.label, value)

		// save
		// index [0; N - 1] -> node
		// mappings for every KRT
		krtNodeMappings[krtLabel] = indexToNode
	}

	// Now we have KRT with values 0...n in KrtNode.value, so we can calculate BinaryLifting and LCA for every KRT
	krtData := make(map[string]*KrtTreeData) // save KRT -> binary lifting pre-calculation

	for krtLabel, krtTreeRoot := range krtTreeLabels {
		krtN := krtNodesCount[krtLabel]

		up, levels := getBinaryLiftingDfs(krtN, krtTreeRoot)

		krtData[krtLabel] = &KrtTreeData{
			binaryLifting: up,
			levels:        levels,
			valueToNode:   krtNodeMappings[krtLabel],
		}
	}

	return DistanceLimitedPathsExist{
		m:                  nil, // todo: remove legacy field of other implementation
		w:                  nil, // todo: remove legacy field of other implementation
		graphNodeToKrtNode: graphNodeToKrtNode,
		ufRootToKrtRoot:    ufRootToKrtRoot,
		indexToKrtIndex:    indexToKrtIndex,
		krtData:            krtData,
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
	// Querying LCA using Binary Lifting is O(log N)
	// If the tree is split into several components, N will be smaller.
	// This passes in 100-120 ms!
	return this.lca_binaryLifting(p, q, limit)

	// this is a stupid O(N) solution with iterating from node to its parent. This is basically the same as DFS through the MST graph
	// without binary-lifting, O(N) search is failing on TLE
	//return this.lca_byParent(p, q, limit)
}

func (this *DistanceLimitedPathsExist) lca_binaryLifting(p int, q int, limit int) bool {
	if this.ufRootToKrtRoot[p] != this.ufRootToKrtRoot[q] {
		// P and Q are in different components -> return false
		return false
	}

	// P and Q are in the same component
	// -> find LCA in the KRT of this component and check whether its weight < limit
	krtLabel := this.ufRootToKrtRoot[p].label
	krt := this.krtData[krtLabel]

	pIndexInKrt := this.indexToKrtIndex[p]
	qIndexInKrt := this.indexToKrtIndex[q]

	lcaIndex := getLca(krt.binaryLifting, krt.levels, pIndexInKrt, qIndexInKrt)
	lcaNode := krt.valueToNode[lcaIndex]
	//fmt.Printf("KRT[%v]: LCA of %v and %v is %v (weight: %v) \n", krtLabel, pIndexInKrt, qIndexInKrt, lcaNode.label, lcaNode.weight)

	return lcaNode.weight < limit
}

func (this *DistanceLimitedPathsExist) lca_byParent(p int, q int, limit int) bool {
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

	value int // [0; n - 1] values, N nodes in the tree. To be numerated [0; n - 1] for the binary lifting and LCA

	left   *KrtNode // not necessary, but for consistency
	right  *KrtNode // not necessary, but for consistency
	parent *KrtNode // to iterate upwards to find the LCA
}

// ========================= My KRT implementation end ========================= //

// ========================= Binary Lifting and LCA implementation begin ========================= //
func getBinaryLiftingDfs(n int, root *KrtNode) (binaryLifting [][]int, levels []int) { // adopted from *TreeNode to *KrtNode for this task
	// see https://www.youtube.com/watch?v=dOAxrhAUIhA

	// todo: generalize how get Val from TreeNode! It's not always TreeNode.Val

	// todo: we should also handle node indexes, not values

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	log := log2(n) + 1

	// up[v][i]
	// v - [0; n-1] - tree node values
	// i - 2^i jumps from 0 to log2(n)
	// todo: probably we need to map not to [int] but to [*TreeNode] ???
	up := createIntMatrix(n, log)

	// set -1 for root, all levels
	for i := range log {
		up[root.value][i] = -1
	}

	// also fill the levels array
	levels = make([]int, n) // for nodes 0 to n - 1

	var dfs func(n *KrtNode, level int)

	// todo: probably we need to map not to [int] but to [*TreeNode]
	dfs = func(n *KrtNode, level int) {
		if n == nil {
			return
		}

		levels[n.value] = level

		// left
		if n.left != nil {
			//fmt.Printf("Root: %v, left: %v \n", root.Val, n.Left.Val)

			// todo: value should be *TreeNode?
			// 2^0 - direct parent
			v := n.left.value
			up[v][0] = n.value

			// fill 2^i ancestor:
			// 2^i = 2^(i - 1) + 2^(i - 1)
			for i := 1; i < log; i++ { // powers 1, 2, ..., log2(n). For 100, we will iterate from 2^1 = 2 to 2^6 = 64
				// at this 2^i jump level, fill all the nodes
				p := up[v][i-1]

				if p == -1 { // reached the parent // todo: do we need this?
					up[v][i] = -1
				} else { // from 2^(i-1) parent, get it 2^(i-1) parent. The sum will sum up to 2^i parent of the current node.
					up[v][i] = up[p][i-1]
				}
			}

			dfs(n.left, level+1)
		}

		// right
		if n.right != nil {
			//fmt.Printf("Root: %v, right: %v \n", root.Val, n.Right.Val)

			// 2^0 - direct parent
			v := n.right.value
			up[v][0] = n.value

			// fill 2^i ancestor:
			// 2^i = 2^(i - 1) + 2^(i - 1)
			for i := 1; i < log; i++ { // powers 1, 2, ..., log2(n). For 100, we will iterate from 2^1 = 2 to 2^6 = 64
				// at this 2^i jump level, fill all the nodes
				p := up[v][i-1]

				if p == -1 { // reached the parent // todo: do we need this?
					up[v][i] = -1
				} else { // from 2^(i-1) parent, get it 2^(i-1) parent. The sum will sum up to 2^i parent of the current node.
					up[v][i] = up[p][i-1]
				}
			}

			dfs(n.right, level+1)
		}
	}

	dfs(root, 0) // root has level 0
	return up, levels
}

func log2(n int) int {
	// todo: log2 should be handled separately, it's undefined
	return int(math.Log2(float64(n)))

	// for positive integers, counting bits can be used:
	// bits.Len(uint(n)) - 1
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func getLca(up [][]int, levels []int, a, b int) int {
	// define what node is deeper in the tree
	lower, upper := a, b

	if levels[a] < levels[b] {
		lower, upper = b, a
	}

	levelDiff := levels[lower] - levels[upper]

	//fmt.Printf("Lower node: %v, upper node: %v, level difference: %v \n", lower, upper, levelDiff)

	// move from the lower (deeper) node to the same level as the upper (shallower) node
	lower = GetKthAncestor(up, lower, levelDiff)

	//fmt.Printf("Lower node moved to the same level %v as upper node %v. Lower node moved up to %v. \n", levels[upper], upper, lower)

	if lower == upper { // at the same level, nodes are the same -> upper node is the LCA
		return lower
	}

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	n := len(up)

	log := log2(n) + 1

	for i := log - 1; i >= 0; i-- {
		// if the ancestor of this level is the same, continue to the next level
		// I.e. this level is LCA or above
		if up[lower][i] == up[upper][i] {
			continue
		}

		// Ancestors of this level is different -> this level is below LCA
		// Move to this level (to the power of 2)
		// LCA will be still above.
		lower = up[lower][i]
		upper = up[upper][i]
	}

	// both nodes will be directly below their LCA
	return up[lower][0]
}

func GetKthAncestor(up [][]int, node int, k int) int {
	// todo: handle -1 specially?
	n := len(up)

	// e.g. for 100, log = 6, and we need powers from 2^0 to 2^6 (from 0 to 64)
	log := log2(n) + 1

	current := node

	for i := log - 1; i >= 0; i-- {
		if current == -1 { // no need to traverse further if we're above the root
			return -1
		}

		// k = 100 -> we'll go
		powerOf2 := 1 << i
		//fmt.Printf("Power of 2^%v = %v \n", i, powerOf2)

		if k >= powerOf2 {
			// go up 2^i levels, decreasing K
			current = up[current][i]

			k -= powerOf2
		}
	}

	return current
}

// ========================= Binary Lifting and LCA implementation end ========================= //

func test(n int, edges [][]int, queries [][]int, expectedResult []bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - count of vertices in the graph: %v \n", n)
	fmt.Printf("Edges: %v \n", edges)
	fmt.Printf("Total queries: %v \n", len(queries))
	fmt.Printf("Queries: %v \n", queries)

	result := distanceLimitedPathsExist(n, edges, queries)

	fmt.Printf("Results of queries: %v \n", result)
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

	edges := [][]int{
		{0, 1, 2},
		{1, 2, 4},
		{2, 0, 8},
		{1, 0, 16},
	}

	queries := [][]int{
		{0, 1, 2},
		{0, 2, 5},
	}

	expected := []bool{
		false,
		true,
	}

	test(n, edges, queries, expected)
}

func test2() {
	n := 5

	edges := [][]int{
		{0, 1, 10},
		{1, 2, 5},
		{2, 3, 9},
		{3, 4, 13},
	}

	queries := [][]int{
		{0, 4, 14},
		{1, 4, 13},
	}

	expected := []bool{
		true,
		false,
	}

	test(n, edges, queries, expected)
}

func main() {
	// 1697. Checking Existence of Edge Length Limited Paths

	// This is a simpler analogue of "1724. Checking Existence of Edge Length Limited Paths II",
	// but the constraints on both N and queries are 10^5 instead of 10^4.

	// The idea is that we know the queries in advance, so we can order them by weight
	// And on every weight reached check whether uf.find(p) == uf.find(q)
	test1()
	test2()
}
