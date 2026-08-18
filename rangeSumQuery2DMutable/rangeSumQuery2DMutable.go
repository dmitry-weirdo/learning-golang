package main

import "fmt"

type NumMatrix struct {
}

func Constructor(matrix [][]int) NumMatrix {
}

func (this *NumMatrix) Update(row int, col int, val int) {
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
}

func testSumRegion(m NumMatrix, r1, c1, r2, c2, expectedResult int) {
	result := m.SumRegion(r1, c1, r2, c2)

	fmt.Printf("Sum of region [%v][%v] - [%v][%v]: %v.", r1, c1, r2, c2, result)
	fmt.Printf("Expected sum: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func testUpdate(m NumMatrix, r, c, value int) {
	m.Update(r, c, value)
	fmt.Printf("Updated m[%v][%v] to %v.", r, c, value)

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
