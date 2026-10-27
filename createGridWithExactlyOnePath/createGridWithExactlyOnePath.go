package main

import (
	"fmt"
	"strings"
)

func createGrid(m int, n int) []string {
	const FREE = "."
	const OBSTACLE = "#"

	firstRow := strings.Repeat(FREE, n)
	nonFirstRow := strings.Repeat(OBSTACLE, n-1) + FREE

	result := make([]string, m)
	result[0] = firstRow

	for i := 1; i < m; i++ {
		result[i] = nonFirstRow
	}

	return result
}

func test(rows, columns int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Rows: %v \n", rows)
	fmt.Printf("Columns: %v \n", columns)

	result := createGrid(rows, columns)

	fmt.Printf("Grid with only one path: %v \n", result)
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
	test(
		2,
		3,
		[]string{
			"...",
			"##.",
		},
	)
}

func test2() {
	test(
		3,
		3,
		[]string{
			"...",
			"##.",
			"##.",
		},
	)
}

func main() {
	// 3963. Create Grid With Exactly One Path
	test1()
	test2()
}
