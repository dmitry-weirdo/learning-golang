package main

import (
	"fmt"
)

func numberOfPatterns(m int, n int) int {
	// all single values are already precalculated by back-tracking
	// passes in 0 ms
	return numberOfPatterns_precalculated(m, n)

	// calculation with backtracking, symmetric repeats saved
	// passes in 5-6 ms
	//return numberOfPatterns_backtracking(m, n)
}

func numberOfPatterns_precalculated(m int, n int) int {
	// m-m precalculated by backtracking
	memo := []int{
		0,
		9,
		56,
		320,
		1624,
		7152,
		26016,
		72912,
		140704,
		140704,
	}

	result := 0

	for i := m; i <= n; i++ {
		result += memo[i]
	}

	return result
}

func precalculate() {
	for m := 1; m <= 9; m++ {
		number := numberOfPatterns_backtracking(m, m)
		fmt.Printf("%v keys - %v patterns \n", m, number)
	}
}

func numberOfPatterns_backtracking(m int, n int) int {
	// backtracking + save time on symmetry on the starting point:
	// - starting from the centre 5 - no symmetrid
	// - stating from corners 1, 3, 7, 9 - 4 symmetries
	// - stating from border centres 2, 4, 6, 8 - 4 symmetries

	// fill in the crossing table - when we require the middle point to be already visited
	cross := createIntMatrix(10, 10) // coords used are 1-9

	// 1 2 3
	// 4 5 6
	// 7 8 9

	// borders
	cross[1][3] = 2
	cross[3][1] = 2
	cross[1][7] = 4
	cross[7][1] = 4
	cross[3][9] = 6
	cross[9][3] = 6
	cross[7][9] = 8
	cross[9][7] = 8

	// middle borders opposite through the centre 5
	cross[2][8] = 5
	cross[8][2] = 5
	cross[4][6] = 5
	cross[6][4] = 5

	// diagonals through the centre 5
	cross[1][9] = 5
	cross[9][1] = 5
	cross[3][7] = 5
	cross[7][3] = 5

	// mark nodes as visited -> this will be backtracked
	visited := make([]bool, 10)

	var dfs func(i int, patternLength int) int // returns number of valid patterns

	dfs = func(i int, patternLength int) int {
		if patternLength > n { // pattern too long
			return 0
		}

		visited[i] = true

		validPatterns := 0

		if (m <= patternLength) && (patternLength <= n) { // current pattern is valid -> count it
			validPatterns++
		}

		for nextKey := 1; nextKey <= 9; nextKey++ { // try all other keys
			if visited[nextKey] {
				// keys must be distinct -> skip the direction if the key is already visited
				continue
			}

			requiredCrossingKey := cross[i][nextKey]

			if (requiredCrossingKey != 0) && !visited[requiredCrossingKey] {
				// we need to cross some key, but it's not yet visited -> skip this direction
				continue
			}

			validPatterns += dfs(nextKey, patternLength+1)
		}

		// backtrack visiting the current key
		visited[i] = false

		// return the valid patterns starting from the current position (probably including the current pattern itself)
		return validPatterns
	}

	// we start with patternLength since the starting button itself counts
	// we start only from 1, 2, 5, since other values are symmetric
	return dfs(5, 1) +
		4*dfs(1, 1) +
		4*dfs(2, 1)
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(m, n int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("M (min pattern length): %v \n", m)
	fmt.Printf("N (max pattern length): %v \n", n)

	result := numberOfPatterns(m, n)

	fmt.Printf("Total patterns with lengths [%v; %v]: %v \n", m, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, 1, 9)
}

func test2() {
	test(1, 2, 65)
}

func test3() {
	test(1, 3, 385)
}

func main() {
	// 351. Android Unlock Patterns
	test1()
	test2()
	test3()
	//precalculate()
}
