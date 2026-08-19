package main

import (
	"fmt"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// todo: bitmask matching must be faster

	// trivial fixing,
	// fails TLE on test-case 48/53,
	// n = 1000000000, reservedSeats = 6916
	return maxNumberOfFamilies_naive(n, reservedSeats)
}

func maxNumberOfFamilies_naive(n int, reservedSeats [][]int) int {
	// n - total rows numbered 1-based

	// collect map to position list from reservedSeats
	// todo: maybe better as array?

	fmt.Printf("Total reserved seats: %v. \n", len(reservedSeats))

	// todo: probably sort by row, column first
	m := make(map[int][]int)
	row := 0
	column := 0

	for _, v := range reservedSeats {
		row = v[0]
		column = v[1]

		// seats 1 and 10 have no value
		if column == 1 || column == 10 {
			continue
		}

		if _, ok := m[row]; ok {
			// list already exists
			m[row] = append(m[row], column)
		} else {
			m[row] = []int{column}
		}
	}

	fmt.Printf("Reserved seats grouped by row: %v \n", m)

	count := 0

	blocked2345 := false
	blocked4567 := false
	blocked6789 := false

	for row = 1; row <= n; row++ {
		if _, ok := m[row]; !ok {
			// no reserved seats -> we can allocate 2
			count += 2
			continue
		}

		// more than 4 seats reserved -> we can allocate nothing
		if len(m[row]) > 4 {
			continue
		}

		blocked2345 = false
		blocked4567 = false
		blocked6789 = false

		// todo: bit masking should be faster than this spaghetti
		for _, r := range m[row] {
			if r < 4 { // 2, 3
				blocked2345 = true
			} else if r <= 5 { // 4, 5
				blocked2345 = true
				blocked4567 = true
			} else if r <= 7 { // 6, 7
				blocked4567 = true
				blocked6789 = true
			} else { // 8, 9
				blocked6789 = true
			}
		}

		if !blocked2345 || !blocked4567 || !blocked6789 {
			// since at least one place is blocked, we cal reserve at most one slot
			count++
		}
	}

	return count
}

func test(n int, m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - number of rows with 10 seats: %v \n", n)
	fmt.Printf("Reserved seats: %v \n", m)

	result := maxNumberOfFamilies(n, m) // todo: replace with your function

	fmt.Printf("Max 4-group allocations allowed: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 3

	reservedSeats := [][]int{
		{1, 2},
		{1, 3},
		{1, 8},
		{2, 6},
		{3, 1},
		{3, 10},
	}

	expected := 4

	test(n, reservedSeats, expected)
}

func test2() {
	n := 2

	reservedSeats := [][]int{
		{2, 1},
		{1, 8},
		{2, 6},
	}

	expected := 2

	test(n, reservedSeats, expected)
}

func test3() {
	n := 4

	reservedSeats := [][]int{
		{4, 3},
		{1, 4},
		{4, 6},
		{1, 7},
	}

	expected := 4

	test(n, reservedSeats, expected)
}

func main() {
	// 1386. Cinema Seat Allocation
	test1()
	test2()
	test3()
}
