package main

import (
	"fmt"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// bitmask matching must be faster

	// Still fails TLE on test-case 48/53,
	// n = 1000000000, reservedSeats = 6916

	// the thing is -> we must not iterate all N rows, we should only calculate rows where there are reserved seats!
	// Yes, fixing this passes in around 10-15 ms
	//return maxNumberOfFamilies_bitMasks(n, reservedSeats)

	// trivial fixing,
	// fails TLE on test-case 48/53,
	// n = 1000000000, reservedSeats = 6916

	// Fixed the iteration to "iterate only rows with meaningful reserved seats"
	// it now passes withing 15-25 ms
	return maxNumberOfFamilies_naive(n, reservedSeats)
}

func maxNumberOfFamilies_bitMasks(n int, reservedSeats [][]int) int {
	// todo: for better readability, we can set bits in the opposite order:
	// seats: 2 3 4 5 6 7 8 9
	// bits:  7 6 5 4 3 2 1 0

	// we have just significant values for 2-9 = 8 bits
	// significant bits are 0000 bits
	// bit 0 (smallest) is seat 2, bit 7 (biggest) is position seat 9:
	// seats: 9 8 7 6 5 4 3 2
	// bits:  7 6 5 4 3 2 1 0

	// to match with the mask, we do the bitwise OR
	// i.e. bits = 1 in the mask will stay 1, i.e. these bits are non-meaningful
	// but bits = 0 will change to 1 if the bit of the row reserved seat is 1.
	const mask2345 = 0b11110000
	const mask4567 = 0b11000011
	const mask6789 = 0b00001111

	// for all reserved seats, collect reserved bit-masks for this row
	m := getBitMasksForRowsWithReservedSeats(reservedSeats)

	reservedRowsCount := len(m)
	rowsWithNoReservedSeatsCount := n - reservedRowsCount

	//fmt.Printf("Bit masks seats grouped by row: %08b \n", m)
	//fmt.Printf("Total rows with any reserved seats in [2; 8] range: %v \n", reservedRowsCount)
	//fmt.Printf("Total rows no reserved seats count in [2; 8] range: %v. Each of this rows will provide 2 reserved positions. \n", rowsWithNoReservedSeatsCount)

	// every row with no reserved places in [2; 8] will provide 2 places -> no need to iterate these rows
	count := 2 * rowsWithNoReservedSeatsCount

	// we're only iterating the rows with meaningful reserved seats
	for _, rowReservedSeatsBitMask := range m {
		// there is at least 1 reserved seat in [2; 9] range we can reserve either 0 or 1 seats

		// if ( mask | reservedSeats == mask )
		// -> No 0 bits changed to 1 with the meaningful reserved seat for this mask
		// -> We can allocate this range.

		// If we can allocate any of the seats, we can allocate only 1 4-place seat for this row.
		// (since at least 1 seat is reserved).
		if (mask2345|rowReservedSeatsBitMask == mask2345) ||
			(mask4567|rowReservedSeatsBitMask == mask4567) ||
			(mask6789|rowReservedSeatsBitMask == mask6789) {
			count += 1
		}
	}

	return count
}

func getBitMasksForRowsWithReservedSeats(reservedSeats [][]int) map[int]int {
	m := make(map[int]int) // 1-based row -> bit mask of reserved seats 2-8 (see the format above)

	row := 0
	column := 0
	bit := 0

	for _, v := range reservedSeats {
		row = v[0]
		column = v[1]

		// seats 1 and 10 have no value
		if column == 1 || column == 10 {
			continue
		}

		// bit for the value.
		// we count seats from 2: 2 -> bit 0, 3 -> bit 1,... 9 -> bit 7
		bit = 1 << (column - 2)

		if _, ok := m[row]; ok {
			// mask for this row already exists -> add a bit to it
			m[row] = m[row] | bit
		} else { // put the first bit for this row
			m[row] = bit
		}
	}

	return m
}

func maxNumberOfFamilies_naive(n int, reservedSeats [][]int) int {
	// n - total rows numbered 1-based
	//fmt.Printf("Total reserved seats: %v. \n", len(reservedSeats))

	m := getRowToReservedSeatsMap(reservedSeats)
	//fmt.Printf("Reserved seats grouped by row: %v \n", m)

	reservedRowsCount := len(m)
	rowsWithNoReservedSeatsCount := n - reservedRowsCount

	//fmt.Printf("Total rows with any reserved seats in [2; 8] range: %v \n", reservedRowsCount)
	//fmt.Printf("Total rows no reserved seats count in [2; 8] range: %v. Each of this rows will provide 2 reserved positions. \n", rowsWithNoReservedSeatsCount)

	// every row with no reserved places in [2; 8] will provide 2 places -> no need to iterate these rows
	count := 2 * rowsWithNoReservedSeatsCount

	blocked2345 := false
	blocked4567 := false
	blocked6789 := false

	// we're only iterating the rows with meaningful reserved seats
	for _, rowReservedSeats := range m {
		// more than 4 seats reserved -> we can allocate nothing
		if len(rowReservedSeats) > 4 {
			continue
		}

		blocked2345 = false
		blocked4567 = false
		blocked6789 = false

		// todo: bit masking should be faster than this spaghetti
		for _, r := range rowReservedSeats {
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

func getRowToReservedSeatsMap(reservedSeats [][]int) map[int][]int {
	// collect map to position list from reservedSeats
	// todo: maybe better as array?
	// todo: probably sort by row, column first?
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
		} else { // create new list
			m[row] = []int{column}
		}
	}

	return m
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
