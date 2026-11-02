package main

import (
	"demo/matrixCommon"
	"fmt"
)

func findLonelyPixel(picture [][]byte) int {
	const BLACK = 'B'
	const WHITE = 'W'

	// count the black pixels for every column and/or every row
	rows, columns := getRowsAndColumnsOfByteMatrix(picture)

	rowCounts := make([]int, rows)
	columnCounts := make([]int, columns)

	for i := range rows {
		for j := range columns {
			if picture[i][j] == BLACK {
				rowCounts[i]++
				columnCounts[j]++
			}
		}
	}

	//fmt.Printf("Row counts: %v \n", rowCounts)
	//fmt.Printf("Column counts: %v \n", columnCounts)

	result := 0

	for i := range rows {
		if rowCounts[i] != 1 { // multiple black pixels in this row -> skip this row
			continue
		}

		for j := range columns {
			if picture[i][j] == BLACK { // found a single black pixel in this row
				if columnCounts[j] == 1 {
					result++
				}

				continue
			}
		}
	}

	return result
}

func getRowsAndColumnsOfByteMatrix(mat [][]byte) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func test(m [][]byte, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of white and black pixels: \n")
	matrixCommon.PrintByteMatrix(m)

	result := findLonelyPixel(m)

	fmt.Printf("Count of black lonely pixels: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}

	test(m, 3)
}

func test2() {
	m := [][]byte{
		{'B', 'B', 'B'},
		{'B', 'B', 'W'},
		{'B', 'B', 'B'},
	}

	test(m, 0)
}

func main() {
	// 531. Lonely Pixel I
	test1()
	test2()
}
