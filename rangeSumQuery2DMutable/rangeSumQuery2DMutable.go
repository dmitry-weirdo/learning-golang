package main

import "fmt"

type NumMatrix struct {
	trees []*FenwickTree
}

func Constructor(matrix [][]int) NumMatrix {
	// for every row, create a Fenwick Tree
	trees := make([]*FenwickTree, len(matrix))

	for i, row := range matrix {
		tree := createFenwickTree(row)
		trees[i] = &tree
	}

	return NumMatrix{trees: trees}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	this.trees[row].Set(col+1, val) // Fenwick Tree indexes are 1-based
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// FW-sum for every row
	sum := 0

	for i := row1; i <= row2; i++ {
		sum += this.trees[i].RangeSum(col1+1, col2+1) // Fenwick Tree indexes are 1-based
	}

	return sum
}

// simple implementation O(m * log n) range sums - array of 1D Fenwick trees, for every row
// we just sum up the range for every row.
// So this is not a true 2D-tree.

type FenwickTree struct {
	t []int
}

func createFenwickTree(a []int) FenwickTree {
	n := len(a)

	t := make([]int, n+1) // t-tree is 1-indexed
	copy(t[1:], a)        // t[0] is not used, we copy from pos 1

	for i := 1; i <= n; i++ {
		j := i + lsb(i) // add to the LSB of i

		if j <= n { // no overflow of array size
			t[j] += t[i]
		}
	}

	return FenwickTree{t: t}
}

func (ft *FenwickTree) prefixSum(i int) int { // sums elements with indexes [1; i], 1-based, inclusive
	sum := 0

	for i > 0 {
		sum += ft.t[i]

		// clear (set to 0) the least significant bit of i
		i &= ^lsb(i)
	}

	return sum
}

func (ft *FenwickTree) RangeSum(i, j int) int { // 1-based indexes, returns sum of [i, j] inclusive
	if i > j {
		panic(fmt.Sprintf("Incorrect range. i = %v > j = %v.", i, j))
	}

	return ft.prefixSum(j) - ft.prefixSum(i-1) // sum[i - 1] to include i
}

func (ft *FenwickTree) Set(i, value int) {
	currentValue := ft.RangeSum(i, i) // ft[i] is NOT value of the array, it's a partial sum [i - lsb(i) + 1, i]

	delta := value - currentValue // calc the value to add

	ft.Add(i, delta)
}

func (ft *FenwickTree) Add(i, delta int) {
	for i < len(ft.t) {
		ft.t[i] += delta

		// add 1 to the LSB, i.e. propagate to parents
		i += lsb(i)
	}
}

func lsb(x int) int { // least significant bit
	return x & -x
}

func testSumRegion(m NumMatrix, r1, c1, r2, c2, expectedResult int) {
	result := m.SumRegion(r1, c1, r2, c2)

	fmt.Printf("Sum of region [%v][%v] - [%v][%v]: %v. \n", r1, c1, r2, c2, result)
	fmt.Printf("Expected sum: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func testUpdate(m NumMatrix, r, c, value int) {
	m.Update(r, c, value)
	fmt.Printf("Updated m[%v][%v] to %v. \n", r, c, value)

	// sum just in this cell should be the new value
	testSumRegion(m, r, c, r, c, value)
}

func test1() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	m := Constructor(matrix)

	testSumRegion(m, 2, 1, 4, 3, 8) // 8 (2 0 1, 1 0 1, 0 3 0)
	testUpdate(m, 3, 2, 2)
	testSumRegion(m, 2, 1, 4, 3, 10) // 10 (2 0 1, 1 2 1, 0 3 0)
}

func main() {
	// 308. Range Sum Query 2D - Mutable
	test1()
}
