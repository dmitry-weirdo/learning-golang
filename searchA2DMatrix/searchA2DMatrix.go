package main

import (
	"demo/matrixCommon"
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	// todo: we can also treat a matrix as one array, with (row = idx / n) and (column = idx % n)
	rows := len(matrix)
	columns := len(matrix[0])

	// first, we binary-search via rows to find a POTENTIAL row
	potentialRowIndex := findPotentialRow(matrix, rows, columns, target)

	fmt.Printf("Potential row index for target value %v: %v \n", target, potentialRowIndex)

	if potentialRowIndex < 0 {
		fmt.Println("There is no potential row. Returning false.")
		return false
	}

	// then, we binary-search within a row
	row := matrix[potentialRowIndex]

	fmt.Printf("Potential row for target value %v: %v \n", target, row)

	left := 0
	right := columns - 1

	for left <= right {
		mid := (left + right) / 2

		if row[mid] > target { // move to left from mid
			right = mid - 1
		} else if row[mid] < target { // move to right from mid
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}

func findPotentialRow(m [][]int, rows int, columns int, target int) int { // returns potential row or -1
	left := 0
	right := rows - 1 // todo: maybe rows?

	for left <= right {
		mid := (left + right) / 2

		row := m[mid]

		if row[0] > target { // move up to upper row
			right = mid - 1
		} else if row[columns-1] < target { // move down to lower row
			left = mid + 1
		} else { // potential row found
			return mid
		}
	}

	return -1
}

func test(matrix [][]int, target int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Println("Matrix: ")
	matrixCommon.PrintIntMatrix(matrix)

	fmt.Printf("Target value: %v \n", target)

	result := searchMatrix(matrix, target)

	fmt.Printf("Target value %v found in matrix: %v \n", target, result)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 3

	expected := true

	test(m, target, expected)
}

func test2() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := -1

	expected := false

	test(m, target, expected)
}

func test3() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 100

	expected := false

	test(m, target, expected)
}

func test4() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 8 // between rows

	expected := false

	test(m, target, expected)
}

func test5() {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 6 // potentially within row 0, but not there

	expected := false

	test(m, target, expected)
}

func main() {
	// 74. Search a 2D Matrix
	test1()
	test2()
	test3()
	test4()
	test5()
}
