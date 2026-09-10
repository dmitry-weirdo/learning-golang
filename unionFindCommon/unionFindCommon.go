package unionFindCommon

import "fmt"

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
