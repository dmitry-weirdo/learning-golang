package main

import "fmt"

var directions = [][]int{
	{0, 1},  // right // we start from going right
	{1, 0},  // bottom
	{0, -1}, // left
	{-1, 0}, // top
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	totalCells := rows * cols

	i := rStart
	j := cStart

	cellsFilled := 1

	pathIndexes := createIntMatrix(totalCells, 2) // every cell is a pair of [i; j]

	// we're guaranteed that the first cell is within the matrix -> we can add it to the result without cellExists check
	pathIndexes[0][0] = i
	pathIndexes[0][1] = j

	segmentLength := 1

	const directionsWithSameSegmentLength = 2
	directionIndex := 0 // start with right direction
	direction := directions[directionIndex]

	// we're going 1 right, 1 down, 2 left, 2 up, 3 right, 3 down, etc
	for cellsFilled < totalCells {
		for range directionsWithSameSegmentLength { // 2 directions
			for range segmentLength { // use current segment length
				i += direction[0]
				j += direction[1]

				if cellExists(rows, cols, i, j) {
					pathIndexes[cellsFilled][0] = i
					pathIndexes[cellsFilled][1] = j

					cellsFilled++
				}
			}

			// change to the next direction
			directionIndex++
			directionIndex = directionIndex % len(directions)
			direction = directions[directionIndex]
		}

		// increase segment length after 2 directions
		segmentLength++
	}

	return pathIndexes
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func test(rows, columns, startRow, startColumn int, expectedResult [][]int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix rows: %v \n", rows)
	fmt.Printf("Matrix columns: %v \n", columns)
	fmt.Printf("Start row: %v \n", startRow)
	fmt.Printf("Start column: %v \n", startColumn)

	result := spiralMatrixIII(rows, columns, startRow, startColumn)

	fmt.Printf("Matrix indexes visited spirally: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	rows := 1
	columns := 4
	startRow := 0
	startColumn := 0

	expected := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	test(rows, columns, startRow, startColumn, expected)
}

func test2() {
	rows := 5
	columns := 6
	startRow := 1
	startColumn := 4

	expected := [][]int{
		{1, 4},
		{1, 5},
		{2, 5},
		{2, 4},
		{2, 3},
		{1, 3},
		{0, 3},
		{0, 4},
		{0, 5},
		{3, 5},
		{3, 4},
		{3, 3},
		{3, 2},
		{2, 2},
		{1, 2},
		{0, 2},
		{4, 5},
		{4, 4},
		{4, 3},
		{4, 2},
		{4, 1},
		{3, 1},
		{2, 1},
		{1, 1},
		{0, 1},
		{4, 0},
		{3, 0},
		{2, 0},
		{1, 0},
		{0, 0},
	}

	test(rows, columns, startRow, startColumn, expected)
}

func main() {
	// 885. Spiral Matrix III
	test1()
	test2()
}
